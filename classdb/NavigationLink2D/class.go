// Code generated by the generate package DO NOT EDIT

// Package NavigationLink2D provides methods for working with NavigationLink2D object instances.
package NavigationLink2D

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
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node2D"
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
import "graphics.gd/variant/Vector2"

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
A link between two positions on [NavigationRegion2D]s that agents can be routed through. These positions can be on the same [NavigationRegion2D] or on two different ones. Links are useful to express navigation methods other than traveling along the surface of the navigation polygon, such as ziplines, teleporters, or gaps that can be jumped across.
*/
type Instance [1]gdclass.NavigationLink2D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNavigationLink2D() Instance
}

/*
Returns the [RID] of this link on the [NavigationServer2D].
*/
func (self Instance) GetRid() RID.NavigationLink2D { //gd:NavigationLink2D.get_rid
	return RID.NavigationLink2D(Advanced(self).GetRid())
}

/*
Sets the [RID] of the navigation map this link should use. By default the link will automatically join the [World2D] default navigation map so this function is only required to override the default map.
*/
func (self Instance) SetNavigationMap(navigation_map RID.NavigationMap2D) { //gd:NavigationLink2D.set_navigation_map
	Advanced(self).SetNavigationMap(RID.Any(navigation_map))
}

/*
Returns the current navigation map [RID] used by this link.
*/
func (self Instance) GetNavigationMap() RID.NavigationMap2D { //gd:NavigationLink2D.get_navigation_map
	return RID.NavigationMap2D(Advanced(self).GetNavigationMap())
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetNavigationLayerValue(layer_number int, value bool) { //gd:NavigationLink2D.set_navigation_layer_value
	Advanced(self).SetNavigationLayerValue(int64(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetNavigationLayerValue(layer_number int) bool { //gd:NavigationLink2D.get_navigation_layer_value
	return bool(Advanced(self).GetNavigationLayerValue(int64(layer_number)))
}

/*
Sets the [member start_position] that is relative to the link from a global [param position].
*/
func (self Instance) SetGlobalStartPosition(position Vector2.XY) { //gd:NavigationLink2D.set_global_start_position
	Advanced(self).SetGlobalStartPosition(Vector2.XY(position))
}

/*
Returns the [member start_position] that is relative to the link as a global position.
*/
func (self Instance) GetGlobalStartPosition() Vector2.XY { //gd:NavigationLink2D.get_global_start_position
	return Vector2.XY(Advanced(self).GetGlobalStartPosition())
}

/*
Sets the [member end_position] that is relative to the link from a global [param position].
*/
func (self Instance) SetGlobalEndPosition(position Vector2.XY) { //gd:NavigationLink2D.set_global_end_position
	Advanced(self).SetGlobalEndPosition(Vector2.XY(position))
}

/*
Returns the [member end_position] that is relative to the link as a global position.
*/
func (self Instance) GetGlobalEndPosition() Vector2.XY { //gd:NavigationLink2D.get_global_end_position
	return Vector2.XY(Advanced(self).GetGlobalEndPosition())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationLink2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationLink2D"))
	casted := Instance{*(*gdclass.NavigationLink2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Enabled() bool {
	return bool(class(self).IsEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) Bidirectional() bool {
	return bool(class(self).IsBidirectional())
}

func (self Instance) SetBidirectional(value bool) {
	class(self).SetBidirectional(value)
}

func (self Instance) NavigationLayers() int {
	return int(int(class(self).GetNavigationLayers()))
}

func (self Instance) SetNavigationLayers(value int) {
	class(self).SetNavigationLayers(int64(value))
}

func (self Instance) StartPosition() Vector2.XY {
	return Vector2.XY(class(self).GetStartPosition())
}

func (self Instance) SetStartPosition(value Vector2.XY) {
	class(self).SetStartPosition(Vector2.XY(value))
}

func (self Instance) EndPosition() Vector2.XY {
	return Vector2.XY(class(self).GetEndPosition())
}

func (self Instance) SetEndPosition(value Vector2.XY) {
	class(self).SetEndPosition(Vector2.XY(value))
}

func (self Instance) EnterCost() Float.X {
	return Float.X(Float.X(class(self).GetEnterCost()))
}

func (self Instance) SetEnterCost(value Float.X) {
	class(self).SetEnterCost(float64(value))
}

func (self Instance) TravelCost() Float.X {
	return Float.X(Float.X(class(self).GetTravelCost()))
}

func (self Instance) SetTravelCost(value Float.X) {
	class(self).SetTravelCost(float64(value))
}

/*
Returns the [RID] of this link on the [NavigationServer2D].
*/
//go:nosplit
func (self class) GetRid() RID.Any { //gd:NavigationLink2D.get_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnabled(enabled bool) { //gd:NavigationLink2D.set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool { //gd:NavigationLink2D.is_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [RID] of the navigation map this link should use. By default the link will automatically join the [World2D] default navigation map so this function is only required to override the default map.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map RID.Any) { //gd:NavigationLink2D.set_navigation_map
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current navigation map [RID] used by this link.
*/
//go:nosplit
func (self class) GetNavigationMap() RID.Any { //gd:NavigationLink2D.get_navigation_map
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBidirectional(bidirectional bool) { //gd:NavigationLink2D.set_bidirectional
	var frame = callframe.New()
	callframe.Arg(frame, bidirectional)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_bidirectional, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsBidirectional() bool { //gd:NavigationLink2D.is_bidirectional
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_is_bidirectional, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationLayers(navigation_layers int64) { //gd:NavigationLink2D.set_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationLayers() int64 { //gd:NavigationLink2D.get_navigation_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number int64, value bool) { //gd:NavigationLink2D.set_navigation_layer_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerValue(layer_number int64) bool { //gd:NavigationLink2D.get_navigation_layer_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStartPosition(position Vector2.XY) { //gd:NavigationLink2D.set_start_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStartPosition() Vector2.XY { //gd:NavigationLink2D.get_start_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEndPosition(position Vector2.XY) { //gd:NavigationLink2D.set_end_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEndPosition() Vector2.XY { //gd:NavigationLink2D.get_end_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [member start_position] that is relative to the link from a global [param position].
*/
//go:nosplit
func (self class) SetGlobalStartPosition(position Vector2.XY) { //gd:NavigationLink2D.set_global_start_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_global_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [member start_position] that is relative to the link as a global position.
*/
//go:nosplit
func (self class) GetGlobalStartPosition() Vector2.XY { //gd:NavigationLink2D.get_global_start_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_global_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [member end_position] that is relative to the link from a global [param position].
*/
//go:nosplit
func (self class) SetGlobalEndPosition(position Vector2.XY) { //gd:NavigationLink2D.set_global_end_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_global_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [member end_position] that is relative to the link as a global position.
*/
//go:nosplit
func (self class) GetGlobalEndPosition() Vector2.XY { //gd:NavigationLink2D.get_global_end_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_global_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnterCost(enter_cost float64) { //gd:NavigationLink2D.set_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, enter_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnterCost() float64 { //gd:NavigationLink2D.get_enter_cost
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTravelCost(travel_cost float64) { //gd:NavigationLink2D.set_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, travel_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTravelCost() float64 { //gd:NavigationLink2D.get_travel_cost
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationLink2D() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNavigationLink2D() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNavigationLink2D() Instance { return self.Super().AsNavigationLink2D() }
func (self class) AsNode2D() Node2D.Advanced            { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode2D() Node2D.Instance    { return self.Super().AsNode2D() }
func (self Instance) AsNode2D() Node2D.Instance         { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("NavigationLink2D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationLink2D{*(*gdclass.NavigationLink2D)(unsafe.Pointer(&ptr))}
	})
}
