//go:build !cgo && !js

package ring

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

// Adopt is a no-op without cgo: there is no C side to drain the ring.
func Adopt() {}

// copies[d] holds the batch of a flush of Main at nesting depth d. A call in
// the batch can call back into Go, which can buffer more than the rest of the
// ring while the batch runs; running from a copy keeps those new entries from
// overwriting the batch's calls that have not run yet.
var copies []*[Size]Entry

// depth counts the flushes of Main that are running.
var depth int

func flush(entries unsafe.Pointer, tail, head uint32) {
	ring := (*[Size]Entry)(entries)
	if entries == unsafe.Pointer(&Main.Entries[0]) {
		if depth == len(copies) {
			copies = append(copies, new([Size]Entry))
		}
		batch := copies[depth]
		for i := tail; i != head; i++ {
			batch[i&Mask] = ring[i&Mask]
		}
		ring = batch
		depth++
		defer func() { depth-- }()
	}
	for i := tail; i != head; i++ {
		CrashIndex = i & Mask
		e := &ring[i&Mask]
		gdextension.Host.Objects.Unsafe.Call(
			gdextension.Object(e.Object),
			gdextension.MethodForClass(e.Method),
			gdextension.CallReturns[any](unsafe.Pointer(&e.Result[0])),
			gdextension.Shape(e.Shape),
			gdextension.CallAccepts[any](unsafe.Pointer(&e.Args[0])),
		)
	}
	CrashIndex = 0xFFFFFFFF
}
