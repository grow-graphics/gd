// Code generated by the generate package DO NOT EDIT

// Package EditorResourcePreviewGenerator provides methods for working with EditorResourcePreviewGenerator object instances.
package EditorResourcePreviewGenerator

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
import "graphics.gd/classdb/Resource"
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
import "graphics.gd/variant/Vector2i"

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
Custom code to generate previews. Check [member EditorSettings.filesystem/file_dialog/thumbnail_size] to find a proper size to generate previews at.

	See [Interface] for methods that can be overridden by a [Class] that extends it.
*/
type Instance [1]gdclass.EditorResourcePreviewGenerator

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorResourcePreviewGenerator() Instance
}
type Interface interface {
	//Returns [code]true[/code] if your generator supports the resource of type [param type].
	Handles(atype string) bool
	//Generate a preview from a given resource with the specified size. This must always be implemented.
	//Returning [code]null[/code] is an OK way to fail and let another generator take care.
	//Care must be taken because this function is always called from a thread (not the main thread).
	//[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
	Generate(resource Resource.Instance, size Vector2i.XY, metadata map[any]any) Texture2D.Instance
	//Generate a preview directly from a path with the specified size. Implementing this is optional, as default code will load and call [method _generate].
	//Returning [code]null[/code] is an OK way to fail and let another generator take care.
	//Care must be taken because this function is always called from a thread (not the main thread).
	//[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
	GenerateFromPath(path string, size Vector2i.XY, metadata map[any]any) Texture2D.Instance
	//If this function returns [code]true[/code], the generator will automatically generate the small previews from the normal preview texture generated by the methods [method _generate] or [method _generate_from_path].
	//By default, it returns [code]false[/code].
	GenerateSmallPreviewAutomatically() bool
	//If this function returns [code]true[/code], the generator will call [method _generate] or [method _generate_from_path] for small previews as well.
	//By default, it returns [code]false[/code].
	CanGenerateSmallPreview() bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Handles(atype string) (_ bool) { return }
func (self implementation) Generate(resource Resource.Instance, size Vector2i.XY, metadata map[any]any) (_ Texture2D.Instance) {
	return
}
func (self implementation) GenerateFromPath(path string, size Vector2i.XY, metadata map[any]any) (_ Texture2D.Instance) {
	return
}
func (self implementation) GenerateSmallPreviewAutomatically() (_ bool) { return }
func (self implementation) CanGenerateSmallPreview() (_ bool)           { return }

/*
Returns [code]true[/code] if your generator supports the resource of type [param type].
*/
func (Instance) _handles(impl func(ptr unsafe.Pointer, atype string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var atype = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(atype))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Generate a preview from a given resource with the specified size. This must always be implemented.
Returning [code]null[/code] is an OK way to fail and let another generator take care.
Care must be taken because this function is always called from a thread (not the main thread).
[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
*/
func (Instance) _generate(impl func(ptr unsafe.Pointer, resource Resource.Instance, size Vector2i.XY, metadata map[any]any) Texture2D.Instance) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		var size = gd.UnsafeGet[Vector2i.XY](p_args, 1)
		var metadata = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(metadata))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, size, gd.DictionaryAs[map[any]any](metadata))
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Generate a preview directly from a path with the specified size. Implementing this is optional, as default code will load and call [method _generate].
Returning [code]null[/code] is an OK way to fail and let another generator take care.
Care must be taken because this function is always called from a thread (not the main thread).
[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
*/
func (Instance) _generate_from_path(impl func(ptr unsafe.Pointer, path string, size Vector2i.XY, metadata map[any]any) Texture2D.Instance) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var size = gd.UnsafeGet[Vector2i.XY](p_args, 1)
		var metadata = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(metadata))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), size, gd.DictionaryAs[map[any]any](metadata))
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
If this function returns [code]true[/code], the generator will automatically generate the small previews from the normal preview texture generated by the methods [method _generate] or [method _generate_from_path].
By default, it returns [code]false[/code].
*/
func (Instance) _generate_small_preview_automatically(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
If this function returns [code]true[/code], the generator will call [method _generate] or [method _generate_from_path] for small previews as well.
By default, it returns [code]false[/code].
*/
func (Instance) _can_generate_small_preview(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorResourcePreviewGenerator

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorResourcePreviewGenerator"))
	casted := Instance{*(*gdclass.EditorResourcePreviewGenerator)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns [code]true[/code] if your generator supports the resource of type [param type].
*/
func (class) _handles(impl func(ptr unsafe.Pointer, atype String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var atype = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(atype))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Generate a preview from a given resource with the specified size. This must always be implemented.
Returning [code]null[/code] is an OK way to fail and let another generator take care.
Care must be taken because this function is always called from a thread (not the main thread).
[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
*/
func (class) _generate(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource, size Vector2i.XY, metadata Dictionary.Any) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		var size = gd.UnsafeGet[Vector2i.XY](p_args, 1)
		var metadata = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(metadata))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, size, metadata)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Generate a preview directly from a path with the specified size. Implementing this is optional, as default code will load and call [method _generate].
Returning [code]null[/code] is an OK way to fail and let another generator take care.
Care must be taken because this function is always called from a thread (not the main thread).
[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
*/
func (class) _generate_from_path(impl func(ptr unsafe.Pointer, path String.Readable, size Vector2i.XY, metadata Dictionary.Any) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var size = gd.UnsafeGet[Vector2i.XY](p_args, 1)
		var metadata = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalDictionary(metadata))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, size, metadata)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
If this function returns [code]true[/code], the generator will automatically generate the small previews from the normal preview texture generated by the methods [method _generate] or [method _generate_from_path].
By default, it returns [code]false[/code].
*/
func (class) _generate_small_preview_automatically(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
If this function returns [code]true[/code], the generator will call [method _generate] or [method _generate_from_path] for small previews as well.
By default, it returns [code]false[/code].
*/
func (class) _can_generate_small_preview(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsEditorResourcePreviewGenerator() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorResourcePreviewGenerator() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsEditorResourcePreviewGenerator() Instance {
	return self.Super().AsEditorResourcePreviewGenerator()
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
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_generate":
		return reflect.ValueOf(self._generate)
	case "_generate_from_path":
		return reflect.ValueOf(self._generate_from_path)
	case "_generate_small_preview_automatically":
		return reflect.ValueOf(self._generate_small_preview_automatically)
	case "_can_generate_small_preview":
		return reflect.ValueOf(self._can_generate_small_preview)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_generate":
		return reflect.ValueOf(self._generate)
	case "_generate_from_path":
		return reflect.ValueOf(self._generate_from_path)
	case "_generate_small_preview_automatically":
		return reflect.ValueOf(self._generate_small_preview_automatically)
	case "_can_generate_small_preview":
		return reflect.ValueOf(self._can_generate_small_preview)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorResourcePreviewGenerator", func(ptr gd.Object) any {
		return [1]gdclass.EditorResourcePreviewGenerator{*(*gdclass.EditorResourcePreviewGenerator)(unsafe.Pointer(&ptr))}
	})
}
