// Code generated by the generate package DO NOT EDIT

// Package NavigationMeshSourceGeometryData3D provides methods for working with NavigationMeshSourceGeometryData3D object instances.
package NavigationMeshSourceGeometryData3D

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
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/AABB"
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
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"

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
Container for parsed source geometry data used in navigation mesh baking.
*/
type Instance [1]gdclass.NavigationMeshSourceGeometryData3D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNavigationMeshSourceGeometryData3D() Instance
}

/*
Appends arrays of [param vertices] and [param indices] at the end of the existing arrays. Adds the existing index as an offset to the appended indices.
*/
func (self Instance) AppendArrays(vertices []float32, indices []int32) { //gd:NavigationMeshSourceGeometryData3D.append_arrays
	Advanced(self).AppendArrays(Packed.New(vertices...), Packed.New(indices...))
}

/*
Clears the internal data.
*/
func (self Instance) Clear() { //gd:NavigationMeshSourceGeometryData3D.clear
	Advanced(self).Clear()
}

/*
Returns [code]true[/code] when parsed source geometry data exists.
*/
func (self Instance) HasData() bool { //gd:NavigationMeshSourceGeometryData3D.has_data
	return bool(Advanced(self).HasData())
}

/*
Adds the geometry data of a [Mesh] resource to the navigation mesh baking data. The mesh must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
func (self Instance) AddMesh(mesh Mesh.Instance, xform Transform3D.BasisOrigin) { //gd:NavigationMeshSourceGeometryData3D.add_mesh
	Advanced(self).AddMesh(mesh, Transform3D.BasisOrigin(xform))
}

/*
Adds an [Array] the size of [constant Mesh.ARRAY_MAX] and with vertices at index [constant Mesh.ARRAY_VERTEX] and indices at index [constant Mesh.ARRAY_INDEX] to the navigation mesh baking data. The array must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
func (self Instance) AddMeshArray(mesh_array []any, xform Transform3D.BasisOrigin) { //gd:NavigationMeshSourceGeometryData3D.add_mesh_array
	Advanced(self).AddMeshArray(gd.EngineArrayFromSlice(mesh_array), Transform3D.BasisOrigin(xform))
}

/*
Adds an array of vertex positions to the geometry data for navigation mesh baking to form triangulated faces. For each face the array must have three vertex positions in clockwise winding order. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
func (self Instance) AddFaces(faces []Vector3.XYZ, xform Transform3D.BasisOrigin) { //gd:NavigationMeshSourceGeometryData3D.add_faces
	Advanced(self).AddFaces(Packed.New(faces...), Transform3D.BasisOrigin(xform))
}

/*
Adds the geometry data of another [NavigationMeshSourceGeometryData3D] to the navigation mesh baking data.
*/
func (self Instance) Merge(other_geometry Instance) { //gd:NavigationMeshSourceGeometryData3D.merge
	Advanced(self).Merge(other_geometry)
}

/*
Adds a projected obstruction shape to the source geometry. The [param vertices] are considered projected on a xz-axes plane, placed at the global y-axis [param elevation] and extruded by [param height]. If [param carve] is [code]true[/code] the carved shape will not be affected by additional offsets (e.g. agent radius) of the navigation mesh baking process.
*/
func (self Instance) AddProjectedObstruction(vertices []Vector3.XYZ, elevation Float.X, height Float.X, carve bool) { //gd:NavigationMeshSourceGeometryData3D.add_projected_obstruction
	Advanced(self).AddProjectedObstruction(Packed.New(vertices...), float64(elevation), float64(height), carve)
}

/*
Clears all projected obstructions.
*/
func (self Instance) ClearProjectedObstructions() { //gd:NavigationMeshSourceGeometryData3D.clear_projected_obstructions
	Advanced(self).ClearProjectedObstructions()
}

/*
Returns an axis-aligned bounding box that covers all the stored geometry data. The bounds are calculated when calling this function with the result cached until further geometry changes are made.
*/
func (self Instance) GetBounds() AABB.PositionSize { //gd:NavigationMeshSourceGeometryData3D.get_bounds
	return AABB.PositionSize(Advanced(self).GetBounds())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationMeshSourceGeometryData3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationMeshSourceGeometryData3D"))
	casted := Instance{*(*gdclass.NavigationMeshSourceGeometryData3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Vertices() []float32 {
	return []float32(slices.Collect(class(self).GetVertices().Values()))
}

func (self Instance) SetVertices(value []float32) {
	class(self).SetVertices(Packed.New(value...))
}

func (self Instance) Indices() []int32 {
	return []int32(slices.Collect(class(self).GetIndices().Values()))
}

func (self Instance) SetIndices(value []int32) {
	class(self).SetIndices(Packed.New(value...))
}

func (self Instance) ProjectedObstructions() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetProjectedObstructions())))
}

func (self Instance) SetProjectedObstructions(value []any) {
	class(self).SetProjectedObstructions(gd.EngineArrayFromSlice(value))
}

/*
Sets the parsed source geometry data vertices. The vertices need to be matched with appropriated indices.
[b]Warning:[/b] Inappropriate data can crash the baking process of the involved third-party libraries.
*/
//go:nosplit
func (self class) SetVertices(vertices Packed.Array[float32]) { //gd:NavigationMeshSourceGeometryData3D.set_vertices
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](vertices)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the parsed source geometry data vertices array.
*/
//go:nosplit
func (self class) GetVertices() Packed.Array[float32] { //gd:NavigationMeshSourceGeometryData3D.get_vertices
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_get_vertices, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets the parsed source geometry data indices. The indices need to be matched with appropriated vertices.
[b]Warning:[/b] Inappropriate data can crash the baking process of the involved third-party libraries.
*/
//go:nosplit
func (self class) SetIndices(indices Packed.Array[int32]) { //gd:NavigationMeshSourceGeometryData3D.set_indices
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedInt32Array, int32](indices)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_set_indices, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the parsed source geometry data indices array.
*/
//go:nosplit
func (self class) GetIndices() Packed.Array[int32] { //gd:NavigationMeshSourceGeometryData3D.get_indices
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_get_indices, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Appends arrays of [param vertices] and [param indices] at the end of the existing arrays. Adds the existing index as an offset to the appended indices.
*/
//go:nosplit
func (self class) AppendArrays(vertices Packed.Array[float32], indices Packed.Array[int32]) { //gd:NavigationMeshSourceGeometryData3D.append_arrays
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](vertices)))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedInt32Array, int32](indices)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_append_arrays, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears the internal data.
*/
//go:nosplit
func (self class) Clear() { //gd:NavigationMeshSourceGeometryData3D.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] when parsed source geometry data exists.
*/
//go:nosplit
func (self class) HasData() bool { //gd:NavigationMeshSourceGeometryData3D.has_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_has_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds the geometry data of a [Mesh] resource to the navigation mesh baking data. The mesh must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddMesh(mesh [1]gdclass.Mesh, xform Transform3D.BasisOrigin) { //gd:NavigationMeshSourceGeometryData3D.add_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	callframe.Arg(frame, gd.Transposed(xform))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_add_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds an [Array] the size of [constant Mesh.ARRAY_MAX] and with vertices at index [constant Mesh.ARRAY_VERTEX] and indices at index [constant Mesh.ARRAY_INDEX] to the navigation mesh baking data. The array must have valid triangulated mesh data to be considered. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddMeshArray(mesh_array Array.Any, xform Transform3D.BasisOrigin) { //gd:NavigationMeshSourceGeometryData3D.add_mesh_array
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(mesh_array)))
	callframe.Arg(frame, gd.Transposed(xform))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_add_mesh_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds an array of vertex positions to the geometry data for navigation mesh baking to form triangulated faces. For each face the array must have three vertex positions in clockwise winding order. Since [NavigationMesh] resources have no transform, all vertex positions need to be offset by the node's transform using [param xform].
*/
//go:nosplit
func (self class) AddFaces(faces Packed.Array[Vector3.XYZ], xform Transform3D.BasisOrigin) { //gd:NavigationMeshSourceGeometryData3D.add_faces
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector3Array, Vector3.XYZ](faces)))
	callframe.Arg(frame, gd.Transposed(xform))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_add_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds the geometry data of another [NavigationMeshSourceGeometryData3D] to the navigation mesh baking data.
*/
//go:nosplit
func (self class) Merge(other_geometry [1]gdclass.NavigationMeshSourceGeometryData3D) { //gd:NavigationMeshSourceGeometryData3D.merge
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(other_geometry[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_merge, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a projected obstruction shape to the source geometry. The [param vertices] are considered projected on a xz-axes plane, placed at the global y-axis [param elevation] and extruded by [param height]. If [param carve] is [code]true[/code] the carved shape will not be affected by additional offsets (e.g. agent radius) of the navigation mesh baking process.
*/
//go:nosplit
func (self class) AddProjectedObstruction(vertices Packed.Array[Vector3.XYZ], elevation float64, height float64, carve bool) { //gd:NavigationMeshSourceGeometryData3D.add_projected_obstruction
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector3Array, Vector3.XYZ](vertices)))
	callframe.Arg(frame, elevation)
	callframe.Arg(frame, height)
	callframe.Arg(frame, carve)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_add_projected_obstruction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears all projected obstructions.
*/
//go:nosplit
func (self class) ClearProjectedObstructions() { //gd:NavigationMeshSourceGeometryData3D.clear_projected_obstructions
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_clear_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the projected obstructions with an Array of Dictionaries with the following key value pairs:
[codeblocks]
[gdscript]
"vertices" : PackedFloat32Array
"elevation" : float
"height" : float
"carve" : bool
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) SetProjectedObstructions(projected_obstructions Array.Any) { //gd:NavigationMeshSourceGeometryData3D.set_projected_obstructions
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(projected_obstructions)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_set_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the projected obstructions as an [Array] of dictionaries. Each [Dictionary] contains the following entries:
- [code]vertices[/code] - A [PackedFloat32Array] that defines the outline points of the projected shape.
- [code]elevation[/code] - A [float] that defines the projected shape placement on the y-axis.
- [code]height[/code] - A [float] that defines how much the projected shape is extruded along the y-axis.
- [code]carve[/code] - A [bool] that defines how the obstacle affects the navigation mesh baking. If [code]true[/code] the projected shape will not be affected by addition offsets, e.g. agent radius.
*/
//go:nosplit
func (self class) GetProjectedObstructions() Array.Any { //gd:NavigationMeshSourceGeometryData3D.get_projected_obstructions
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_get_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an axis-aligned bounding box that covers all the stored geometry data. The bounds are calculated when calling this function with the result cached until further geometry changes are made.
*/
//go:nosplit
func (self class) GetBounds() AABB.PositionSize { //gd:NavigationMeshSourceGeometryData3D.get_bounds
	var frame = callframe.New()
	var r_ret = callframe.Ret[AABB.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData3D.Bind_get_bounds, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationMeshSourceGeometryData3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsNavigationMeshSourceGeometryData3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsNavigationMeshSourceGeometryData3D() Instance {
	return self.Super().AsNavigationMeshSourceGeometryData3D()
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsResource() Resource.Instance { return self.Super().AsResource() }
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
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
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("NavigationMeshSourceGeometryData3D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationMeshSourceGeometryData3D{*(*gdclass.NavigationMeshSourceGeometryData3D)(unsafe.Pointer(&ptr))}
	})
}
