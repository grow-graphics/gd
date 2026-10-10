package gd_test

import (
	"fmt"
	"testing"

	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Resource"
)

// TestRingNestedFlush queues a batch of two void calls on the main thread:
// Resource.emit_changed, then a rename of marker. The "changed" handler
// queues more than a ring (256) of renames of another node, so a nested
// flush runs while the outer batch is half done. The nested calls must not
// take the place of the outer batch's remaining call.
func TestRingNestedFlush(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		const nested = 300
		res := Resource.New()
		marker := Node.New()
		other := Node.New()
		marker.SetName("Before")
		res.OnChanged(func() {
			for i := range nested {
				other.SetName(fmt.Sprintf("N%d", i))
			}
		})
		_ = marker.Name() // a call with a result runs everything queued so far
		res.EmitChanged() // queued: its handler runs while the batch runs
		marker.SetName("After")
		if got := marker.Name(); got != "After" {
			t.Errorf("marker name: got %q, want %q (a nested flush overwrote a queued call)", got, "After")
		}
		if got, want := other.Name(), fmt.Sprintf("N%d", nested-1); got != want {
			t.Errorf("other name: got %q, want %q", got, want)
		}
		marker.QueueFree()
		other.QueueFree()
	})
}

// TestRingNestedFlushFullBatch runs a batch that fills the whole ring, with
// a call in the middle of it whose handler makes more calls than a ring
// holds. There is no room left for the handler's calls while the batch runs,
// and none of the batch's remaining calls may be lost to them.
func TestRingNestedFlushFullBatch(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		const nested = 300
		res := Resource.New()
		marker := Node.New()
		other := Node.New()
		res.OnChanged(func() {
			for i := range nested {
				other.SetName(fmt.Sprintf("N%d", i))
			}
		})
		_ = marker.Name()
		for i := range 120 {
			marker.SetName(fmt.Sprintf("A%d", i))
		}
		res.EmitChanged()
		for i := range 200 { // the ring fills, and flushes, part way through
			marker.SetName(fmt.Sprintf("B%d", i))
		}
		if got := marker.Name(); got != "B199" {
			t.Errorf("marker name: got %q, want %q", got, "B199")
		}
		if got, want := other.Name(), fmt.Sprintf("N%d", nested-1); got != want {
			t.Errorf("other name: got %q, want %q", got, want)
		}
		marker.QueueFree()
		other.QueueFree()
	})
}

// TestRingNestedFlushDeep nests the handlers several levels deep, each level
// queueing calls before and after it triggers the next one.
func TestRingNestedFlushDeep(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		const levels = 6
		const calls = 150
		var res [levels]Resource.Instance
		var nodes [levels]Node.Instance
		for i := range levels {
			res[i] = Resource.New()
			nodes[i] = Node.New()
		}
		for i := range levels {
			res[i].OnChanged(func() {
				for n := range calls {
					nodes[i].SetName(fmt.Sprintf("L%d-%d", i, n))
				}
				if i+1 < levels {
					res[i+1].EmitChanged()
				}
				for n := range calls {
					nodes[i].SetName(fmt.Sprintf("L%d-%d", i, calls+n))
				}
			})
		}
		_ = nodes[0].Name()
		res[0].EmitChanged()
		for i := range levels {
			if got, want := nodes[i].Name(), fmt.Sprintf("L%d-%d", i, 2*calls-1); got != want {
				t.Errorf("level %d: got %q, want %q", i, got, want)
			}
			nodes[i].QueueFree()
		}
	})
}

// TestRingNestedFlushLeftover reads a value right after a queued call whose
// handler queues another call. The handler's call must run before the read.
func TestRingNestedFlushLeftover(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		res := Resource.New()
		other := Node.New()
		other.SetName("Before")
		res.OnChanged(func() { other.SetName("After") })
		_ = other.Name()
		res.EmitChanged()
		if got := other.Name(); got != "After" {
			t.Errorf("got %q, want %q", got, "After")
		}
		other.QueueFree()
	})
}
