// Package gdmemory provides functions for transferring data between Go and the graphics engine.
//
// This package is primarily used on platforms where the extension is running in a different
// address space, ie. web/wasm.
package gdmemory

import (
	"sync"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

// buffers is one set of engine-side buffers for the calls made at one
// callback depth.
//
// A call can still be using its buffers when the engine calls back into Go:
// Object.emit_signal hands each connected handler pointers into the argument
// buffer of the emit_signal call. So an engine call made from inside a
// callback must not reuse the buffers of the call that is running. Each
// callback depth (see depth) has its own set. Calls at the same depth run one
// after another, so they share that depth's set.
type buffers struct {
	arguments gdextension.Pointer
	receiver  gdextension.Pointer
	results   [2]gdextension.Pointer
	current   int
}

// frames[d] holds the buffers of callback depth d, allocated on first use and
// kept for the life of the program.
var frames []*buffers

func frame() *buffers {
	d := depth()
	for len(frames) <= d {
		b := &buffers{
			arguments: gdextension.Host.Memory.Malloc(64 * 64),
			receiver:  gdextension.Host.Memory.Malloc(64 * 64 / 16),
		}
		for i := range b.results {
			b.results[i] = gdextension.Host.Memory.Malloc(64 * 64)
			gdextension.Host.Memory.Clear(b.results[i], 64*64)
		}
		frames = append(frames, b)
	}
	return frames[d]
}

var setup = sync.OnceFunc(func() { frame() })

// CopyArguments copies arguments from args to the arguments buffer, respecting Go's alignment rules.
func CopyArguments[T any](shape gdextension.Shape, args gdextension.CallAccepts[T]) gdextension.Pointer {
	setup()
	if args == nil {
		return 0
	}
	return copyIntoEngine(shape.SizeArguments(), unsafe.Pointer(args), frame().arguments)
}

// CopyReceiver copies the receiver into the receiver buffer, respecting Go's alignment rules.
func CopyReceiver[T any](shape gdextension.Shape, self gdextension.CallMutates[T]) gdextension.Pointer {
	setup()
	if self == nil {
		return 0
	}
	return copyIntoEngine((shape & 0b11110000).SizeArguments(), unsafe.Pointer(self), frame().receiver)
}

func copyIntoEngine(bytes int, args unsafe.Pointer, into gdextension.Pointer) gdextension.Pointer {
	if into == 0 {
		panic("nil pointer dereference")
	}
	buf := unsafe.Slice((*byte)(args), bytes)
	ptr := into
	off := gdextension.Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			gdextension.Host.Memory.Edit.Uint64(ptr+off, *(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gdextension.Host.Memory.Edit.Uint32(ptr+off, *(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gdextension.Host.Memory.Edit.Uint16(ptr+off, *(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gdextension.Host.Memory.Edit.Byte(ptr+off, *(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return ptr
}

func MakeResult(shape gdextension.Shape) gdextension.Pointer {
	setup()
	// alternating between two buffers for results
	b := frame()
	if b.current == 0 {
		b.current = 1
	} else {
		b.current = 0
	}
	return b.results[b.current]
}

func LoadResult[T ~unsafe.Pointer](shape gdextension.Shape, result T, from gdextension.Pointer) {
	setup()
	if from == 0 {
		panic("nil pointer dereference")
	}
	data := unsafe.Pointer(result)
	done := 0
	size := shape.SizeResult()
	if size == 0 {
		return
	}
	defer gdextension.Host.Memory.Clear(from, shape.SizeResult())
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(data, done)) = gdextension.Host.Memory.Load.Uint32(from + gdextension.Pointer(done))
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(data, done)) = gdextension.Host.Memory.Load.Uint16(from + gdextension.Pointer(done))
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(data, done)) = gdextension.Host.Memory.Load.Byte(from + gdextension.Pointer(done))
			done += 1
			size -= 1
		default:
			return
		}
	}
}

func CopyVariants[T any](args gdextension.CallAccepts[T], n int) gdextension.Pointer {
	setup()
	var offset int
	var data = unsafe.Pointer(args)
	arguments := frame().arguments
	for i := range n {
		gdextension.Host.Memory.Edit.Bits128(arguments+gdextension.Pointer(offset), *(*[2]uint64)(unsafe.Add(data, uintptr(i*24))))
		gdextension.Host.Memory.Edit.Uint64(arguments+gdextension.Pointer(offset+16), *(*uint64)(unsafe.Add(data, uintptr(i*24+16))))
		offset += 24
	}
	return arguments
}

func Int64frombits(bits uint64) int64 {
	return *(*int64)(unsafe.Pointer(&bits))
}

func CopyBufferToEngine(buf []byte) gdextension.Pointer {
	ptr := gdextension.Host.Memory.Malloc(len(buf))
	off := gdextension.Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			gdextension.Host.Memory.Edit.Uint64(ptr+off, *(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gdextension.Host.Memory.Edit.Uint32(ptr+off, *(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gdextension.Host.Memory.Edit.Uint16(ptr+off, *(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gdextension.Host.Memory.Edit.Byte(ptr+off, *(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return ptr
}

func CopyBufferToGo(ptr gdextension.Pointer, buf []byte) {
	if ptr == 0 {
		panic("nil pointer dereference")
	}
	off := 0
	for len(buf) > 0 {
		switch {
		case len(buf) >= 4:
			*(*uint32)(unsafe.Pointer(&buf[0])) = gdextension.Host.Memory.Load.Uint32(ptr + gdextension.Pointer(off))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = gdextension.Host.Memory.Load.Uint16(ptr + gdextension.Pointer(off))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			*(*uint8)(unsafe.Pointer(&buf[0])) = gdextension.Host.Memory.Load.Byte(ptr + gdextension.Pointer(off))
			buf = buf[1:]
			off += 1
		}
	}
	gdextension.Host.Memory.Free(ptr)
}
