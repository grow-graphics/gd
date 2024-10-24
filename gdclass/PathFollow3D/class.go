package PathFollow3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This node takes its parent [Path3D], and returns the coordinates of a point within it, given a distance from the first vertex.
It is useful for making other nodes follow a path, without coding the movement pattern. For that, the nodes must be children of this node. The descendant nodes will then move accordingly when setting the [member progress] in this node.

*/
type Go [1]classdb.PathFollow3D

/*
Correct the [param transform]. [param rotation_mode] implicitly specifies how posture (forward, up and sideway direction) is calculated.
*/
func (self Go) CorrectPosture(transform gd.Transform3D, rotation_mode classdb.PathFollow3DRotationMode) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(class(self).CorrectPosture(gc, transform, rotation_mode))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PathFollow3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PathFollow3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Progress() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetProgress()))
}

func (self Go) SetProgress(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProgress(gd.Float(value))
}

func (self Go) ProgressRatio() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetProgressRatio()))
}

func (self Go) SetProgressRatio(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProgressRatio(gd.Float(value))
}

func (self Go) HOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetHOffset()))
}

func (self Go) SetHOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHOffset(gd.Float(value))
}

func (self Go) VOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVOffset()))
}

func (self Go) SetVOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVOffset(gd.Float(value))
}

func (self Go) RotationMode() classdb.PathFollow3DRotationMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.PathFollow3DRotationMode(class(self).GetRotationMode())
}

func (self Go) SetRotationMode(value classdb.PathFollow3DRotationMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRotationMode(value)
}

func (self Go) UseModelFront() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingModelFront())
}

func (self Go) SetUseModelFront(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseModelFront(value)
}

func (self Go) CubicInterp() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetCubicInterpolation())
}

func (self Go) SetCubicInterp(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCubicInterpolation(value)
}

func (self Go) Loop() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasLoop())
}

func (self Go) SetLoop(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoop(value)
}

func (self Go) TiltEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsTiltEnabled())
}

func (self Go) SetTiltEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTiltEnabled(value)
}

//go:nosplit
func (self class) SetProgress(progress gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, progress)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProgress() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_get_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHOffset(h_offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, h_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_h_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_get_h_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVOffset(v_offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, v_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_v_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_get_v_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProgressRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_progress_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProgressRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_get_progress_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRotationMode(rotation_mode classdb.PathFollow3DRotationMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rotation_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_rotation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRotationMode() classdb.PathFollow3DRotationMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PathFollow3DRotationMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_get_rotation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCubicInterpolation(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_cubic_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCubicInterpolation() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_get_cubic_interpolation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseModelFront(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_use_model_front, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingModelFront() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_is_using_model_front, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoop(loop bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasLoop() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTiltEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_set_tilt_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTiltEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PathFollow3D.Bind_is_tilt_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Correct the [param transform]. [param rotation_mode] implicitly specifies how posture (forward, up and sideway direction) is calculated.
*/
//go:nosplit
func (self class) CorrectPosture(ctx gd.Lifetime, transform gd.Transform3D, rotation_mode classdb.PathFollow3DRotationMode) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	callframe.Arg(frame, rotation_mode)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.PathFollow3D.Bind_correct_posture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPathFollow3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPathFollow3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("PathFollow3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type RotationMode = classdb.PathFollow3DRotationMode

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