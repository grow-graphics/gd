//go:build js

package ring

import "unsafe"

// A flushed call can call back into Go, and that Go code can buffer and flush
// more calls while the outer flush is still running in the engine. Each flush
// nesting depth therefore copies into, and runs from, its own shadow ring: a
// nested flush can never write over the entries of a flush that has not
// finished. Flushes at the same depth run one after another, so they share it.
//
// shadowRings[d] is the Godot-memory copy of the ring for a flush at
// nesting depth d. Allocated on first use, kept for the life of the program.
var shadowRings []uint32

// flushDepth counts the flushes that are running (wasm is single-threaded).
var flushDepth int

//go:wasmimport gd ring_flush
func wasm_gd_ring_flush(entries uint32, stride uint32, tail uint32, head uint32)

//go:wasmimport gd memory_malloc
func wasm_gd_memory_malloc(size uint32) uint32

//go:wasmimport gd bulk_copy
func wasm_gd_bulk_copy(godot_dst uint32, go_src uint32, length uint32)

// Adopt is a no-op on wasm: the ring is flushed through the wasm imports.
func Adopt() {}

func flush(entries unsafe.Pointer, tail, head uint32) {
	entrySize := uint32(unsafe.Sizeof(Entry{}))
	if flushDepth == len(shadowRings) {
		shadowRings = append(shadowRings, wasm_gd_memory_malloc(Size*entrySize))
	}
	shadowRing := shadowRings[flushDepth]
	flushDepth++
	defer func() { flushDepth-- }()
	start := tail & Mask
	end := head & Mask
	base := uint32(uintptr(entries))
	if start < end {
		// Contiguous range
		wasm_gd_bulk_copy(
			shadowRing+start*entrySize,
			base+start*entrySize,
			(end-start)*entrySize,
		)
	} else {
		// Wraps around: copy [start..Size) then [0..end)
		wasm_gd_bulk_copy(
			shadowRing+start*entrySize,
			base+start*entrySize,
			(Size-start)*entrySize,
		)
		if end > 0 {
			wasm_gd_bulk_copy(
				shadowRing,
				base,
				end*entrySize,
			)
		}
	}
	wasm_gd_ring_flush(shadowRing, entrySize, tail, head)
}
