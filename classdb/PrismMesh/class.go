// Code generated by the generate package DO NOT EDIT

// Package PrismMesh provides methods for working with PrismMesh object instances.
package PrismMesh

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
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/PrimitiveMesh"
import "graphics.gd/classdb/Resource"
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
import "graphics.gd/variant/Vector3"

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
Class representing a prism-shaped [PrimitiveMesh].
*/
type Instance [1]gdclass.PrismMesh

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPrismMesh() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PrismMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PrismMesh"))
	casted := Instance{*(*gdclass.PrismMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) LeftToRight() Float.X {
	return Float.X(Float.X(class(self).GetLeftToRight()))
}

func (self Instance) SetLeftToRight(value Float.X) {
	class(self).SetLeftToRight(float64(value))
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(Vector3.XYZ(value))
}

func (self Instance) SubdivideWidth() int {
	return int(int(class(self).GetSubdivideWidth()))
}

func (self Instance) SetSubdivideWidth(value int) {
	class(self).SetSubdivideWidth(int64(value))
}

func (self Instance) SubdivideHeight() int {
	return int(int(class(self).GetSubdivideHeight()))
}

func (self Instance) SetSubdivideHeight(value int) {
	class(self).SetSubdivideHeight(int64(value))
}

func (self Instance) SubdivideDepth() int {
	return int(int(class(self).GetSubdivideDepth()))
}

func (self Instance) SetSubdivideDepth(value int) {
	class(self).SetSubdivideDepth(int64(value))
}

//go:nosplit
func (self class) SetLeftToRight(left_to_right float64) { //gd:PrismMesh.set_left_to_right
	var frame = callframe.New()
	callframe.Arg(frame, left_to_right)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_set_left_to_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLeftToRight() float64 { //gd:PrismMesh.get_left_to_right
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_get_left_to_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSize(size Vector3.XYZ) { //gd:PrismMesh.set_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() Vector3.XYZ { //gd:PrismMesh.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdivideWidth(segments int64) { //gd:PrismMesh.set_subdivide_width
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_set_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdivideWidth() int64 { //gd:PrismMesh.get_subdivide_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_get_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdivideHeight(segments int64) { //gd:PrismMesh.set_subdivide_height
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_set_subdivide_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdivideHeight() int64 { //gd:PrismMesh.get_subdivide_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_get_subdivide_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdivideDepth(segments int64) { //gd:PrismMesh.set_subdivide_depth
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_set_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdivideDepth() int64 { //gd:PrismMesh.get_subdivide_depth
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PrismMesh.Bind_get_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPrismMesh() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPrismMesh() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsPrismMesh() Instance { return self.Super().AsPrismMesh() }
func (self class) AsPrimitiveMesh() PrimitiveMesh.Advanced {
	return *((*PrimitiveMesh.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsPrimitiveMesh() PrimitiveMesh.Instance {
	return self.Super().AsPrimitiveMesh()
}
func (self Instance) AsPrimitiveMesh() PrimitiveMesh.Instance {
	return *((*PrimitiveMesh.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMesh() Mesh.Advanced         { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsMesh() Mesh.Instance { return self.Super().AsMesh() }
func (self Instance) AsMesh() Mesh.Instance      { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsResource() Resource.Instance { return self.Super().AsResource() }
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsRefCounted() [1]gd.RefCounted { return self.Super().AsRefCounted() }
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PrimitiveMesh.Advanced(self.AsPrimitiveMesh()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PrimitiveMesh.Instance(self.AsPrimitiveMesh()), name)
	}
}
func init() {
	gdclass.Register("PrismMesh", func(ptr gd.Object) any { return [1]gdclass.PrismMesh{*(*gdclass.PrismMesh)(unsafe.Pointer(&ptr))} })
}
