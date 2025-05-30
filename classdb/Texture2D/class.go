// Code generated by the generate package DO NOT EDIT

// Package Texture2D provides methods for working with Texture2D object instances.
package Texture2D

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
import "graphics.gd/classdb/Image"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Texture"
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
import "graphics.gd/variant/Rect2"
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
A texture works by registering an image in the video hardware, which then can be used in 3D models or 2D [Sprite2D] or GUI [Control].
Textures are often created by loading them from a file. See [method @GDScript.load].
[Texture2D] is a base for other resources. It cannot be used directly.
[b]Note:[/b] The maximum texture size is 16384×16384 pixels due to graphics hardware limitations. Larger textures may fail to import.

	See [Interface] for methods that can be overridden by a [Class] that extends it.
*/
type Instance [1]gdclass.Texture2D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.Texture2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTexture2D() Instance
}
type Interface interface {
	//Called when the [Texture2D]'s width is queried.
	GetWidth() int
	//Called when the [Texture2D]'s height is queried.
	GetHeight() int
	//Called when a pixel's opaque state in the [Texture2D] is queried at the specified [code](x, y)[/code] position.
	IsPixelOpaque(x int, y int) bool
	//Called when the presence of an alpha channel in the [Texture2D] is queried.
	HasAlpha() bool
	//Called when the entire [Texture2D] is requested to be drawn over a [CanvasItem], with the top-left offset specified in [param pos]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
	//[b]Note:[/b] This is only used in 2D rendering, not 3D.
	Draw(to_canvas_item RID.Any, pos Vector2.XY, modulate Color.RGBA, transpose bool)
	//Called when the [Texture2D] is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
	//[b]Note:[/b] This is only used in 2D rendering, not 3D.
	DrawRect(to_canvas_item RID.Any, rect Rect2.PositionSize, tile bool, modulate Color.RGBA, transpose bool)
	//Called when a part of the [Texture2D] specified by [param src_rect]'s coordinates is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
	//[b]Note:[/b] This is only used in 2D rendering, not 3D.
	DrawRectRegion(to_canvas_item RID.Any, rect Rect2.PositionSize, src_rect Rect2.PositionSize, modulate Color.RGBA, transpose bool, clip_uv bool)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetWidth() (_ int)                   { return }
func (self implementation) GetHeight() (_ int)                  { return }
func (self implementation) IsPixelOpaque(x int, y int) (_ bool) { return }
func (self implementation) HasAlpha() (_ bool)                  { return }
func (self implementation) Draw(to_canvas_item RID.Any, pos Vector2.XY, modulate Color.RGBA, transpose bool) {
	return
}
func (self implementation) DrawRect(to_canvas_item RID.Any, rect Rect2.PositionSize, tile bool, modulate Color.RGBA, transpose bool) {
	return
}
func (self implementation) DrawRectRegion(to_canvas_item RID.Any, rect Rect2.PositionSize, src_rect Rect2.PositionSize, modulate Color.RGBA, transpose bool, clip_uv bool) {
	return
}

/*
Called when the [Texture2D]'s width is queried.
*/
func (Instance) _get_width(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Called when the [Texture2D]'s height is queried.
*/
func (Instance) _get_height(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Called when a pixel's opaque state in the [Texture2D] is queried at the specified [code](x, y)[/code] position.
*/
func (Instance) _is_pixel_opaque(impl func(ptr unsafe.Pointer, x int, y int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var x = gd.UnsafeGet[int64](p_args, 0)
		var y = gd.UnsafeGet[int64](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(x), int(y))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the presence of an alpha channel in the [Texture2D] is queried.
*/
func (Instance) _has_alpha(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the entire [Texture2D] is requested to be drawn over a [CanvasItem], with the top-left offset specified in [param pos]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (Instance) _draw(impl func(ptr unsafe.Pointer, to_canvas_item RID.Any, pos Vector2.XY, modulate Color.RGBA, transpose bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var to_canvas_item = gd.UnsafeGet[RID.Any](p_args, 0)
		var pos = gd.UnsafeGet[Vector2.XY](p_args, 1)
		var modulate = gd.UnsafeGet[Color.RGBA](p_args, 2)
		var transpose = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, pos, modulate, transpose)
	}
}

/*
Called when the [Texture2D] is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (Instance) _draw_rect(impl func(ptr unsafe.Pointer, to_canvas_item RID.Any, rect Rect2.PositionSize, tile bool, modulate Color.RGBA, transpose bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var to_canvas_item = gd.UnsafeGet[RID.Any](p_args, 0)
		var rect = gd.UnsafeGet[Rect2.PositionSize](p_args, 1)
		var tile = gd.UnsafeGet[bool](p_args, 2)
		var modulate = gd.UnsafeGet[Color.RGBA](p_args, 3)
		var transpose = gd.UnsafeGet[bool](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, rect, tile, modulate, transpose)
	}
}

/*
Called when a part of the [Texture2D] specified by [param src_rect]'s coordinates is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (Instance) _draw_rect_region(impl func(ptr unsafe.Pointer, to_canvas_item RID.Any, rect Rect2.PositionSize, src_rect Rect2.PositionSize, modulate Color.RGBA, transpose bool, clip_uv bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var to_canvas_item = gd.UnsafeGet[RID.Any](p_args, 0)
		var rect = gd.UnsafeGet[Rect2.PositionSize](p_args, 1)
		var src_rect = gd.UnsafeGet[Rect2.PositionSize](p_args, 2)
		var modulate = gd.UnsafeGet[Color.RGBA](p_args, 3)
		var transpose = gd.UnsafeGet[bool](p_args, 4)
		var clip_uv = gd.UnsafeGet[bool](p_args, 5)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, rect, src_rect, modulate, transpose, clip_uv)
	}
}

/*
Returns the texture width in pixels.
*/
func (self Instance) GetWidth() int { //gd:Texture2D.get_width
	return int(int(Advanced(self).GetWidth()))
}

/*
Returns the texture height in pixels.
*/
func (self Instance) GetHeight() int { //gd:Texture2D.get_height
	return int(int(Advanced(self).GetHeight()))
}

/*
Returns the texture size in pixels.
*/
func (self Instance) GetSize() Vector2.XY { //gd:Texture2D.get_size
	return Vector2.XY(Advanced(self).GetSize())
}

/*
Returns [code]true[/code] if this [Texture2D] has an alpha channel.
*/
func (self Instance) HasAlpha() bool { //gd:Texture2D.has_alpha
	return bool(Advanced(self).HasAlpha())
}

/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API at the specified [param position].
*/
func (self Instance) Draw(canvas_item RID.CanvasItem, position Vector2.XY) { //gd:Texture2D.draw
	Advanced(self).Draw(RID.Any(canvas_item), Vector2.XY(position), Color.RGBA(gd.Color{1, 1, 1, 1}), false)
}

/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API at the specified [param position].
*/
func (self Expanded) Draw(canvas_item RID.CanvasItem, position Vector2.XY, modulate Color.RGBA, transpose bool) { //gd:Texture2D.draw
	Advanced(self).Draw(RID.Any(canvas_item), Vector2.XY(position), Color.RGBA(modulate), transpose)
}

/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API.
*/
func (self Instance) DrawRect(canvas_item RID.CanvasItem, rect Rect2.PositionSize, tile bool) { //gd:Texture2D.draw_rect
	Advanced(self).DrawRect(RID.Any(canvas_item), Rect2.PositionSize(rect), tile, Color.RGBA(gd.Color{1, 1, 1, 1}), false)
}

/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API.
*/
func (self Expanded) DrawRect(canvas_item RID.CanvasItem, rect Rect2.PositionSize, tile bool, modulate Color.RGBA, transpose bool) { //gd:Texture2D.draw_rect
	Advanced(self).DrawRect(RID.Any(canvas_item), Rect2.PositionSize(rect), tile, Color.RGBA(modulate), transpose)
}

/*
Draws a part of the texture using a [CanvasItem] with the [RenderingServer] API.
*/
func (self Instance) DrawRectRegion(canvas_item RID.CanvasItem, rect Rect2.PositionSize, src_rect Rect2.PositionSize) { //gd:Texture2D.draw_rect_region
	Advanced(self).DrawRectRegion(RID.Any(canvas_item), Rect2.PositionSize(rect), Rect2.PositionSize(src_rect), Color.RGBA(gd.Color{1, 1, 1, 1}), false, true)
}

/*
Draws a part of the texture using a [CanvasItem] with the [RenderingServer] API.
*/
func (self Expanded) DrawRectRegion(canvas_item RID.CanvasItem, rect Rect2.PositionSize, src_rect Rect2.PositionSize, modulate Color.RGBA, transpose bool, clip_uv bool) { //gd:Texture2D.draw_rect_region
	Advanced(self).DrawRectRegion(RID.Any(canvas_item), Rect2.PositionSize(rect), Rect2.PositionSize(src_rect), Color.RGBA(modulate), transpose, clip_uv)
}

/*
Returns an [Image] that is a copy of data from this [Texture2D] (a new [Image] is created each time). [Image]s can be accessed and manipulated directly.
[b]Note:[/b] This will return [code]null[/code] if this [Texture2D] is invalid.
[b]Note:[/b] This will fetch the texture data from the GPU, which might cause performance problems when overused. Avoid calling [method get_image] every frame, especially on large textures.
*/
func (self Instance) GetImage() Image.Instance { //gd:Texture2D.get_image
	return Image.Instance(Advanced(self).GetImage())
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture2D]).
*/
func (self Instance) CreatePlaceholder() Resource.Instance { //gd:Texture2D.create_placeholder
	return Resource.Instance(Advanced(self).CreatePlaceholder())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Texture2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture2D"))
	casted := Instance{*(*gdclass.Texture2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Called when the [Texture2D]'s width is queried.
*/
func (class) _get_width(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [Texture2D]'s height is queried.
*/
func (class) _get_height(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when a pixel's opaque state in the [Texture2D] is queried at the specified [code](x, y)[/code] position.
*/
func (class) _is_pixel_opaque(impl func(ptr unsafe.Pointer, x int64, y int64) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var x = gd.UnsafeGet[int64](p_args, 0)
		var y = gd.UnsafeGet[int64](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, x, y)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the presence of an alpha channel in the [Texture2D] is queried.
*/
func (class) _has_alpha(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the entire [Texture2D] is requested to be drawn over a [CanvasItem], with the top-left offset specified in [param pos]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (class) _draw(impl func(ptr unsafe.Pointer, to_canvas_item RID.Any, pos Vector2.XY, modulate Color.RGBA, transpose bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var to_canvas_item = gd.UnsafeGet[RID.Any](p_args, 0)
		var pos = gd.UnsafeGet[Vector2.XY](p_args, 1)
		var modulate = gd.UnsafeGet[Color.RGBA](p_args, 2)
		var transpose = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, pos, modulate, transpose)
	}
}

/*
Called when the [Texture2D] is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (class) _draw_rect(impl func(ptr unsafe.Pointer, to_canvas_item RID.Any, rect Rect2.PositionSize, tile bool, modulate Color.RGBA, transpose bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var to_canvas_item = gd.UnsafeGet[RID.Any](p_args, 0)
		var rect = gd.UnsafeGet[Rect2.PositionSize](p_args, 1)
		var tile = gd.UnsafeGet[bool](p_args, 2)
		var modulate = gd.UnsafeGet[Color.RGBA](p_args, 3)
		var transpose = gd.UnsafeGet[bool](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, rect, tile, modulate, transpose)
	}
}

/*
Called when a part of the [Texture2D] specified by [param src_rect]'s coordinates is requested to be drawn onto [CanvasItem]'s specified [param rect]. [param modulate] specifies a multiplier for the colors being drawn, while [param transpose] specifies whether drawing should be performed in column-major order instead of row-major order (resulting in 90-degree clockwise rotation).
[b]Note:[/b] This is only used in 2D rendering, not 3D.
*/
func (class) _draw_rect_region(impl func(ptr unsafe.Pointer, to_canvas_item RID.Any, rect Rect2.PositionSize, src_rect Rect2.PositionSize, modulate Color.RGBA, transpose bool, clip_uv bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var to_canvas_item = gd.UnsafeGet[RID.Any](p_args, 0)
		var rect = gd.UnsafeGet[Rect2.PositionSize](p_args, 1)
		var src_rect = gd.UnsafeGet[Rect2.PositionSize](p_args, 2)
		var modulate = gd.UnsafeGet[Color.RGBA](p_args, 3)
		var transpose = gd.UnsafeGet[bool](p_args, 4)
		var clip_uv = gd.UnsafeGet[bool](p_args, 5)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, to_canvas_item, rect, src_rect, modulate, transpose, clip_uv)
	}
}

/*
Returns the texture width in pixels.
*/
//go:nosplit
func (self class) GetWidth() int64 { //gd:Texture2D.get_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the texture height in pixels.
*/
//go:nosplit
func (self class) GetHeight() int64 { //gd:Texture2D.get_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the texture size in pixels.
*/
//go:nosplit
func (self class) GetSize() Vector2.XY { //gd:Texture2D.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this [Texture2D] has an alpha channel.
*/
//go:nosplit
func (self class) HasAlpha() bool { //gd:Texture2D.has_alpha
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_has_alpha, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API at the specified [param position].
*/
//go:nosplit
func (self class) Draw(canvas_item RID.Any, position Vector2.XY, modulate Color.RGBA, transpose bool) { //gd:Texture2D.draw
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, position)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws the texture using a [CanvasItem] with the [RenderingServer] API.
*/
//go:nosplit
func (self class) DrawRect(canvas_item RID.Any, rect Rect2.PositionSize, tile bool, modulate Color.RGBA, transpose bool) { //gd:Texture2D.draw_rect
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, tile)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_draw_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws a part of the texture using a [CanvasItem] with the [RenderingServer] API.
*/
//go:nosplit
func (self class) DrawRectRegion(canvas_item RID.Any, rect Rect2.PositionSize, src_rect Rect2.PositionSize, modulate Color.RGBA, transpose bool, clip_uv bool) { //gd:Texture2D.draw_rect_region
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, rect)
	callframe.Arg(frame, src_rect)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, transpose)
	callframe.Arg(frame, clip_uv)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_draw_rect_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an [Image] that is a copy of data from this [Texture2D] (a new [Image] is created each time). [Image]s can be accessed and manipulated directly.
[b]Note:[/b] This will return [code]null[/code] if this [Texture2D] is invalid.
[b]Note:[/b] This will fetch the texture data from the GPU, which might cause performance problems when overused. Avoid calling [method get_image] every frame, especially on large textures.
*/
//go:nosplit
func (self class) GetImage() [1]gdclass.Image { //gd:Texture2D.get_image
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_get_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a placeholder version of this resource ([PlaceholderTexture2D]).
*/
//go:nosplit
func (self class) CreatePlaceholder() [1]gdclass.Resource { //gd:Texture2D.create_placeholder
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture2D.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsTexture2D() Advanced               { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture2D() Instance            { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsTexture2D() Instance       { return self.Super().AsTexture2D() }
func (self class) AsTexture() Texture.Advanced         { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsTexture() Texture.Instance { return self.Super().AsTexture() }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
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
	case "_get_width":
		return reflect.ValueOf(self._get_width)
	case "_get_height":
		return reflect.ValueOf(self._get_height)
	case "_is_pixel_opaque":
		return reflect.ValueOf(self._is_pixel_opaque)
	case "_has_alpha":
		return reflect.ValueOf(self._has_alpha)
	case "_draw":
		return reflect.ValueOf(self._draw)
	case "_draw_rect":
		return reflect.ValueOf(self._draw_rect)
	case "_draw_rect_region":
		return reflect.ValueOf(self._draw_rect_region)
	default:
		return gd.VirtualByName(Texture.Advanced(self.AsTexture()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_width":
		return reflect.ValueOf(self._get_width)
	case "_get_height":
		return reflect.ValueOf(self._get_height)
	case "_is_pixel_opaque":
		return reflect.ValueOf(self._is_pixel_opaque)
	case "_has_alpha":
		return reflect.ValueOf(self._has_alpha)
	case "_draw":
		return reflect.ValueOf(self._draw)
	case "_draw_rect":
		return reflect.ValueOf(self._draw_rect)
	case "_draw_rect_region":
		return reflect.ValueOf(self._draw_rect_region)
	default:
		return gd.VirtualByName(Texture.Instance(self.AsTexture()), name)
	}
}
func init() {
	gdclass.Register("Texture2D", func(ptr gd.Object) any { return [1]gdclass.Texture2D{*(*gdclass.Texture2D)(unsafe.Pointer(&ptr))} })
}
