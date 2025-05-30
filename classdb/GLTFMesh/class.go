// Code generated by the generate package DO NOT EDIT

// Package GLTFMesh provides methods for working with GLTFMesh object instances.
package GLTFMesh

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
import "graphics.gd/classdb/ImporterMesh"
import "graphics.gd/classdb/Material"
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
GLTFMesh handles 3D mesh data imported from glTF files. It includes properties for blend channels, blend weights, instance materials, and the mesh itself.
*/
type Instance [1]gdclass.GLTFMesh

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFMesh() Instance
}

/*
Gets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the glTF file), and the return value can be anything you set. If nothing was set, the return value is [code]null[/code].
*/
func (self Instance) GetAdditionalData(extension_name string) any { //gd:GLTFMesh.get_additional_data
	return any(Advanced(self).GetAdditionalData(String.Name(String.New(extension_name))).Interface())
}

/*
Sets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the glTF file), and the second argument can be anything you want.
*/
func (self Instance) SetAdditionalData(extension_name string, additional_data any) { //gd:GLTFMesh.set_additional_data
	Advanced(self).SetAdditionalData(String.Name(String.New(extension_name)), variant.New(additional_data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFMesh"))
	casted := Instance{*(*gdclass.GLTFMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) OriginalName() string {
	return string(class(self).GetOriginalName().String())
}

func (self Instance) SetOriginalName(value string) {
	class(self).SetOriginalName(String.New(value))
}

func (self Instance) Mesh() ImporterMesh.Instance {
	return ImporterMesh.Instance(class(self).GetMesh())
}

func (self Instance) SetMesh(value ImporterMesh.Instance) {
	class(self).SetMesh(value)
}

func (self Instance) BlendWeights() []float32 {
	return []float32(slices.Collect(class(self).GetBlendWeights().Values()))
}

func (self Instance) SetBlendWeights(value []float32) {
	class(self).SetBlendWeights(Packed.New(value...))
}

func (self Instance) InstanceMaterials() []Material.Instance {
	return []Material.Instance(gd.ArrayAs[[]Material.Instance](gd.InternalArray(class(self).GetInstanceMaterials())))
}

func (self Instance) SetInstanceMaterials(value []Material.Instance) {
	class(self).SetInstanceMaterials(gd.ArrayFromSlice[Array.Contains[[1]gdclass.Material]](value))
}

//go:nosplit
func (self class) GetOriginalName() String.Readable { //gd:GLTFMesh.get_original_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_original_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOriginalName(original_name String.Readable) { //gd:GLTFMesh.set_original_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(original_name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_original_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() [1]gdclass.ImporterMesh { //gd:GLTFMesh.get_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ImporterMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.ImporterMesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMesh(mesh [1]gdclass.ImporterMesh) { //gd:GLTFMesh.set_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendWeights() Packed.Array[float32] { //gd:GLTFMesh.get_blend_weights
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_blend_weights, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendWeights(blend_weights Packed.Array[float32]) { //gd:GLTFMesh.set_blend_weights
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](blend_weights)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_blend_weights, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInstanceMaterials() Array.Contains[[1]gdclass.Material] { //gd:GLTFMesh.get_instance_materials
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_instance_materials, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Material]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInstanceMaterials(instance_materials Array.Contains[[1]gdclass.Material]) { //gd:GLTFMesh.set_instance_materials
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(instance_materials)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_instance_materials, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the glTF file), and the return value can be anything you set. If nothing was set, the return value is [code]null[/code].
*/
//go:nosplit
func (self class) GetAdditionalData(extension_name String.Name) variant.Any { //gd:GLTFMesh.get_additional_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(extension_name)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets additional arbitrary data in this [GLTFMesh] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the glTF file), and the second argument can be anything you want.
*/
//go:nosplit
func (self class) SetAdditionalData(extension_name String.Name, additional_data variant.Any) { //gd:GLTFMesh.set_additional_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(extension_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(additional_data)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFMesh.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsGLTFMesh() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFMesh() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsGLTFMesh() Instance { return self.Super().AsGLTFMesh() }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("GLTFMesh", func(ptr gd.Object) any { return [1]gdclass.GLTFMesh{*(*gdclass.GLTFMesh)(unsafe.Pointer(&ptr))} })
}
