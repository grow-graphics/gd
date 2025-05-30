// Code generated by the generate package DO NOT EDIT

// Package RDPipelineMultisampleState provides methods for working with RDPipelineMultisampleState object instances.
package RDPipelineMultisampleState

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
import "graphics.gd/classdb/Rendering"
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
[RDPipelineMultisampleState] is used to control how multisample or supersample antialiasing is being performed when rendering using [RenderingDevice].
*/
type Instance [1]gdclass.RDPipelineMultisampleState

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDPipelineMultisampleState() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDPipelineMultisampleState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineMultisampleState"))
	casted := Instance{*(*gdclass.RDPipelineMultisampleState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SampleCount() Rendering.TextureSamples {
	return Rendering.TextureSamples(class(self).GetSampleCount())
}

func (self Instance) SetSampleCount(value Rendering.TextureSamples) {
	class(self).SetSampleCount(value)
}

func (self Instance) EnableSampleShading() bool {
	return bool(class(self).GetEnableSampleShading())
}

func (self Instance) SetEnableSampleShading(value bool) {
	class(self).SetEnableSampleShading(value)
}

func (self Instance) MinSampleShading() Float.X {
	return Float.X(Float.X(class(self).GetMinSampleShading()))
}

func (self Instance) SetMinSampleShading(value Float.X) {
	class(self).SetMinSampleShading(float64(value))
}

func (self Instance) EnableAlphaToCoverage() bool {
	return bool(class(self).GetEnableAlphaToCoverage())
}

func (self Instance) SetEnableAlphaToCoverage(value bool) {
	class(self).SetEnableAlphaToCoverage(value)
}

func (self Instance) EnableAlphaToOne() bool {
	return bool(class(self).GetEnableAlphaToOne())
}

func (self Instance) SetEnableAlphaToOne(value bool) {
	class(self).SetEnableAlphaToOne(value)
}

func (self Instance) SampleMasks() []int {
	return []int(gd.ArrayAs[[]int](gd.InternalArray(class(self).GetSampleMasks())))
}

func (self Instance) SetSampleMasks(value []int) {
	class(self).SetSampleMasks(gd.ArrayFromSlice[Array.Contains[int64]](value))
}

//go:nosplit
func (self class) SetSampleCount(p_member Rendering.TextureSamples) { //gd:RDPipelineMultisampleState.set_sample_count
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_sample_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSampleCount() Rendering.TextureSamples { //gd:RDPipelineMultisampleState.get_sample_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[Rendering.TextureSamples](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_sample_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableSampleShading(p_member bool) { //gd:RDPipelineMultisampleState.set_enable_sample_shading
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_enable_sample_shading, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableSampleShading() bool { //gd:RDPipelineMultisampleState.get_enable_sample_shading
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_enable_sample_shading, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinSampleShading(p_member float64) { //gd:RDPipelineMultisampleState.set_min_sample_shading
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_min_sample_shading, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinSampleShading() float64 { //gd:RDPipelineMultisampleState.get_min_sample_shading
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_min_sample_shading, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableAlphaToCoverage(p_member bool) { //gd:RDPipelineMultisampleState.set_enable_alpha_to_coverage
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_enable_alpha_to_coverage, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableAlphaToCoverage() bool { //gd:RDPipelineMultisampleState.get_enable_alpha_to_coverage
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_enable_alpha_to_coverage, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableAlphaToOne(p_member bool) { //gd:RDPipelineMultisampleState.set_enable_alpha_to_one
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_enable_alpha_to_one, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableAlphaToOne() bool { //gd:RDPipelineMultisampleState.get_enable_alpha_to_one
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_enable_alpha_to_one, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSampleMasks(masks Array.Contains[int64]) { //gd:RDPipelineMultisampleState.set_sample_masks
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(masks)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_sample_masks, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSampleMasks() Array.Contains[int64] { //gd:RDPipelineMultisampleState.get_sample_masks
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_sample_masks, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[int64]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsRDPipelineMultisampleState() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineMultisampleState() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsRDPipelineMultisampleState() Instance {
	return self.Super().AsRDPipelineMultisampleState()
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
	gdclass.Register("RDPipelineMultisampleState", func(ptr gd.Object) any {
		return [1]gdclass.RDPipelineMultisampleState{*(*gdclass.RDPipelineMultisampleState)(unsafe.Pointer(&ptr))}
	})
}
