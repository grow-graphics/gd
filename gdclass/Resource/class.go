package Resource

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Resource is the base class for all Godot-specific resource types, serving primarily as data containers. Since they inherit from [RefCounted], resources are reference-counted and freed when no longer in use. They can also be nested within other resources, and saved on disk. [PackedScene], one of the most common [Object]s in a Godot project, is also a resource, uniquely capable of storing and instantiating the [Node]s it contains as many times as desired.
In GDScript, resources can loaded from disk by their [member resource_path] using [method @GDScript.load] or [method @GDScript.preload].
The engine keeps a global cache of all loaded resources, referenced by paths (see [method ResourceLoader.has_cached]). A resource will be cached when loaded for the first time and removed from cache once all references are released. When a resource is cached, subsequent loads using its path will return the cached reference.
[b]Note:[/b] In C#, resources will not be freed instantly after they are no longer in use. Instead, garbage collection will run periodically and will free resources that are no longer in use. This means that unused resources will remain in memory for a while before being removed.
	// Resource methods that can be overridden by a [Class] that extends it.
	type Resource interface {
		//Override this method to customize the newly duplicated resource created from [method PackedScene.instantiate], if the original's [member resource_local_to_scene] is set to [code]true[/code].
		//[b]Example:[/b] Set a random [code]damage[/code] value to every local resource from an instantiated scene.
		//[codeblock]
		//extends Resource
		//
		//var damage = 0
		//
		//func _setup_local_to_scene():
		//    damage = randi_range(10, 40)
		//[/codeblock]
		SetupLocalToScene() 
	}

*/
type Go [1]classdb.Resource

/*
Override this method to customize the newly duplicated resource created from [method PackedScene.instantiate], if the original's [member resource_local_to_scene] is set to [code]true[/code].
[b]Example:[/b] Set a random [code]damage[/code] value to every local resource from an instantiated scene.
[codeblock]
extends Resource

var damage = 0

func _setup_local_to_scene():
    damage = randi_range(10, 40)
[/codeblock]
*/
func (Go) _setup_local_to_scene(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Sets the [member resource_path] to [param path], potentially overriding an existing cache entry for this path. Further attempts to load an overridden resource by path will instead return this resource.
*/
func (self Go) TakeOverPath(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TakeOverPath(gc.String(path))
}

/*
Returns the [RID] of this resource (or an empty RID). Many resources (such as [Texture2D], [Mesh], and so on) are high-level abstractions of resources stored in a specialized server ([DisplayServer], [RenderingServer], etc.), so this function will return the original [RID].
*/
func (self Go) GetRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRid())
}

/*
If [member resource_local_to_scene] is set to [code]true[/code] and the resource has been loaded from a [PackedScene] instantiation, returns the root [Node] of the scene where this resource is used. Otherwise, returns [code]null[/code].
*/
func (self Go) GetLocalScene() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node(class(self).GetLocalScene(gc))
}

/*
Calls [method _setup_local_to_scene]. If [member resource_local_to_scene] is set to [code]true[/code], this method is automatically called from [method PackedScene.instantiate] by the newly duplicated resource within the scene instance.
*/
func (self Go) SetupLocalToScene() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetupLocalToScene()
}

/*
Generates a unique identifier for a resource to be contained inside a [PackedScene], based on the current date, time, and a random value. The returned string is only composed of letters ([code]a[/code] to [code]y[/code]) and numbers ([code]0[/code] to [code]8[/code]). See also [member resource_scene_unique_id].
*/
func (self Go) GenerateSceneUniqueId() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GenerateSceneUniqueId(gc).String())
}

/*
Emits the [signal changed] signal. This method is called automatically for some built-in resources.
[b]Note:[/b] For custom resources, it's recommended to call this method whenever a meaningful change occurs, such as a modified property. This ensures that custom [Object]s depending on the resource are properly updated.
[codeblock]
var damage:
    set(new_value):
        if damage != new_value:
            damage = new_value
            emit_changed()
[/codeblock]
*/
func (self Go) EmitChanged() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EmitChanged()
}

/*
Duplicates this resource, returning a new resource with its [code]export[/code]ed or [constant PROPERTY_USAGE_STORAGE] properties copied from the original.
If [param subresources] is [code]false[/code], a shallow copy is returned; nested resources within subresources are not duplicated and are shared with the original resource (with one exception; see below). If [param subresources] is [code]true[/code], a deep copy is returned; nested subresources will be duplicated and are not shared (with two exceptions; see below).
[param subresources] is usually respected, with the following exceptions:
- Subresource properties with the [constant PROPERTY_USAGE_ALWAYS_DUPLICATE] flag are always duplicated.
- Subresource properties with the [constant PROPERTY_USAGE_NEVER_DUPLICATE] flag are never duplicated.
- Subresources inside [Array] and [Dictionary] properties are never duplicated.
[b]Note:[/b] For custom resources, this method will fail if [method Object._init] has been defined with required parameters.
*/
func (self Go) Duplicate() gdclass.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Resource(class(self).Duplicate(gc, false))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Resource
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Resource"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ResourceLocalToScene() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsLocalToScene())
}

func (self Go) SetResourceLocalToScene(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLocalToScene(value)
}

func (self Go) ResourcePath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetPath(gc).String())
}

func (self Go) SetResourcePath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPath(gc.String(value))
}

func (self Go) ResourceName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetName(gc).String())
}

func (self Go) SetResourceName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetName(gc.String(value))
}

func (self Go) ResourceSceneUniqueId() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSceneUniqueId(gc).String())
}

func (self Go) SetResourceSceneUniqueId(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSceneUniqueId(gc.String(value))
}

/*
Override this method to customize the newly duplicated resource created from [method PackedScene.instantiate], if the original's [member resource_local_to_scene] is set to [code]true[/code].
[b]Example:[/b] Set a random [code]damage[/code] value to every local resource from an instantiated scene.
[codeblock]
extends Resource

var damage = 0

func _setup_local_to_scene():
    damage = randi_range(10, 40)
[/codeblock]
*/
func (class) _setup_local_to_scene(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetPath(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_set_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member resource_path] to [param path], potentially overriding an existing cache entry for this path. Further attempts to load an overridden resource by path will instead return this resource.
*/
//go:nosplit
func (self class) TakeOverPath(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_take_over_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetName(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_set_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [RID] of this resource (or an empty RID). Many resources (such as [Texture2D], [Mesh], and so on) are high-level abstractions of resources stored in a specialized server ([DisplayServer], [RenderingServer], etc.), so this function will return the original [RID].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLocalToScene(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_set_local_to_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLocalToScene() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_is_local_to_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [member resource_local_to_scene] is set to [code]true[/code] and the resource has been loaded from a [PackedScene] instantiation, returns the root [Node] of the scene where this resource is used. Otherwise, returns [code]null[/code].
*/
//go:nosplit
func (self class) GetLocalScene(ctx gd.Lifetime) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_get_local_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Calls [method _setup_local_to_scene]. If [member resource_local_to_scene] is set to [code]true[/code], this method is automatically called from [method PackedScene.instantiate] by the newly duplicated resource within the scene instance.
*/
//go:nosplit
func (self class) SetupLocalToScene()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_setup_local_to_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates a unique identifier for a resource to be contained inside a [PackedScene], based on the current date, time, and a random value. The returned string is only composed of letters ([code]a[/code] to [code]y[/code]) and numbers ([code]0[/code] to [code]8[/code]). See also [member resource_scene_unique_id].
*/
//go:nosplit
func (self class) GenerateSceneUniqueId(ctx gd.Lifetime) gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.Resource.Bind_generate_scene_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSceneUniqueId(id gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(id))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_set_scene_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSceneUniqueId(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_get_scene_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Emits the [signal changed] signal. This method is called automatically for some built-in resources.
[b]Note:[/b] For custom resources, it's recommended to call this method whenever a meaningful change occurs, such as a modified property. This ensures that custom [Object]s depending on the resource are properly updated.
[codeblock]
var damage:
    set(new_value):
        if damage != new_value:
            damage = new_value
            emit_changed()
[/codeblock]
*/
//go:nosplit
func (self class) EmitChanged()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_emit_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Duplicates this resource, returning a new resource with its [code]export[/code]ed or [constant PROPERTY_USAGE_STORAGE] properties copied from the original.
If [param subresources] is [code]false[/code], a shallow copy is returned; nested resources within subresources are not duplicated and are shared with the original resource (with one exception; see below). If [param subresources] is [code]true[/code], a deep copy is returned; nested subresources will be duplicated and are not shared (with two exceptions; see below).
[param subresources] is usually respected, with the following exceptions:
- Subresource properties with the [constant PROPERTY_USAGE_ALWAYS_DUPLICATE] flag are always duplicated.
- Subresource properties with the [constant PROPERTY_USAGE_NEVER_DUPLICATE] flag are never duplicated.
- Subresources inside [Array] and [Dictionary] properties are never duplicated.
[b]Note:[/b] For custom resources, this method will fail if [method Object._init] has been defined with required parameters.
*/
//go:nosplit
func (self class) Duplicate(ctx gd.Lifetime, subresources bool) gdclass.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, subresources)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Resource.Bind_duplicate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self Go) OnChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("changed"), gc.Callable(cb), 0)
}


func (self Go) OnSetupLocalToSceneRequested(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("setup_local_to_scene_requested"), gc.Callable(cb), 0)
}


func (self class) AsResource() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_setup_local_to_scene": return reflect.ValueOf(self._setup_local_to_scene);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_setup_local_to_scene": return reflect.ValueOf(self._setup_local_to_scene);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("Resource", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}