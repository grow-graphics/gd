package CSGPolygon3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/CSGPrimitive3D"
import "grow.graphics/gd/gdclass/CSGShape3D"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An array of 2D points is extruded to quickly and easily create a variety of 3D meshes. See also [CSGMesh3D] for using 3D meshes as CSG nodes.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.

*/
type Go [1]classdb.CSGPolygon3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CSGPolygon3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CSGPolygon3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Polygon() []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector2(class(self).GetPolygon(gc).AsSlice())
}

func (self Go) SetPolygon(value []gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPolygon(gc.PackedVector2Slice(value))
}

func (self Go) Mode() classdb.CSGPolygon3DMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CSGPolygon3DMode(class(self).GetMode())
}

func (self Go) SetMode(value classdb.CSGPolygon3DMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMode(value)
}

func (self Go) Depth() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDepth()))
}

func (self Go) SetDepth(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepth(gd.Float(value))
}

func (self Go) SpinDegrees() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSpinDegrees()))
}

func (self Go) SetSpinDegrees(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpinDegrees(gd.Float(value))
}

func (self Go) SpinSides() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSpinSides()))
}

func (self Go) SetSpinSides(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpinSides(gd.Int(value))
}

func (self Go) PathNode() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetPathNode(gc).String())
}

func (self Go) SetPathNode(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathNode(gc.String(value).NodePath(gc))
}

func (self Go) PathIntervalType() classdb.CSGPolygon3DPathIntervalType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CSGPolygon3DPathIntervalType(class(self).GetPathIntervalType())
}

func (self Go) SetPathIntervalType(value classdb.CSGPolygon3DPathIntervalType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathIntervalType(value)
}

func (self Go) PathInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPathInterval()))
}

func (self Go) SetPathInterval(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathInterval(gd.Float(value))
}

func (self Go) PathSimplifyAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPathSimplifyAngle()))
}

func (self Go) SetPathSimplifyAngle(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathSimplifyAngle(gd.Float(value))
}

func (self Go) PathRotation() classdb.CSGPolygon3DPathRotation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CSGPolygon3DPathRotation(class(self).GetPathRotation())
}

func (self Go) SetPathRotation(value classdb.CSGPolygon3DPathRotation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathRotation(value)
}

func (self Go) PathLocal() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPathLocal())
}

func (self Go) SetPathLocal(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathLocal(value)
}

func (self Go) PathContinuousU() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPathContinuousU())
}

func (self Go) SetPathContinuousU(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathContinuousU(value)
}

func (self Go) PathUDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPathUDistance()))
}

func (self Go) SetPathUDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathUDistance(gd.Float(value))
}

func (self Go) PathJoined() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPathJoined())
}

func (self Go) SetPathJoined(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathJoined(value)
}

func (self Go) SmoothFaces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetSmoothFaces())
}

func (self Go) SetSmoothFaces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSmoothFaces(value)
}

func (self Go) Material() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetMaterial(gc))
}

func (self Go) SetMaterial(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterial(value)
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(polygon))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPolygon(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMode(mode classdb.CSGPolygon3DMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMode() classdb.CSGPolygon3DMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CSGPolygon3DMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepth(depth gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpinDegrees(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_spin_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpinDegrees() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_spin_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpinSides(spin_sides gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, spin_sides)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_spin_sides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpinSides() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_spin_sides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathNode(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathIntervalType(interval_type classdb.CSGPolygon3DPathIntervalType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interval_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_interval_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathIntervalType() classdb.CSGPolygon3DPathIntervalType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CSGPolygon3DPathIntervalType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_interval_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathInterval(interval gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathInterval() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathSimplifyAngle(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_simplify_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathSimplifyAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_simplify_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathRotation(path_rotation classdb.CSGPolygon3DPathRotation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, path_rotation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathRotation() classdb.CSGPolygon3DPathRotation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CSGPolygon3DPathRotation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathLocal(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPathLocal() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_is_path_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathContinuousU(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_continuous_u, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPathContinuousU() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_is_path_continuous_u, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathUDistance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_u_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathUDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_u_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathJoined(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_joined, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPathJoined() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_is_path_joined, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaterial(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSmoothFaces(smooth_faces bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, smooth_faces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSmoothFaces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCSGPolygon3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCSGPolygon3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsCSGPrimitive3D() CSGPrimitive3D.GD { return *((*CSGPrimitive3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCSGPrimitive3D() CSGPrimitive3D.Go { return *((*CSGPrimitive3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCSGShape3D() CSGShape3D.GD { return *((*CSGShape3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCSGShape3D() CSGShape3D.Go { return *((*CSGShape3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.GD { return *((*GeometryInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() GeometryInstance3D.Go { return *((*GeometryInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCSGPrimitive3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCSGPrimitive3D(), name)
	}
}
func init() {classdb.Register("CSGPolygon3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Mode = classdb.CSGPolygon3DMode

const (
/*The [member polygon] shape is extruded along the negative Z axis.*/
	ModeDepth Mode = 0
/*The [member polygon] shape is extruded by rotating it around the Y axis.*/
	ModeSpin Mode = 1
/*The [member polygon] shape is extruded along the [Path3D] specified in [member path_node].*/
	ModePath Mode = 2
)
type PathRotation = classdb.CSGPolygon3DPathRotation

const (
/*The [member polygon] shape is not rotated.
[b]Note:[/b] Requires the path Z coordinates to continually decrease to ensure viable shapes.*/
	PathRotationPolygon PathRotation = 0
/*The [member polygon] shape is rotated along the path, but it is not rotated around the path axis.
[b]Note:[/b] Requires the path Z coordinates to continually decrease to ensure viable shapes.*/
	PathRotationPath PathRotation = 1
/*The [member polygon] shape follows the path and its rotations around the path axis.*/
	PathRotationPathFollow PathRotation = 2
)
type PathIntervalType = classdb.CSGPolygon3DPathIntervalType

const (
/*When [member mode] is set to [constant MODE_PATH], [member path_interval] will determine the distance, in meters, each interval of the path will extrude.*/
	PathIntervalDistance PathIntervalType = 0
/*When [member mode] is set to [constant MODE_PATH], [member path_interval] will subdivide the polygons along the path.*/
	PathIntervalSubdivide PathIntervalType = 1
)