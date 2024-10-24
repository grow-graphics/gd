package AStarGrid2D

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
[AStarGrid2D] is a variant of [AStar2D] that is specialized for partial 2D grids. It is simpler to use because it doesn't require you to manually create points and connect them together. This class also supports multiple types of heuristics, modes for diagonal movement, and a jumping mode to speed up calculations.
To use [AStarGrid2D], you only need to set the [member region] of the grid, optionally set the [member cell_size], and then call the [method update] method:
[codeblocks]
[gdscript]
var astar_grid = AStarGrid2D.new()
astar_grid.region = Rect2i(0, 0, 32, 32)
astar_grid.cell_size = Vector2(16, 16)
astar_grid.update()
print(astar_grid.get_id_path(Vector2i(0, 0), Vector2i(3, 4))) # prints (0, 0), (1, 1), (2, 2), (3, 3), (3, 4)
print(astar_grid.get_point_path(Vector2i(0, 0), Vector2i(3, 4))) # prints (0, 0), (16, 16), (32, 32), (48, 48), (48, 64)
[/gdscript]
[csharp]
AStarGrid2D astarGrid = new AStarGrid2D();
astarGrid.Region = new Rect2I(0, 0, 32, 32);
astarGrid.CellSize = new Vector2I(16, 16);
astarGrid.Update();
GD.Print(astarGrid.GetIdPath(Vector2I.Zero, new Vector2I(3, 4))); // prints (0, 0), (1, 1), (2, 2), (3, 3), (3, 4)
GD.Print(astarGrid.GetPointPath(Vector2I.Zero, new Vector2I(3, 4))); // prints (0, 0), (16, 16), (32, 32), (48, 48), (48, 64)
[/csharp]
[/codeblocks]
To remove a point from the pathfinding grid, it must be set as "solid" with [method set_point_solid].
	// AStarGrid2D methods that can be overridden by a [Class] that extends it.
	type AStarGrid2D interface {
		//Called when estimating the cost between a point and the path's ending point.
		//Note that this function is hidden in the default [AStarGrid2D] class.
		EstimateCost(from_id gd.Vector2i, to_id gd.Vector2i) float64
		//Called when computing the cost between two connected points.
		//Note that this function is hidden in the default [AStarGrid2D] class.
		ComputeCost(from_id gd.Vector2i, to_id gd.Vector2i) float64
	}

*/
type Go [1]classdb.AStarGrid2D

/*
Called when estimating the cost between a point and the path's ending point.
Note that this function is hidden in the default [AStarGrid2D] class.
*/
func (Go) _estimate_cost(impl func(ptr unsafe.Pointer, from_id gd.Vector2i, to_id gd.Vector2i) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var from_id = gd.UnsafeGet[gd.Vector2i](p_args,0)
		var to_id = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_id, to_id)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}

/*
Called when computing the cost between two connected points.
Note that this function is hidden in the default [AStarGrid2D] class.
*/
func (Go) _compute_cost(impl func(ptr unsafe.Pointer, from_id gd.Vector2i, to_id gd.Vector2i) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var from_id = gd.UnsafeGet[gd.Vector2i](p_args,0)
		var to_id = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_id, to_id)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}

/*
Returns [code]true[/code] if the [param x] and [param y] is a valid grid coordinate (id), i.e. if it is inside [member region]. Equivalent to [code]region.has_point(Vector2i(x, y))[/code].
*/
func (self Go) IsInBounds(x int, y int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsInBounds(gd.Int(x), gd.Int(y)))
}

/*
Returns [code]true[/code] if the [param id] vector is a valid grid coordinate, i.e. if it is inside [member region]. Equivalent to [code]region.has_point(id)[/code].
*/
func (self Go) IsInBoundsv(id gd.Vector2i) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsInBoundsv(id))
}

/*
Indicates that the grid parameters were changed and [method update] needs to be called.
*/
func (self Go) IsDirty() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsDirty())
}

/*
Updates the internal state of the grid according to the parameters to prepare it to search the path. Needs to be called if parameters like [member region], [member cell_size] or [member offset] are changed. [method is_dirty] will return [code]true[/code] if this is the case and this needs to be called.
[b]Note:[/b] All point data (solidity and weight scale) will be cleared.
*/
func (self Go) Update() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Update()
}

/*
Disables or enables the specified point for pathfinding. Useful for making an obstacle. By default, all points are enabled.
[b]Note:[/b] Calling [method update] is not needed after the call of this function.
*/
func (self Go) SetPointSolid(id gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPointSolid(id, true)
}

/*
Returns [code]true[/code] if a point is disabled for pathfinding. By default, all points are enabled.
*/
func (self Go) IsPointSolid(id gd.Vector2i) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPointSolid(id))
}

/*
Sets the [param weight_scale] for the point with the given [param id]. The [param weight_scale] is multiplied by the result of [method _compute_cost] when determining the overall cost of traveling across a segment from a neighboring point to this point.
[b]Note:[/b] Calling [method update] is not needed after the call of this function.
*/
func (self Go) SetPointWeightScale(id gd.Vector2i, weight_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPointWeightScale(id, gd.Float(weight_scale))
}

/*
Returns the weight scale of the point associated with the given [param id].
*/
func (self Go) GetPointWeightScale(id gd.Vector2i) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetPointWeightScale(id)))
}

/*
Fills the given [param region] on the grid with the specified value for the solid flag.
[b]Note:[/b] Calling [method update] is not needed after the call of this function.
*/
func (self Go) FillSolidRegion(region gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FillSolidRegion(region, true)
}

/*
Fills the given [param region] on the grid with the specified value for the weight scale.
[b]Note:[/b] Calling [method update] is not needed after the call of this function.
*/
func (self Go) FillWeightScaleRegion(region gd.Rect2i, weight_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FillWeightScaleRegion(region, gd.Float(weight_scale))
}

/*
Clears the grid and sets the [member region] to [code]Rect2i(0, 0, 0, 0)[/code].
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Returns the position of the point associated with the given [param id].
*/
func (self Go) GetPointPosition(id gd.Vector2i) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetPointPosition(id))
}

/*
Returns an array with the points that are in the path found by [AStarGrid2D] between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
[b]Note:[/b] This method is not thread-safe. If called from a [Thread], it will return an empty array and will print an error message.
*/
func (self Go) GetPointPath(from_id gd.Vector2i, to_id gd.Vector2i) []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return []gd.Vector2(class(self).GetPointPath(gc, from_id, to_id, false).AsSlice())
}

/*
Returns an array with the IDs of the points that form the path found by AStar2D between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
*/
func (self Go) GetIdPath(from_id gd.Vector2i, to_id gd.Vector2i) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](class(self).GetIdPath(gc, from_id, to_id, false))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AStarGrid2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AStarGrid2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Region() gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Rect2i(class(self).GetRegion())
}

func (self Go) SetRegion(value gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRegion(value)
}

func (self Go) Size() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2i(class(self).GetSize())
}

func (self Go) SetSize(value gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSize(value)
}

func (self Go) Offset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetOffset())
}

func (self Go) SetOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOffset(value)
}

func (self Go) CellSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetCellSize())
}

func (self Go) SetCellSize(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellSize(value)
}

func (self Go) CellShape() classdb.AStarGrid2DCellShape {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AStarGrid2DCellShape(class(self).GetCellShape())
}

func (self Go) SetCellShape(value classdb.AStarGrid2DCellShape) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellShape(value)
}

func (self Go) JumpingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsJumpingEnabled())
}

func (self Go) SetJumpingEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetJumpingEnabled(value)
}

func (self Go) DefaultComputeHeuristic() classdb.AStarGrid2DHeuristic {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AStarGrid2DHeuristic(class(self).GetDefaultComputeHeuristic())
}

func (self Go) SetDefaultComputeHeuristic(value classdb.AStarGrid2DHeuristic) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDefaultComputeHeuristic(value)
}

func (self Go) DefaultEstimateHeuristic() classdb.AStarGrid2DHeuristic {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AStarGrid2DHeuristic(class(self).GetDefaultEstimateHeuristic())
}

func (self Go) SetDefaultEstimateHeuristic(value classdb.AStarGrid2DHeuristic) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDefaultEstimateHeuristic(value)
}

func (self Go) DiagonalMode() classdb.AStarGrid2DDiagonalMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AStarGrid2DDiagonalMode(class(self).GetDiagonalMode())
}

func (self Go) SetDiagonalMode(value classdb.AStarGrid2DDiagonalMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDiagonalMode(value)
}

/*
Called when estimating the cost between a point and the path's ending point.
Note that this function is hidden in the default [AStarGrid2D] class.
*/
func (class) _estimate_cost(impl func(ptr unsafe.Pointer, from_id gd.Vector2i, to_id gd.Vector2i) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var from_id = gd.UnsafeGet[gd.Vector2i](p_args,0)
		var to_id = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_id, to_id)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called when computing the cost between two connected points.
Note that this function is hidden in the default [AStarGrid2D] class.
*/
func (class) _compute_cost(impl func(ptr unsafe.Pointer, from_id gd.Vector2i, to_id gd.Vector2i) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var from_id = gd.UnsafeGet[gd.Vector2i](p_args,0)
		var to_id = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_id, to_id)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetRegion(region gd.Rect2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRegion() gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSize(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellSize(cell_size gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cell_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellShape(cell_shape classdb.AStarGrid2DCellShape)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cell_shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_cell_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellShape() classdb.AStarGrid2DCellShape {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AStarGrid2DCellShape](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_cell_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [param x] and [param y] is a valid grid coordinate (id), i.e. if it is inside [member region]. Equivalent to [code]region.has_point(Vector2i(x, y))[/code].
*/
//go:nosplit
func (self class) IsInBounds(x gd.Int, y gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_is_in_bounds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [param id] vector is a valid grid coordinate, i.e. if it is inside [member region]. Equivalent to [code]region.has_point(id)[/code].
*/
//go:nosplit
func (self class) IsInBoundsv(id gd.Vector2i) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_is_in_boundsv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Indicates that the grid parameters were changed and [method update] needs to be called.
*/
//go:nosplit
func (self class) IsDirty() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_is_dirty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the internal state of the grid according to the parameters to prepare it to search the path. Needs to be called if parameters like [member region], [member cell_size] or [member offset] are changed. [method is_dirty] will return [code]true[/code] if this is the case and this needs to be called.
[b]Note:[/b] All point data (solidity and weight scale) will be cleared.
*/
//go:nosplit
func (self class) Update()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetJumpingEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_jumping_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsJumpingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_is_jumping_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiagonalMode(mode classdb.AStarGrid2DDiagonalMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_diagonal_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDiagonalMode() classdb.AStarGrid2DDiagonalMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AStarGrid2DDiagonalMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_diagonal_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultComputeHeuristic(heuristic classdb.AStarGrid2DHeuristic)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, heuristic)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_default_compute_heuristic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultComputeHeuristic() classdb.AStarGrid2DHeuristic {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AStarGrid2DHeuristic](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_default_compute_heuristic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultEstimateHeuristic(heuristic classdb.AStarGrid2DHeuristic)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, heuristic)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_default_estimate_heuristic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultEstimateHeuristic() classdb.AStarGrid2DHeuristic {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AStarGrid2DHeuristic](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_default_estimate_heuristic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Disables or enables the specified point for pathfinding. Useful for making an obstacle. By default, all points are enabled.
[b]Note:[/b] Calling [method update] is not needed after the call of this function.
*/
//go:nosplit
func (self class) SetPointSolid(id gd.Vector2i, solid bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, solid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_point_solid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a point is disabled for pathfinding. By default, all points are enabled.
*/
//go:nosplit
func (self class) IsPointSolid(id gd.Vector2i) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_is_point_solid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param weight_scale] for the point with the given [param id]. The [param weight_scale] is multiplied by the result of [method _compute_cost] when determining the overall cost of traveling across a segment from a neighboring point to this point.
[b]Note:[/b] Calling [method update] is not needed after the call of this function.
*/
//go:nosplit
func (self class) SetPointWeightScale(id gd.Vector2i, weight_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, weight_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_set_point_weight_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the weight scale of the point associated with the given [param id].
*/
//go:nosplit
func (self class) GetPointWeightScale(id gd.Vector2i) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_point_weight_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Fills the given [param region] on the grid with the specified value for the solid flag.
[b]Note:[/b] Calling [method update] is not needed after the call of this function.
*/
//go:nosplit
func (self class) FillSolidRegion(region gd.Rect2i, solid bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, solid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_fill_solid_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Fills the given [param region] on the grid with the specified value for the weight scale.
[b]Note:[/b] Calling [method update] is not needed after the call of this function.
*/
//go:nosplit
func (self class) FillWeightScaleRegion(region gd.Rect2i, weight_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, weight_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_fill_weight_scale_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the grid and sets the [member region] to [code]Rect2i(0, 0, 0, 0)[/code].
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the point associated with the given [param id].
*/
//go:nosplit
func (self class) GetPointPosition(id gd.Vector2i) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array with the points that are in the path found by [AStarGrid2D] between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
[b]Note:[/b] This method is not thread-safe. If called from a [Thread], it will return an empty array and will print an error message.
*/
//go:nosplit
func (self class) GetPointPath(ctx gd.Lifetime, from_id gd.Vector2i, to_id gd.Vector2i, allow_partial_path bool) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, allow_partial_path)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_point_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with the IDs of the points that form the path found by AStar2D between the given points. The array is ordered from the starting point to the ending point of the path.
If there is no valid path to the target, and [param allow_partial_path] is [code]true[/code], returns a path to the point closest to the target that can be reached.
*/
//go:nosplit
func (self class) GetIdPath(ctx gd.Lifetime, from_id gd.Vector2i, to_id gd.Vector2i, allow_partial_path bool) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_id)
	callframe.Arg(frame, to_id)
	callframe.Arg(frame, allow_partial_path)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AStarGrid2D.Bind_get_id_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
func (self class) AsAStarGrid2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAStarGrid2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_estimate_cost": return reflect.ValueOf(self._estimate_cost);
	case "_compute_cost": return reflect.ValueOf(self._compute_cost);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_estimate_cost": return reflect.ValueOf(self._estimate_cost);
	case "_compute_cost": return reflect.ValueOf(self._compute_cost);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("AStarGrid2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Heuristic = classdb.AStarGrid2DHeuristic

const (
/*The [url=https://en.wikipedia.org/wiki/Euclidean_distance]Euclidean heuristic[/url] to be used for the pathfinding using the following formula:
[codeblock]
dx = abs(to_id.x - from_id.x)
dy = abs(to_id.y - from_id.y)
result = sqrt(dx * dx + dy * dy)
[/codeblock]
[b]Note:[/b] This is also the internal heuristic used in [AStar3D] and [AStar2D] by default (with the inclusion of possible z-axis coordinate).*/
	HeuristicEuclidean Heuristic = 0
/*The [url=https://en.wikipedia.org/wiki/Taxicab_geometry]Manhattan heuristic[/url] to be used for the pathfinding using the following formula:
[codeblock]
dx = abs(to_id.x - from_id.x)
dy = abs(to_id.y - from_id.y)
result = dx + dy
[/codeblock]
[b]Note:[/b] This heuristic is intended to be used with 4-side orthogonal movements, provided by setting the [member diagonal_mode] to [constant DIAGONAL_MODE_NEVER].*/
	HeuristicManhattan Heuristic = 1
/*The Octile heuristic to be used for the pathfinding using the following formula:
[codeblock]
dx = abs(to_id.x - from_id.x)
dy = abs(to_id.y - from_id.y)
f = sqrt(2) - 1
result = (dx < dy) ? f * dx + dy : f * dy + dx;
[/codeblock]*/
	HeuristicOctile Heuristic = 2
/*The [url=https://en.wikipedia.org/wiki/Chebyshev_distance]Chebyshev heuristic[/url] to be used for the pathfinding using the following formula:
[codeblock]
dx = abs(to_id.x - from_id.x)
dy = abs(to_id.y - from_id.y)
result = max(dx, dy)
[/codeblock]*/
	HeuristicChebyshev Heuristic = 3
/*Represents the size of the [enum Heuristic] enum.*/
	HeuristicMax Heuristic = 4
)
type DiagonalMode = classdb.AStarGrid2DDiagonalMode

const (
/*The pathfinding algorithm will ignore solid neighbors around the target cell and allow passing using diagonals.*/
	DiagonalModeAlways DiagonalMode = 0
/*The pathfinding algorithm will ignore all diagonals and the way will be always orthogonal.*/
	DiagonalModeNever DiagonalMode = 1
/*The pathfinding algorithm will avoid using diagonals if at least two obstacles have been placed around the neighboring cells of the specific path segment.*/
	DiagonalModeAtLeastOneWalkable DiagonalMode = 2
/*The pathfinding algorithm will avoid using diagonals if any obstacle has been placed around the neighboring cells of the specific path segment.*/
	DiagonalModeOnlyIfNoObstacles DiagonalMode = 3
/*Represents the size of the [enum DiagonalMode] enum.*/
	DiagonalModeMax DiagonalMode = 4
)
type CellShape = classdb.AStarGrid2DCellShape

const (
/*Rectangular cell shape.*/
	CellShapeSquare CellShape = 0
/*Diamond cell shape (for isometric look). Cell coordinates layout where the horizontal axis goes up-right, and the vertical one goes down-right.*/
	CellShapeIsometricRight CellShape = 1
/*Diamond cell shape (for isometric look). Cell coordinates layout where the horizontal axis goes down-right, and the vertical one goes down-left.*/
	CellShapeIsometricDown CellShape = 2
/*Represents the size of the [enum CellShape] enum.*/
	CellShapeMax CellShape = 3
)