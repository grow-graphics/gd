package MeshDataTool

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
MeshDataTool provides access to individual vertices in a [Mesh]. It allows users to read and edit vertex data of meshes. It also creates an array of faces and edges.
To use MeshDataTool, load a mesh with [method create_from_surface]. When you are finished editing the data commit the data to a mesh with [method commit_to_surface].
Below is an example of how MeshDataTool may be used.
[codeblocks]
[gdscript]
var mesh = ArrayMesh.new()
mesh.add_surface_from_arrays(Mesh.PRIMITIVE_TRIANGLES, BoxMesh.new().get_mesh_arrays())
var mdt = MeshDataTool.new()
mdt.create_from_surface(mesh, 0)
for i in range(mdt.get_vertex_count()):
    var vertex = mdt.get_vertex(i)
    # In this example we extend the mesh by one unit, which results in separated faces as it is flat shaded.
    vertex += mdt.get_vertex_normal(i)
    # Save your change.
    mdt.set_vertex(i, vertex)
mesh.clear_surfaces()
mdt.commit_to_surface(mesh)
var mi = MeshInstance.new()
mi.mesh = mesh
add_child(mi)
[/gdscript]
[csharp]
var mesh = new ArrayMesh();
mesh.AddSurfaceFromArrays(Mesh.PrimitiveType.Triangles, new BoxMesh().GetMeshArrays());
var mdt = new MeshDataTool();
mdt.CreateFromSurface(mesh, 0);
for (var i = 0; i < mdt.GetVertexCount(); i++)
{
    Vector3 vertex = mdt.GetVertex(i);
    // In this example we extend the mesh by one unit, which results in separated faces as it is flat shaded.
    vertex += mdt.GetVertexNormal(i);
    // Save your change.
    mdt.SetVertex(i, vertex);
}
mesh.ClearSurfaces();
mdt.CommitToSurface(mesh);
var mi = new MeshInstance();
mi.Mesh = mesh;
AddChild(mi);
[/csharp]
[/codeblocks]
See also [ArrayMesh], [ImmediateMesh] and [SurfaceTool] for procedural geometry generation.
[b]Note:[/b] Godot uses clockwise [url=https://learnopengl.com/Advanced-OpenGL/Face-culling]winding order[/url] for front faces of triangle primitive modes.

*/
type Go [1]classdb.MeshDataTool

/*
Clears all data currently in MeshDataTool.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Uses specified surface of given [Mesh] to populate data for MeshDataTool.
Requires [Mesh] with primitive type [constant Mesh.PRIMITIVE_TRIANGLES].
*/
func (self Go) CreateFromSurface(mesh gdclass.ArrayMesh, surface int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).CreateFromSurface(mesh, gd.Int(surface)))
}

/*
Adds a new surface to specified [Mesh] with edited data.
*/
func (self Go) CommitToSurface(mesh gdclass.ArrayMesh) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).CommitToSurface(mesh, gd.Int(0)))
}

/*
Returns the [Mesh]'s format as a combination of the [enum Mesh.ArrayFormat] flags. For example, a mesh containing both vertices and normals would return a format of [code]3[/code] because [constant Mesh.ARRAY_FORMAT_VERTEX] is [code]1[/code] and [constant Mesh.ARRAY_FORMAT_NORMAL] is [code]2[/code].
*/
func (self Go) GetFormat() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFormat()))
}

/*
Returns the total number of vertices in [Mesh].
*/
func (self Go) GetVertexCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetVertexCount()))
}

/*
Returns the number of edges in this [Mesh].
*/
func (self Go) GetEdgeCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetEdgeCount()))
}

/*
Returns the number of faces in this [Mesh].
*/
func (self Go) GetFaceCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFaceCount()))
}

/*
Sets the position of the given vertex.
*/
func (self Go) SetVertex(idx int, vertex gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertex(gd.Int(idx), vertex)
}

/*
Returns the position of the given vertex.
*/
func (self Go) GetVertex(idx int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetVertex(gd.Int(idx)))
}

/*
Sets the normal of the given vertex.
*/
func (self Go) SetVertexNormal(idx int, normal gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexNormal(gd.Int(idx), normal)
}

/*
Returns the normal of the given vertex.
*/
func (self Go) GetVertexNormal(idx int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetVertexNormal(gd.Int(idx)))
}

/*
Sets the tangent of the given vertex.
*/
func (self Go) SetVertexTangent(idx int, tangent gd.Plane) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexTangent(gd.Int(idx), tangent)
}

/*
Returns the tangent of the given vertex.
*/
func (self Go) GetVertexTangent(idx int) gd.Plane {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Plane(class(self).GetVertexTangent(gd.Int(idx)))
}

/*
Sets the UV of the given vertex.
*/
func (self Go) SetVertexUv(idx int, uv gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexUv(gd.Int(idx), uv)
}

/*
Returns the UV of the given vertex.
*/
func (self Go) GetVertexUv(idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetVertexUv(gd.Int(idx)))
}

/*
Sets the UV2 of the given vertex.
*/
func (self Go) SetVertexUv2(idx int, uv2 gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexUv2(gd.Int(idx), uv2)
}

/*
Returns the UV2 of the given vertex.
*/
func (self Go) GetVertexUv2(idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetVertexUv2(gd.Int(idx)))
}

/*
Sets the color of the given vertex.
*/
func (self Go) SetVertexColor(idx int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexColor(gd.Int(idx), color)
}

/*
Returns the color of the given vertex.
*/
func (self Go) GetVertexColor(idx int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetVertexColor(gd.Int(idx)))
}

/*
Sets the bones of the given vertex.
*/
func (self Go) SetVertexBones(idx int, bones []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexBones(gd.Int(idx), gc.PackedInt32Slice(bones))
}

/*
Returns the bones of the given vertex.
*/
func (self Go) GetVertexBones(idx int) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetVertexBones(gc, gd.Int(idx)).AsSlice())
}

/*
Sets the bone weights of the given vertex.
*/
func (self Go) SetVertexWeights(idx int, weights []float32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexWeights(gd.Int(idx), gc.PackedFloat32Slice(weights))
}

/*
Returns bone weights of the given vertex.
*/
func (self Go) GetVertexWeights(idx int) []float32 {
	gc := gd.GarbageCollector(); _ = gc
	return []float32(class(self).GetVertexWeights(gc, gd.Int(idx)).AsSlice())
}

/*
Sets the metadata associated with the given vertex.
*/
func (self Go) SetVertexMeta(idx int, meta gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexMeta(gd.Int(idx), meta)
}

/*
Returns the metadata associated with the given vertex.
*/
func (self Go) GetVertexMeta(idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetVertexMeta(gc, gd.Int(idx)))
}

/*
Returns an array of edges that share the given vertex.
*/
func (self Go) GetVertexEdges(idx int) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetVertexEdges(gc, gd.Int(idx)).AsSlice())
}

/*
Returns an array of faces that share the given vertex.
*/
func (self Go) GetVertexFaces(idx int) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetVertexFaces(gc, gd.Int(idx)).AsSlice())
}

/*
Returns index of specified vertex connected to given edge.
Vertex argument can only be 0 or 1 because edges are comprised of two vertices.
*/
func (self Go) GetEdgeVertex(idx int, vertex int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetEdgeVertex(gd.Int(idx), gd.Int(vertex))))
}

/*
Returns array of faces that touch given edge.
*/
func (self Go) GetEdgeFaces(idx int) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetEdgeFaces(gc, gd.Int(idx)).AsSlice())
}

/*
Sets the metadata of the given edge.
*/
func (self Go) SetEdgeMeta(idx int, meta gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEdgeMeta(gd.Int(idx), meta)
}

/*
Returns meta information assigned to given edge.
*/
func (self Go) GetEdgeMeta(idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetEdgeMeta(gc, gd.Int(idx)))
}

/*
Returns the specified vertex index of the given face.
Vertex argument must be either 0, 1, or 2 because faces contain three vertices.
[b]Example:[/b]
[codeblocks]
[gdscript]
var index = mesh_data_tool.get_face_vertex(0, 1) # Gets the index of the second vertex of the first face.
var position = mesh_data_tool.get_vertex(index)
var normal = mesh_data_tool.get_vertex_normal(index)
[/gdscript]
[csharp]
int index = meshDataTool.GetFaceVertex(0, 1); // Gets the index of the second vertex of the first face.
Vector3 position = meshDataTool.GetVertex(index);
Vector3 normal = meshDataTool.GetVertexNormal(index);
[/csharp]
[/codeblocks]
*/
func (self Go) GetFaceVertex(idx int, vertex int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFaceVertex(gd.Int(idx), gd.Int(vertex))))
}

/*
Returns specified edge associated with given face.
Edge argument must be either 0, 1, or 2 because a face only has three edges.
*/
func (self Go) GetFaceEdge(idx int, edge int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFaceEdge(gd.Int(idx), gd.Int(edge))))
}

/*
Sets the metadata of the given face.
*/
func (self Go) SetFaceMeta(idx int, meta gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFaceMeta(gd.Int(idx), meta)
}

/*
Returns the metadata associated with the given face.
*/
func (self Go) GetFaceMeta(idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetFaceMeta(gc, gd.Int(idx)))
}

/*
Calculates and returns the face normal of the given face.
*/
func (self Go) GetFaceNormal(idx int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetFaceNormal(gd.Int(idx)))
}

/*
Sets the material to be used by newly-constructed [Mesh].
*/
func (self Go) SetMaterial(material gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterial(material)
}

/*
Returns the material assigned to the [Mesh].
*/
func (self Go) GetMaterial() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Material(class(self).GetMaterial(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MeshDataTool
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MeshDataTool"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Clears all data currently in MeshDataTool.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Uses specified surface of given [Mesh] to populate data for MeshDataTool.
Requires [Mesh] with primitive type [constant Mesh.PRIMITIVE_TRIANGLES].
*/
//go:nosplit
func (self class) CreateFromSurface(mesh gdclass.ArrayMesh, surface gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_create_from_surface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new surface to specified [Mesh] with edited data.
*/
//go:nosplit
func (self class) CommitToSurface(mesh gdclass.ArrayMesh, compression_flags gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	callframe.Arg(frame, compression_flags)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_commit_to_surface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Mesh]'s format as a combination of the [enum Mesh.ArrayFormat] flags. For example, a mesh containing both vertices and normals would return a format of [code]3[/code] because [constant Mesh.ARRAY_FORMAT_VERTEX] is [code]1[/code] and [constant Mesh.ARRAY_FORMAT_NORMAL] is [code]2[/code].
*/
//go:nosplit
func (self class) GetFormat() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total number of vertices in [Mesh].
*/
//go:nosplit
func (self class) GetVertexCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of edges in this [Mesh].
*/
//go:nosplit
func (self class) GetEdgeCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_edge_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of faces in this [Mesh].
*/
//go:nosplit
func (self class) GetFaceCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_face_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the position of the given vertex.
*/
//go:nosplit
func (self class) SetVertex(idx gd.Int, vertex gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the given vertex.
*/
//go:nosplit
func (self class) GetVertex(idx gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the normal of the given vertex.
*/
//go:nosplit
func (self class) SetVertexNormal(idx gd.Int, normal gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, normal)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the normal of the given vertex.
*/
//go:nosplit
func (self class) GetVertexNormal(idx gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the tangent of the given vertex.
*/
//go:nosplit
func (self class) SetVertexTangent(idx gd.Int, tangent gd.Plane)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, tangent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tangent of the given vertex.
*/
//go:nosplit
func (self class) GetVertexTangent(idx gd.Int) gd.Plane {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Plane](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the UV of the given vertex.
*/
//go:nosplit
func (self class) SetVertexUv(idx gd.Int, uv gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, uv)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the UV of the given vertex.
*/
//go:nosplit
func (self class) GetVertexUv(idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the UV2 of the given vertex.
*/
//go:nosplit
func (self class) SetVertexUv2(idx gd.Int, uv2 gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, uv2)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex_uv2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the UV2 of the given vertex.
*/
//go:nosplit
func (self class) GetVertexUv2(idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_uv2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the color of the given vertex.
*/
//go:nosplit
func (self class) SetVertexColor(idx gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the color of the given vertex.
*/
//go:nosplit
func (self class) GetVertexColor(idx gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the bones of the given vertex.
*/
//go:nosplit
func (self class) SetVertexBones(idx gd.Int, bones gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(bones))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the bones of the given vertex.
*/
//go:nosplit
func (self class) GetVertexBones(ctx gd.Lifetime, idx gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone weights of the given vertex.
*/
//go:nosplit
func (self class) SetVertexWeights(idx gd.Int, weights gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(weights))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns bone weights of the given vertex.
*/
//go:nosplit
func (self class) GetVertexWeights(ctx gd.Lifetime, idx gd.Int) gd.PackedFloat32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the metadata associated with the given vertex.
*/
//go:nosplit
func (self class) SetVertexMeta(idx gd.Int, meta gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(meta))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_vertex_meta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the metadata associated with the given vertex.
*/
//go:nosplit
func (self class) GetVertexMeta(ctx gd.Lifetime, idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_meta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of edges that share the given vertex.
*/
//go:nosplit
func (self class) GetVertexEdges(ctx gd.Lifetime, idx gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_edges, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of faces that share the given vertex.
*/
//go:nosplit
func (self class) GetVertexFaces(ctx gd.Lifetime, idx gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_vertex_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns index of specified vertex connected to given edge.
Vertex argument can only be 0 or 1 because edges are comprised of two vertices.
*/
//go:nosplit
func (self class) GetEdgeVertex(idx gd.Int, vertex gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_edge_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns array of faces that touch given edge.
*/
//go:nosplit
func (self class) GetEdgeFaces(ctx gd.Lifetime, idx gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_edge_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the metadata of the given edge.
*/
//go:nosplit
func (self class) SetEdgeMeta(idx gd.Int, meta gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(meta))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_edge_meta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns meta information assigned to given edge.
*/
//go:nosplit
func (self class) GetEdgeMeta(ctx gd.Lifetime, idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_edge_meta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the specified vertex index of the given face.
Vertex argument must be either 0, 1, or 2 because faces contain three vertices.
[b]Example:[/b]
[codeblocks]
[gdscript]
var index = mesh_data_tool.get_face_vertex(0, 1) # Gets the index of the second vertex of the first face.
var position = mesh_data_tool.get_vertex(index)
var normal = mesh_data_tool.get_vertex_normal(index)
[/gdscript]
[csharp]
int index = meshDataTool.GetFaceVertex(0, 1); // Gets the index of the second vertex of the first face.
Vector3 position = meshDataTool.GetVertex(index);
Vector3 normal = meshDataTool.GetVertexNormal(index);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetFaceVertex(idx gd.Int, vertex gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_face_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns specified edge associated with given face.
Edge argument must be either 0, 1, or 2 because a face only has three edges.
*/
//go:nosplit
func (self class) GetFaceEdge(idx gd.Int, edge gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, edge)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_face_edge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the metadata of the given face.
*/
//go:nosplit
func (self class) SetFaceMeta(idx gd.Int, meta gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(meta))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_face_meta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the metadata associated with the given face.
*/
//go:nosplit
func (self class) GetFaceMeta(ctx gd.Lifetime, idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_face_meta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Calculates and returns the face normal of the given face.
*/
//go:nosplit
func (self class) GetFaceNormal(idx gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_face_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the material to be used by newly-constructed [Mesh].
*/
//go:nosplit
func (self class) SetMaterial(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the material assigned to the [Mesh].
*/
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshDataTool.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsMeshDataTool() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMeshDataTool() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("MeshDataTool", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}