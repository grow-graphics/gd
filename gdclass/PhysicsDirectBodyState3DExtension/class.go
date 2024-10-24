package PhysicsDirectBodyState3DExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PhysicsDirectBodyState3D"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class extends [PhysicsDirectBodyState3D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectBodyState3D].
	// PhysicsDirectBodyState3DExtension methods that can be overridden by a [Class] that extends it.
	type PhysicsDirectBodyState3DExtension interface {
		GetTotalGravity() gd.Vector3
		GetTotalLinearDamp() float64
		GetTotalAngularDamp() float64
		GetCenterOfMass() gd.Vector3
		GetCenterOfMassLocal() gd.Vector3
		GetPrincipalInertiaAxes() gd.Basis
		GetInverseMass() float64
		GetInverseInertia() gd.Vector3
		GetInverseInertiaTensor() gd.Basis
		SetLinearVelocity(velocity gd.Vector3) 
		GetLinearVelocity() gd.Vector3
		SetAngularVelocity(velocity gd.Vector3) 
		GetAngularVelocity() gd.Vector3
		SetTransform(transform gd.Transform3D) 
		GetTransform() gd.Transform3D
		GetVelocityAtLocalPosition(local_position gd.Vector3) gd.Vector3
		ApplyCentralImpulse(impulse gd.Vector3) 
		ApplyImpulse(impulse gd.Vector3, position gd.Vector3) 
		ApplyTorqueImpulse(impulse gd.Vector3) 
		ApplyCentralForce(force gd.Vector3) 
		ApplyForce(force gd.Vector3, position gd.Vector3) 
		ApplyTorque(torque gd.Vector3) 
		AddConstantCentralForce(force gd.Vector3) 
		AddConstantForce(force gd.Vector3, position gd.Vector3) 
		AddConstantTorque(torque gd.Vector3) 
		SetConstantForce(force gd.Vector3) 
		GetConstantForce() gd.Vector3
		SetConstantTorque(torque gd.Vector3) 
		GetConstantTorque() gd.Vector3
		SetSleepState(enabled bool) 
		IsSleeping() bool
		GetContactCount() int
		GetContactLocalPosition(contact_idx int) gd.Vector3
		GetContactLocalNormal(contact_idx int) gd.Vector3
		GetContactImpulse(contact_idx int) gd.Vector3
		GetContactLocalShape(contact_idx int) int
		GetContactLocalVelocityAtPosition(contact_idx int) gd.Vector3
		GetContactCollider(contact_idx int) gd.RID
		GetContactColliderPosition(contact_idx int) gd.Vector3
		GetContactColliderId(contact_idx int) int
		GetContactColliderObject(contact_idx int) gd.Object
		GetContactColliderShape(contact_idx int) int
		GetContactColliderVelocityAtPosition(contact_idx int) gd.Vector3
		GetStep() float64
		IntegrateForces() 
		GetSpaceState() gdclass.PhysicsDirectSpaceState3D
	}

*/
type Go [1]classdb.PhysicsDirectBodyState3DExtension
func (Go) _get_total_gravity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_total_linear_damp(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Go) _get_total_angular_damp(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Go) _get_center_of_mass(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_principal_inertia_axes(impl func(ptr unsafe.Pointer) gd.Basis, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_inverse_mass(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Go) _get_inverse_inertia(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_inverse_inertia_tensor(impl func(ptr unsafe.Pointer) gd.Basis, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var velocity = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
		gc.End()
	}
}
func (Go) _get_linear_velocity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var velocity = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
		gc.End()
	}
}
func (Go) _get_angular_velocity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _set_transform(impl func(ptr unsafe.Pointer, transform gd.Transform3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, transform)
		gc.End()
	}
}
func (Go) _get_transform(impl func(ptr unsafe.Pointer) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position gd.Vector3) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var local_position = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
		gc.End()
	}
}
func (Go) _apply_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse, position)
		gc.End()
	}
}
func (Go) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
		gc.End()
	}
}
func (Go) _apply_central_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		gc.End()
	}
}
func (Go) _apply_force(impl func(ptr unsafe.Pointer, force gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
		gc.End()
	}
}
func (Go) _apply_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		gc.End()
	}
}
func (Go) _add_constant_central_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		gc.End()
	}
}
func (Go) _add_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
		gc.End()
	}
}
func (Go) _add_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		gc.End()
	}
}
func (Go) _set_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		gc.End()
	}
}
func (Go) _get_constant_force(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _set_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		gc.End()
	}
}
func (Go) _get_constant_torque(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var enabled = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enabled)
		gc.End()
	}
}
func (Go) _is_sleeping(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_contact_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx int) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		gc.End()
	}
}
func (Go) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Go) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Go) _get_step(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Go) _integrate_forces(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Go) _get_space_state(impl func(ptr unsafe.Pointer) gdclass.PhysicsDirectSpaceState3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsDirectBodyState3DExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PhysicsDirectBodyState3DExtension"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (class) _get_total_gravity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_total_linear_damp(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_total_angular_damp(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_center_of_mass(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_principal_inertia_axes(impl func(ptr unsafe.Pointer) gd.Basis, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_inverse_mass(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_inverse_inertia(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_inverse_inertia_tensor(impl func(ptr unsafe.Pointer) gd.Basis, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var velocity = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
		ctx.End()
	}
}

func (class) _get_linear_velocity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var velocity = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
		ctx.End()
	}
}

func (class) _get_angular_velocity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_transform(impl func(ptr unsafe.Pointer, transform gd.Transform3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, transform)
		ctx.End()
	}
}

func (class) _get_transform(impl func(ptr unsafe.Pointer) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position gd.Vector3) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var local_position = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
		ctx.End()
	}
}

func (class) _apply_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse, position)
		ctx.End()
	}
}

func (class) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
		ctx.End()
	}
}

func (class) _apply_central_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

func (class) _apply_force(impl func(ptr unsafe.Pointer, force gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
		ctx.End()
	}
}

func (class) _apply_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

func (class) _add_constant_central_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

func (class) _add_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
		ctx.End()
	}
}

func (class) _add_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

func (class) _set_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

func (class) _get_constant_force(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

func (class) _get_constant_torque(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var enabled = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enabled)
		ctx.End()
	}
}

func (class) _is_sleeping(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		ctx.End()
	}
}

func (class) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_step(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _integrate_forces(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _get_space_state(impl func(ptr unsafe.Pointer) gdclass.PhysicsDirectSpaceState3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

func (self class) AsPhysicsDirectBodyState3DExtension() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsDirectBodyState3DExtension() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsDirectBodyState3D() PhysicsDirectBodyState3D.GD { return *((*PhysicsDirectBodyState3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsDirectBodyState3D() PhysicsDirectBodyState3D.Go { return *((*PhysicsDirectBodyState3D.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_total_gravity": return reflect.ValueOf(self._get_total_gravity);
	case "_get_total_linear_damp": return reflect.ValueOf(self._get_total_linear_damp);
	case "_get_total_angular_damp": return reflect.ValueOf(self._get_total_angular_damp);
	case "_get_center_of_mass": return reflect.ValueOf(self._get_center_of_mass);
	case "_get_center_of_mass_local": return reflect.ValueOf(self._get_center_of_mass_local);
	case "_get_principal_inertia_axes": return reflect.ValueOf(self._get_principal_inertia_axes);
	case "_get_inverse_mass": return reflect.ValueOf(self._get_inverse_mass);
	case "_get_inverse_inertia": return reflect.ValueOf(self._get_inverse_inertia);
	case "_get_inverse_inertia_tensor": return reflect.ValueOf(self._get_inverse_inertia_tensor);
	case "_set_linear_velocity": return reflect.ValueOf(self._set_linear_velocity);
	case "_get_linear_velocity": return reflect.ValueOf(self._get_linear_velocity);
	case "_set_angular_velocity": return reflect.ValueOf(self._set_angular_velocity);
	case "_get_angular_velocity": return reflect.ValueOf(self._get_angular_velocity);
	case "_set_transform": return reflect.ValueOf(self._set_transform);
	case "_get_transform": return reflect.ValueOf(self._get_transform);
	case "_get_velocity_at_local_position": return reflect.ValueOf(self._get_velocity_at_local_position);
	case "_apply_central_impulse": return reflect.ValueOf(self._apply_central_impulse);
	case "_apply_impulse": return reflect.ValueOf(self._apply_impulse);
	case "_apply_torque_impulse": return reflect.ValueOf(self._apply_torque_impulse);
	case "_apply_central_force": return reflect.ValueOf(self._apply_central_force);
	case "_apply_force": return reflect.ValueOf(self._apply_force);
	case "_apply_torque": return reflect.ValueOf(self._apply_torque);
	case "_add_constant_central_force": return reflect.ValueOf(self._add_constant_central_force);
	case "_add_constant_force": return reflect.ValueOf(self._add_constant_force);
	case "_add_constant_torque": return reflect.ValueOf(self._add_constant_torque);
	case "_set_constant_force": return reflect.ValueOf(self._set_constant_force);
	case "_get_constant_force": return reflect.ValueOf(self._get_constant_force);
	case "_set_constant_torque": return reflect.ValueOf(self._set_constant_torque);
	case "_get_constant_torque": return reflect.ValueOf(self._get_constant_torque);
	case "_set_sleep_state": return reflect.ValueOf(self._set_sleep_state);
	case "_is_sleeping": return reflect.ValueOf(self._is_sleeping);
	case "_get_contact_count": return reflect.ValueOf(self._get_contact_count);
	case "_get_contact_local_position": return reflect.ValueOf(self._get_contact_local_position);
	case "_get_contact_local_normal": return reflect.ValueOf(self._get_contact_local_normal);
	case "_get_contact_impulse": return reflect.ValueOf(self._get_contact_impulse);
	case "_get_contact_local_shape": return reflect.ValueOf(self._get_contact_local_shape);
	case "_get_contact_local_velocity_at_position": return reflect.ValueOf(self._get_contact_local_velocity_at_position);
	case "_get_contact_collider": return reflect.ValueOf(self._get_contact_collider);
	case "_get_contact_collider_position": return reflect.ValueOf(self._get_contact_collider_position);
	case "_get_contact_collider_id": return reflect.ValueOf(self._get_contact_collider_id);
	case "_get_contact_collider_object": return reflect.ValueOf(self._get_contact_collider_object);
	case "_get_contact_collider_shape": return reflect.ValueOf(self._get_contact_collider_shape);
	case "_get_contact_collider_velocity_at_position": return reflect.ValueOf(self._get_contact_collider_velocity_at_position);
	case "_get_step": return reflect.ValueOf(self._get_step);
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	case "_get_space_state": return reflect.ValueOf(self._get_space_state);
	default: return gd.VirtualByName(self.AsPhysicsDirectBodyState3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_total_gravity": return reflect.ValueOf(self._get_total_gravity);
	case "_get_total_linear_damp": return reflect.ValueOf(self._get_total_linear_damp);
	case "_get_total_angular_damp": return reflect.ValueOf(self._get_total_angular_damp);
	case "_get_center_of_mass": return reflect.ValueOf(self._get_center_of_mass);
	case "_get_center_of_mass_local": return reflect.ValueOf(self._get_center_of_mass_local);
	case "_get_principal_inertia_axes": return reflect.ValueOf(self._get_principal_inertia_axes);
	case "_get_inverse_mass": return reflect.ValueOf(self._get_inverse_mass);
	case "_get_inverse_inertia": return reflect.ValueOf(self._get_inverse_inertia);
	case "_get_inverse_inertia_tensor": return reflect.ValueOf(self._get_inverse_inertia_tensor);
	case "_set_linear_velocity": return reflect.ValueOf(self._set_linear_velocity);
	case "_get_linear_velocity": return reflect.ValueOf(self._get_linear_velocity);
	case "_set_angular_velocity": return reflect.ValueOf(self._set_angular_velocity);
	case "_get_angular_velocity": return reflect.ValueOf(self._get_angular_velocity);
	case "_set_transform": return reflect.ValueOf(self._set_transform);
	case "_get_transform": return reflect.ValueOf(self._get_transform);
	case "_get_velocity_at_local_position": return reflect.ValueOf(self._get_velocity_at_local_position);
	case "_apply_central_impulse": return reflect.ValueOf(self._apply_central_impulse);
	case "_apply_impulse": return reflect.ValueOf(self._apply_impulse);
	case "_apply_torque_impulse": return reflect.ValueOf(self._apply_torque_impulse);
	case "_apply_central_force": return reflect.ValueOf(self._apply_central_force);
	case "_apply_force": return reflect.ValueOf(self._apply_force);
	case "_apply_torque": return reflect.ValueOf(self._apply_torque);
	case "_add_constant_central_force": return reflect.ValueOf(self._add_constant_central_force);
	case "_add_constant_force": return reflect.ValueOf(self._add_constant_force);
	case "_add_constant_torque": return reflect.ValueOf(self._add_constant_torque);
	case "_set_constant_force": return reflect.ValueOf(self._set_constant_force);
	case "_get_constant_force": return reflect.ValueOf(self._get_constant_force);
	case "_set_constant_torque": return reflect.ValueOf(self._set_constant_torque);
	case "_get_constant_torque": return reflect.ValueOf(self._get_constant_torque);
	case "_set_sleep_state": return reflect.ValueOf(self._set_sleep_state);
	case "_is_sleeping": return reflect.ValueOf(self._is_sleeping);
	case "_get_contact_count": return reflect.ValueOf(self._get_contact_count);
	case "_get_contact_local_position": return reflect.ValueOf(self._get_contact_local_position);
	case "_get_contact_local_normal": return reflect.ValueOf(self._get_contact_local_normal);
	case "_get_contact_impulse": return reflect.ValueOf(self._get_contact_impulse);
	case "_get_contact_local_shape": return reflect.ValueOf(self._get_contact_local_shape);
	case "_get_contact_local_velocity_at_position": return reflect.ValueOf(self._get_contact_local_velocity_at_position);
	case "_get_contact_collider": return reflect.ValueOf(self._get_contact_collider);
	case "_get_contact_collider_position": return reflect.ValueOf(self._get_contact_collider_position);
	case "_get_contact_collider_id": return reflect.ValueOf(self._get_contact_collider_id);
	case "_get_contact_collider_object": return reflect.ValueOf(self._get_contact_collider_object);
	case "_get_contact_collider_shape": return reflect.ValueOf(self._get_contact_collider_shape);
	case "_get_contact_collider_velocity_at_position": return reflect.ValueOf(self._get_contact_collider_velocity_at_position);
	case "_get_step": return reflect.ValueOf(self._get_step);
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	case "_get_space_state": return reflect.ValueOf(self._get_space_state);
	default: return gd.VirtualByName(self.AsPhysicsDirectBodyState3D(), name)
	}
}
func init() {classdb.Register("PhysicsDirectBodyState3DExtension", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}