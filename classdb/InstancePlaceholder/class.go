// Code generated by the generate package DO NOT EDIT

// Package InstancePlaceholder provides methods for working with InstancePlaceholder object instances.
package InstancePlaceholder

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
import "graphics.gd/classdb/PackedScene"
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
Turning on the option [b]Load As Placeholder[/b] for an instantiated scene in the editor causes it to be replaced by an [InstancePlaceholder] when running the game, this will not replace the node in the editor. This makes it possible to delay actually loading the scene until calling [method create_instance]. This is useful to avoid loading large scenes all at once by loading parts of it selectively.
The [InstancePlaceholder] does not have a transform. This causes any child nodes to be positioned relatively to the [Viewport] from point (0,0), rather than their parent as displayed in the editor. Replacing the placeholder with a scene with a transform will transform children relatively to their parent again.
*/
type Instance [1]gdclass.InstancePlaceholder

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.InstancePlaceholder

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInstancePlaceholder() Instance
}

/*
Returns the list of properties that will be applied to the node when [method create_instance] is called.
If [param with_order] is [code]true[/code], a key named [code].order[/code] (note the leading period) is added to the dictionary. This [code].order[/code] key is an [Array] of [String] property names specifying the order in which properties will be applied (with index 0 being the first).
*/
func (self Instance) GetStoredValues() map[string]interface{} { //gd:InstancePlaceholder.get_stored_values
	return map[string]interface{}(gd.DictionaryAs[map[string]interface{}](Advanced(self).GetStoredValues(false)))
}

/*
Returns the list of properties that will be applied to the node when [method create_instance] is called.
If [param with_order] is [code]true[/code], a key named [code].order[/code] (note the leading period) is added to the dictionary. This [code].order[/code] key is an [Array] of [String] property names specifying the order in which properties will be applied (with index 0 being the first).
*/
func (self Expanded) GetStoredValues(with_order bool) map[string]interface{} { //gd:InstancePlaceholder.get_stored_values
	return map[string]interface{}(gd.DictionaryAs[map[string]interface{}](Advanced(self).GetStoredValues(with_order)))
}

/*
Call this method to actually load in the node. The created node will be placed as a sibling [i]above[/i] the [InstancePlaceholder] in the scene tree. The [Node]'s reference is also returned for convenience.
[b]Note:[/b] [method create_instance] is not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
func (self Instance) CreateInstance() Node.Instance { //gd:InstancePlaceholder.create_instance
	return Node.Instance(Advanced(self).CreateInstance(false, [1]PackedScene.Instance{}[0]))
}

/*
Call this method to actually load in the node. The created node will be placed as a sibling [i]above[/i] the [InstancePlaceholder] in the scene tree. The [Node]'s reference is also returned for convenience.
[b]Note:[/b] [method create_instance] is not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
func (self Expanded) CreateInstance(replace bool, custom_scene PackedScene.Instance) Node.Instance { //gd:InstancePlaceholder.create_instance
	return Node.Instance(Advanced(self).CreateInstance(replace, custom_scene))
}

/*
Gets the path to the [PackedScene] resource file that is loaded by default when calling [method create_instance]. Not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
func (self Instance) GetInstancePath() string { //gd:InstancePlaceholder.get_instance_path
	return string(Advanced(self).GetInstancePath().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InstancePlaceholder

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InstancePlaceholder"))
	casted := Instance{*(*gdclass.InstancePlaceholder)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns the list of properties that will be applied to the node when [method create_instance] is called.
If [param with_order] is [code]true[/code], a key named [code].order[/code] (note the leading period) is added to the dictionary. This [code].order[/code] key is an [Array] of [String] property names specifying the order in which properties will be applied (with index 0 being the first).
*/
//go:nosplit
func (self class) GetStoredValues(with_order bool) Dictionary.Any { //gd:InstancePlaceholder.get_stored_values
	var frame = callframe.New()
	callframe.Arg(frame, with_order)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InstancePlaceholder.Bind_get_stored_values, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Call this method to actually load in the node. The created node will be placed as a sibling [i]above[/i] the [InstancePlaceholder] in the scene tree. The [Node]'s reference is also returned for convenience.
[b]Note:[/b] [method create_instance] is not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
//go:nosplit
func (self class) CreateInstance(replace bool, custom_scene [1]gdclass.PackedScene) [1]gdclass.Node { //gd:InstancePlaceholder.create_instance
	var frame = callframe.New()
	callframe.Arg(frame, replace)
	callframe.Arg(frame, pointers.Get(custom_scene[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InstancePlaceholder.Bind_create_instance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerWithOwnershipTransferredToGo[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Gets the path to the [PackedScene] resource file that is loaded by default when calling [method create_instance]. Not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
//go:nosplit
func (self class) GetInstancePath() String.Readable { //gd:InstancePlaceholder.get_instance_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InstancePlaceholder.Bind_get_instance_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsInstancePlaceholder() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInstancePlaceholder() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsInstancePlaceholder() Instance {
	return self.Super().AsInstancePlaceholder()
}
func (self class) AsNode() Node.Advanced         { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance      { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("InstancePlaceholder", func(ptr gd.Object) any {
		return [1]gdclass.InstancePlaceholder{*(*gdclass.InstancePlaceholder)(unsafe.Pointer(&ptr))}
	})
}
