// Code generated by the generate package DO NOT EDIT

// Package AnimatableBody2D provides methods for working with AnimatableBody2D object instances.
package AnimatableBody2D

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
import "graphics.gd/classdb/StaticBody2D"
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
An animatable 2D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform2D].
When [AnimatableBody2D] is moved, its linear and angular velocity are estimated and used to affect other physics bodies in its path. This makes it useful for moving platforms, doors, and other moving objects.
*/
type Instance [1]gdclass.AnimatableBody2D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimatableBody2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimatableBody2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimatableBody2D"))
	casted := Instance{*(*gdclass.AnimatableBody2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) SyncToPhysics() bool {
	return bool(class(self).IsSyncToPhysicsEnabled())
}

func (self Instance) SetSyncToPhysics(value bool) {
	class(self).SetSyncToPhysics(value)
}

//go:nosplit
func (self class) SetSyncToPhysics(enable bool) { //gd:AnimatableBody2D.set_sync_to_physics
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatableBody2D.Bind_set_sync_to_physics, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSyncToPhysicsEnabled() bool { //gd:AnimatableBody2D.is_sync_to_physics_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatableBody2D.Bind_is_sync_to_physics_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimatableBody2D() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimatableBody2D() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsAnimatableBody2D() Instance { return self.Super().AsAnimatableBody2D() }
func (self class) AsStaticBody2D() StaticBody2D.Advanced {
	return *((*StaticBody2D.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsStaticBody2D() StaticBody2D.Instance {
	return self.Super().AsStaticBody2D()
}
func (self Instance) AsStaticBody2D() StaticBody2D.Instance {
	return *((*StaticBody2D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(StaticBody2D.Advanced(self.AsStaticBody2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StaticBody2D.Instance(self.AsStaticBody2D()), name)
	}
}
func init() {
	gdclass.Register("AnimatableBody2D", func(ptr gd.Object) any {
		return [1]gdclass.AnimatableBody2D{*(*gdclass.AnimatableBody2D)(unsafe.Pointer(&ptr))}
	})
}
