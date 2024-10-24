package SoftBody3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/MeshInstance3D"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A deformable 3D physics mesh. Used to create elastic or deformable objects such as cloth, rubber, or other flexible materials.
Additionally, [SoftBody3D] is subject to wind forces defined in [Area3D] (see [member Area3D.wind_source_path], [member Area3D.wind_force_magnitude], and [member Area3D.wind_attenuation_factor]).
[b]Note:[/b] There are many known bugs in [SoftBody3D]. Therefore, it's not recommended to use them for things that can affect gameplay (such as trampolines).

*/
type Go [1]classdb.SoftBody3D

/*
Returns the internal [RID] used by the [PhysicsServer3D] for this body.
*/
func (self Go) GetPhysicsRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetPhysicsRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCollisionLayerValue(gd.Int(layer_number)))
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

/*
Returns local translation of a vertex in the surface array.
*/
func (self Go) GetPointTransform(point_index int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetPointTransform(gd.Int(point_index)))
}

/*
Sets the pinned state of a surface vertex. When set to [code]true[/code], the optional [param attachment_path] can define a [Node3D] the pinned vertex will be attached to.
*/
func (self Go) SetPointPinned(point_index int, pinned bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPointPinned(gd.Int(point_index), pinned, gc.String("").NodePath(gc))
}

/*
Returns [code]true[/code] if vertex is set to pinned.
*/
func (self Go) IsPointPinned(point_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPointPinned(gd.Int(point_index)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SoftBody3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SoftBody3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) CollisionLayer() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionLayer()))
}

func (self Go) SetCollisionLayer(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionLayer(gd.Int(value))
}

func (self Go) CollisionMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetCollisionMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) ParentCollisionIgnore() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetParentCollisionIgnore(gc).String())
}

func (self Go) SetParentCollisionIgnore(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParentCollisionIgnore(gc.String(value).NodePath(gc))
}

func (self Go) SimulationPrecision() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSimulationPrecision()))
}

func (self Go) SetSimulationPrecision(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSimulationPrecision(gd.Int(value))
}

func (self Go) TotalMass() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTotalMass()))
}

func (self Go) SetTotalMass(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTotalMass(gd.Float(value))
}

func (self Go) LinearStiffness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLinearStiffness()))
}

func (self Go) SetLinearStiffness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearStiffness(gd.Float(value))
}

func (self Go) PressureCoefficient() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPressureCoefficient()))
}

func (self Go) SetPressureCoefficient(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPressureCoefficient(gd.Float(value))
}

func (self Go) DampingCoefficient() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDampingCoefficient()))
}

func (self Go) SetDampingCoefficient(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDampingCoefficient(gd.Float(value))
}

func (self Go) DragCoefficient() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDragCoefficient()))
}

func (self Go) SetDragCoefficient(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragCoefficient(gd.Float(value))
}

func (self Go) RayPickable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsRayPickable())
}

func (self Go) SetRayPickable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRayPickable(value)
}

func (self Go) DisableMode() classdb.SoftBody3DDisableMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.SoftBody3DDisableMode(class(self).GetDisableMode())
}

func (self Go) SetDisableMode(value classdb.SoftBody3DDisableMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisableMode(value)
}

/*
Returns the internal [RID] used by the [PhysicsServer3D] for this body.
*/
//go:nosplit
func (self class) GetPhysicsRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_physics_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionMask(collision_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionLayer(collision_layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionLayerValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionLayerValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParentCollisionIgnore(parent_collision_ignore gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(parent_collision_ignore))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_parent_collision_ignore, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParentCollisionIgnore(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_parent_collision_ignore, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisableMode(mode classdb.SoftBody3DDisableMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_disable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDisableMode() classdb.SoftBody3DDisableMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SoftBody3DDisableMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_disable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_collision_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_add_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_remove_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSimulationPrecision(simulation_precision gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, simulation_precision)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_simulation_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimulationPrecision() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_simulation_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTotalMass(mass gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_total_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTotalMass() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_total_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearStiffness(linear_stiffness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_stiffness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_linear_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearStiffness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_linear_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPressureCoefficient(pressure_coefficient gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressure_coefficient)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_pressure_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPressureCoefficient() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_pressure_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDampingCoefficient(damping_coefficient gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, damping_coefficient)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_damping_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDampingCoefficient() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_damping_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragCoefficient(drag_coefficient gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, drag_coefficient)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_drag_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDragCoefficient() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_drag_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns local translation of a vertex in the surface array.
*/
//go:nosplit
func (self class) GetPointTransform(point_index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_get_point_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the pinned state of a surface vertex. When set to [code]true[/code], the optional [param attachment_path] can define a [Node3D] the pinned vertex will be attached to.
*/
//go:nosplit
func (self class) SetPointPinned(point_index gd.Int, pinned bool, attachment_path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point_index)
	callframe.Arg(frame, pinned)
	callframe.Arg(frame, mmm.Get(attachment_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_point_pinned, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if vertex is set to pinned.
*/
//go:nosplit
func (self class) IsPointPinned(point_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_is_point_pinned, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRayPickable(ray_pickable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ray_pickable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_set_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRayPickable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SoftBody3D.Bind_is_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSoftBody3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSoftBody3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMeshInstance3D() MeshInstance3D.GD { return *((*MeshInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMeshInstance3D() MeshInstance3D.Go { return *((*MeshInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.GD { return *((*GeometryInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() GeometryInstance3D.Go { return *((*GeometryInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMeshInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMeshInstance3D(), name)
	}
}
func init() {classdb.Register("SoftBody3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type DisableMode = classdb.SoftBody3DDisableMode

const (
/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], remove from the physics simulation to stop all physics interactions with this [SoftBody3D].
Automatically re-added to the physics simulation when the [Node] is processed again.*/
	DisableModeRemove DisableMode = 0
/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], do not affect the physics simulation.*/
	DisableModeKeepActive DisableMode = 1
)