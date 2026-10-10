//go:build !generate

package gd_test

import (
	"fmt"
	"sync"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/noescape"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

// PanicProcessNode's _process panics on demand.
type PanicProcessNode struct {
	Node.Extension[PanicProcessNode] `gd:"PanicProcessNode"`

	boom bool
}

func (b *PanicProcessNode) Process(delta Float.X) {
	if b.boom {
		panic("a panic in a virtual method (expected by TestVirtualPanic)")
	}
}

var registerPanicProcess = sync.OnceFunc(func() { classdb.Register[PanicProcessNode]() })

// TestVirtualPanic panics inside a virtual method that the engine calls
// while it runs a batch of queued calls. The panic must stop at the
// callback, like one in a signal handler does: it must not unwind through
// the engine into the Go code that made the call, and the calls queued
// behind the one that panicked must still run.
func TestVirtualPanic(t *testing.T) {
	registerPanicProcess()
	runOnMain(t, func(t testing.TB) {
		node := new(PanicProcessNode)
		object := gd.ObjectChecked(node.AsObject())
		class, method := gd.NewStringName("Object"), gd.NewStringName("notification")
		defer class.Free()
		defer method.Free()
		notify := gdextension.Host.Objects.Method.Lookup(pointers.Get(class), pointers.Get(method), notificationHash)
		if notify == 0 {
			t.Fatal("Object.notification method bind not found: the hash has drifted from this engine build")
		}
		args := notifyArgs{what: int64(Node.NotificationProcess)}
		marker := Node.New()
		marker.SetName("Before")
		_ = marker.Name()
		node.boom = true
		for round := range 3 {
			want := fmt.Sprintf("After%d", round)
			var reached any
			var got string
			func() {
				defer func() { reached = recover() }()
				noescape.Call[struct{}](object, notify, shapeNotification, &args) // queued: runs Process
				marker.SetName(want)                                              // queued behind it
				got = marker.Name()                                               // runs the batch
			}()
			if reached != nil {
				t.Fatalf("round %d: the panic unwound through the engine: %v", round, reached)
			}
			if got != want {
				t.Fatalf("round %d: the call queued behind the panic was lost: got %q, want %q", round, got, want)
			}
		}
		node.boom = false
		for i := range 600 { // more than two laps of the ring
			want := fmt.Sprintf("N%d", i)
			marker.SetName(want)
			if i%7 == 0 {
				if got := marker.Name(); got != want {
					t.Fatalf("after the panics, call %d: got %q, want %q", i, got, want)
				}
			}
		}
		marker.QueueFree()
		node.AsNode().QueueFree()
	})
}

// PanicPropertyNode's dynamic property hooks panic.
type PanicPropertyNode struct {
	Node.Extension[PanicPropertyNode] `gd:"PanicPropertyNode"`

	gets, sets int
}

func (p *PanicPropertyNode) Get(name string) any {
	p.gets++
	panic("a panic in a property getter (expected by TestPropertyPanic)")
}

func (p *PanicPropertyNode) Set(name string, value any) bool {
	p.sets++
	panic("a panic in a property setter (expected by TestPropertyPanic)")
}

var registerPanicProperty = sync.OnceFunc(func() { classdb.Register[PanicPropertyNode]() })

// TestPropertyPanic panics inside the property hooks of an extension class.
// Like a panic in a virtual method, it must stop where the engine called
// into Go: the engine then treats the property as one the hooks do not
// handle, and the Go code that asked for it carries on.
func TestPropertyPanic(t *testing.T) {
	registerPanicProperty()
	runOnMain(t, func(t testing.TB) {
		node := new(PanicPropertyNode)
		node.AsNode().SetName("Probe")
		marker := Node.New()
		for round := range 3 {
			want := fmt.Sprintf("After%d", round)
			var reached any
			var name any
			var got string
			func() {
				defer func() { reached = recover() }()
				name = Object.Get(node, "name")         // the getter panics, the engine answers
				Object.Set(node, "no_such_property", 1) // queued: the setter panics
				marker.SetName(want)                    // queued behind it
				got = marker.Name()                     // runs the batch
			}()
			if reached != nil {
				t.Fatalf("round %d: the panic unwound through the engine: %v", round, reached)
			}
			if name != "Probe" {
				t.Fatalf("round %d: got name %v, want %q", round, name, "Probe")
			}
			if got != want {
				t.Fatalf("round %d: the call queued behind the panic was lost: got %q, want %q", round, got, want)
			}
		}
		if node.gets == 0 || node.sets == 0 {
			t.Fatalf("the hooks did not run: %d gets, %d sets", node.gets, node.sets)
		}
		marker.QueueFree()
		node.AsNode().QueueFree()
	})
}
