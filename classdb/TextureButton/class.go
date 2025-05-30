// Code generated by the generate package DO NOT EDIT

// Package TextureButton provides methods for working with TextureButton object instances.
package TextureButton

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
import "graphics.gd/classdb/BaseButton"
import "graphics.gd/classdb/BitMap"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Texture2D"
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
[TextureButton] has the same functionality as [Button], except it uses sprites instead of Godot's [Theme] resource. It is faster to create, but it doesn't support localization like more complex [Control]s.
See also [BaseButton] which contains common properties and methods associated with this node.
[b]Note:[/b] Setting a texture for the "normal" state ([member texture_normal]) is recommended. If [member texture_normal] is not set, the [TextureButton] will still receive input events and be clickable, but the user will not be able to see it unless they activate another one of its states with a texture assigned (e.g., hover over it to show [member texture_hover]).
*/
type Instance [1]gdclass.TextureButton

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTextureButton() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextureButton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextureButton"))
	casted := Instance{*(*gdclass.TextureButton)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TextureNormal() Texture2D.Instance {
	return Texture2D.Instance(class(self).GetTextureNormal())
}

func (self Instance) SetTextureNormal(value Texture2D.Instance) {
	class(self).SetTextureNormal(value)
}

func (self Instance) TexturePressed() Texture2D.Instance {
	return Texture2D.Instance(class(self).GetTexturePressed())
}

func (self Instance) SetTexturePressed(value Texture2D.Instance) {
	class(self).SetTexturePressed(value)
}

func (self Instance) TextureHover() Texture2D.Instance {
	return Texture2D.Instance(class(self).GetTextureHover())
}

func (self Instance) SetTextureHover(value Texture2D.Instance) {
	class(self).SetTextureHover(value)
}

func (self Instance) TextureDisabled() Texture2D.Instance {
	return Texture2D.Instance(class(self).GetTextureDisabled())
}

func (self Instance) SetTextureDisabled(value Texture2D.Instance) {
	class(self).SetTextureDisabled(value)
}

func (self Instance) TextureFocused() Texture2D.Instance {
	return Texture2D.Instance(class(self).GetTextureFocused())
}

func (self Instance) SetTextureFocused(value Texture2D.Instance) {
	class(self).SetTextureFocused(value)
}

func (self Instance) TextureClickMask() BitMap.Instance {
	return BitMap.Instance(class(self).GetClickMask())
}

func (self Instance) SetTextureClickMask(value BitMap.Instance) {
	class(self).SetClickMask(value)
}

func (self Instance) IgnoreTextureSize() bool {
	return bool(class(self).GetIgnoreTextureSize())
}

func (self Instance) SetIgnoreTextureSize(value bool) {
	class(self).SetIgnoreTextureSize(value)
}

func (self Instance) StretchMode() StretchMode {
	return StretchMode(class(self).GetStretchMode())
}

func (self Instance) SetStretchMode(value StretchMode) {
	class(self).SetStretchMode(value)
}

func (self Instance) FlipH() bool {
	return bool(class(self).IsFlippedH())
}

func (self Instance) SetFlipH(value bool) {
	class(self).SetFlipH(value)
}

func (self Instance) FlipV() bool {
	return bool(class(self).IsFlippedV())
}

func (self Instance) SetFlipV(value bool) {
	class(self).SetFlipV(value)
}

//go:nosplit
func (self class) SetTextureNormal(texture [1]gdclass.Texture2D) { //gd:TextureButton.set_texture_normal
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTexturePressed(texture [1]gdclass.Texture2D) { //gd:TextureButton.set_texture_pressed
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureHover(texture [1]gdclass.Texture2D) { //gd:TextureButton.set_texture_hover
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_hover, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureDisabled(texture [1]gdclass.Texture2D) { //gd:TextureButton.set_texture_disabled
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureFocused(texture [1]gdclass.Texture2D) { //gd:TextureButton.set_texture_focused
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_focused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetClickMask(mask [1]gdclass.BitMap) { //gd:TextureButton.set_click_mask
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mask[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_click_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetIgnoreTextureSize(ignore bool) { //gd:TextureButton.set_ignore_texture_size
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_ignore_texture_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetStretchMode(mode StretchMode) { //gd:TextureButton.set_stretch_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFlipH(enable bool) { //gd:TextureButton.set_flip_h
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedH() bool { //gd:TextureButton.is_flipped_h
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(enable bool) { //gd:TextureButton.set_flip_v
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedV() bool { //gd:TextureButton.is_flipped_v
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureNormal() [1]gdclass.Texture2D { //gd:TextureButton.get_texture_normal
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTexturePressed() [1]gdclass.Texture2D { //gd:TextureButton.get_texture_pressed
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureHover() [1]gdclass.Texture2D { //gd:TextureButton.get_texture_hover
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_hover, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureDisabled() [1]gdclass.Texture2D { //gd:TextureButton.get_texture_disabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureFocused() [1]gdclass.Texture2D { //gd:TextureButton.get_texture_focused
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_focused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetClickMask() [1]gdclass.BitMap { //gd:TextureButton.get_click_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_click_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.BitMap{gd.PointerWithOwnershipTransferredToGo[gdclass.BitMap](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetIgnoreTextureSize() bool { //gd:TextureButton.get_ignore_texture_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_ignore_texture_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetStretchMode() StretchMode { //gd:TextureButton.get_stretch_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[StretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextureButton() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextureButton() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsTextureButton() Instance { return self.Super().AsTextureButton() }
func (self class) AsBaseButton() BaseButton.Advanced {
	return *((*BaseButton.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsBaseButton() BaseButton.Instance { return self.Super().AsBaseButton() }
func (self Instance) AsBaseButton() BaseButton.Instance {
	return *((*BaseButton.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced         { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsControl() Control.Instance { return self.Super().AsControl() }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(BaseButton.Advanced(self.AsBaseButton()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(BaseButton.Instance(self.AsBaseButton()), name)
	}
}
func init() {
	gdclass.Register("TextureButton", func(ptr gd.Object) any {
		return [1]gdclass.TextureButton{*(*gdclass.TextureButton)(unsafe.Pointer(&ptr))}
	})
}

type StretchMode int //gd:TextureButton.StretchMode

const (
	/*Scale to fit the node's bounding rectangle.*/
	StretchScale StretchMode = 0
	/*Tile inside the node's bounding rectangle.*/
	StretchTile StretchMode = 1
	/*The texture keeps its original size and stays in the bounding rectangle's top-left corner.*/
	StretchKeep StretchMode = 2
	/*The texture keeps its original size and stays centered in the node's bounding rectangle.*/
	StretchKeepCentered StretchMode = 3
	/*Scale the texture to fit the node's bounding rectangle, but maintain the texture's aspect ratio.*/
	StretchKeepAspect StretchMode = 4
	/*Scale the texture to fit the node's bounding rectangle, center it, and maintain its aspect ratio.*/
	StretchKeepAspectCentered StretchMode = 5
	/*Scale the texture so that the shorter side fits the bounding rectangle. The other side clips to the node's limits.*/
	StretchKeepAspectCovered StretchMode = 6
)
