package NavigationPathQueryParameters2D

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
By changing various properties of this object, such as the start and target position, you can configure path queries to the [NavigationServer2D].

*/
type Go [1]classdb.NavigationPathQueryParameters2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NavigationPathQueryParameters2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("NavigationPathQueryParameters2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Map() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
		return gd.RID(class(self).GetMap())
}

func (self Go) SetMap(value gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMap(value)
}

func (self Go) StartPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetStartPosition())
}

func (self Go) SetStartPosition(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStartPosition(value)
}

func (self Go) TargetPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetTargetPosition())
}

func (self Go) SetTargetPosition(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetPosition(value)
}

func (self Go) NavigationLayers() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetNavigationLayers()))
}

func (self Go) SetNavigationLayers(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Go) PathfindingAlgorithm() classdb.NavigationPathQueryParameters2DPathfindingAlgorithm {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationPathQueryParameters2DPathfindingAlgorithm(class(self).GetPathfindingAlgorithm())
}

func (self Go) SetPathfindingAlgorithm(value classdb.NavigationPathQueryParameters2DPathfindingAlgorithm) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathfindingAlgorithm(value)
}

func (self Go) PathPostprocessing() classdb.NavigationPathQueryParameters2DPathPostProcessing {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationPathQueryParameters2DPathPostProcessing(class(self).GetPathPostprocessing())
}

func (self Go) SetPathPostprocessing(value classdb.NavigationPathQueryParameters2DPathPostProcessing) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathPostprocessing(value)
}

func (self Go) MetadataFlags() classdb.NavigationPathQueryParameters2DPathMetadataFlags {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationPathQueryParameters2DPathMetadataFlags(class(self).GetMetadataFlags())
}

func (self Go) SetMetadataFlags(value classdb.NavigationPathQueryParameters2DPathMetadataFlags) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMetadataFlags(value)
}

func (self Go) SimplifyPath() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetSimplifyPath())
}

func (self Go) SetSimplifyPath(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSimplifyPath(value)
}

func (self Go) SimplifyEpsilon() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSimplifyEpsilon()))
}

func (self Go) SetSimplifyEpsilon(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSimplifyEpsilon(gd.Float(value))
}

//go:nosplit
func (self class) SetPathfindingAlgorithm(pathfinding_algorithm classdb.NavigationPathQueryParameters2DPathfindingAlgorithm)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pathfinding_algorithm)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathfindingAlgorithm() classdb.NavigationPathQueryParameters2DPathfindingAlgorithm {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters2DPathfindingAlgorithm](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathPostprocessing(path_postprocessing classdb.NavigationPathQueryParameters2DPathPostProcessing)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, path_postprocessing)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathPostprocessing() classdb.NavigationPathQueryParameters2DPathPostProcessing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters2DPathPostProcessing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMap(mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMap() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStartPosition(start_position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, start_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStartPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetPosition(target_position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, target_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMetadataFlags(flags classdb.NavigationPathQueryParameters2DPathMetadataFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMetadataFlags() classdb.NavigationPathQueryParameters2DPathMetadataFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters2DPathMetadataFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSimplifyPath(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimplifyPath() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSimplifyEpsilon(epsilon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, epsilon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_set_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimplifyEpsilon() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters2D.Bind_get_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationPathQueryParameters2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNavigationPathQueryParameters2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("NavigationPathQueryParameters2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type PathfindingAlgorithm = classdb.NavigationPathQueryParameters2DPathfindingAlgorithm

const (
/*The path query uses the default A* pathfinding algorithm.*/
	PathfindingAlgorithmAstar PathfindingAlgorithm = 0
)
type PathPostProcessing = classdb.NavigationPathQueryParameters2DPathPostProcessing

const (
/*Applies a funnel algorithm to the raw path corridor found by the pathfinding algorithm. This will result in the shortest path possible inside the path corridor. This postprocessing very much depends on the navigation mesh polygon layout and the created corridor. Especially tile- or gridbased layouts can face artificial corners with diagonal movement due to a jagged path corridor imposed by the cell shapes.*/
	PathPostprocessingCorridorfunnel PathPostProcessing = 0
/*Centers every path position in the middle of the traveled navigation mesh polygon edge. This creates better paths for tile- or gridbased layouts that restrict the movement to the cells center.*/
	PathPostprocessingEdgecentered PathPostProcessing = 1
)
type PathMetadataFlags = classdb.NavigationPathQueryParameters2DPathMetadataFlags

const (
/*Don't include any additional metadata about the returned path.*/
	PathMetadataIncludeNone PathMetadataFlags = 0
/*Include the type of navigation primitive (region or link) that each point of the path goes through.*/
	PathMetadataIncludeTypes PathMetadataFlags = 1
/*Include the [RID]s of the regions and links that each point of the path goes through.*/
	PathMetadataIncludeRids PathMetadataFlags = 2
/*Include the [code]ObjectID[/code]s of the [Object]s which manage the regions and links each point of the path goes through.*/
	PathMetadataIncludeOwners PathMetadataFlags = 4
/*Include all available metadata about the returned path.*/
	PathMetadataIncludeAll PathMetadataFlags = 7
)