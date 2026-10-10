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
