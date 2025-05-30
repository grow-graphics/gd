// Code generated by the generate package DO NOT EDIT

// Package SpringBoneCollisionPlane3D provides methods for working with SpringBoneCollisionPlane3D object instances.
package SpringBoneCollisionPlane3D

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
import "graphics.gd/classdb/SpringBoneCollision3D"
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
A infinite plane collision that interacts with [SpringBoneSimulator3D]. It is an infinite size XZ plane, and the +Y direction is treated as normal.
*/
type Instance [1]gdclass.SpringBoneCollisionPlane3D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSpringBoneCollisionPlane3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SpringBoneCollisionPlane3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SpringBoneCollisionPlane3D"))
	casted := Instance{*(*gdclass.SpringBoneCollisionPlane3D)(unsafe.Pointer(&object))}
	return casted
}

func (self class) AsSpringBoneCollisionPlane3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSpringBoneCollisionPlane3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsSpringBoneCollisionPlane3D() Instance {
	return self.Super().AsSpringBoneCollisionPlane3D()
}
func (self class) AsSpringBoneCollision3D() SpringBoneCollision3D.Advanced {
	return *((*SpringBoneCollision3D.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsSpringBoneCollision3D() SpringBoneCollision3D.Instance {
	return self.Super().AsSpringBoneCollision3D()
}
func (self Instance) AsSpringBoneCollision3D() SpringBoneCollision3D.Instance {
	return *((*SpringBoneCollision3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(SpringBoneCollision3D.Advanced(self.AsSpringBoneCollision3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SpringBoneCollision3D.Instance(self.AsSpringBoneCollision3D()), name)
	}
}
func init() {
	gdclass.Register("SpringBoneCollisionPlane3D", func(ptr gd.Object) any {
		return [1]gdclass.SpringBoneCollisionPlane3D{*(*gdclass.SpringBoneCollisionPlane3D)(unsafe.Pointer(&ptr))}
	})
}
