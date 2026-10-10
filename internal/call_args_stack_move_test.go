//go:build !generate

package gd_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	gd "graphics.gd/internal"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
)

// StackMoveGetter declines every dynamic property, so that the engine goes
// on to look the property up itself: Object::get takes the name by reference
// and still uses it after the getter has returned. Before it declines it
// either passes the question on to the next node, or (the last node) moves
// the goroutine stack and gives Go reasons to reuse the old one.
type StackMoveGetter struct {
	Node.Extension[StackMoveGetter] `gd:"StackMoveGetter"`

	next *StackMoveGetter
}

var stackMoveSink []byte

func (s *StackMoveGetter) Get(name string) any {
	if s.next != nil {
		Object.Get(s.next, "name")
		return nil
	}
	growStack(4000)
	var wg sync.WaitGroup
	for range 64 {
		wg.Go(func() { growStack(48) })
	}
	wg.Wait()
	for range 64 {
		b := make([]byte, 64<<10)
		for i := range b {
			b[i] = 0xAA
		}
		stackMoveSink = b
	}
	return nil
}

var registerStackMoveGetter = sync.OnceFunc(func() { classdb.Register[StackMoveGetter]() })

// TestCallArgsStackMove makes engine calls whose callbacks move the calling
// goroutine's stack while the engine still holds the calls' arguments by
// reference. The chain of nodes is deeper than the resident call packs, so
// the innermost calls take the plain cgo path in every build. The engine
// must not read its arguments from the stack that was left behind.
func TestCallArgsStackMove(t *testing.T) {
	registerStackMoveGetter()
	runOnMain(t, func(t testing.TB) {
		const depth = 10
		var nodes [depth]*StackMoveGetter
		for i := range nodes {
			nodes[i] = new(StackMoveGetter)
			nodes[i].AsNode().SetName("Probe")
		}
		for i := 0; i < depth-1; i++ {
			nodes[i].next = nodes[i+1]
		}
		for round := range 30 {
			runtime.GC() // lets the stack shrink, so that every round moves it again
			if got := Object.Get(nodes[0], "name"); got != "Probe" {
				t.Fatalf("round %d: got %v, want %q", round, got, "Probe")
			}
		}
		for _, node := range nodes {
			node.AsNode().QueueFree()
		}
	})
}

// TestCallableResultStackMove calls a Go function through the engine's
// Callable (callv), and the function moves the goroutine stack before it
// returns. The engine writes the result after that: it must still arrive.
func TestCallableResultStackMove(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		engine := gd.InternalCallable(Callable.New(func() int {
			growStack(4000)
			var wg sync.WaitGroup
			for range 64 {
				wg.Go(func() { growStack(48) })
			}
			wg.Wait()
			return 42
		}))
		for round := range 5 {
			runtime.GC() // lets the stack shrink, so that every round moves it again
			if got := engine.Call().ConvenientInterface(); fmt.Sprint(got) != "42" {
				t.Fatalf("round %d: got %v, want 42", round, got)
			}
		}
	})
}
