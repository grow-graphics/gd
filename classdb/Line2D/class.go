// Code generated by the generate package DO NOT EDIT

// Package Line2D provides methods for working with Line2D object instances.
package Line2D

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
import "graphics.gd/classdb/Curve"
import "graphics.gd/classdb/Gradient"
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
This node draws a 2D polyline, i.e. a shape consisting of several points connected by segments. [Line2D] is not a mathematical polyline, i.e. the segments are not infinitely thin. It is intended for rendering and it can be colored and optionally textured.
[b]Warning:[/b] Certain configurations may be impossible to draw nicely, such as very sharp angles. In these situations, the node uses fallback drawing logic to look decent.
[b]Note:[/b] [Line2D] is drawn using a 2D mesh.
*/
type Instance [1]gdclass.Line2D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.Line2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLine2D() Instance
}

/*
Overwrites the position of the point at the given [param index] with the supplied [param position].
*/
func (self Instance) SetPointPosition(index int, position Vector2.XY) { //gd:Line2D.set_point_position
	Advanced(self).SetPointPosition(int64(index), Vector2.XY(position))
}

/*
Returns the position of the point at index [param index].
*/
func (self Instance) GetPointPosition(index int) Vector2.XY { //gd:Line2D.get_point_position
	return Vector2.XY(Advanced(self).GetPointPosition(int64(index)))
}

/*
Returns the number of points in the polyline.
*/
func (self Instance) GetPointCount() int { //gd:Line2D.get_point_count
	return int(int(Advanced(self).GetPointCount()))
}

/*
Adds a point with the specified [param position] relative to the polyline's own position. If no [param index] is provided, the new point will be added to the end of the points array.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. The indices of the points after the new point get increased by 1. The provided [param index] must not exceed the number of existing points in the polyline. See [method get_point_count].
*/
func (self Instance) AddPoint(position Vector2.XY) { //gd:Line2D.add_point
	Advanced(self).AddPoint(Vector2.XY(position), int64(-1))
}

/*
Adds a point with the specified [param position] relative to the polyline's own position. If no [param index] is provided, the new point will be added to the end of the points array.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. The indices of the points after the new point get increased by 1. The provided [param index] must not exceed the number of existing points in the polyline. See [method get_point_count].
*/
func (self Expanded) AddPoint(position Vector2.XY, index int) { //gd:Line2D.add_point
	Advanced(self).AddPoint(Vector2.XY(position), int64(index))
}

/*
Removes the point at index [param index] from the polyline.
*/
func (self Instance) RemovePoint(index int) { //gd:Line2D.remove_point
	Advanced(self).RemovePoint(int64(index))
}

/*
Removes all points from the polyline, making it empty.
*/
func (self Instance) ClearPoints() { //gd:Line2D.clear_points
	Advanced(self).ClearPoints()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Line2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Line2D"))
	casted := Instance{*(*gdclass.Line2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Points() []Vector2.XY {
	return []Vector2.XY(slices.Collect(class(self).GetPoints().Values()))
}

func (self Instance) SetPoints(value []Vector2.XY) {
	class(self).SetPoints(Packed.New(value...))
}

func (self Instance) Closed() bool {
	return bool(class(self).IsClosed())
}

func (self Instance) SetClosed(value bool) {
	class(self).SetClosed(value)
}

func (self Instance) Width() Float.X {
	return Float.X(Float.X(class(self).GetWidth()))
}

func (self Instance) SetWidth(value Float.X) {
	class(self).SetWidth(float64(value))
}

func (self Instance) WidthCurve() Curve.Instance {
	return Curve.Instance(class(self).GetCurve())
}

func (self Instance) SetWidthCurve(value Curve.Instance) {
	class(self).SetCurve(value)
}

func (self Instance) DefaultColor() Color.RGBA {
	return Color.RGBA(class(self).GetDefaultColor())
}

func (self Instance) SetDefaultColor(value Color.RGBA) {
	class(self).SetDefaultColor(Color.RGBA(value))
}

func (self Instance) Gradient() Gradient.Instance {
	return Gradient.Instance(class(self).GetGradient())
}

func (self Instance) SetGradient(value Gradient.Instance) {
	class(self).SetGradient(value)
}

func (self Instance) Texture() Texture2D.Instance {
	return Texture2D.Instance(class(self).GetTexture())
}

func (self Instance) SetTexture(value Texture2D.Instance) {
	class(self).SetTexture(value)
}

func (self Instance) TextureMode() LineTextureMode {
	return LineTextureMode(class(self).GetTextureMode())
}

func (self Instance) SetTextureMode(value LineTextureMode) {
	class(self).SetTextureMode(value)
}

func (self Instance) JointMode() LineJointMode {
	return LineJointMode(class(self).GetJointMode())
}

func (self Instance) SetJointMode(value LineJointMode) {
	class(self).SetJointMode(value)
}

func (self Instance) BeginCapMode() LineCapMode {
	return LineCapMode(class(self).GetBeginCapMode())
}

func (self Instance) SetBeginCapMode(value LineCapMode) {
	class(self).SetBeginCapMode(value)
}

func (self Instance) EndCapMode() LineCapMode {
	return LineCapMode(class(self).GetEndCapMode())
}

func (self Instance) SetEndCapMode(value LineCapMode) {
	class(self).SetEndCapMode(value)
}

func (self Instance) SharpLimit() Float.X {
	return Float.X(Float.X(class(self).GetSharpLimit()))
}

func (self Instance) SetSharpLimit(value Float.X) {
	class(self).SetSharpLimit(float64(value))
}

func (self Instance) RoundPrecision() int {
	return int(int(class(self).GetRoundPrecision()))
}

func (self Instance) SetRoundPrecision(value int) {
	class(self).SetRoundPrecision(int64(value))
}

func (self Instance) Antialiased() bool {
	return bool(class(self).GetAntialiased())
}

func (self Instance) SetAntialiased(value bool) {
	class(self).SetAntialiased(value)
}

//go:nosplit
func (self class) SetPoints(points Packed.Array[Vector2.XY]) { //gd:Line2D.set_points
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](points)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPoints() Packed.Array[Vector2.XY] { //gd:Line2D.get_points
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Overwrites the position of the point at the given [param index] with the supplied [param position].
*/
//go:nosplit
func (self class) SetPointPosition(index int64, position Vector2.XY) { //gd:Line2D.set_point_position
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the point at index [param index].
*/
//go:nosplit
func (self class) GetPointPosition(index int64) Vector2.XY { //gd:Line2D.get_point_position
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of points in the polyline.
*/
//go:nosplit
func (self class) GetPointCount() int64 { //gd:Line2D.get_point_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a point with the specified [param position] relative to the polyline's own position. If no [param index] is provided, the new point will be added to the end of the points array.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. The indices of the points after the new point get increased by 1. The provided [param index] must not exceed the number of existing points in the polyline. See [method get_point_count].
*/
//go:nosplit
func (self class) AddPoint(position Vector2.XY, index int64) { //gd:Line2D.add_point
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the point at index [param index] from the polyline.
*/
//go:nosplit
func (self class) RemovePoint(index int64) { //gd:Line2D.remove_point
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all points from the polyline, making it empty.
*/
//go:nosplit
func (self class) ClearPoints() { //gd:Line2D.clear_points
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_clear_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetClosed(closed bool) { //gd:Line2D.set_closed
	var frame = callframe.New()
	callframe.Arg(frame, closed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_closed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsClosed() bool { //gd:Line2D.is_closed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_is_closed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(width float64) { //gd:Line2D.set_width
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWidth() float64 { //gd:Line2D.get_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurve(curve [1]gdclass.Curve) { //gd:Line2D.set_curve
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() [1]gdclass.Curve { //gd:Line2D.get_curve
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultColor(color Color.RGBA) { //gd:Line2D.set_default_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_default_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultColor() Color.RGBA { //gd:Line2D.get_default_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_default_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGradient(color [1]gdclass.Gradient) { //gd:Line2D.set_gradient
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(color[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_gradient, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGradient() [1]gdclass.Gradient { //gd:Line2D.get_gradient
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_gradient, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Gradient{gd.PointerWithOwnershipTransferredToGo[gdclass.Gradient](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) { //gd:Line2D.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D { //gd:Line2D.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureMode(mode LineTextureMode) { //gd:Line2D.set_texture_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_texture_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureMode() LineTextureMode { //gd:Line2D.get_texture_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[LineTextureMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_texture_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointMode(mode LineJointMode) { //gd:Line2D.set_joint_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_joint_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointMode() LineJointMode { //gd:Line2D.get_joint_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[LineJointMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_joint_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBeginCapMode(mode LineCapMode) { //gd:Line2D.set_begin_cap_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_begin_cap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBeginCapMode() LineCapMode { //gd:Line2D.get_begin_cap_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[LineCapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_begin_cap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEndCapMode(mode LineCapMode) { //gd:Line2D.set_end_cap_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_end_cap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEndCapMode() LineCapMode { //gd:Line2D.get_end_cap_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[LineCapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_end_cap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSharpLimit(limit float64) { //gd:Line2D.set_sharp_limit
	var frame = callframe.New()
	callframe.Arg(frame, limit)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_sharp_limit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSharpLimit() float64 { //gd:Line2D.get_sharp_limit
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_sharp_limit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoundPrecision(precision int64) { //gd:Line2D.set_round_precision
	var frame = callframe.New()
	callframe.Arg(frame, precision)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_round_precision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoundPrecision() int64 { //gd:Line2D.get_round_precision
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_round_precision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAntialiased(antialiased bool) { //gd:Line2D.set_antialiased
	var frame = callframe.New()
	callframe.Arg(frame, antialiased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_antialiased, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAntialiased() bool { //gd:Line2D.get_antialiased
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_antialiased, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLine2D() Advanced                { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLine2D() Instance             { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsLine2D() Instance        { return self.Super().AsLine2D() }
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
	gdclass.Register("Line2D", func(ptr gd.Object) any { return [1]gdclass.Line2D{*(*gdclass.Line2D)(unsafe.Pointer(&ptr))} })
}

type LineJointMode int //gd:Line2D.LineJointMode

const (
	/*Makes the polyline's joints pointy, connecting the sides of the two segments by extending them until they intersect. If the rotation of a joint is too big (based on [member sharp_limit]), the joint falls back to [constant LINE_JOINT_BEVEL] to prevent very long miters.*/
	LineJointSharp LineJointMode = 0
	/*Makes the polyline's joints bevelled/chamfered, connecting the sides of the two segments with a simple line.*/
	LineJointBevel LineJointMode = 1
	/*Makes the polyline's joints rounded, connecting the sides of the two segments with an arc. The detail of this arc depends on [member round_precision].*/
	LineJointRound LineJointMode = 2
)

type LineCapMode int //gd:Line2D.LineCapMode

const (
	/*Draws no line cap.*/
	LineCapNone LineCapMode = 0
	/*Draws the line cap as a box, slightly extending the first/last segment.*/
	LineCapBox LineCapMode = 1
	/*Draws the line cap as a semicircle attached to the first/last segment.*/
	LineCapRound LineCapMode = 2
)

type LineTextureMode int //gd:Line2D.LineTextureMode

const (
	/*Takes the left pixels of the texture and renders them over the whole polyline.*/
	LineTextureNone LineTextureMode = 0
	/*Tiles the texture over the polyline. [member CanvasItem.texture_repeat] of the [Line2D] node must be [constant CanvasItem.TEXTURE_REPEAT_ENABLED] or [constant CanvasItem.TEXTURE_REPEAT_MIRROR] for it to work properly.*/
	LineTextureTile LineTextureMode = 1
	/*Stretches the texture across the polyline. [member CanvasItem.texture_repeat] of the [Line2D] node must be [constant CanvasItem.TEXTURE_REPEAT_DISABLED] for best results.*/
	LineTextureStretch LineTextureMode = 2
)
