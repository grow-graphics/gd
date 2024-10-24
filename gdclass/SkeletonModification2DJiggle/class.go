package SkeletonModification2DJiggle

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SkeletonModification2D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This modification moves a series of bones, typically called a bone chain, towards a target. What makes this modification special is that it calculates the velocity and acceleration for each bone in the bone chain, and runs a very light physics-like calculation using the inputted values. This allows the bones to overshoot the target and "jiggle" around. It can be configured to act more like a spring, or sway around like cloth might.
This modification is useful for adding additional motion to things like hair, the edges of clothing, and more. It has several settings to that allow control over how the joint moves when the target moves.
[b]Note:[/b] The Jiggle modifier has [code]jiggle_joints[/code], which are the data objects that hold the data for each joint in the Jiggle chain. This is different from than [Bone2D] nodes! Jiggle joints hold the data needed for each [Bone2D] in the bone chain used by the Jiggle modification.

*/
type Go [1]classdb.SkeletonModification2DJiggle

/*
If [code]true[/code], the Jiggle modifier will take colliders into account, keeping them from entering into these collision objects.
*/
func (self Go) SetUseColliders(use_colliders bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseColliders(use_colliders)
}

/*
Returns whether the jiggle modifier is taking physics colliders into account when solving.
*/
func (self Go) GetUseColliders() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetUseColliders())
}

/*
Sets the collision mask that the Jiggle modifier will use when reacting to colliders, if the Jiggle modifier is set to take colliders into account.
*/
func (self Go) SetCollisionMask(collision_mask int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMask(gd.Int(collision_mask))
}

/*
Returns the collision mask used by the Jiggle modifier when collisions are enabled.
*/
func (self Go) GetCollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCollisionMask()))
}

/*
Sets the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Go) SetJiggleJointBone2dNode(joint_idx int, bone2d_node string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleJointBone2dNode(gd.Int(joint_idx), gc.String(bone2d_node).NodePath(gc))
}

/*
Returns the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Go) GetJiggleJointBone2dNode(joint_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetJiggleJointBone2dNode(gc, gd.Int(joint_idx)).String())
}

/*
Sets the bone index, [param bone_idx], of the Jiggle joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the Jiggle joint based on data provided by the linked skeleton.
*/
func (self Go) SetJiggleJointBoneIndex(joint_idx int, bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Go) GetJiggleJointBoneIndex(joint_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetJiggleJointBoneIndex(gd.Int(joint_idx))))
}

/*
Sets whether the Jiggle joint at [param joint_idx] should override the default Jiggle joint settings. Setting this to [code]true[/code] will make the joint use its own settings rather than the default ones attached to the modification.
*/
func (self Go) SetJiggleJointOverride(joint_idx int, override bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleJointOverride(gd.Int(joint_idx), override)
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is overriding the default Jiggle joint data defined in the modification.
*/
func (self Go) GetJiggleJointOverride(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetJiggleJointOverride(gd.Int(joint_idx)))
}

/*
Sets the of stiffness of the Jiggle joint at [param joint_idx].
*/
func (self Go) SetJiggleJointStiffness(joint_idx int, stiffness float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleJointStiffness(gd.Int(joint_idx), gd.Float(stiffness))
}

/*
Returns the stiffness of the Jiggle joint at [param joint_idx].
*/
func (self Go) GetJiggleJointStiffness(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetJiggleJointStiffness(gd.Int(joint_idx))))
}

/*
Sets the of mass of the Jiggle joint at [param joint_idx].
*/
func (self Go) SetJiggleJointMass(joint_idx int, mass float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleJointMass(gd.Int(joint_idx), gd.Float(mass))
}

/*
Returns the amount of mass of the jiggle joint at [param joint_idx].
*/
func (self Go) GetJiggleJointMass(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetJiggleJointMass(gd.Int(joint_idx))))
}

/*
Sets the amount of damping of the Jiggle joint at [param joint_idx].
*/
func (self Go) SetJiggleJointDamping(joint_idx int, damping float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleJointDamping(gd.Int(joint_idx), gd.Float(damping))
}

/*
Returns the amount of damping of the Jiggle joint at [param joint_idx].
*/
func (self Go) GetJiggleJointDamping(joint_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetJiggleJointDamping(gd.Int(joint_idx))))
}

/*
Sets whether the Jiggle joint at [param joint_idx] should use gravity.
*/
func (self Go) SetJiggleJointUseGravity(joint_idx int, use_gravity bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleJointUseGravity(gd.Int(joint_idx), use_gravity)
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is using gravity or not.
*/
func (self Go) GetJiggleJointUseGravity(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetJiggleJointUseGravity(gd.Int(joint_idx)))
}

/*
Sets the gravity vector of the Jiggle joint at [param joint_idx].
*/
func (self Go) SetJiggleJointGravity(joint_idx int, gravity gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleJointGravity(gd.Int(joint_idx), gravity)
}

/*
Returns a [Vector2] representing the amount of gravity the Jiggle joint at [param joint_idx] is influenced by.
*/
func (self Go) GetJiggleJointGravity(joint_idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetJiggleJointGravity(gd.Int(joint_idx)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModification2DJiggle
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonModification2DJiggle"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TargetNodepath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTargetNode(gc).String())
}

func (self Go) SetTargetNodepath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetNode(gc.String(value).NodePath(gc))
}

func (self Go) JiggleDataChainLength() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetJiggleDataChainLength()))
}

func (self Go) SetJiggleDataChainLength(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJiggleDataChainLength(gd.Int(value))
}

func (self Go) Stiffness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetStiffness()))
}

func (self Go) SetStiffness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStiffness(gd.Float(value))
}

func (self Go) Mass() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMass()))
}

func (self Go) SetMass(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMass(gd.Float(value))
}

func (self Go) Damping() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDamping()))
}

func (self Go) SetDamping(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDamping(gd.Float(value))
}

func (self Go) UseGravity() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUseGravity())
}

func (self Go) SetUseGravity(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseGravity(value)
}

func (self Go) Gravity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetGravity())
}

func (self Go) SetGravity(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravity(value)
}

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(target_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJiggleDataChainLength(length gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJiggleDataChainLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStiffness(stiffness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stiffness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStiffness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMass(mass gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMass() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDamping(damping gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, damping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDamping() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseGravity(use_gravity bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseGravity() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravity(gravity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the Jiggle modifier will take colliders into account, keeping them from entering into these collision objects.
*/
//go:nosplit
func (self class) SetUseColliders(use_colliders bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_colliders)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_use_colliders, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the jiggle modifier is taking physics colliders into account when solving.
*/
//go:nosplit
func (self class) GetUseColliders() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_use_colliders, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the collision mask that the Jiggle modifier will use when reacting to colliders, if the Jiggle modifier is set to take colliders into account.
*/
//go:nosplit
func (self class) SetCollisionMask(collision_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the collision mask used by the Jiggle modifier when collisions are enabled.
*/
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointBone2dNode(joint_idx gd.Int, bone2d_node gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mmm.Get(bone2d_node))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointBone2dNode(ctx gd.Lifetime, joint_idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone index, [param bone_idx], of the Jiggle joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the Jiggle joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetJiggleJointBoneIndex(joint_idx gd.Int, bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointBoneIndex(joint_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the Jiggle joint at [param joint_idx] should override the default Jiggle joint settings. Setting this to [code]true[/code] will make the joint use its own settings rather than the default ones attached to the modification.
*/
//go:nosplit
func (self class) SetJiggleJointOverride(joint_idx gd.Int, override bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, override)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a boolean that indicates whether the joint at [param joint_idx] is overriding the default Jiggle joint data defined in the modification.
*/
//go:nosplit
func (self class) GetJiggleJointOverride(joint_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the of stiffness of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointStiffness(joint_idx gd.Int, stiffness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, stiffness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the stiffness of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointStiffness(joint_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the of mass of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointMass(joint_idx gd.Int, mass gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the amount of mass of the jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointMass(joint_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the amount of damping of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointDamping(joint_idx gd.Int, damping gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, damping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the amount of damping of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointDamping(joint_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the Jiggle joint at [param joint_idx] should use gravity.
*/
//go:nosplit
func (self class) SetJiggleJointUseGravity(joint_idx gd.Int, use_gravity bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, use_gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a boolean that indicates whether the joint at [param joint_idx] is using gravity or not.
*/
//go:nosplit
func (self class) GetJiggleJointUseGravity(joint_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the gravity vector of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointGravity(joint_idx gd.Int, gravity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Vector2] representing the amount of gravity the Jiggle joint at [param joint_idx] is influenced by.
*/
//go:nosplit
func (self class) GetJiggleJointGravity(joint_idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DJiggle() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2DJiggle() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsSkeletonModification2D() SkeletonModification2D.GD { return *((*SkeletonModification2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2D() SkeletonModification2D.Go { return *((*SkeletonModification2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}
func init() {classdb.Register("SkeletonModification2DJiggle", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}