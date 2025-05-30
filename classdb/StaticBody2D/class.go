// Code generated by the generate package DO NOT EDIT

// Package StaticBody2D provides methods for working with StaticBody2D object instances.
package StaticBody2D

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
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/CollisionObject2D"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/PhysicsBody2D"
import "graphics.gd/classdb/PhysicsMaterial"
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
import "graphics.gd/variant/Vector2"

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
A static 2D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform2D].
When [StaticBody2D] is moved, it is teleported to its new position without affecting other physics bodies in its path. If this is not desired, use [AnimatableBody2D] instead.
[StaticBody2D] is useful for completely static objects like floors and walls, as well as moving surfaces like conveyor belts and circular revolving platforms (by using [member constant_linear_velocity] and [member constant_angular_velocity]).
*/
type Instance [1]gdclass.StaticBody2D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStaticBody2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StaticBody2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StaticBody2D"))
	casted := Instance{*(*gdclass.StaticBody2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) PhysicsMaterialOverride() PhysicsMaterial.Instance {
	return PhysicsMaterial.Instance(class(self).GetPhysicsMaterialOverride())
}

func (self Instance) SetPhysicsMaterialOverride(value PhysicsMaterial.Instance) {
	class(self).SetPhysicsMaterialOverride(value)
}

func (self Instance) ConstantLinearVelocity() Vector2.XY {
	return Vector2.XY(class(self).GetConstantLinearVelocity())
}

func (self Instance) SetConstantLinearVelocity(value Vector2.XY) {
	class(self).SetConstantLinearVelocity(Vector2.XY(value))
}

func (self Instance) ConstantAngularVelocity() Float.X {
	return Float.X(Float.X(class(self).GetConstantAngularVelocity()))
}

func (self Instance) SetConstantAngularVelocity(value Float.X) {
	class(self).SetConstantAngularVelocity(float64(value))
}

//go:nosplit
func (self class) SetConstantLinearVelocity(vel Vector2.XY) { //gd:StaticBody2D.set_constant_linear_velocity
	var frame = callframe.New()
	callframe.Arg(frame, vel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody2D.Bind_set_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetConstantAngularVelocity(vel float64) { //gd:StaticBody2D.set_constant_angular_velocity
	var frame = callframe.New()
	callframe.Arg(frame, vel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody2D.Bind_set_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstantLinearVelocity() Vector2.XY { //gd:StaticBody2D.get_constant_linear_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody2D.Bind_get_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetConstantAngularVelocity() float64 { //gd:StaticBody2D.get_constant_angular_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody2D.Bind_get_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsMaterialOverride(physics_material_override [1]gdclass.PhysicsMaterial) { //gd:StaticBody2D.set_physics_material_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(physics_material_override[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody2D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsMaterialOverride() [1]gdclass.PhysicsMaterial { //gd:StaticBody2D.get_physics_material_override
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody2D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PhysicsMaterial{gd.PointerWithOwnershipTransferredToGo[gdclass.PhysicsMaterial](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsStaticBody2D() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStaticBody2D() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsStaticBody2D() Instance { return self.Super().AsStaticBody2D() }
func (self class) AsPhysicsBody2D() PhysicsBody2D.Advanced {
	return *((*PhysicsBody2D.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsPhysicsBody2D() PhysicsBody2D.Instance {
	return self.Super().AsPhysicsBody2D()
}
func (self Instance) AsPhysicsBody2D() PhysicsBody2D.Instance {
	return *((*PhysicsBody2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCollisionObject2D() CollisionObject2D.Advanced {
	return *((*CollisionObject2D.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsCollisionObject2D() CollisionObject2D.Instance {
	return self.Super().AsCollisionObject2D()
}
func (self Instance) AsCollisionObject2D() CollisionObject2D.Instance {
	return *((*CollisionObject2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode2D() Node2D.Advanced         { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode2D() Node2D.Instance { return self.Super().AsNode2D() }
func (self Instance) AsNode2D() Node2D.Instance      { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsCanvasItem() CanvasItem.Instance { return self.Super().AsCanvasItem() }
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced         { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance      { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PhysicsBody2D.Advanced(self.AsPhysicsBody2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PhysicsBody2D.Instance(self.AsPhysicsBody2D()), name)
	}
}
func init() {
	gdclass.Register("StaticBody2D", func(ptr gd.Object) any {
		return [1]gdclass.StaticBody2D{*(*gdclass.StaticBody2D)(unsafe.Pointer(&ptr))}
	})
}
