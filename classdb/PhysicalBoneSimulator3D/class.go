// Code generated by the generate package DO NOT EDIT

// Package PhysicalBoneSimulator3D provides methods for working with PhysicalBoneSimulator3D object instances.
package PhysicalBoneSimulator3D

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
import "graphics.gd/classdb/SkeletonModifier3D"
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
Node that can be the parent of [PhysicalBone3D] and can apply the simulation results to [Skeleton3D].
*/
type Instance [1]gdclass.PhysicalBoneSimulator3D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.PhysicalBoneSimulator3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicalBoneSimulator3D() Instance
}

/*
Returns a boolean that indicates whether the [PhysicalBoneSimulator3D] is running and simulating.
*/
func (self Instance) IsSimulatingPhysics() bool { //gd:PhysicalBoneSimulator3D.is_simulating_physics
	return bool(Advanced(self).IsSimulatingPhysics())
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
func (self Instance) PhysicalBonesStopSimulation() { //gd:PhysicalBoneSimulator3D.physical_bones_stop_simulation
	Advanced(self).PhysicalBonesStopSimulation()
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
func (self Instance) PhysicalBonesStartSimulation() { //gd:PhysicalBoneSimulator3D.physical_bones_start_simulation
	Advanced(self).PhysicalBonesStartSimulation(gd.ArrayFromSlice[Array.Contains[String.Name]]([1][]string{}[0]))
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
func (self Expanded) PhysicalBonesStartSimulation(bones []string) { //gd:PhysicalBoneSimulator3D.physical_bones_start_simulation
	Advanced(self).PhysicalBonesStartSimulation(gd.ArrayFromSlice[Array.Contains[String.Name]](bones))
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Instance) PhysicalBonesAddCollisionException(exception RID.Body3D) { //gd:PhysicalBoneSimulator3D.physical_bones_add_collision_exception
	Advanced(self).PhysicalBonesAddCollisionException(RID.Any(exception))
}

/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
func (self Instance) PhysicalBonesRemoveCollisionException(exception RID.Body3D) { //gd:PhysicalBoneSimulator3D.physical_bones_remove_collision_exception
	Advanced(self).PhysicalBonesRemoveCollisionException(RID.Any(exception))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicalBoneSimulator3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalBoneSimulator3D"))
	casted := Instance{*(*gdclass.PhysicalBoneSimulator3D)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns a boolean that indicates whether the [PhysicalBoneSimulator3D] is running and simulating.
*/
//go:nosplit
func (self class) IsSimulatingPhysics() bool { //gd:PhysicalBoneSimulator3D.is_simulating_physics
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to stop simulating.
*/
//go:nosplit
func (self class) PhysicalBonesStopSimulation() { //gd:PhysicalBoneSimulator3D.physical_bones_stop_simulation
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_physical_bones_stop_simulation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Tells the [PhysicalBone3D] nodes in the Skeleton to start simulating and reacting to the physics world.
Optionally, a list of bone names can be passed-in, allowing only the passed-in bones to be simulated.
*/
//go:nosplit
func (self class) PhysicalBonesStartSimulation(bones Array.Contains[String.Name]) { //gd:PhysicalBoneSimulator3D.physical_bones_start_simulation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(bones)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_physical_bones_start_simulation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesAddCollisionException(exception RID.Any) { //gd:PhysicalBoneSimulator3D.physical_bones_add_collision_exception
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_physical_bones_add_collision_exception, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a collision exception to the physical bone.
Works just like the [RigidBody3D] node.
*/
//go:nosplit
func (self class) PhysicalBonesRemoveCollisionException(exception RID.Any) { //gd:PhysicalBoneSimulator3D.physical_bones_remove_collision_exception
	var frame = callframe.New()
	callframe.Arg(frame, exception)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBoneSimulator3D.Bind_physical_bones_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsPhysicalBoneSimulator3D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicalBoneSimulator3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsPhysicalBoneSimulator3D() Instance {
	return self.Super().AsPhysicalBoneSimulator3D()
}
func (self class) AsSkeletonModifier3D() SkeletonModifier3D.Advanced {
	return *((*SkeletonModifier3D.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsSkeletonModifier3D() SkeletonModifier3D.Instance {
	return self.Super().AsSkeletonModifier3D()
}
func (self Instance) AsSkeletonModifier3D() SkeletonModifier3D.Instance {
	return *((*SkeletonModifier3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced         { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode3D() Node3D.Instance { return self.Super().AsNode3D() }
func (self Instance) AsNode3D() Node3D.Instance      { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced             { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance     { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance          { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModifier3D.Advanced(self.AsSkeletonModifier3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModifier3D.Instance(self.AsSkeletonModifier3D()), name)
	}
}
func init() {
	gdclass.Register("PhysicalBoneSimulator3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicalBoneSimulator3D{*(*gdclass.PhysicalBoneSimulator3D)(unsafe.Pointer(&ptr))}
	})
}
