package NavigationRegion2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A traversable 2D region based on a [NavigationPolygon] that [NavigationAgent2D]s can use for pathfinding.
Two regions can be connected to each other if they share a similar edge. You can set the minimum distance between two vertices required to connect two edges by using [method NavigationServer2D.map_set_edge_connection_margin].
[b]Note:[/b] Overlapping two regions' navigation polygons is not enough for connecting two regions. They must share a similar edge.
The pathfinding cost of entering a region from another region can be controlled with the [member enter_cost] value.
[b]Note:[/b] This value is not added to the path cost when the start position is already inside this region.
The pathfinding cost of traveling distances inside this region can be controlled with the [member travel_cost] multiplier.
[b]Note:[/b] This node caches changes to its properties, so if you make changes to the underlying region [RID] in [NavigationServer2D], they will not be reflected in this node's properties.

*/
type Go [1]classdb.NavigationRegion2D

/*
Returns the [RID] of this region on the [NavigationServer2D]. Combined with [method NavigationServer2D.map_get_closest_point_owner] can be used to identify the [NavigationRegion2D] closest to a point on the merged navigation map.
*/
func (self Go) GetRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRid())
}

/*
Sets the [RID] of the navigation map this region should use. By default the region will automatically join the [World2D] default navigation map so this function is only required to override the default map.
*/
func (self Go) SetNavigationMap(navigation_map gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationMap(navigation_map)
}

/*
Returns the current navigation map [RID] used by this region.
*/
func (self Go) GetNavigationMap() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetNavigationMap())
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Go) SetNavigationLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetNavigationLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetNavigationLayerValue(gd.Int(layer_number)))
}

/*
Returns the [RID] of this region on the [NavigationServer2D].
*/
func (self Go) GetRegionRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRegionRid())
}

/*
Bakes the [NavigationPolygon]. If [param on_thread] is set to [code]true[/code] (default), the baking is done on a separate thread.
*/
func (self Go) BakeNavigationPolygon() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).BakeNavigationPolygon(true)
}

/*
Returns [code]true[/code] when the [NavigationPolygon] is being baked on a background thread.
*/
func (self Go) IsBaking() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsBaking())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NavigationRegion2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("NavigationRegion2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) NavigationPolygon() gdclass.NavigationPolygon {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.NavigationPolygon(class(self).GetNavigationPolygon(gc))
}

func (self Go) SetNavigationPolygon(value gdclass.NavigationPolygon) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationPolygon(value)
}

func (self Go) Enabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEnabled())
}

func (self Go) SetEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnabled(value)
}

func (self Go) UseEdgeConnections() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUseEdgeConnections())
}

func (self Go) SetUseEdgeConnections(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseEdgeConnections(value)
}

func (self Go) NavigationLayers() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetNavigationLayers()))
}

func (self Go) SetNavigationLayers(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Go) EnterCost() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEnterCost()))
}

func (self Go) SetEnterCost(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnterCost(gd.Float(value))
}

func (self Go) TravelCost() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTravelCost()))
}

func (self Go) SetTravelCost(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTravelCost(gd.Float(value))
}

/*
Returns the [RID] of this region on the [NavigationServer2D]. Combined with [method NavigationServer2D.map_get_closest_point_owner] can be used to identify the [NavigationRegion2D] closest to a point on the merged navigation map.
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationPolygon(navigation_polygon gdclass.NavigationPolygon)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_polygon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_set_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationPolygon(ctx gd.Lifetime) gdclass.NavigationPolygon {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.NavigationPolygon
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [RID] of the navigation map this region should use. By default the region will automatically join the [World2D] default navigation map so this function is only required to override the default map.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current navigation map [RID] used by this region.
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseEdgeConnections(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseEdgeConnections() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [RID] of this region on the [NavigationServer2D].
*/
//go:nosplit
func (self class) GetRegionRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_region_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnterCost(enter_cost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enter_cost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnterCost() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTravelCost(travel_cost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, travel_cost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTravelCost() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Bakes the [NavigationPolygon]. If [param on_thread] is set to [code]true[/code] (default), the baking is done on a separate thread.
*/
//go:nosplit
func (self class) BakeNavigationPolygon(on_thread bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, on_thread)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_bake_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] when the [NavigationPolygon] is being baked on a background thread.
*/
//go:nosplit
func (self class) IsBaking() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion2D.Bind_is_baking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnNavigationPolygonChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("navigation_polygon_changed"), gc.Callable(cb), 0)
}


func (self Go) OnBakeFinished(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("bake_finished"), gc.Callable(cb), 0)
}


func (self class) AsNavigationRegion2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNavigationRegion2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("NavigationRegion2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}