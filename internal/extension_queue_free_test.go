//go:build !generate

package gd_test

import (
	"runtime"
	"testing"
	"time"
)

// TestExtensionQueueFreeThenCollected makes extension nodes in Go, asks the
// engine to free them (queue_free) and drops them, so that the garbage
// collector's cleanup and the engine both want to free the same object in
// the same frame. Each object must be freed once.
func TestExtensionQueueFreeThenCollected(t *testing.T) {
	registerPanicProcess()
	for range 5 {
		runOnMain(t, func(t testing.TB) {
			for range 64 {
				node := new(PanicProcessNode)
				node.AsNode().QueueFree()
			}
			runtime.GC()
			// let the cleanups run before the frame ends
			if runtime.GOOS == "js" {
				// a callback cannot block on wasm, the engine is waiting
				// on this very call: yield to the other goroutines.
				for range 100 {
					runtime.Gosched()
				}
			} else {
				time.Sleep(time.Millisecond)
			}
			runtime.GC()
		})
		for range 3 { // the frames in which the frees happen
			runOnMain(t, func(t testing.TB) {})
		}
	}
}
