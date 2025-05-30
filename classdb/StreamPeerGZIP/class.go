// Code generated by the generate package DO NOT EDIT

// Package StreamPeerGZIP provides methods for working with StreamPeerGZIP object instances.
package StreamPeerGZIP

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Angle"
import "graphics.gd/variant/Euler"
import "graphics.gd/classdb/StreamPeer"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID

type _ gdclass.Node

var _ gd.Object
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ Angle.Radians
var _ Euler.Radians
var _ = slices.Delete[[]struct{}, struct{}]

/*
ID is a typed object ID (reference) to an instance of this class, use it to store references to objects with
unknown lifetimes, as an ID will not panic on use if the underlying object has been destroyed.
*/
type ID Object.ID

func (id ID) Instance() (Instance, bool) { return Object.As[Instance](Object.ID(id).Instance()) }

/*
Extension can be embedded in a new struct to create an extension of this class.
T should be the type that is embedding this [Extension]
*/
type Extension[T gdclass.Interface] struct{ gdclass.Extension[T, Instance] }

/*
This class allows to compress or decompress data using GZIP/deflate in a streaming fashion. This is particularly useful when compressing or decompressing files that have to be sent through the network without needing to allocate them all in memory.
After starting the stream via [method start_compression] (or [method start_decompression]), calling [method StreamPeer.put_partial_data] on this stream will compress (or decompress) the data, writing it to the internal buffer. Calling [method StreamPeer.get_available_bytes] will return the pending bytes in the internal buffer, and [method StreamPeer.get_partial_data] will retrieve the compressed (or decompressed) bytes from it. When the stream is over, you must call [method finish] to ensure the internal buffer is properly flushed (make sure to call [method StreamPeer.get_available_bytes] on last time to check if more data needs to be read after that).
*/
type Instance [1]gdclass.StreamPeerGZIP

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.StreamPeerGZIP

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStreamPeerGZIP() Instance
}

/*
Start the stream in compression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
func (self Instance) StartCompression() error { //gd:StreamPeerGZIP.start_compression
	return error(gd.ToError(Advanced(self).StartCompression(false, int64(65535))))
}

/*
Start the stream in compression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
func (self Expanded) StartCompression(use_deflate bool, buffer_size int) error { //gd:StreamPeerGZIP.start_compression
	return error(gd.ToError(Advanced(self).StartCompression(use_deflate, int64(buffer_size))))
}

/*
Start the stream in decompression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
func (self Instance) StartDecompression() error { //gd:StreamPeerGZIP.start_decompression
	return error(gd.ToError(Advanced(self).StartDecompression(false, int64(65535))))
}

/*
Start the stream in decompression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
func (self Expanded) StartDecompression(use_deflate bool, buffer_size int) error { //gd:StreamPeerGZIP.start_decompression
	return error(gd.ToError(Advanced(self).StartDecompression(use_deflate, int64(buffer_size))))
}

/*
Finalizes the stream, compressing or decompressing any buffered chunk left.
*/
func (self Instance) Finish() error { //gd:StreamPeerGZIP.finish
	return error(gd.ToError(Advanced(self).Finish()))
}

/*
Clears this stream, resetting the internal state.
*/
func (self Instance) Clear() { //gd:StreamPeerGZIP.clear
	Advanced(self).Clear()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StreamPeerGZIP

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeerGZIP"))
	casted := Instance{*(*gdclass.StreamPeerGZIP)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Start the stream in compression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
//go:nosplit
func (self class) StartCompression(use_deflate bool, buffer_size int64) Error.Code { //gd:StreamPeerGZIP.start_compression
	var frame = callframe.New()
	callframe.Arg(frame, use_deflate)
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerGZIP.Bind_start_compression, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Start the stream in decompression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
//go:nosplit
func (self class) StartDecompression(use_deflate bool, buffer_size int64) Error.Code { //gd:StreamPeerGZIP.start_decompression
	var frame = callframe.New()
	callframe.Arg(frame, use_deflate)
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerGZIP.Bind_start_decompression, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Finalizes the stream, compressing or decompressing any buffered chunk left.
*/
//go:nosplit
func (self class) Finish() Error.Code { //gd:StreamPeerGZIP.finish
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerGZIP.Bind_finish, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears this stream, resetting the internal state.
*/
//go:nosplit
func (self class) Clear() { //gd:StreamPeerGZIP.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerGZIP.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsStreamPeerGZIP() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeerGZIP() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsStreamPeerGZIP() Instance { return self.Super().AsStreamPeerGZIP() }
func (self class) AsStreamPeer() StreamPeer.Advanced {
	return *((*StreamPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsStreamPeer() StreamPeer.Instance { return self.Super().AsStreamPeer() }
func (self Instance) AsStreamPeer() StreamPeer.Instance {
	return *((*StreamPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsRefCounted() [1]gd.RefCounted { return self.Super().AsRefCounted() }
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StreamPeer.Advanced(self.AsStreamPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StreamPeer.Instance(self.AsStreamPeer()), name)
	}
}
func init() {
	gdclass.Register("StreamPeerGZIP", func(ptr gd.Object) any {
		return [1]gdclass.StreamPeerGZIP{*(*gdclass.StreamPeerGZIP)(unsafe.Pointer(&ptr))}
	})
}
