// Package NavigationPathQueryParameters3D provides methods for working with NavigationPathQueryParameters3D object instances.
package NavigationPathQueryParameters3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
import "graphics.gd/variant/Vector3"

var _ Object.ID
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
var _ = slices.Delete[[]struct{}, struct{}]

/*
By changing various properties of this object, such as the start and target position, you can configure path queries to the [NavigationServer3D].
*/
type Instance [1]gdclass.NavigationPathQueryParameters3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNavigationPathQueryParameters3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationPathQueryParameters3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationPathQueryParameters3D"))
	casted := Instance{*(*gdclass.NavigationPathQueryParameters3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Map() RID.Any {
	return RID.Any(class(self).GetMap())
}

func (self Instance) SetMap(value RID.Any) {
	class(self).SetMap(RID.Any(value))
}

func (self Instance) StartPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetStartPosition())
}

func (self Instance) SetStartPosition(value Vector3.XYZ) {
	class(self).SetStartPosition(Vector3.XYZ(value))
}

func (self Instance) TargetPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetTargetPosition())
}

func (self Instance) SetTargetPosition(value Vector3.XYZ) {
	class(self).SetTargetPosition(Vector3.XYZ(value))
}

func (self Instance) NavigationLayers() int {
	return int(int(class(self).GetNavigationLayers()))
}

func (self Instance) SetNavigationLayers(value int) {
	class(self).SetNavigationLayers(int64(value))
}

func (self Instance) PathfindingAlgorithm() gdclass.NavigationPathQueryParameters3DPathfindingAlgorithm {
	return gdclass.NavigationPathQueryParameters3DPathfindingAlgorithm(class(self).GetPathfindingAlgorithm())
}

func (self Instance) SetPathfindingAlgorithm(value gdclass.NavigationPathQueryParameters3DPathfindingAlgorithm) {
	class(self).SetPathfindingAlgorithm(value)
}

func (self Instance) PathPostprocessing() gdclass.NavigationPathQueryParameters3DPathPostProcessing {
	return gdclass.NavigationPathQueryParameters3DPathPostProcessing(class(self).GetPathPostprocessing())
}

func (self Instance) SetPathPostprocessing(value gdclass.NavigationPathQueryParameters3DPathPostProcessing) {
	class(self).SetPathPostprocessing(value)
}

func (self Instance) MetadataFlags() gdclass.NavigationPathQueryParameters3DPathMetadataFlags {
	return gdclass.NavigationPathQueryParameters3DPathMetadataFlags(class(self).GetMetadataFlags())
}

func (self Instance) SetMetadataFlags(value gdclass.NavigationPathQueryParameters3DPathMetadataFlags) {
	class(self).SetMetadataFlags(value)
}

func (self Instance) SimplifyPath() bool {
	return bool(class(self).GetSimplifyPath())
}

func (self Instance) SetSimplifyPath(value bool) {
	class(self).SetSimplifyPath(value)
}

func (self Instance) SimplifyEpsilon() Float.X {
	return Float.X(Float.X(class(self).GetSimplifyEpsilon()))
}

func (self Instance) SetSimplifyEpsilon(value Float.X) {
	class(self).SetSimplifyEpsilon(float64(value))
}

//go:nosplit
func (self class) SetPathfindingAlgorithm(pathfinding_algorithm gdclass.NavigationPathQueryParameters3DPathfindingAlgorithm) { //gd:NavigationPathQueryParameters3D.set_pathfinding_algorithm
	var frame = callframe.New()
	callframe.Arg(frame, pathfinding_algorithm)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathfindingAlgorithm() gdclass.NavigationPathQueryParameters3DPathfindingAlgorithm { //gd:NavigationPathQueryParameters3D.get_pathfinding_algorithm
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters3DPathfindingAlgorithm](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathPostprocessing(path_postprocessing gdclass.NavigationPathQueryParameters3DPathPostProcessing) { //gd:NavigationPathQueryParameters3D.set_path_postprocessing
	var frame = callframe.New()
	callframe.Arg(frame, path_postprocessing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathPostprocessing() gdclass.NavigationPathQueryParameters3DPathPostProcessing { //gd:NavigationPathQueryParameters3D.get_path_postprocessing
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters3DPathPostProcessing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMap(mapping RID.Any) { //gd:NavigationPathQueryParameters3D.set_map
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMap() RID.Any { //gd:NavigationPathQueryParameters3D.get_map
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStartPosition(start_position Vector3.XYZ) { //gd:NavigationPathQueryParameters3D.set_start_position
	var frame = callframe.New()
	callframe.Arg(frame, start_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStartPosition() Vector3.XYZ { //gd:NavigationPathQueryParameters3D.get_start_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetPosition(target_position Vector3.XYZ) { //gd:NavigationPathQueryParameters3D.set_target_position
	var frame = callframe.New()
	callframe.Arg(frame, target_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetPosition() Vector3.XYZ { //gd:NavigationPathQueryParameters3D.get_target_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationLayers(navigation_layers int64) { //gd:NavigationPathQueryParameters3D.set_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationLayers() int64 { //gd:NavigationPathQueryParameters3D.get_navigation_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMetadataFlags(flags gdclass.NavigationPathQueryParameters3DPathMetadataFlags) { //gd:NavigationPathQueryParameters3D.set_metadata_flags
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMetadataFlags() gdclass.NavigationPathQueryParameters3DPathMetadataFlags { //gd:NavigationPathQueryParameters3D.get_metadata_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters3DPathMetadataFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSimplifyPath(enabled bool) { //gd:NavigationPathQueryParameters3D.set_simplify_path
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_simplify_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimplifyPath() bool { //gd:NavigationPathQueryParameters3D.get_simplify_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_simplify_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSimplifyEpsilon(epsilon float64) { //gd:NavigationPathQueryParameters3D.set_simplify_epsilon
	var frame = callframe.New()
	callframe.Arg(frame, epsilon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimplifyEpsilon() float64 { //gd:NavigationPathQueryParameters3D.get_simplify_epsilon
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationPathQueryParameters3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsNavigationPathQueryParameters3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("NavigationPathQueryParameters3D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationPathQueryParameters3D{*(*gdclass.NavigationPathQueryParameters3D)(unsafe.Pointer(&ptr))}
	})
}

type PathfindingAlgorithm = gdclass.NavigationPathQueryParameters3DPathfindingAlgorithm //gd:NavigationPathQueryParameters3D.PathfindingAlgorithm

const (
	/*The path query uses the default A* pathfinding algorithm.*/
	PathfindingAlgorithmAstar PathfindingAlgorithm = 0
)

type PathPostProcessing = gdclass.NavigationPathQueryParameters3DPathPostProcessing //gd:NavigationPathQueryParameters3D.PathPostProcessing

const (
	/*Applies a funnel algorithm to the raw path corridor found by the pathfinding algorithm. This will result in the shortest path possible inside the path corridor. This postprocessing very much depends on the navigation mesh polygon layout and the created corridor. Especially tile- or gridbased layouts can face artificial corners with diagonal movement due to a jagged path corridor imposed by the cell shapes.*/
	PathPostprocessingCorridorfunnel PathPostProcessing = 0
	/*Centers every path position in the middle of the traveled navigation mesh polygon edge. This creates better paths for tile- or gridbased layouts that restrict the movement to the cells center.*/
	PathPostprocessingEdgecentered PathPostProcessing = 1
	/*Applies no postprocessing and returns the raw path corridor as found by the pathfinding algorithm.*/
	PathPostprocessingNone PathPostProcessing = 2
)

type PathMetadataFlags = gdclass.NavigationPathQueryParameters3DPathMetadataFlags //gd:NavigationPathQueryParameters3D.PathMetadataFlags

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
