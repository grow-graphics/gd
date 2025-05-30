// Code generated by the generate package DO NOT EDIT

// Package Polygon2D provides methods for working with Polygon2D object instances.
package Polygon2D

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
import "graphics.gd/classdb/Texture2D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
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
A Polygon2D is defined by a set of points. Each point is connected to the next, with the final point being connected to the first, resulting in a closed polygon. Polygon2Ds can be filled with color (solid or gradient) or filled with a given texture.
*/
type Instance [1]gdclass.Polygon2D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPolygon2D() Instance
}

/*
Adds a bone with the specified [param path] and [param weights].
*/
func (self Instance) AddBone(path string, weights []float32) { //gd:Polygon2D.add_bone
	Advanced(self).AddBone(Path.ToNode(String.New(path)), Packed.New(weights...))
}

/*
Returns the number of bones in this [Polygon2D].
*/
func (self Instance) GetBoneCount() int { //gd:Polygon2D.get_bone_count
	return int(int(Advanced(self).GetBoneCount()))
}

/*
Returns the path to the node associated with the specified bone.
*/
func (self Instance) GetBonePath(index int) string { //gd:Polygon2D.get_bone_path
	return string(Advanced(self).GetBonePath(int64(index)).String())
}

/*
Returns the weight values of the specified bone.
*/
func (self Instance) GetBoneWeights(index int) []float32 { //gd:Polygon2D.get_bone_weights
	return []float32(slices.Collect(Advanced(self).GetBoneWeights(int64(index)).Values()))
}

/*
Removes the specified bone from this [Polygon2D].
*/
func (self Instance) EraseBone(index int) { //gd:Polygon2D.erase_bone
	Advanced(self).EraseBone(int64(index))
}

/*
Removes all bones from this [Polygon2D].
*/
func (self Instance) ClearBones() { //gd:Polygon2D.clear_bones
	Advanced(self).ClearBones()
}

/*
Sets the path to the node associated with the specified bone.
*/
func (self Instance) SetBonePath(index int, path string) { //gd:Polygon2D.set_bone_path
	Advanced(self).SetBonePath(int64(index), Path.ToNode(String.New(path)))
}

/*
Sets the weight values for the specified bone.
*/
func (self Instance) SetBoneWeights(index int, weights []float32) { //gd:Polygon2D.set_bone_weights
	Advanced(self).SetBoneWeights(int64(index), Packed.New(weights...))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Polygon2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Polygon2D"))
	casted := Instance{*(*gdclass.Polygon2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetColor(Color.RGBA(value))
}

func (self Instance) Offset() Vector2.XY {
	return Vector2.XY(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector2.XY) {
	class(self).SetOffset(Vector2.XY(value))
}

func (self Instance) Antialiased() bool {
	return bool(class(self).GetAntialiased())
}

func (self Instance) SetAntialiased(value bool) {
	class(self).SetAntialiased(value)
}

func (self Instance) Texture() Texture2D.Instance {
	return Texture2D.Instance(class(self).GetTexture())
}

func (self Instance) SetTexture(value Texture2D.Instance) {
	class(self).SetTexture(value)
}

func (self Instance) TextureOffset() Vector2.XY {
	return Vector2.XY(class(self).GetTextureOffset())
}

func (self Instance) SetTextureOffset(value Vector2.XY) {
	class(self).SetTextureOffset(Vector2.XY(value))
}

func (self Instance) TextureScale() Vector2.XY {
	return Vector2.XY(class(self).GetTextureScale())
}

func (self Instance) SetTextureScale(value Vector2.XY) {
	class(self).SetTextureScale(Vector2.XY(value))
}

func (self Instance) TextureRotation() Angle.Radians {
	return Angle.Radians(Float.X(class(self).GetTextureRotation()))
}

func (self Instance) SetTextureRotation(value Angle.Radians) {
	class(self).SetTextureRotation(float64(value))
}

func (self Instance) Skeleton() string {
	return string(class(self).GetSkeleton().String())
}

func (self Instance) SetSkeleton(value string) {
	class(self).SetSkeleton(Path.ToNode(String.New(value)))
}

func (self Instance) InvertEnabled() bool {
	return bool(class(self).GetInvertEnabled())
}

func (self Instance) SetInvertEnabled(value bool) {
	class(self).SetInvertEnabled(value)
}

func (self Instance) InvertBorder() Float.X {
	return Float.X(Float.X(class(self).GetInvertBorder()))
}

func (self Instance) SetInvertBorder(value Float.X) {
	class(self).SetInvertBorder(float64(value))
}

func (self Instance) Polygon() []Vector2.XY {
	return []Vector2.XY(slices.Collect(class(self).GetPolygon().Values()))
}

func (self Instance) SetPolygon(value []Vector2.XY) {
	class(self).SetPolygon(Packed.New(value...))
}

func (self Instance) Uv() []Vector2.XY {
	return []Vector2.XY(slices.Collect(class(self).GetUv().Values()))
}

func (self Instance) SetUv(value []Vector2.XY) {
	class(self).SetUv(Packed.New(value...))
}

func (self Instance) VertexColors() []Color.RGBA {
	return []Color.RGBA(slices.Collect(class(self).GetVertexColors().Values()))
}

func (self Instance) SetVertexColors(value []Color.RGBA) {
	class(self).SetVertexColors(Packed.New(value...))
}

func (self Instance) Polygons() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetPolygons())))
}

func (self Instance) SetPolygons(value []any) {
	class(self).SetPolygons(gd.EngineArrayFromSlice(value))
}

func (self Instance) InternalVertexCount() int {
	return int(int(class(self).GetInternalVertexCount()))
}

func (self Instance) SetInternalVertexCount(value int) {
	class(self).SetInternalVertexCount(int64(value))
}

//go:nosplit
func (self class) SetPolygon(polygon Packed.Array[Vector2.XY]) { //gd:Polygon2D.set_polygon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygon() Packed.Array[Vector2.XY] { //gd:Polygon2D.get_polygon
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv(uv Packed.Array[Vector2.XY]) { //gd:Polygon2D.set_uv
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](uv)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv() Packed.Array[Vector2.XY] { //gd:Polygon2D.get_uv
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color Color.RGBA) { //gd:Polygon2D.set_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() Color.RGBA { //gd:Polygon2D.get_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPolygons(polygons Array.Any) { //gd:Polygon2D.set_polygons
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(polygons)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_polygons, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygons() Array.Any { //gd:Polygon2D.get_polygons
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_polygons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertexColors(vertex_colors Packed.Array[Color.RGBA]) { //gd:Polygon2D.set_vertex_colors
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedColorArray, Color.RGBA](vertex_colors)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_vertex_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVertexColors() Packed.Array[Color.RGBA] { //gd:Polygon2D.get_vertex_colors
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_vertex_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Color.RGBA](Array.Through(gd.PackedProxy[gd.PackedColorArray, Color.RGBA]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) { //gd:Polygon2D.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D { //gd:Polygon2D.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureOffset(texture_offset Vector2.XY) { //gd:Polygon2D.set_texture_offset
	var frame = callframe.New()
	callframe.Arg(frame, texture_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_texture_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureOffset() Vector2.XY { //gd:Polygon2D.get_texture_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_texture_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureRotation(texture_rotation float64) { //gd:Polygon2D.set_texture_rotation
	var frame = callframe.New()
	callframe.Arg(frame, texture_rotation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_texture_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRotation() float64 { //gd:Polygon2D.get_texture_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_texture_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureScale(texture_scale Vector2.XY) { //gd:Polygon2D.set_texture_scale
	var frame = callframe.New()
	callframe.Arg(frame, texture_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_texture_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureScale() Vector2.XY { //gd:Polygon2D.get_texture_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_texture_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInvertEnabled(invert bool) { //gd:Polygon2D.set_invert_enabled
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_invert_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInvertEnabled() bool { //gd:Polygon2D.get_invert_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_invert_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAntialiased(antialiased bool) { //gd:Polygon2D.set_antialiased
	var frame = callframe.New()
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_antialiased, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAntialiased() bool { //gd:Polygon2D.get_antialiased
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_antialiased, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInvertBorder(invert_border float64) { //gd:Polygon2D.set_invert_border
	var frame = callframe.New()
	callframe.Arg(frame, invert_border)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_invert_border, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInvertBorder() float64 { //gd:Polygon2D.get_invert_border
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_invert_border, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset Vector2.XY) { //gd:Polygon2D.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() Vector2.XY { //gd:Polygon2D.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a bone with the specified [param path] and [param weights].
*/
//go:nosplit
func (self class) AddBone(path Path.ToNode, weights Packed.Array[float32]) { //gd:Polygon2D.add_bone
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](weights)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_add_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of bones in this [Polygon2D].
*/
//go:nosplit
func (self class) GetBoneCount() int64 { //gd:Polygon2D.get_bone_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_bone_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the path to the node associated with the specified bone.
*/
//go:nosplit
func (self class) GetBonePath(index int64) Path.ToNode { //gd:Polygon2D.get_bone_path
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_bone_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the weight values of the specified bone.
*/
//go:nosplit
func (self class) GetBoneWeights(index int64) Packed.Array[float32] { //gd:Polygon2D.get_bone_weights
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_bone_weights, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Removes the specified bone from this [Polygon2D].
*/
//go:nosplit
func (self class) EraseBone(index int64) { //gd:Polygon2D.erase_bone
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_erase_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all bones from this [Polygon2D].
*/
//go:nosplit
func (self class) ClearBones() { //gd:Polygon2D.clear_bones
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_clear_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the path to the node associated with the specified bone.
*/
//go:nosplit
func (self class) SetBonePath(index int64, path Path.ToNode) { //gd:Polygon2D.set_bone_path
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_bone_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the weight values for the specified bone.
*/
//go:nosplit
func (self class) SetBoneWeights(index int64, weights Packed.Array[float32]) { //gd:Polygon2D.set_bone_weights
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](weights)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_bone_weights, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSkeleton(skeleton Path.ToNode) { //gd:Polygon2D.set_skeleton
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(skeleton)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeleton() Path.ToNode { //gd:Polygon2D.get_skeleton
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInternalVertexCount(internal_vertex_count int64) { //gd:Polygon2D.set_internal_vertex_count
	var frame = callframe.New()
	callframe.Arg(frame, internal_vertex_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_internal_vertex_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInternalVertexCount() int64 { //gd:Polygon2D.get_internal_vertex_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_internal_vertex_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPolygon2D() Advanced             { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPolygon2D() Instance          { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsPolygon2D() Instance     { return self.Super().AsPolygon2D() }
func (self class) AsNode2D() Node2D.Advanced         { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode2D() Node2D.Instance { return self.Super().AsNode2D() }
func (self Instance) AsNode2D() Node2D.Instance      { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Polygon2D", func(ptr gd.Object) any { return [1]gdclass.Polygon2D{*(*gdclass.Polygon2D)(unsafe.Pointer(&ptr))} })
}
