package PhysicsBody3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/CollisionObject3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[PhysicsBody3D] is an abstract base class for 3D game objects affected by physics. All 3D physics bodies inherit from it.
[b]Warning:[/b] With a non-uniform scale, this node will likely not behave as expected. It is advised to keep its scale the same on all axes and adjust its collision shape(s) instead.

*/
type Go [1]classdb.PhysicsBody3D

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
The body will stop if it collides. Returns a [KinematicCollision3D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody3D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody3D] for improving floor detection during floor snapping.
[param max_collisions] allows to retrieve more than one collision result.
*/
func (self Go) MoveAndCollide(motion gd.Vector3) gdclass.KinematicCollision3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.KinematicCollision3D(class(self).MoveAndCollide(gc, motion, false, gd.Float(0.001), false, gd.Int(1)))
}

/*
Checks for collisions without moving the body. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Virtually sets the node's position, scale and rotation to that of the given [Transform3D], then tries to move the body along the vector [param motion]. Returns [code]true[/code] if a collision would stop the body from moving along the whole path.
[param collision] is an optional object of type [KinematicCollision3D], which contains additional information about the collision when stopped, or when touching another body along the motion.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody3D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is useful for checking whether the body would [i]touch[/i] any other bodies.
[param max_collisions] allows to retrieve more than one collision result.
*/
func (self Go) TestMove(from gd.Transform3D, motion gd.Vector3) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).TestMove(from, motion, ([1]gdclass.KinematicCollision3D{}[0]), gd.Float(0.001), false, gd.Int(1)))
}

/*
Returns the gravity vector computed from all sources that can affect the body, including all gravity overrides from [Area3D] nodes and the global world gravity.
*/
func (self Go) GetGravity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetGravity())
}

/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
func (self Go) GetCollisionExceptions() gd.ArrayOf[gdclass.PhysicsBody3D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.PhysicsBody3D](class(self).GetCollisionExceptions(gc))
}

/*
Adds a body to the list of bodies that this body can't collide with.
*/
func (self Go) AddCollisionExceptionWith(body gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddCollisionExceptionWith(body)
}

/*
Removes a body from the list of bodies that this body can't collide with.
*/
func (self Go) RemoveCollisionExceptionWith(body gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveCollisionExceptionWith(body)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsBody3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicsBody3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) AxisLockLinearX() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAxisLock(1))
}

func (self Go) SetAxisLockLinearX(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAxisLock(1, value)
}

func (self Go) AxisLockLinearY() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAxisLock(2))
}

func (self Go) SetAxisLockLinearY(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAxisLock(2, value)
}

func (self Go) AxisLockLinearZ() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAxisLock(4))
}

func (self Go) SetAxisLockLinearZ(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAxisLock(4, value)
}

func (self Go) AxisLockAngularX() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAxisLock(8))
}

func (self Go) SetAxisLockAngularX(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAxisLock(8, value)
}

func (self Go) AxisLockAngularY() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAxisLock(16))
}

func (self Go) SetAxisLockAngularY(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAxisLock(16, value)
}

func (self Go) AxisLockAngularZ() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAxisLock(32))
}

func (self Go) SetAxisLockAngularZ(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAxisLock(32, value)
}

/*
Moves the body along the vector [param motion]. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
The body will stop if it collides. Returns a [KinematicCollision3D], which contains information about the collision when stopped, or when touching another body along the motion.
If [param test_only] is [code]true[/code], the body does not move but the would-be collision information is given.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody3D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is used e.g. by [CharacterBody3D] for improving floor detection during floor snapping.
[param max_collisions] allows to retrieve more than one collision result.
*/
//go:nosplit
func (self class) MoveAndCollide(ctx gd.Lifetime, motion gd.Vector3, test_only bool, safe_margin gd.Float, recovery_as_collision bool, max_collisions gd.Int) gdclass.KinematicCollision3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	callframe.Arg(frame, test_only)
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	callframe.Arg(frame, max_collisions)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody3D.Bind_move_and_collide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.KinematicCollision3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Checks for collisions without moving the body. In order to be frame rate independent in [method Node._physics_process] or [method Node._process], [param motion] should be computed using [code]delta[/code].
Virtually sets the node's position, scale and rotation to that of the given [Transform3D], then tries to move the body along the vector [param motion]. Returns [code]true[/code] if a collision would stop the body from moving along the whole path.
[param collision] is an optional object of type [KinematicCollision3D], which contains additional information about the collision when stopped, or when touching another body along the motion.
[param safe_margin] is the extra margin used for collision recovery (see [member CharacterBody3D.safe_margin] for more details).
If [param recovery_as_collision] is [code]true[/code], any depenetration from the recovery phase is also reported as a collision; this is useful for checking whether the body would [i]touch[/i] any other bodies.
[param max_collisions] allows to retrieve more than one collision result.
*/
//go:nosplit
func (self class) TestMove(from gd.Transform3D, motion gd.Vector3, collision gdclass.KinematicCollision3D, safe_margin gd.Float, recovery_as_collision bool, max_collisions gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, motion)
	callframe.Arg(frame, mmm.Get(collision[0].AsPointer())[0])
	callframe.Arg(frame, safe_margin)
	callframe.Arg(frame, recovery_as_collision)
	callframe.Arg(frame, max_collisions)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody3D.Bind_test_move, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the gravity vector computed from all sources that can affect the body, including all gravity overrides from [Area3D] nodes and the global world gravity.
*/
//go:nosplit
func (self class) GetGravity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody3D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Locks or unlocks the specified linear or rotational [param axis] depending on the value of [param lock].
*/
//go:nosplit
func (self class) SetAxisLock(axis classdb.PhysicsServer3DBodyAxis, lock bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	callframe.Arg(frame, lock)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody3D.Bind_set_axis_lock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified linear or rotational [param axis] is locked.
*/
//go:nosplit
func (self class) GetAxisLock(axis classdb.PhysicsServer3DBodyAxis) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody3D.Bind_get_axis_lock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
//go:nosplit
func (self class) GetCollisionExceptions(ctx gd.Lifetime) gd.ArrayOf[gdclass.PhysicsBody3D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody3D.Bind_get_collision_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.PhysicsBody3D](ret)
}
/*
Adds a body to the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) AddCollisionExceptionWith(body gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody3D.Bind_add_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a body from the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) RemoveCollisionExceptionWith(body gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsBody3D.Bind_remove_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPhysicsBody3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsBody3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject3D() CollisionObject3D.GD { return *((*CollisionObject3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionObject3D() CollisionObject3D.Go { return *((*CollisionObject3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCollisionObject3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCollisionObject3D(), name)
	}
}
func init() {classdb.Register("PhysicsBody3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}