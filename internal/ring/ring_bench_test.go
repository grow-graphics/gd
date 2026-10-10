package ring

import (
	"runtime"
	"testing"
	"unsafe"

	_ "graphics.gd/internal/ring/internal/stubflush" // stands in for the engine's C flush
)

// shapeVector2 is one Vector2 argument and no result: the shape of
// Node2D.set_position, which is the call a scene full of moving nodes makes
// once per node per frame, and so the one worth measuring.
const shapeVector2 = 5 << 4

// BenchmarkBufferDrain models what a frame actually does to the ring: a node
// buffers one call, and the drain executes it as the callback returns — gd.c's
// gd_ring_drain runs after every engine->Go callback, so enqueue and drain
// alternate one entry at a time.
//
// That alternation is why [Entry] is not sized for the cache. It is tempting
// to look at the entry (392 bytes, with the PC field a further 360 bytes past
// the header) and the ring (Size entries, ~98KB, swept hundreds of times a
// frame) and conclude that packing an entry into one 64-byte line would keep
// the ring inside L1 and pay for itself. It does not: an entry is read
// microseconds after it is written and is still in L1 whatever the stride, so
// the ring is a store-and-forward buffer of depth one here, not a working set
// — nothing is ever revisited after eviction. Measured on this benchmark, a
// 64-byte entry came out marginally SLOWER than the 392-byte one (3.37 vs
// 3.27 ns/op on BufferBatch), and declining oversized argument packs so the
// entry could stay small cost a further 6%. Shrink the entry for memory if
// that ever matters; do not expect throughput from it.
func BenchmarkBufferDrain(b *testing.B) {
	var arg [16]byte
	var sink uintptr
	Main.head, Main.tail, Main.floor = 0, 0, 0
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Main.Buffer(uintptr(i), 42, shapeVector2, unsafe.Pointer(&arg), 7)
		for Main.tail != Main.head {
			e := &Main.Entries[Main.tail&Mask]
			sink += e.Object + e.Method + uintptr(e.Shape) + e.PC + uintptr(e.Args[0])
			Main.tail++
			Main.floor = Main.tail
		}
	}
	runtime.KeepAlive(sink)
}

// BenchmarkBufferBatch is the same walk without the interleaved drain: the
// ring fills before anything reads it, which is what a caller outside a
// callback does. It isolates the enqueue side.
func BenchmarkBufferBatch(b *testing.B) {
	var arg [16]byte
	Main.head, Main.tail, Main.floor = 0, 0, 0
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if Main.head-Main.tail >= Size-1 {
			Main.tail, Main.floor = Main.head, Main.head // discard; the flush is not what is measured
		}
		Main.Buffer(uintptr(i), 42, shapeVector2, unsafe.Pointer(&arg), 7)
	}
}
