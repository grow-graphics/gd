package KinematicCollision2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Holds collision data from the movement of a [PhysicsBody2D], usually from [method PhysicsBody2D.move_and_collide]. When a [PhysicsBody2D] is moved, it stops if it detects a collision with another body. If a collision is detected, a [KinematicCollision2D] object is returned.
The collision data includes the colliding object, the remaining motion, and the collision position. This data can be used to determine a custom response to the collision.

*/
type Go [1]classdb.KinematicCollision2D

/*
Returns the point of collision in global coordinates.
*/
func (self Go) GetPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetPosition())
}

/*
Returns the colliding body's shape's normal at the point of collision.
*/
func (self Go) GetNormal() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetNormal())
}

/*
Returns the moving object's travel before collision.
*/
func (self Go) GetTravel() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetTravel())
}

/*
Returns the moving object's remaining movement vector.
*/
func (self Go) GetRemainder() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetRemainder())
}

/*
Returns the collision angle according to [param up_direction], which is [constant Vector2.UP] by default. This value is always positive.
*/
func (self Go) GetAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetAngle(gd.Vector2{0, -1})))
}

/*
Returns the colliding body's length of overlap along the collision normal.
*/
func (self Go) GetDepth() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetDepth()))
}

/*
Returns the moving object's colliding shape.
*/
func (self Go) GetLocalShape() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetLocalShape(gc))
}

/*
Returns the colliding body's attached [Object].
*/
func (self Go) GetCollider() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetCollider(gc))
}

/*
Returns the unique instance ID of the colliding body's attached [Object]. See [method Object.get_instance_id].
*/
func (self Go) GetColliderId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetColliderId()))
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer2D].
*/
func (self Go) GetColliderRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetColliderRid())
}

/*
Returns the colliding body's shape.
*/
func (self Go) GetColliderShape() gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Object(class(self).GetColliderShape(gc))
}

/*
Returns the colliding body's shape index. See [CollisionObject2D].
*/
func (self Go) GetColliderShapeIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetColliderShapeIndex()))
}

/*
Returns the colliding body's velocity.
*/
func (self Go) GetColliderVelocity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetColliderVelocity())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.KinematicCollision2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("KinematicCollision2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the point of collision in global coordinates.
*/
//go:nosplit
func (self class) GetPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's shape's normal at the point of collision.
*/
//go:nosplit
func (self class) GetNormal() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the moving object's travel before collision.
*/
//go:nosplit
func (self class) GetTravel() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the moving object's remaining movement vector.
*/
//go:nosplit
func (self class) GetRemainder() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_remainder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision angle according to [param up_direction], which is [constant Vector2.UP] by default. This value is always positive.
*/
//go:nosplit
func (self class) GetAngle(up_direction gd.Vector2) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's length of overlap along the collision normal.
*/
//go:nosplit
func (self class) GetDepth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the moving object's colliding shape.
*/
//go:nosplit
func (self class) GetLocalShape(ctx gd.Lifetime) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_local_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the colliding body's attached [Object].
*/
//go:nosplit
func (self class) GetCollider(ctx gd.Lifetime) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the unique instance ID of the colliding body's attached [Object]. See [method Object.get_instance_id].
*/
//go:nosplit
func (self class) GetColliderId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_collider_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's [RID] used by the [PhysicsServer2D].
*/
//go:nosplit
func (self class) GetColliderRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's shape.
*/
//go:nosplit
func (self class) GetColliderShape(ctx gd.Lifetime) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the colliding body's shape index. See [CollisionObject2D].
*/
//go:nosplit
func (self class) GetColliderShapeIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_collider_shape_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the colliding body's velocity.
*/
//go:nosplit
func (self class) GetColliderVelocity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.KinematicCollision2D.Bind_get_collider_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsKinematicCollision2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsKinematicCollision2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("KinematicCollision2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}