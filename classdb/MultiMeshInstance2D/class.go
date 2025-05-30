// Code generated by the generate package DO NOT EDIT

// Package MultiMeshInstance2D provides methods for working with MultiMeshInstance2D object instances.
package MultiMeshInstance2D

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
import "graphics.gd/classdb/MultiMesh"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/Texture2D"
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
[MultiMeshInstance2D] is a specialized node to instance a [MultiMesh] resource in 2D.
Usage is the same as [MultiMeshInstance3D].
*/
type Instance [1]gdclass.MultiMeshInstance2D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMultiMeshInstance2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiMeshInstance2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiMeshInstance2D"))
	casted := Instance{*(*gdclass.MultiMeshInstance2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Multimesh() MultiMesh.Instance {
	return MultiMesh.Instance(class(self).GetMultimesh())
}

func (self Instance) SetMultimesh(value MultiMesh.Instance) {
	class(self).SetMultimesh(value)
}

func (self Instance) Texture() Texture2D.Instance {
	return Texture2D.Instance(class(self).GetTexture())
}

func (self Instance) SetTexture(value Texture2D.Instance) {
	class(self).SetTexture(value)
}

//go:nosplit
func (self class) SetMultimesh(multimesh [1]gdclass.MultiMesh) { //gd:MultiMeshInstance2D.set_multimesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(multimesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance2D.Bind_set_multimesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMultimesh() [1]gdclass.MultiMesh { //gd:MultiMeshInstance2D.get_multimesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance2D.Bind_get_multimesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MultiMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.MultiMesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) { //gd:MultiMeshInstance2D.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D { //gd:MultiMeshInstance2D.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMeshInstance2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self Instance) OnTextureChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("texture_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsMultiMeshInstance2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiMeshInstance2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsMultiMeshInstance2D() Instance {
	return self.Super().AsMultiMeshInstance2D()
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("MultiMeshInstance2D", func(ptr gd.Object) any {
		return [1]gdclass.MultiMeshInstance2D{*(*gdclass.MultiMeshInstance2D)(unsafe.Pointer(&ptr))}
	})
}
