// Code generated by the generate package DO NOT EDIT

// Package PathFollow3D provides methods for working with PathFollow3D object instances.
package PathFollow3D

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
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node3D"
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
import "graphics.gd/variant/Transform3D"

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
This node takes its parent [Path3D], and returns the coordinates of a point within it, given a distance from the first vertex.
It is useful for making other nodes follow a path, without coding the movement pattern. For that, the nodes must be children of this node. The descendant nodes will then move accordingly when setting the [member progress] in this node.
*/
type Instance [1]gdclass.PathFollow3D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPathFollow3D() Instance
}

/*
Correct the [param transform]. [param rotation_mode] implicitly specifies how posture (forward, up and sideway direction) is calculated.
*/
func CorrectPosture(transform Transform3D.BasisOrigin, rotation_mode RotationMode) Transform3D.BasisOrigin { //gd:PathFollow3D.correct_posture
	self := Instance{}
	return Transform3D.BasisOrigin(Advanced(self).CorrectPosture(Transform3D.BasisOrigin(transform), rotation_mode))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PathFollow3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PathFollow3D"))
	casted := Instance{*(*gdclass.PathFollow3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Progress() Float.X {
	return Float.X(Float.X(class(self).GetProgress()))
}

func (self Instance) SetProgress(value Float.X) {
	class(self).SetProgress(float64(value))
}

func (self Instance) ProgressRatio() Float.X {
	return Float.X(Float.X(class(self).GetProgressRatio()))
}

func (self Instance) SetProgressRatio(value Float.X) {
	class(self).SetProgressRatio(float64(value))
}

func (self Instance) HOffset() Float.X {
	return Float.X(Float.X(class(self).GetHOffset()))
}

func (self Instance) SetHOffset(value Float.X) {
	class(self).SetHOffset(float64(value))
}

func (self Instance) VOffset() Float.X {
	return Float.X(Float.X(class(self).GetVOffset()))
}

func (self Instance) SetVOffset(value Float.X) {
	class(self).SetVOffset(float64(value))
}

func (self Instance) RotationMode() RotationMode {
	return RotationMode(class(self).GetRotationMode())
}

func (self Instance) SetRotationMode(value RotationMode) {
	class(self).SetRotationMode(value)
}

func (self Instance) UseModelFront() bool {
	return bool(class(self).IsUsingModelFront())
}

func (self Instance) SetUseModelFront(value bool) {
	class(self).SetUseModelFront(value)
}

func (self Instance) CubicInterp() bool {
	return bool(class(self).GetCubicInterpolation())
}

func (self Instance) SetCubicInterp(value bool) {
	class(self).SetCubicInterpolation(value)
}

func (self Instance) Loop() bool {
	return bool(class(self).HasLoop())
}

func (self Instance) SetLoop(value bool) {
	class(self).SetLoop(value)
}

func (self Instance) TiltEnabled() bool {
	return bool(class(self).IsTiltEnabled())
}

func (self Instance) SetTiltEnabled(value bool) {
	class(self).SetTiltEnabled(value)
}

//go:nosplit
func (self class) SetProgress(progress float64) { //gd:PathFollow3D.set_progress
	var frame = callframe.New()
	callframe.Arg(frame, progress)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProgress() float64 { //gd:PathFollow3D.get_progress
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHOffset(h_offset float64) { //gd:PathFollow3D.set_h_offset
	var frame = callframe.New()
	callframe.Arg(frame, h_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_h_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHOffset() float64 { //gd:PathFollow3D.get_h_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_h_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVOffset(v_offset float64) { //gd:PathFollow3D.set_v_offset
	var frame = callframe.New()
	callframe.Arg(frame, v_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_v_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVOffset() float64 { //gd:PathFollow3D.get_v_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_v_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProgressRatio(ratio float64) { //gd:PathFollow3D.set_progress_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_progress_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProgressRatio() float64 { //gd:PathFollow3D.get_progress_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_progress_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotationMode(rotation_mode RotationMode) { //gd:PathFollow3D.set_rotation_mode
	var frame = callframe.New()
	callframe.Arg(frame, rotation_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_rotation_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotationMode() RotationMode { //gd:PathFollow3D.get_rotation_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[RotationMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_rotation_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCubicInterpolation(enabled bool) { //gd:PathFollow3D.set_cubic_interpolation
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_cubic_interpolation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCubicInterpolation() bool { //gd:PathFollow3D.get_cubic_interpolation
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_get_cubic_interpolation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseModelFront(enabled bool) { //gd:PathFollow3D.set_use_model_front
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_use_model_front, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingModelFront() bool { //gd:PathFollow3D.is_using_model_front
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_is_using_model_front, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoop(loop bool) { //gd:PathFollow3D.set_loop
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) HasLoop() bool { //gd:PathFollow3D.has_loop
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTiltEnabled(enabled bool) { //gd:PathFollow3D.set_tilt_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_set_tilt_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsTiltEnabled() bool { //gd:PathFollow3D.is_tilt_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PathFollow3D.Bind_is_tilt_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Correct the [param transform]. [param rotation_mode] implicitly specifies how posture (forward, up and sideway direction) is calculated.
*/
//go:nosplit
func (self class) CorrectPosture(transform Transform3D.BasisOrigin, rotation_mode RotationMode) Transform3D.BasisOrigin { //gd:PathFollow3D.correct_posture
	var frame = callframe.New()
	callframe.Arg(frame, gd.Transposed(transform))
	callframe.Arg(frame, rotation_mode)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCallStatic(gd.Global.Methods.PathFollow3D.Bind_correct_posture, frame.Array(0), r_ret.Addr())
	var ret = gd.Transposed(r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsPathFollow3D() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPathFollow3D() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsPathFollow3D() Instance  { return self.Super().AsPathFollow3D() }
func (self class) AsNode3D() Node3D.Advanced         { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode3D() Node3D.Instance { return self.Super().AsNode3D() }
func (self Instance) AsNode3D() Node3D.Instance      { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced             { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance     { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance          { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("PathFollow3D", func(ptr gd.Object) any {
		return [1]gdclass.PathFollow3D{*(*gdclass.PathFollow3D)(unsafe.Pointer(&ptr))}
	})
}

type RotationMode int //gd:PathFollow3D.RotationMode

const (
	/*Forbids the PathFollow3D to rotate.*/
	RotationNone RotationMode = 0
	/*Allows the PathFollow3D to rotate in the Y axis only.*/
	RotationY RotationMode = 1
	/*Allows the PathFollow3D to rotate in both the X, and Y axes.*/
	RotationXy RotationMode = 2
	/*Allows the PathFollow3D to rotate in any axis.*/
	RotationXyz RotationMode = 3
	/*Uses the up vector information in a [Curve3D] to enforce orientation. This rotation mode requires the [Path3D]'s [member Curve3D.up_vector_enabled] property to be set to [code]true[/code].*/
	RotationOriented RotationMode = 4
)
