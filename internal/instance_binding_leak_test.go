package gd_test

import (
	"testing"

	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/Performance"
	"graphics.gd/variant/Object"
)

// TestInstanceBindingFetchNoLeak casts one plain (non-extension) Node many
// times. Each cast looks up the object's Go instance binding. A lookup must
// not store anything on the object, so engine static memory must not grow
// with the number of casts.
func TestInstanceBindingFetchNoLeak(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		const casts = 20000
		node := Node.New()
		before := Performance.GetMonitor(Performance.MemoryStatic)
		misses := 0
		for range casts {
			if _, ok := Object.As[Node3D.Instance](node); !ok {
				misses++
			}
		}
		grew := Performance.GetMonitor(Performance.MemoryStatic) - before
		node.QueueFree()
		if misses != casts {
			t.Fatalf("%d of %d casts of a Node to Node3D succeeded", casts-misses, casts)
		}
		if before == 0 {
			t.Skip("this engine build does not report static memory")
		}
		if grew > 64*1024 {
			t.Errorf("%d casts grew engine static memory by %.0f bytes, want < 65536", casts, grew)
		}
	})
}
