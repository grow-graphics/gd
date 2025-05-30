// Code generated by the generate package DO NOT EDIT

// Package FontVariation provides methods for working with FontVariation object instances.
package FontVariation

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
import "graphics.gd/classdb/Font"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/TextServer"
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
import "graphics.gd/variant/Transform2D"

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
Provides OpenType variations, simulated bold / slant, and additional font settings like OpenType features and extra spacing.
To use simulated bold font variant:
[codeblocks]
[gdscript]
var fv = FontVariation.new()
fv.base_font = load("res://BarlowCondensed-Regular.ttf")
fv.variation_embolden = 1.2
$Label.add_theme_font_override("font", fv)
$Label.add_theme_font_size_override("font_size", 64)
[/gdscript]
[csharp]
var fv = new FontVariation();
fv.SetBaseFont(ResourceLoader.Load<FontFile>("res://BarlowCondensed-Regular.ttf"));
fv.SetVariationEmbolden(1.2);
GetNode("Label").AddThemeFontOverride("font", fv);
GetNode("Label").AddThemeFontSizeOverride("font_size", 64);
[/csharp]
[/codeblocks]
To set the coordinate of multiple variation axes:
[codeblock]
var fv = FontVariation.new();
var ts = TextServerManager.get_primary_interface()
fv.base_font = load("res://BarlowCondensed-Regular.ttf")
fv.variation_opentype = { ts.name_to_tag("wght"): 900, ts.name_to_tag("custom_hght"): 900 }
[/codeblock]
*/
type Instance [1]gdclass.FontVariation

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFontVariation() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FontVariation

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FontVariation"))
	casted := Instance{*(*gdclass.FontVariation)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BaseFont() Font.Instance {
	return Font.Instance(class(self).GetBaseFont())
}

func (self Instance) SetBaseFont(value Font.Instance) {
	class(self).SetBaseFont(value)
}

func (self Instance) VariationOpentype() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetVariationOpentype()))
}

func (self Instance) SetVariationOpentype(value map[any]any) {
	class(self).SetVariationOpentype(gd.DictionaryFromMap(value))
}

func (self Instance) VariationFaceIndex() int {
	return int(int(class(self).GetVariationFaceIndex()))
}

func (self Instance) SetVariationFaceIndex(value int) {
	class(self).SetVariationFaceIndex(int64(value))
}

func (self Instance) VariationEmbolden() Float.X {
	return Float.X(Float.X(class(self).GetVariationEmbolden()))
}

func (self Instance) SetVariationEmbolden(value Float.X) {
	class(self).SetVariationEmbolden(float64(value))
}

func (self Instance) VariationTransform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetVariationTransform())
}

func (self Instance) SetVariationTransform(value Transform2D.OriginXY) {
	class(self).SetVariationTransform(Transform2D.OriginXY(value))
}

func (self Instance) SetOpentypeFeatures(value map[any]any) {
	class(self).SetOpentypeFeatures(gd.DictionaryFromMap(value))
}

func (self Instance) SetSpacingGlyph(value int) {
	class(self).SetSpacing(0, int64(value))
}

func (self Instance) SetSpacingSpace(value int) {
	class(self).SetSpacing(1, int64(value))
}

func (self Instance) SetSpacingTop(value int) {
	class(self).SetSpacing(2, int64(value))
}

func (self Instance) SetSpacingBottom(value int) {
	class(self).SetSpacing(3, int64(value))
}

func (self Instance) BaselineOffset() Float.X {
	return Float.X(Float.X(class(self).GetBaselineOffset()))
}

func (self Instance) SetBaselineOffset(value Float.X) {
	class(self).SetBaselineOffset(float64(value))
}

//go:nosplit
func (self class) SetBaseFont(font [1]gdclass.Font) { //gd:FontVariation.set_base_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_base_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBaseFont() [1]gdclass.Font { //gd:FontVariation.get_base_font
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_base_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariationOpentype(coords Dictionary.Any) { //gd:FontVariation.set_variation_opentype
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(coords)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_variation_opentype, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariationOpentype() Dictionary.Any { //gd:FontVariation.get_variation_opentype
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_variation_opentype, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariationEmbolden(strength float64) { //gd:FontVariation.set_variation_embolden
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_variation_embolden, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariationEmbolden() float64 { //gd:FontVariation.get_variation_embolden
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_variation_embolden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariationFaceIndex(face_index int64) { //gd:FontVariation.set_variation_face_index
	var frame = callframe.New()
	callframe.Arg(frame, face_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_variation_face_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariationFaceIndex() int64 { //gd:FontVariation.get_variation_face_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_variation_face_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVariationTransform(transform Transform2D.OriginXY) { //gd:FontVariation.set_variation_transform
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_variation_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVariationTransform() Transform2D.OriginXY { //gd:FontVariation.get_variation_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[Transform2D.OriginXY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_variation_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOpentypeFeatures(features Dictionary.Any) { //gd:FontVariation.set_opentype_features
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(features)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_opentype_features, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) SetSpacing(spacing TextServer.SpacingType, value int64) { //gd:FontVariation.set_spacing
	var frame = callframe.New()
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetBaselineOffset(baseline_offset float64) { //gd:FontVariation.set_baseline_offset
	var frame = callframe.New()
	callframe.Arg(frame, baseline_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_set_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBaselineOffset() float64 { //gd:FontVariation.get_baseline_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontVariation.Bind_get_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFontVariation() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFontVariation() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsFontVariation() Instance { return self.Super().AsFontVariation() }
func (self class) AsFont() Font.Advanced             { return *((*Font.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsFont() Font.Instance     { return self.Super().AsFont() }
func (self Instance) AsFont() Font.Instance          { return *((*Font.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Font.Advanced(self.AsFont()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Font.Instance(self.AsFont()), name)
	}
}
func init() {
	gdclass.Register("FontVariation", func(ptr gd.Object) any {
		return [1]gdclass.FontVariation{*(*gdclass.FontVariation)(unsafe.Pointer(&ptr))}
	})
}
