// Code generated by the generate package DO NOT EDIT

// Package GridMap provides methods for working with GridMap object instances.
package GridMap

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
import "graphics.gd/classdb/MeshLibrary"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/PhysicsMaterial"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Basis"
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
import "graphics.gd/variant/Vector3i"

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
GridMap lets you place meshes on a grid interactively. It works both from the editor and from scripts, which can help you create in-game level editors.
GridMaps use a [MeshLibrary] which contains a list of tiles. Each tile is a mesh with materials plus optional collision and navigation shapes.
A GridMap contains a collection of cells. Each grid cell refers to a tile in the [MeshLibrary]. All cells in the map have the same dimensions.
Internally, a GridMap is split into a sparse collection of octants for efficient rendering and physics processing. Every octant has the same dimensions and can contain several cells.
[b]Note:[/b] GridMap doesn't extend [VisualInstance3D] and therefore can't be hidden or cull masked based on [member VisualInstance3D.layers]. If you make a light not affect the first layer, the whole GridMap won't be lit by the light in question.
*/
type Instance [1]gdclass.GridMap

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.GridMap

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGridMap() Instance
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionMaskValue(layer_number int, value bool) { //gd:GridMap.set_collision_mask_value
	Advanced(self).SetCollisionMaskValue(int64(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionMaskValue(layer_number int) bool { //gd:GridMap.get_collision_mask_value
	return bool(Advanced(self).GetCollisionMaskValue(int64(layer_number)))
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionLayerValue(layer_number int, value bool) { //gd:GridMap.set_collision_layer_value
	Advanced(self).SetCollisionLayerValue(int64(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionLayerValue(layer_number int) bool { //gd:GridMap.get_collision_layer_value
	return bool(Advanced(self).GetCollisionLayerValue(int64(layer_number)))
}

/*
Sets the [RID] of the navigation map this GridMap node should use for its cell baked navigation meshes.
*/
func (self Instance) SetNavigationMap(navigation_map RID.NavigationMap3D) { //gd:GridMap.set_navigation_map
	Advanced(self).SetNavigationMap(RID.Any(navigation_map))
}

/*
Returns the [RID] of the navigation map this GridMap node uses for its cell baked navigation meshes.
This function returns always the map set on the GridMap node and not the map on the NavigationServer. If the map is changed directly with the NavigationServer API the GridMap node will not be aware of the map change.
*/
func (self Instance) GetNavigationMap() RID.NavigationMap3D { //gd:GridMap.get_navigation_map
	return RID.NavigationMap3D(Advanced(self).GetNavigationMap())
}

/*
Sets the mesh index for the cell referenced by its grid coordinates.
A negative item index such as [constant INVALID_CELL_ITEM] will clear the cell.
Optionally, the item's orientation can be passed. For valid orientation values, see [method get_orthogonal_index_from_basis].
*/
func (self Instance) SetCellItem(position Vector3i.XYZ, item CellItem) { //gd:GridMap.set_cell_item
	Advanced(self).SetCellItem(Vector3i.XYZ(position), int64(item), int64(0))
}

/*
Sets the mesh index for the cell referenced by its grid coordinates.
A negative item index such as [constant INVALID_CELL_ITEM] will clear the cell.
Optionally, the item's orientation can be passed. For valid orientation values, see [method get_orthogonal_index_from_basis].
*/
func (self Expanded) SetCellItem(position Vector3i.XYZ, item CellItem, orientation int) { //gd:GridMap.set_cell_item
	Advanced(self).SetCellItem(Vector3i.XYZ(position), int64(item), int64(orientation))
}

/*
The [MeshLibrary] item index located at the given grid coordinates. If the cell is empty, [constant INVALID_CELL_ITEM] will be returned.
*/
func (self Instance) GetCellItem(position Vector3i.XYZ) CellItem { //gd:GridMap.get_cell_item
	return CellItem(int(Advanced(self).GetCellItem(Vector3i.XYZ(position))))
}

/*
The orientation of the cell at the given grid coordinates. [code]-1[/code] is returned if the cell is empty.
*/
func (self Instance) GetCellItemOrientation(position Vector3i.XYZ) CellItem { //gd:GridMap.get_cell_item_orientation
	return CellItem(int(Advanced(self).GetCellItemOrientation(Vector3i.XYZ(position))))
}

/*
Returns the basis that gives the specified cell its orientation.
*/
func (self Instance) GetCellItemBasis(position Vector3i.XYZ) Basis.XYZ { //gd:GridMap.get_cell_item_basis
	return Basis.XYZ(Advanced(self).GetCellItemBasis(Vector3i.XYZ(position)))
}

/*
Returns one of 24 possible rotations that lie along the vectors (x,y,z) with each component being either -1, 0, or 1. For further details, refer to the Godot source code.
*/
func (self Instance) GetBasisWithOrthogonalIndex(index int) Basis.XYZ { //gd:GridMap.get_basis_with_orthogonal_index
	return Basis.XYZ(Advanced(self).GetBasisWithOrthogonalIndex(int64(index)))
}

/*
This function considers a discretization of rotations into 24 points on unit sphere, lying along the vectors (x,y,z) with each component being either -1, 0, or 1, and returns the index (in the range from 0 to 23) of the point best representing the orientation of the object. For further details, refer to the Godot source code.
*/
func (self Instance) GetOrthogonalIndexFromBasis(basis Basis.XYZ) int { //gd:GridMap.get_orthogonal_index_from_basis
	return int(int(Advanced(self).GetOrthogonalIndexFromBasis(Basis.XYZ(basis))))
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node3D.to_local] before passing it to this method. See also [method map_to_local].
*/
func (self Instance) LocalToMap(local_position Vector3.XYZ) Vector3i.XYZ { //gd:GridMap.local_to_map
	return Vector3i.XYZ(Advanced(self).LocalToMap(Vector3.XYZ(local_position)))
}

/*
Returns the position of a grid cell in the GridMap's local coordinate space. To convert the returned value into global coordinates, use [method Node3D.to_global]. See also [method local_to_map].
*/
func (self Instance) MapToLocal(map_position Vector3i.XYZ) Vector3.XYZ { //gd:GridMap.map_to_local
	return Vector3.XYZ(Advanced(self).MapToLocal(Vector3i.XYZ(map_position)))
}

/*
This method does nothing.
*/
func (self Instance) ResourceChanged(resource Resource.Instance) { //gd:GridMap.resource_changed
	Advanced(self).ResourceChanged(resource)
}

/*
Clear all cells.
*/
func (self Instance) Clear() { //gd:GridMap.clear
	Advanced(self).Clear()
}

/*
Returns an array of [Vector3] with the non-empty cell coordinates in the grid map.
*/
func (self Instance) GetUsedCells() []Vector3i.XYZ { //gd:GridMap.get_used_cells
	return []Vector3i.XYZ(gd.ArrayAs[[]Vector3i.XYZ](gd.InternalArray(Advanced(self).GetUsedCells())))
}

/*
Returns an array of all cells with the given item index specified in [param item].
*/
func (self Instance) GetUsedCellsByItem(item CellItem) []Vector3i.XYZ { //gd:GridMap.get_used_cells_by_item
	return []Vector3i.XYZ(gd.ArrayAs[[]Vector3i.XYZ](gd.InternalArray(Advanced(self).GetUsedCellsByItem(int64(item)))))
}

/*
Returns an array of [Transform3D] and [Mesh] references corresponding to the non-empty cells in the grid. The transforms are specified in local space.
*/
func (self Instance) GetMeshes() []any { //gd:GridMap.get_meshes
	return []any(gd.ArrayAs[[]any](gd.InternalArray(Advanced(self).GetMeshes())))
}

/*
Returns an array of [ArrayMesh]es and [Transform3D] references of all bake meshes that exist within the current GridMap.
*/
func (self Instance) GetBakeMeshes() []any { //gd:GridMap.get_bake_meshes
	return []any(gd.ArrayAs[[]any](gd.InternalArray(Advanced(self).GetBakeMeshes())))
}

/*
Returns [RID] of a baked mesh with the given [param idx].
*/
func (self Instance) GetBakeMeshInstance(idx int) RID.Mesh { //gd:GridMap.get_bake_mesh_instance
	return RID.Mesh(Advanced(self).GetBakeMeshInstance(int64(idx)))
}

/*
Clears all baked meshes. See [method make_baked_meshes].
*/
func (self Instance) ClearBakedMeshes() { //gd:GridMap.clear_baked_meshes
	Advanced(self).ClearBakedMeshes()
}

/*
Bakes lightmap data for all meshes in the assigned [MeshLibrary].
*/
func (self Instance) MakeBakedMeshes() { //gd:GridMap.make_baked_meshes
	Advanced(self).MakeBakedMeshes(false, float64(0.1))
}

/*
Bakes lightmap data for all meshes in the assigned [MeshLibrary].
*/
func (self Expanded) MakeBakedMeshes(gen_lightmap_uv bool, lightmap_uv_texel_size Float.X) { //gd:GridMap.make_baked_meshes
	Advanced(self).MakeBakedMeshes(gen_lightmap_uv, float64(lightmap_uv_texel_size))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GridMap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GridMap"))
	casted := Instance{*(*gdclass.GridMap)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) MeshLibrary() MeshLibrary.Instance {
	return MeshLibrary.Instance(class(self).GetMeshLibrary())
}

func (self Instance) SetMeshLibrary(value MeshLibrary.Instance) {
	class(self).SetMeshLibrary(value)
}

func (self Instance) PhysicsMaterial() PhysicsMaterial.Instance {
	return PhysicsMaterial.Instance(class(self).GetPhysicsMaterial())
}

func (self Instance) SetPhysicsMaterial(value PhysicsMaterial.Instance) {
	class(self).SetPhysicsMaterial(value)
}

func (self Instance) CellSize() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetCellSize())
}

func (self Instance) SetCellSize(value Vector3.XYZ) {
	class(self).SetCellSize(Vector3.XYZ(value))
}

func (self Instance) CellOctantSize() int {
	return int(int(class(self).GetOctantSize()))
}

func (self Instance) SetCellOctantSize(value int) {
	class(self).SetOctantSize(int64(value))
}

func (self Instance) CellCenterX() bool {
	return bool(class(self).GetCenterX())
}

func (self Instance) SetCellCenterX(value bool) {
	class(self).SetCenterX(value)
}

func (self Instance) CellCenterY() bool {
	return bool(class(self).GetCenterY())
}

func (self Instance) SetCellCenterY(value bool) {
	class(self).SetCenterY(value)
}

func (self Instance) CellCenterZ() bool {
	return bool(class(self).GetCenterZ())
}

func (self Instance) SetCellCenterZ(value bool) {
	class(self).SetCenterZ(value)
}

func (self Instance) CellScale() Float.X {
	return Float.X(Float.X(class(self).GetCellScale()))
}

func (self Instance) SetCellScale(value Float.X) {
	class(self).SetCellScale(float64(value))
}

func (self Instance) CollisionLayer() int {
	return int(int(class(self).GetCollisionLayer()))
}

func (self Instance) SetCollisionLayer(value int) {
	class(self).SetCollisionLayer(int64(value))
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(int64(value))
}

func (self Instance) CollisionPriority() Float.X {
	return Float.X(Float.X(class(self).GetCollisionPriority()))
}

func (self Instance) SetCollisionPriority(value Float.X) {
	class(self).SetCollisionPriority(float64(value))
}

func (self Instance) BakeNavigation() bool {
	return bool(class(self).IsBakingNavigation())
}

func (self Instance) SetBakeNavigation(value bool) {
	class(self).SetBakeNavigation(value)
}

//go:nosplit
func (self class) SetCollisionLayer(layer int64) { //gd:GridMap.set_collision_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionLayer() int64 { //gd:GridMap.get_collision_layer
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(mask int64) { //gd:GridMap.set_collision_mask
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() int64 { //gd:GridMap.get_collision_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number int64, value bool) { //gd:GridMap.set_collision_mask_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number int64) bool { //gd:GridMap.get_collision_mask_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionLayerValue(layer_number int64, value bool) { //gd:GridMap.set_collision_layer_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionLayerValue(layer_number int64) bool { //gd:GridMap.get_collision_layer_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionPriority(priority float64) { //gd:GridMap.set_collision_priority
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionPriority() float64 { //gd:GridMap.get_collision_priority
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsMaterial(material [1]gdclass.PhysicsMaterial) { //gd:GridMap.set_physics_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_physics_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsMaterial() [1]gdclass.PhysicsMaterial { //gd:GridMap.get_physics_material
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_physics_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PhysicsMaterial{gd.PointerWithOwnershipTransferredToGo[gdclass.PhysicsMaterial](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeNavigation(bake_navigation bool) { //gd:GridMap.set_bake_navigation
	var frame = callframe.New()
	callframe.Arg(frame, bake_navigation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_bake_navigation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsBakingNavigation() bool { //gd:GridMap.is_baking_navigation
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_is_baking_navigation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [RID] of the navigation map this GridMap node should use for its cell baked navigation meshes.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map RID.Any) { //gd:GridMap.set_navigation_map
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [RID] of the navigation map this GridMap node uses for its cell baked navigation meshes.
This function returns always the map set on the GridMap node and not the map on the NavigationServer. If the map is changed directly with the NavigationServer API the GridMap node will not be aware of the map change.
*/
//go:nosplit
func (self class) GetNavigationMap() RID.Any { //gd:GridMap.get_navigation_map
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMeshLibrary(mesh_library [1]gdclass.MeshLibrary) { //gd:GridMap.set_mesh_library
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh_library[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_mesh_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMeshLibrary() [1]gdclass.MeshLibrary { //gd:GridMap.get_mesh_library
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_mesh_library, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MeshLibrary{gd.PointerWithOwnershipTransferredToGo[gdclass.MeshLibrary](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellSize(size Vector3.XYZ) { //gd:GridMap.set_cell_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellSize() Vector3.XYZ { //gd:GridMap.get_cell_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellScale(scale float64) { //gd:GridMap.set_cell_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_cell_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellScale() float64 { //gd:GridMap.get_cell_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOctantSize(size int64) { //gd:GridMap.set_octant_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_octant_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOctantSize() int64 { //gd:GridMap.get_octant_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_octant_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the mesh index for the cell referenced by its grid coordinates.
A negative item index such as [constant INVALID_CELL_ITEM] will clear the cell.
Optionally, the item's orientation can be passed. For valid orientation values, see [method get_orthogonal_index_from_basis].
*/
//go:nosplit
func (self class) SetCellItem(position Vector3i.XYZ, item int64, orientation int64) { //gd:GridMap.set_cell_item
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, item)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_cell_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
The [MeshLibrary] item index located at the given grid coordinates. If the cell is empty, [constant INVALID_CELL_ITEM] will be returned.
*/
//go:nosplit
func (self class) GetCellItem(position Vector3i.XYZ) int64 { //gd:GridMap.get_cell_item
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The orientation of the cell at the given grid coordinates. [code]-1[/code] is returned if the cell is empty.
*/
//go:nosplit
func (self class) GetCellItemOrientation(position Vector3i.XYZ) int64 { //gd:GridMap.get_cell_item_orientation
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_item_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the basis that gives the specified cell its orientation.
*/
//go:nosplit
func (self class) GetCellItemBasis(position Vector3i.XYZ) Basis.XYZ { //gd:GridMap.get_cell_item_basis
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[Basis.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_cell_item_basis, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Basis.Transposed(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns one of 24 possible rotations that lie along the vectors (x,y,z) with each component being either -1, 0, or 1. For further details, refer to the Godot source code.
*/
//go:nosplit
func (self class) GetBasisWithOrthogonalIndex(index int64) Basis.XYZ { //gd:GridMap.get_basis_with_orthogonal_index
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[Basis.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_basis_with_orthogonal_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Basis.Transposed(r_ret.Get())
	frame.Free()
	return ret
}

/*
This function considers a discretization of rotations into 24 points on unit sphere, lying along the vectors (x,y,z) with each component being either -1, 0, or 1, and returns the index (in the range from 0 to 23) of the point best representing the orientation of the object. For further details, refer to the Godot source code.
*/
//go:nosplit
func (self class) GetOrthogonalIndexFromBasis(basis Basis.XYZ) int64 { //gd:GridMap.get_orthogonal_index_from_basis
	var frame = callframe.New()
	callframe.Arg(frame, Basis.Transposed(basis))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_orthogonal_index_from_basis, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node3D.to_local] before passing it to this method. See also [method map_to_local].
*/
//go:nosplit
func (self class) LocalToMap(local_position Vector3.XYZ) Vector3i.XYZ { //gd:GridMap.local_to_map
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Ret[Vector3i.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_local_to_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position of a grid cell in the GridMap's local coordinate space. To convert the returned value into global coordinates, use [method Node3D.to_global]. See also [method local_to_map].
*/
//go:nosplit
func (self class) MapToLocal(map_position Vector3i.XYZ) Vector3.XYZ { //gd:GridMap.map_to_local
	var frame = callframe.New()
	callframe.Arg(frame, map_position)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_map_to_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) ResourceChanged(resource [1]gdclass.Resource) { //gd:GridMap.resource_changed
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_resource_changed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCenterX(enable bool) { //gd:GridMap.set_center_x
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_center_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterX() bool { //gd:GridMap.get_center_x
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_center_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterY(enable bool) { //gd:GridMap.set_center_y
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_center_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterY() bool { //gd:GridMap.get_center_y
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_center_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterZ(enable bool) { //gd:GridMap.set_center_z
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_set_center_z, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterZ() bool { //gd:GridMap.get_center_z
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_center_z, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clear all cells.
*/
//go:nosplit
func (self class) Clear() { //gd:GridMap.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of [Vector3] with the non-empty cell coordinates in the grid map.
*/
//go:nosplit
func (self class) GetUsedCells() Array.Contains[Vector3i.XYZ] { //gd:GridMap.get_used_cells
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Vector3i.XYZ]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of all cells with the given item index specified in [param item].
*/
//go:nosplit
func (self class) GetUsedCellsByItem(item int64) Array.Contains[Vector3i.XYZ] { //gd:GridMap.get_used_cells_by_item
	var frame = callframe.New()
	callframe.Arg(frame, item)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_used_cells_by_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Vector3i.XYZ]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of [Transform3D] and [Mesh] references corresponding to the non-empty cells in the grid. The transforms are specified in local space.
*/
//go:nosplit
func (self class) GetMeshes() Array.Any { //gd:GridMap.get_meshes
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of [ArrayMesh]es and [Transform3D] references of all bake meshes that exist within the current GridMap.
*/
//go:nosplit
func (self class) GetBakeMeshes() Array.Any { //gd:GridMap.get_bake_meshes
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_bake_meshes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [RID] of a baked mesh with the given [param idx].
*/
//go:nosplit
func (self class) GetBakeMeshInstance(idx int64) RID.Any { //gd:GridMap.get_bake_mesh_instance
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_get_bake_mesh_instance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears all baked meshes. See [method make_baked_meshes].
*/
//go:nosplit
func (self class) ClearBakedMeshes() { //gd:GridMap.clear_baked_meshes
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_clear_baked_meshes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Bakes lightmap data for all meshes in the assigned [MeshLibrary].
*/
//go:nosplit
func (self class) MakeBakedMeshes(gen_lightmap_uv bool, lightmap_uv_texel_size float64) { //gd:GridMap.make_baked_meshes
	var frame = callframe.New()
	callframe.Arg(frame, gen_lightmap_uv)
	callframe.Arg(frame, lightmap_uv_texel_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GridMap.Bind_make_baked_meshes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnCellSizeChanged(cb func(cell_size Vector3.XYZ)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("cell_size_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}

func (self class) AsGridMap() Advanced               { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGridMap() Instance            { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsGridMap() Instance       { return self.Super().AsGridMap() }
func (self class) AsNode3D() Node3D.Advanced         { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode3D() Node3D.Instance { return self.Super().AsNode3D() }
func (self Instance) AsNode3D() Node3D.Instance      { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced             { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance     { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance          { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("GridMap", func(ptr gd.Object) any { return [1]gdclass.GridMap{*(*gdclass.GridMap)(unsafe.Pointer(&ptr))} })
}

type CellItem int

const InvalidCellItem CellItem = -1 //gd:GridMap.INVALID_CELL_ITEM
