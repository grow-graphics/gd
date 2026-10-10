//go:build !wasm

package gdmemory

// depth is always 0 outside wasm: there the engine reads arguments from Go
// memory, and these buffers are not used.
func depth() int { return 0 }
