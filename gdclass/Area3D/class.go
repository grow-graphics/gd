package Area3D

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
[Area3D] is a region of 3D space defined by one or multiple [CollisionShape3D] or [CollisionPolygon3D] child nodes. It detects when other [CollisionObject3D]s enter or exit it, and it also keeps track of which collision objects haven't exited it yet (i.e. which one are overlapping it).
This node can also locally alter or override physics parameters (gravity, damping) and route audio to custom audio buses.
[b]Note:[/b] Areas and bodies created with [PhysicsServer3D] might not interact as expected with [Area3D]s, and might not emit signals or track objects correctly.
[b]Warning:[/b] Using a [ConcavePolygonShape3D] inside a [CollisionShape3D] child of this node (created e.g. by using the [b]Create Trimesh Collision Sibling[/b] option in the [b]Mesh[/b] menu that appears when selecting a [MeshInstance3D] node) may give unexpected results, since this collision shape is hollow. If this is not desired, it has to be split into multiple [ConvexPolygonShape3D]s or primitive shapes like [BoxShape3D], or in some cases it may be replaceable by a [CollisionPolygon3D].

*/
type Go [1]classdb.Area3D

/*
Returns a list of intersecting [PhysicsBody3D]s and [GridMap]s. The overlapping body's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Go) GetOverlappingBodies() gd.ArrayOf[gdclass.Node3D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Node3D](class(self).GetOverlappingBodies(gc))
}

/*
Returns a list of intersecting [Area3D]s. The overlapping area's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Go) GetOverlappingAreas() gd.ArrayOf[gdclass.Area3D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Area3D](class(self).GetOverlappingAreas(gc))
}

/*
Returns [code]true[/code] if intersecting any [PhysicsBody3D]s or [GridMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Go) HasOverlappingBodies() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasOverlappingBodies())
}

/*
Returns [code]true[/code] if intersecting any [Area3D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Go) HasOverlappingAreas() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasOverlappingAreas())
}

/*
Returns [code]true[/code] if the given physics body intersects or overlaps this [Area3D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
The [param body] argument can either be a [PhysicsBody3D] or a [GridMap] instance. While GridMaps are not physics body themselves, they register their tiles with collision shapes as a virtual physics body.
*/
func (self Go) OverlapsBody(body gdclass.Node) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).OverlapsBody(body))
}

/*
Returns [code]true[/code] if the given [Area3D] intersects or overlaps this [Area3D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
func (self Go) OverlapsArea(area gdclass.Node) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).OverlapsArea(area))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Area3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Area3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Monitoring() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMonitoring())
}

func (self Go) SetMonitoring(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMonitoring(value)
}

func (self Go) Monitorable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMonitorable())
}

func (self Go) SetMonitorable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMonitorable(value)
}

func (self Go) Priority() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPriority()))
}

func (self Go) SetPriority(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPriority(gd.Int(value))
}

func (self Go) GravitySpaceOverride() classdb.Area3DSpaceOverride {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Area3DSpaceOverride(class(self).GetGravitySpaceOverrideMode())
}

func (self Go) SetGravitySpaceOverride(value classdb.Area3DSpaceOverride) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravitySpaceOverrideMode(value)
}

func (self Go) GravityPoint() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsGravityAPoint())
}

func (self Go) SetGravityPoint(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityIsPoint(value)
}

func (self Go) GravityPointUnitDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGravityPointUnitDistance()))
}

func (self Go) SetGravityPointUnitDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityPointUnitDistance(gd.Float(value))
}

func (self Go) GravityPointCenter() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetGravityPointCenter())
}

func (self Go) SetGravityPointCenter(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityPointCenter(value)
}

func (self Go) GravityDirection() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetGravityDirection())
}

func (self Go) SetGravityDirection(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityDirection(value)
}

func (self Go) Gravity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGravity()))
}

func (self Go) SetGravity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravity(gd.Float(value))
}

func (self Go) LinearDampSpaceOverride() classdb.Area3DSpaceOverride {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Area3DSpaceOverride(class(self).GetLinearDampSpaceOverrideMode())
}

func (self Go) SetLinearDampSpaceOverride(value classdb.Area3DSpaceOverride) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearDampSpaceOverrideMode(value)
}

func (self Go) LinearDamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLinearDamp()))
}

func (self Go) SetLinearDamp(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Go) AngularDampSpaceOverride() classdb.Area3DSpaceOverride {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Area3DSpaceOverride(class(self).GetAngularDampSpaceOverrideMode())
}

func (self Go) SetAngularDampSpaceOverride(value classdb.Area3DSpaceOverride) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularDampSpaceOverrideMode(value)
}

func (self Go) AngularDamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAngularDamp()))
}

func (self Go) SetAngularDamp(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Go) WindForceMagnitude() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetWindForceMagnitude()))
}

func (self Go) SetWindForceMagnitude(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWindForceMagnitude(gd.Float(value))
}

func (self Go) WindAttenuationFactor() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetWindAttenuationFactor()))
}

func (self Go) SetWindAttenuationFactor(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWindAttenuationFactor(gd.Float(value))
}

func (self Go) WindSourcePath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetWindSourcePath(gc).String())
}

func (self Go) SetWindSourcePath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWindSourcePath(gc.String(value).NodePath(gc))
}

func (self Go) AudioBusOverride() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsOverridingAudioBus())
}

func (self Go) SetAudioBusOverride(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAudioBusOverride(value)
}

func (self Go) AudioBusName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetAudioBusName(gc).String())
}

func (self Go) SetAudioBusName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAudioBusName(gc.StringName(value))
}

func (self Go) ReverbBusEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingReverbBus())
}

func (self Go) SetReverbBusEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseReverbBus(value)
}

func (self Go) ReverbBusName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetReverbBusName(gc).String())
}

func (self Go) SetReverbBusName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReverbBusName(gc.StringName(value))
}

func (self Go) ReverbBusAmount() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetReverbAmount()))
}

func (self Go) SetReverbBusAmount(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReverbAmount(gd.Float(value))
}

func (self Go) ReverbBusUniformity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetReverbUniformity()))
}

func (self Go) SetReverbBusUniformity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReverbUniformity(gd.Float(value))
}

//go:nosplit
func (self class) SetGravitySpaceOverrideMode(space_override_mode classdb.Area3DSpaceOverride)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravitySpaceOverrideMode() classdb.Area3DSpaceOverride {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area3DSpaceOverride](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityIsPoint(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_gravity_is_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGravityAPoint() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_is_gravity_a_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityPointUnitDistance(distance_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityPointUnitDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityPointCenter(center gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, center)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityPointCenter() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityDirection(direction gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityDirection() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravity(gravity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDampSpaceOverrideMode(space_override_mode classdb.Area3DSpaceOverride)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDampSpaceOverrideMode() classdb.Area3DSpaceOverride {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area3DSpaceOverride](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDampSpaceOverrideMode(space_override_mode classdb.Area3DSpaceOverride)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDampSpaceOverrideMode() classdb.Area3DSpaceOverride {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area3DSpaceOverride](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPriority(priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPriority() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWindForceMagnitude(wind_force_magnitude gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wind_force_magnitude)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_wind_force_magnitude, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWindForceMagnitude() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_wind_force_magnitude, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWindAttenuationFactor(wind_attenuation_factor gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wind_attenuation_factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_wind_attenuation_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWindAttenuationFactor() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_wind_attenuation_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWindSourcePath(wind_source_path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(wind_source_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_wind_source_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWindSourcePath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_wind_source_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMonitorable(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_monitorable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMonitorable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_is_monitorable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMonitoring(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_monitoring, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMonitoring() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_is_monitoring, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a list of intersecting [PhysicsBody3D]s and [GridMap]s. The overlapping body's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingBodies(ctx gd.Lifetime) gd.ArrayOf[gdclass.Node3D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Node3D](ret)
}
/*
Returns a list of intersecting [Area3D]s. The overlapping area's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingAreas(ctx gd.Lifetime) gd.ArrayOf[gdclass.Area3D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Area3D](ret)
}
/*
Returns [code]true[/code] if intersecting any [PhysicsBody3D]s or [GridMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingBodies() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_has_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if intersecting any [Area3D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingAreas() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_has_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given physics body intersects or overlaps this [Area3D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
The [param body] argument can either be a [PhysicsBody3D] or a [GridMap] instance. While GridMaps are not physics body themselves, they register their tiles with collision shapes as a virtual physics body.
*/
//go:nosplit
func (self class) OverlapsBody(body gdclass.Node) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_overlaps_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given [Area3D] intersects or overlaps this [Area3D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) OverlapsArea(area gdclass.Node) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(area[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_overlaps_area, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAudioBusOverride(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_audio_bus_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsOverridingAudioBus() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_is_overriding_audio_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAudioBusName(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAudioBusName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseReverbBus(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_use_reverb_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingReverbBus() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_is_using_reverb_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReverbBusName(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_reverb_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReverbBusName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_reverb_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReverbAmount(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_reverb_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReverbAmount() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_reverb_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReverbUniformity(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_set_reverb_uniformity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReverbUniformity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area3D.Bind_get_reverb_uniformity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnBodyShapeEntered(cb func(body_rid gd.RID, body gdclass.Node3D, body_shape_index int, local_shape_index int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("body_shape_entered"), gc.Callable(cb), 0)
}


func (self Go) OnBodyShapeExited(cb func(body_rid gd.RID, body gdclass.Node3D, body_shape_index int, local_shape_index int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("body_shape_exited"), gc.Callable(cb), 0)
}


func (self Go) OnBodyEntered(cb func(body gdclass.Node3D)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("body_entered"), gc.Callable(cb), 0)
}


func (self Go) OnBodyExited(cb func(body gdclass.Node3D)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("body_exited"), gc.Callable(cb), 0)
}


func (self Go) OnAreaShapeEntered(cb func(area_rid gd.RID, area gdclass.Area3D, area_shape_index int, local_shape_index int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("area_shape_entered"), gc.Callable(cb), 0)
}


func (self Go) OnAreaShapeExited(cb func(area_rid gd.RID, area gdclass.Area3D, area_shape_index int, local_shape_index int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("area_shape_exited"), gc.Callable(cb), 0)
}


func (self Go) OnAreaEntered(cb func(area gdclass.Area3D)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("area_entered"), gc.Callable(cb), 0)
}


func (self Go) OnAreaExited(cb func(area gdclass.Area3D)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("area_exited"), gc.Callable(cb), 0)
}


func (self class) AsArea3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsArea3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("Area3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type SpaceOverride = classdb.Area3DSpaceOverride

const (
/*This area does not affect gravity/damping.*/
	SpaceOverrideDisabled SpaceOverride = 0
/*This area adds its gravity/damping values to whatever has been calculated so far (in [member priority] order).*/
	SpaceOverrideCombine SpaceOverride = 1
/*This area adds its gravity/damping values to whatever has been calculated so far (in [member priority] order), ignoring any lower priority areas.*/
	SpaceOverrideCombineReplace SpaceOverride = 2
/*This area replaces any gravity/damping, even the defaults, ignoring any lower priority areas.*/
	SpaceOverrideReplace SpaceOverride = 3
/*This area replaces any gravity/damping calculated so far (in [member priority] order), but keeps calculating the rest of the areas.*/
	SpaceOverrideReplaceCombine SpaceOverride = 4
)