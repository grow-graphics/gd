// Code generated by the generate package DO NOT EDIT

// Package GDExtensionManager provides methods for working with GDExtensionManager object instances.
package GDExtensionManager

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Angle"
import "graphics.gd/variant/Euler"
import "graphics.gd/classdb/GDExtension"
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
The GDExtensionManager loads, initializes, and keeps track of all available [GDExtension] libraries in the project.
[b]Note:[/b] Do not worry about GDExtension unless you know what you are doing.
*/
type Instance [1]gdclass.GDExtensionManager

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

var self [1]gdclass.GDExtensionManager
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.GDExtensionManager)
	self = *(*[1]gdclass.GDExtensionManager)(unsafe.Pointer(&obj))
}

/*
Loads an extension by absolute file path. The [param path] needs to point to a valid [GDExtension]. Returns [constant LOAD_STATUS_OK] if successful.
*/
func LoadExtension(path string) LoadStatus { //gd:GDExtensionManager.load_extension
	once.Do(singleton)
	return LoadStatus(Advanced().LoadExtension(String.New(path)))
}

/*
Reloads the extension at the given file path. The [param path] needs to point to a valid [GDExtension], otherwise this method may return either [constant LOAD_STATUS_NOT_LOADED] or [constant LOAD_STATUS_FAILED].
[b]Note:[/b] You can only reload extensions in the editor. In release builds, this method always fails and returns [constant LOAD_STATUS_FAILED].
*/
func ReloadExtension(path string) LoadStatus { //gd:GDExtensionManager.reload_extension
	once.Do(singleton)
	return LoadStatus(Advanced().ReloadExtension(String.New(path)))
}

/*
Unloads an extension by file path. The [param path] needs to point to an already loaded [GDExtension], otherwise this method returns [constant LOAD_STATUS_NOT_LOADED].
*/
func UnloadExtension(path string) LoadStatus { //gd:GDExtensionManager.unload_extension
	once.Do(singleton)
	return LoadStatus(Advanced().UnloadExtension(String.New(path)))
}

/*
Returns [code]true[/code] if the extension at the given file [param path] has already been loaded successfully. See also [method get_loaded_extensions].
*/
func IsExtensionLoaded(path string) bool { //gd:GDExtensionManager.is_extension_loaded
	once.Do(singleton)
	return bool(Advanced().IsExtensionLoaded(String.New(path)))
}

/*
Returns the file paths of all currently loaded extensions.
*/
func GetLoadedExtensions() []string { //gd:GDExtensionManager.get_loaded_extensions
	once.Do(singleton)
	return []string(Advanced().GetLoadedExtensions().Strings())
}

/*
Returns the [GDExtension] at the given file [param path], or [code]null[/code] if it has not been loaded or does not exist.
*/
func GetExtension(path string) GDExtension.Instance { //gd:GDExtensionManager.get_extension
	once.Do(singleton)
	return GDExtension.Instance(Advanced().GetExtension(String.New(path)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.GDExtensionManager

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }

/*
Loads an extension by absolute file path. The [param path] needs to point to a valid [GDExtension]. Returns [constant LOAD_STATUS_OK] if successful.
*/
//go:nosplit
func (self class) LoadExtension(path String.Readable) LoadStatus { //gd:GDExtensionManager.load_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[LoadStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_load_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Reloads the extension at the given file path. The [param path] needs to point to a valid [GDExtension], otherwise this method may return either [constant LOAD_STATUS_NOT_LOADED] or [constant LOAD_STATUS_FAILED].
[b]Note:[/b] You can only reload extensions in the editor. In release builds, this method always fails and returns [constant LOAD_STATUS_FAILED].
*/
//go:nosplit
func (self class) ReloadExtension(path String.Readable) LoadStatus { //gd:GDExtensionManager.reload_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[LoadStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_reload_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Unloads an extension by file path. The [param path] needs to point to an already loaded [GDExtension], otherwise this method returns [constant LOAD_STATUS_NOT_LOADED].
*/
//go:nosplit
func (self class) UnloadExtension(path String.Readable) LoadStatus { //gd:GDExtensionManager.unload_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[LoadStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_unload_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the extension at the given file [param path] has already been loaded successfully. See also [method get_loaded_extensions].
*/
//go:nosplit
func (self class) IsExtensionLoaded(path String.Readable) bool { //gd:GDExtensionManager.is_extension_loaded
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_is_extension_loaded, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the file paths of all currently loaded extensions.
*/
//go:nosplit
func (self class) GetLoadedExtensions() Packed.Strings { //gd:GDExtensionManager.get_loaded_extensions
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_get_loaded_extensions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the [GDExtension] at the given file [param path], or [code]null[/code] if it has not been loaded or does not exist.
*/
//go:nosplit
func (self class) GetExtension(path String.Readable) [1]gdclass.GDExtension { //gd:GDExtensionManager.get_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtensionManager.Bind_get_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GDExtension{gd.PointerWithOwnershipTransferredToGo[gdclass.GDExtension](r_ret.Get())}
	frame.Free()
	return ret
}
func OnExtensionsReloaded(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("extensions_reloaded"), gd.NewCallable(cb), 0)
}

func OnExtensionLoaded(cb func(extension GDExtension.Instance)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("extension_loaded"), gd.NewCallable(cb), 0)
}

func OnExtensionUnloading(cb func(extension GDExtension.Instance)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("extension_unloading"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("GDExtensionManager", func(ptr gd.Object) any {
		return [1]gdclass.GDExtensionManager{*(*gdclass.GDExtensionManager)(unsafe.Pointer(&ptr))}
	})
}

type LoadStatus int //gd:GDExtensionManager.LoadStatus

const (
	/*The extension has loaded successfully.*/
	LoadStatusOk LoadStatus = 0
	/*The extension has failed to load, possibly because it does not exist or has missing dependencies.*/
	LoadStatusFailed LoadStatus = 1
	/*The extension has already been loaded.*/
	LoadStatusAlreadyLoaded LoadStatus = 2
	/*The extension has not been loaded.*/
	LoadStatusNotLoaded LoadStatus = 3
	/*The extension requires the application to restart to fully load.*/
	LoadStatusNeedsRestart LoadStatus = 4
)
