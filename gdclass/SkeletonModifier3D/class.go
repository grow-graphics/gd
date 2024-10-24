package SkeletonModifier3D

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
[SkeletonModifier3D] retrieves a target [Skeleton3D] by having a [Skeleton3D] parent.
If there is [AnimationMixer], modification always performs after playback process of the [AnimationMixer].
This node should be used to implement custom IK solvers, constraints, or skeleton physics.
	// SkeletonModifier3D methods that can be overridden by a [Class] that extends it.
	type SkeletonModifier3D interface {
		//Override this virtual method to implement a custom skeleton modifier. You should do things like get the [Skeleton3D]'s current pose and apply the pose here.
		//[method _process_modification] must not apply [member influence] to bone poses because the [Skeleton3D] automatically applies influence to all bone poses set by the modifier.
		ProcessModification() 
	}

*/
type Go [1]classdb.SkeletonModifier3D

/*
Override this virtual method to implement a custom skeleton modifier. You should do things like get the [Skeleton3D]'s current pose and apply the pose here.
[method _process_modification] must not apply [member influence] to bone poses because the [Skeleton3D] automatically applies influence to all bone poses set by the modifier.
*/
func (Go) _process_modification(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Get parent [Skeleton3D] node if found.
*/
func (self Go) GetSkeleton() gdclass.Skeleton3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Skeleton3D(class(self).GetSkeleton(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModifier3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonModifier3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Active() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsActive())
}

func (self Go) SetActive(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetActive(value)
}

func (self Go) Influence() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetInfluence()))
}

func (self Go) SetInfluence(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInfluence(gd.Float(value))
}

/*
Override this virtual method to implement a custom skeleton modifier. You should do things like get the [Skeleton3D]'s current pose and apply the pose here.
[method _process_modification] must not apply [member influence] to bone poses because the [Skeleton3D] automatically applies influence to all bone poses set by the modifier.
*/
func (class) _process_modification(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Get parent [Skeleton3D] node if found.
*/
//go:nosplit
func (self class) GetSkeleton(ctx gd.Lifetime) gdclass.Skeleton3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModifier3D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Skeleton3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModifier3D.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModifier3D.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInfluence(influence gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, influence)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModifier3D.Bind_set_influence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInfluence() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModifier3D.Bind_get_influence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnModificationProcessed(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("modification_processed"), gc.Callable(cb), 0)
}


func (self class) AsSkeletonModifier3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModifier3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process_modification": return reflect.ValueOf(self._process_modification);
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_process_modification": return reflect.ValueOf(self._process_modification);
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("SkeletonModifier3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}