// Package RDVertexAttribute provides methods for working with RDVertexAttribute object instances.
package RDVertexAttribute

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
var _ = slices.Delete[[]struct{}, struct{}]

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDVertexAttribute

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDVertexAttribute() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDVertexAttribute

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDVertexAttribute"))
	casted := Instance{*(*gdclass.RDVertexAttribute)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Location() int {
	return int(int(class(self).GetLocation()))
}

func (self Instance) SetLocation(value int) {
	class(self).SetLocation(int64(value))
}

func (self Instance) Offset() int {
	return int(int(class(self).GetOffset()))
}

func (self Instance) SetOffset(value int) {
	class(self).SetOffset(int64(value))
}

func (self Instance) Format() gdclass.RenderingDeviceDataFormat {
	return gdclass.RenderingDeviceDataFormat(class(self).GetFormat())
}

func (self Instance) SetFormat(value gdclass.RenderingDeviceDataFormat) {
	class(self).SetFormat(value)
}

func (self Instance) Stride() int {
	return int(int(class(self).GetStride()))
}

func (self Instance) SetStride(value int) {
	class(self).SetStride(int64(value))
}

func (self Instance) Frequency() gdclass.RenderingDeviceVertexFrequency {
	return gdclass.RenderingDeviceVertexFrequency(class(self).GetFrequency())
}

func (self Instance) SetFrequency(value gdclass.RenderingDeviceVertexFrequency) {
	class(self).SetFrequency(value)
}

//go:nosplit
func (self class) SetLocation(p_member int64) { //gd:RDVertexAttribute.set_location
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_set_location, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLocation() int64 { //gd:RDVertexAttribute.get_location
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_get_location, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(p_member int64) { //gd:RDVertexAttribute.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() int64 { //gd:RDVertexAttribute.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFormat(p_member gdclass.RenderingDeviceDataFormat) { //gd:RDVertexAttribute.set_format
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormat() gdclass.RenderingDeviceDataFormat { //gd:RDVertexAttribute.get_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceDataFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStride(p_member int64) { //gd:RDVertexAttribute.set_stride
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_set_stride, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStride() int64 { //gd:RDVertexAttribute.get_stride
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_get_stride, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrequency(p_member gdclass.RenderingDeviceVertexFrequency) { //gd:RDVertexAttribute.set_frequency
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_set_frequency, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrequency() gdclass.RenderingDeviceVertexFrequency { //gd:RDVertexAttribute.get_frequency
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceVertexFrequency](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDVertexAttribute.Bind_get_frequency, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDVertexAttribute() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDVertexAttribute() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RDVertexAttribute", func(ptr gd.Object) any {
		return [1]gdclass.RDVertexAttribute{*(*gdclass.RDVertexAttribute)(unsafe.Pointer(&ptr))}
	})
}
