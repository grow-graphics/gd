// Code generated by the generate package DO NOT EDIT

// Package RenderDataExtension provides methods for working with RenderDataExtension object instances.
package RenderDataExtension

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
import "graphics.gd/classdb/RenderData"
import "graphics.gd/classdb/RenderSceneBuffers"
import "graphics.gd/classdb/RenderSceneData"
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
This class allows for a RenderData implementation to be made in GDExtension.

	See [Interface] for methods that can be overridden by a [Class] that extends it.
*/
type Instance [1]gdclass.RenderDataExtension

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRenderDataExtension() Instance
}
type Interface interface {
	//Implement this in GDExtension to return the implementation's [RenderSceneBuffers] object.
	GetRenderSceneBuffers() RenderSceneBuffers.Instance
	//Implement this in GDExtension to return the implementation's [RenderSceneDataExtension] object.
	GetRenderSceneData() RenderSceneData.Instance
	//Implement this in GDExtension to return the [RID] of the implementation's environment object.
	GetEnvironment() RID.Any
	//Implement this in GDExtension to return the [RID] for the implementation's camera attributes object.
	GetCameraAttributes() RID.Any
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetRenderSceneBuffers() (_ RenderSceneBuffers.Instance) { return }
func (self implementation) GetRenderSceneData() (_ RenderSceneData.Instance)       { return }
func (self implementation) GetEnvironment() (_ RID.Any)                            { return }
func (self implementation) GetCameraAttributes() (_ RID.Any)                       { return }

/*
Implement this in GDExtension to return the implementation's [RenderSceneBuffers] object.
*/
func (Instance) _get_render_scene_buffers(impl func(ptr unsafe.Pointer) RenderSceneBuffers.Instance) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Implement this in GDExtension to return the implementation's [RenderSceneDataExtension] object.
*/
func (Instance) _get_render_scene_data(impl func(ptr unsafe.Pointer) RenderSceneData.Instance) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Implement this in GDExtension to return the [RID] of the implementation's environment object.
*/
func (Instance) _get_environment(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}

/*
Implement this in GDExtension to return the [RID] for the implementation's camera attributes object.
*/
func (Instance) _get_camera_attributes(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderDataExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderDataExtension"))
	casted := Instance{*(*gdclass.RenderDataExtension)(unsafe.Pointer(&object))}
	return casted
}

/*
Implement this in GDExtension to return the implementation's [RenderSceneBuffers] object.
*/
func (class) _get_render_scene_buffers(impl func(ptr unsafe.Pointer) [1]gdclass.RenderSceneBuffers) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Implement this in GDExtension to return the implementation's [RenderSceneDataExtension] object.
*/
func (class) _get_render_scene_data(impl func(ptr unsafe.Pointer) [1]gdclass.RenderSceneData) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Implement this in GDExtension to return the [RID] of the implementation's environment object.
*/
func (class) _get_environment(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the [RID] for the implementation's camera attributes object.
*/
func (class) _get_camera_attributes(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsRenderDataExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderDataExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsRenderDataExtension() Instance {
	return self.Super().AsRenderDataExtension()
}
func (self class) AsRenderData() RenderData.Advanced {
	return *((*RenderData.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsRenderData() RenderData.Instance { return self.Super().AsRenderData() }
func (self Instance) AsRenderData() RenderData.Instance {
	return *((*RenderData.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_render_scene_buffers":
		return reflect.ValueOf(self._get_render_scene_buffers)
	case "_get_render_scene_data":
		return reflect.ValueOf(self._get_render_scene_data)
	case "_get_environment":
		return reflect.ValueOf(self._get_environment)
	case "_get_camera_attributes":
		return reflect.ValueOf(self._get_camera_attributes)
	default:
		return gd.VirtualByName(RenderData.Advanced(self.AsRenderData()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_render_scene_buffers":
		return reflect.ValueOf(self._get_render_scene_buffers)
	case "_get_render_scene_data":
		return reflect.ValueOf(self._get_render_scene_data)
	case "_get_environment":
		return reflect.ValueOf(self._get_environment)
	case "_get_camera_attributes":
		return reflect.ValueOf(self._get_camera_attributes)
	default:
		return gd.VirtualByName(RenderData.Instance(self.AsRenderData()), name)
	}
}
func init() {
	gdclass.Register("RenderDataExtension", func(ptr gd.Object) any {
		return [1]gdclass.RenderDataExtension{*(*gdclass.RenderDataExtension)(unsafe.Pointer(&ptr))}
	})
}
