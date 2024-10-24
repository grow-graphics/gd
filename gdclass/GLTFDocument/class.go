package GLTFDocument

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
GLTFDocument supports reading data from a glTF file, buffer, or Godot scene. This data can then be written to the filesystem, buffer, or used to create a Godot scene.
All of the data in a GLTF scene is stored in the [GLTFState] class. GLTFDocument processes state objects, but does not contain any scene data itself. GLTFDocument has member variables to store export configuration settings such as the image format, but is otherwise stateless. Multiple scenes can be processed with the same settings using the same GLTFDocument object and different [GLTFState] objects.
GLTFDocument can be extended with arbitrary functionality by extending the [GLTFDocumentExtension] class and registering it with GLTFDocument via [method register_gltf_document_extension]. This allows for custom data to be imported and exported.

*/
type Go [1]classdb.GLTFDocument

/*
Takes a path to a GLTF file and imports the data at that file path to the given [GLTFState] object through the [param state] parameter.
[b]Note:[/b] The [param base_path] tells [method append_from_file] where to find dependencies and can be empty.
*/
func (self Go) AppendFromFile(path string, state gdclass.GLTFState) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).AppendFromFile(gc.String(path), state, gd.Int(0), gc.String("")))
}

/*
Takes a [PackedByteArray] defining a GLTF and imports the data to the given [GLTFState] object through the [param state] parameter.
[b]Note:[/b] The [param base_path] tells [method append_from_buffer] where to find dependencies and can be empty.
*/
func (self Go) AppendFromBuffer(bytes []byte, base_path string, state gdclass.GLTFState) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).AppendFromBuffer(gc.PackedByteSlice(bytes), gc.String(base_path), state, gd.Int(0)))
}

/*
Takes a Godot Engine scene node and exports it and its descendants to the given [GLTFState] object through the [param state] parameter.
*/
func (self Go) AppendFromScene(node gdclass.Node, state gdclass.GLTFState) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).AppendFromScene(node, state, gd.Int(0)))
}

/*
Takes a [GLTFState] object through the [param state] parameter and returns a Godot Engine scene node.
The [param bake_fps] parameter overrides the bake_fps in [param state].
*/
func (self Go) GenerateScene(state gdclass.GLTFState) gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node(class(self).GenerateScene(gc, state, gd.Float(30), false, true))
}

/*
Takes a [GLTFState] object through the [param state] parameter and returns a GLTF [PackedByteArray].
*/
func (self Go) GenerateBuffer(state gdclass.GLTFState) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).GenerateBuffer(gc, state).Bytes())
}

/*
Takes a [GLTFState] object through the [param state] parameter and writes a glTF file to the filesystem.
[b]Note:[/b] The extension of the glTF file determines if it is a .glb binary file or a .gltf text file.
*/
func (self Go) WriteToFilesystem(state gdclass.GLTFState, path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).WriteToFilesystem(state, gc.String(path)))
}

/*
Registers the given [GLTFDocumentExtension] instance with GLTFDocument. If [param first_priority] is true, this extension will be run first. Otherwise, it will be run last.
[b]Note:[/b] Like GLTFDocument itself, all GLTFDocumentExtension classes must be stateless in order to function properly. If you need to store data, use the [code]set_additional_data[/code] and [code]get_additional_data[/code] methods in [GLTFState] or [GLTFNode].
*/
func (self Go) RegisterGltfDocumentExtension(extension gdclass.GLTFDocumentExtension) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RegisterGltfDocumentExtension(gc, extension, false)
}

/*
Unregisters the given [GLTFDocumentExtension] instance.
*/
func (self Go) UnregisterGltfDocumentExtension(extension gdclass.GLTFDocumentExtension) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UnregisterGltfDocumentExtension(gc, extension)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFDocument
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFDocument"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ImageFormat() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetImageFormat(gc).String())
}

func (self Go) SetImageFormat(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetImageFormat(gc.String(value))
}

func (self Go) LossyQuality() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLossyQuality()))
}

func (self Go) SetLossyQuality(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLossyQuality(gd.Float(value))
}

func (self Go) RootNodeMode() classdb.GLTFDocumentRootNodeMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GLTFDocumentRootNodeMode(class(self).GetRootNodeMode())
}

func (self Go) SetRootNodeMode(value classdb.GLTFDocumentRootNodeMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRootNodeMode(value)
}

//go:nosplit
func (self class) SetImageFormat(image_format gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image_format))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_set_image_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetImageFormat(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_get_image_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLossyQuality(lossy_quality gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lossy_quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_set_lossy_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLossyQuality() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_get_lossy_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRootNodeMode(root_node_mode classdb.GLTFDocumentRootNodeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, root_node_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_set_root_node_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootNodeMode() classdb.GLTFDocumentRootNodeMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GLTFDocumentRootNodeMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_get_root_node_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Takes a path to a GLTF file and imports the data at that file path to the given [GLTFState] object through the [param state] parameter.
[b]Note:[/b] The [param base_path] tells [method append_from_file] where to find dependencies and can be empty.
*/
//go:nosplit
func (self class) AppendFromFile(path gd.String, state gdclass.GLTFState, flags gd.Int, base_path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(state[0].AsPointer())[0])
	callframe.Arg(frame, flags)
	callframe.Arg(frame, mmm.Get(base_path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_append_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Takes a [PackedByteArray] defining a GLTF and imports the data to the given [GLTFState] object through the [param state] parameter.
[b]Note:[/b] The [param base_path] tells [method append_from_buffer] where to find dependencies and can be empty.
*/
//go:nosplit
func (self class) AppendFromBuffer(bytes gd.PackedByteArray, base_path gd.String, state gdclass.GLTFState, flags gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bytes))
	callframe.Arg(frame, mmm.Get(base_path))
	callframe.Arg(frame, mmm.Get(state[0].AsPointer())[0])
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_append_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Takes a Godot Engine scene node and exports it and its descendants to the given [GLTFState] object through the [param state] parameter.
*/
//go:nosplit
func (self class) AppendFromScene(node gdclass.Node, state gdclass.GLTFState, flags gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(state[0].AsPointer())[0])
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_append_from_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Takes a [GLTFState] object through the [param state] parameter and returns a Godot Engine scene node.
The [param bake_fps] parameter overrides the bake_fps in [param state].
*/
//go:nosplit
func (self class) GenerateScene(ctx gd.Lifetime, state gdclass.GLTFState, bake_fps gd.Float, trimming bool, remove_immutable_tracks bool) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(state[0].AsPointer())[0])
	callframe.Arg(frame, bake_fps)
	callframe.Arg(frame, trimming)
	callframe.Arg(frame, remove_immutable_tracks)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_generate_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Takes a [GLTFState] object through the [param state] parameter and returns a GLTF [PackedByteArray].
*/
//go:nosplit
func (self class) GenerateBuffer(ctx gd.Lifetime, state gdclass.GLTFState) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(state[0].AsPointer())[0])
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_generate_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Takes a [GLTFState] object through the [param state] parameter and writes a glTF file to the filesystem.
[b]Note:[/b] The extension of the glTF file determines if it is a .glb binary file or a .gltf text file.
*/
//go:nosplit
func (self class) WriteToFilesystem(state gdclass.GLTFState, path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(state[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFDocument.Bind_write_to_filesystem, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Registers the given [GLTFDocumentExtension] instance with GLTFDocument. If [param first_priority] is true, this extension will be run first. Otherwise, it will be run last.
[b]Note:[/b] Like GLTFDocument itself, all GLTFDocumentExtension classes must be stateless in order to function properly. If you need to store data, use the [code]set_additional_data[/code] and [code]get_additional_data[/code] methods in [GLTFState] or [GLTFNode].
*/
//go:nosplit
func (self class) RegisterGltfDocumentExtension(ctx gd.Lifetime, extension gdclass.GLTFDocumentExtension, first_priority bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension[0].AsPointer())[0])
	callframe.Arg(frame, first_priority)
	var r_ret callframe.Nil
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFDocument.Bind_register_gltf_document_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters the given [GLTFDocumentExtension] instance.
*/
//go:nosplit
func (self class) UnregisterGltfDocumentExtension(ctx gd.Lifetime, extension gdclass.GLTFDocumentExtension)  {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension[0].AsPointer())[0])
	var r_ret callframe.Nil
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFDocument.Bind_unregister_gltf_document_extension, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFDocument() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFDocument() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFDocument", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type RootNodeMode = classdb.GLTFDocumentRootNodeMode

const (
/*Treat the Godot scene's root node as the root node of the glTF file, and mark it as the single root node via the [code]GODOT_single_root[/code] glTF extension. This will be parsed the same as [constant ROOT_NODE_MODE_KEEP_ROOT] if the implementation does not support [code]GODOT_single_root[/code].*/
	RootNodeModeSingleRoot RootNodeMode = 0
/*Treat the Godot scene's root node as the root node of the glTF file, but do not mark it as anything special. An extra root node will be generated when importing into Godot. This uses only vanilla glTF features. This is equivalent to the behavior in Godot 4.1 and earlier.*/
	RootNodeModeKeepRoot RootNodeMode = 1
/*Treat the Godot scene's root node as the name of the glTF scene, and add all of its children as root nodes of the glTF file. This uses only vanilla glTF features. This avoids an extra root node, but only the name of the Godot scene's root node will be preserved, as it will not be saved as a node.*/
	RootNodeModeMultiRoot RootNodeMode = 2
)