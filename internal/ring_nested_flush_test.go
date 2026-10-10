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
