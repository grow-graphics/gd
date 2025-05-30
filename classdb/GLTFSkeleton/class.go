// Code generated by the generate package DO NOT EDIT

// Package GLTFSkeleton provides methods for working with GLTFSkeleton object instances.
package GLTFSkeleton

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
import "graphics.gd/classdb/BoneAttachment3D"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Skeleton3D"
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
type Instance [1]gdclass.GLTFSkeleton

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFSkeleton() Instance
}

func (self Instance) GetGodotSkeleton() Skeleton3D.Instance { //gd:GLTFSkeleton.get_godot_skeleton
	return Skeleton3D.Instance(Advanced(self).GetGodotSkeleton())
}
func (self Instance) GetBoneAttachmentCount() int { //gd:GLTFSkeleton.get_bone_attachment_count
	return int(int(Advanced(self).GetBoneAttachmentCount()))
}
func (self Instance) GetBoneAttachment(idx int) BoneAttachment3D.Instance { //gd:GLTFSkeleton.get_bone_attachment
	return BoneAttachment3D.Instance(Advanced(self).GetBoneAttachment(int64(idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFSkeleton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSkeleton"))
	casted := Instance{*(*gdclass.GLTFSkeleton)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Joints() []int32 {
	return []int32(slices.Collect(class(self).GetJoints().Values()))
}

func (self Instance) SetJoints(value []int32) {
	class(self).SetJoints(Packed.New(value...))
}

func (self Instance) Roots() []int32 {
	return []int32(slices.Collect(class(self).GetRoots().Values()))
}

func (self Instance) SetRoots(value []int32) {
	class(self).SetRoots(Packed.New(value...))
}

func (self Instance) UniqueNames() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetUniqueNames())))
}

func (self Instance) SetUniqueNames(value []string) {
	class(self).SetUniqueNames(gd.ArrayFromSlice[Array.Contains[String.Readable]](value))
}

func (self Instance) GodotBoneNode() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetGodotBoneNode()))
}

func (self Instance) SetGodotBoneNode(value map[any]any) {
	class(self).SetGodotBoneNode(gd.DictionaryFromMap(value))
}

//go:nosplit
func (self class) GetJoints() Packed.Array[int32] { //gd:GLTFSkeleton.get_joints
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_joints, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJoints(joints Packed.Array[int32]) { //gd:GLTFSkeleton.set_joints
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedInt32Array, int32](joints)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_joints, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoots() Packed.Array[int32] { //gd:GLTFSkeleton.get_roots
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_roots, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoots(roots Packed.Array[int32]) { //gd:GLTFSkeleton.set_roots
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedInt32Array, int32](roots)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_roots, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGodotSkeleton() [1]gdclass.Skeleton3D { //gd:GLTFSkeleton.get_godot_skeleton
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_godot_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skeleton3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Skeleton3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUniqueNames() Array.Contains[String.Readable] { //gd:GLTFSkeleton.get_unique_names
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_unique_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUniqueNames(unique_names Array.Contains[String.Readable]) { //gd:GLTFSkeleton.set_unique_names
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(unique_names)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_unique_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [Dictionary] that maps skeleton bone indices to the indices of glTF nodes. This property is unused during import, and only set during export. In a glTF file, a bone is a node, so Godot converts skeleton bones to glTF nodes.
*/
//go:nosplit
func (self class) GetGodotBoneNode() Dictionary.Any { //gd:GLTFSkeleton.get_godot_bone_node
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets a [Dictionary] that maps skeleton bone indices to the indices of glTF nodes. This property is unused during import, and only set during export. In a glTF file, a bone is a node, so Godot converts skeleton bones to glTF nodes.
*/
//go:nosplit
func (self class) SetGodotBoneNode(godot_bone_node Dictionary.Any) { //gd:GLTFSkeleton.set_godot_bone_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(godot_bone_node)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneAttachmentCount() int64 { //gd:GLTFSkeleton.get_bone_attachment_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_bone_attachment_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBoneAttachment(idx int64) [1]gdclass.BoneAttachment3D { //gd:GLTFSkeleton.get_bone_attachment
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_bone_attachment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.BoneAttachment3D{gd.PointerWithOwnershipTransferredToGo[gdclass.BoneAttachment3D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsGLTFSkeleton() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFSkeleton() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsGLTFSkeleton() Instance { return self.Super().AsGLTFSkeleton() }
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
	gdclass.Register("GLTFSkeleton", func(ptr gd.Object) any {
		return [1]gdclass.GLTFSkeleton{*(*gdclass.GLTFSkeleton)(unsafe.Pointer(&ptr))}
	})
}
