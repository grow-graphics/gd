package Font

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Abstract base class for different font types. It has methods for drawing text and font character introspection.

*/
type Go [1]classdb.Font

/*
Returns [TextServer] RID of the font cache for specific variation.
*/
func (self Go) FindVariation(variation_coordinates gd.Dictionary) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).FindVariation(variation_coordinates, gd.Int(0), gd.Float(0.0), gd.NewTransform2D(1, 0, 0, 1, 0, 0), gd.Int(0), gd.Int(0), gd.Int(0), gd.Int(0), gd.Float(0.0)))
}

/*
Returns [Array] of valid [Font] [RID]s, which can be passed to the [TextServer] methods.
*/
func (self Go) GetRids() gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](class(self).GetRids(gc))
}

/*
Returns the total average font height (ascent plus descent) in pixels.
[b]Note:[/b] Real height of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the height of empty line).
*/
func (self Go) GetHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetHeight(gd.Int(16))))
}

/*
Returns the average font ascent (number of pixels above the baseline).
[b]Note:[/b] Real ascent of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the ascent of empty line).
*/
func (self Go) GetAscent() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetAscent(gd.Int(16))))
}

/*
Returns the average font descent (number of pixels below the baseline).
[b]Note:[/b] Real descent of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the descent of empty line).
*/
func (self Go) GetDescent() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetDescent(gd.Int(16))))
}

/*
Returns average pixel offset of the underline below the baseline.
[b]Note:[/b] Real underline position of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate.
*/
func (self Go) GetUnderlinePosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetUnderlinePosition(gd.Int(16))))
}

/*
Returns average thickness of the underline.
[b]Note:[/b] Real underline thickness of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate.
*/
func (self Go) GetUnderlineThickness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetUnderlineThickness(gd.Int(16))))
}

/*
Returns font family name.
*/
func (self Go) GetFontName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFontName(gc).String())
}

/*
Returns font style name.
*/
func (self Go) GetFontStyleName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFontStyleName(gc).String())
}

/*
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
func (self Go) GetOtNameStrings() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetOtNameStrings(gc))
}

/*
Returns font style flags, see [enum TextServer.FontStyle].
*/
func (self Go) GetFontStyle() classdb.TextServerFontStyle {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerFontStyle(class(self).GetFontStyle())
}

/*
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (self Go) GetFontWeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFontWeight()))
}

/*
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (self Go) GetFontStretch() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFontStretch()))
}

/*
Returns the spacing for the given [code]type[/code] (see [enum TextServer.SpacingType]).
*/
func (self Go) GetSpacing(spacing classdb.TextServerSpacingType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSpacing(spacing)))
}

/*
Returns a set of OpenType feature tags. More info: [url=https://docs.microsoft.com/en-us/typography/opentype/spec/featuretags]OpenType feature tags[/url].
*/
func (self Go) GetOpentypeFeatures() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetOpentypeFeatures(gc))
}

/*
Sets LRU cache capacity for [code]draw_*[/code] methods.
*/
func (self Go) SetCacheCapacity(single_line int, multi_line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCacheCapacity(gd.Int(single_line), gd.Int(multi_line))
}

/*
Returns the size of a bounding box of a single-line string, taking kerning, advance and subpixel positioning into account. See also [method get_multiline_string_size] and [method draw_string].
For example, to get the string size as displayed by a single-line Label, use:
[codeblocks]
[gdscript]
var string_size = $Label.get_theme_font("font").get_string_size($Label.text, HORIZONTAL_ALIGNMENT_LEFT, -1, $Label.get_theme_font_size("font_size"))
[/gdscript]
[csharp]
Label label = GetNode<Label>("Label");
Vector2 stringSize = label.GetThemeFont("font").GetStringSize(label.Text, HorizontalAlignment.Left, -1, label.GetThemeFontSize("font_size"));
[/csharp]
[/codeblocks]
[b]Note:[/b] Since kerning, advance and subpixel positioning are taken into account by [method get_string_size], using separate [method get_string_size] calls on substrings of a string then adding the results together will return a different result compared to using a single [method get_string_size] call on the full string.
[b]Note:[/b] Real height of the string is context-dependent and can be significantly different from the value returned by [method get_height].
*/
func (self Go) GetStringSize(text string) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetStringSize(gc.String(text), 0, gd.Float(-1), gd.Int(16), 3, 0, 0))
}

/*
Returns the size of a bounding box of a string broken into the lines, taking kerning and advance into account.
See also [method draw_multiline_string].
*/
func (self Go) GetMultilineStringSize(text string) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetMultilineStringSize(gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), 3, 3, 0, 0))
}

/*
Draw [param text] into a canvas item using the font, at a given position, with [param modulate] color, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_string].
*/
func (self Go) DrawString(canvas_item gd.RID, pos gd.Vector2, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawString(canvas_item, pos, gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Color{1, 1, 1, 1}, 3, 0, 0)
}

/*
Breaks [param text] into lines using rules specified by [param brk_flags] and draws it into a canvas item using the font, at a given position, with [param modulate] color, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline of the first line, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_multiline_string].
*/
func (self Go) DrawMultilineString(canvas_item gd.RID, pos gd.Vector2, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMultilineString(canvas_item, pos, gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), gd.Color{1, 1, 1, 1}, 3, 3, 0, 0)
}

/*
Draw [param text] outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_string_outline].
*/
func (self Go) DrawStringOutline(canvas_item gd.RID, pos gd.Vector2, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawStringOutline(canvas_item, pos, gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Int(1), gd.Color{1, 1, 1, 1}, 3, 0, 0)
}

/*
Breaks [param text] to the lines using rules specified by [param brk_flags] and draws text outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline of the first line, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_multiline_string_outline].
*/
func (self Go) DrawMultilineStringOutline(canvas_item gd.RID, pos gd.Vector2, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawMultilineStringOutline(canvas_item, pos, gc.String(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), gd.Int(1), gd.Color{1, 1, 1, 1}, 3, 3, 0, 0)
}

/*
Returns the size of a character. Does not take kerning into account.
[b]Note:[/b] Do not use this function to calculate width of the string character by character, use [method get_string_size] or [TextLine] instead. The height returned is the font height (see also [method get_height]) and has no relation to the glyph height.
*/
func (self Go) GetCharSize(char int, font_size int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetCharSize(gd.Int(char), gd.Int(font_size)))
}

/*
Draw a single Unicode character [param char] into a canvas item using the font, at a given position, with [param modulate] color. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
[b]Note:[/b] Do not use this function to draw strings character by character, use [method draw_string] or [TextLine] instead.
*/
func (self Go) DrawChar(canvas_item gd.RID, pos gd.Vector2, char int, font_size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).DrawChar(canvas_item, pos, gd.Int(char), gd.Int(font_size), gd.Color{1, 1, 1, 1})))
}

/*
Draw a single Unicode character [param char] outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
[b]Note:[/b] Do not use this function to draw strings character by character, use [method draw_string] or [TextLine] instead.
*/
func (self Go) DrawCharOutline(canvas_item gd.RID, pos gd.Vector2, char int, font_size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).DrawCharOutline(canvas_item, pos, gd.Int(char), gd.Int(font_size), gd.Int(-1), gd.Color{1, 1, 1, 1})))
}

/*
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
func (self Go) HasChar(char int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasChar(gd.Int(char)))
}

/*
Returns a string containing all the characters available in the font.
If a given character is included in more than one font data source, it appears only once in the returned string.
*/
func (self Go) GetSupportedChars() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSupportedChars(gc).String())
}

/*
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
func (self Go) IsLanguageSupported(language string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLanguageSupported(gc.String(language)))
}

/*
Returns [code]true[/code], if font supports given script ([url=https://en.wikipedia.org/wiki/ISO_15924]ISO 15924[/url] code).
*/
func (self Go) IsScriptSupported(script string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsScriptSupported(gc.String(script)))
}

/*
Returns list of OpenType features supported by font.
*/
func (self Go) GetSupportedFeatureList() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetSupportedFeatureList(gc))
}

/*
Returns list of supported [url=https://docs.microsoft.com/en-us/typography/opentype/spec/dvaraxisreg]variation coordinates[/url], each coordinate is returned as [code]tag: Vector3i(min_value,max_value,default_value)[/code].
Font variations allow for continuous change of glyph characteristics along some given design axis, such as weight, width or slant.
To print available variation axes of a variable font:
[codeblock]
var fv = FontVariation.new()
fv.base_font = load("res://RobotoFlex.ttf")
var variation_list = fv.get_supported_variation_list()
for tag in variation_list:
    var name = TextServerManager.get_primary_interface().tag_to_name(tag)
    var values = variation_list[tag]
    print("variation axis: %s (%d)\n\tmin, max, default: %s" % [name, tag, values])
[/codeblock]
[b]Note:[/b] To set and get variation coordinates of a [FontVariation], use [member FontVariation.variation_opentype].
*/
func (self Go) GetSupportedVariationList() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetSupportedVariationList(gc))
}

/*
Returns number of faces in the TrueType / OpenType collection.
*/
func (self Go) GetFaceCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFaceCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Font
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Font"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Fallbacks() gd.ArrayOf[gdclass.Font] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.Font](class(self).GetFallbacks(gc))
}

func (self Go) SetFallbacks(value gd.ArrayOf[gdclass.Font]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFallbacks(value)
}

//go:nosplit
func (self class) SetFallbacks(fallbacks gd.ArrayOf[gdclass.Font])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(fallbacks))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_set_fallbacks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFallbacks(ctx gd.Lifetime) gd.ArrayOf[gdclass.Font] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_fallbacks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Font](ret)
}
/*
Returns [TextServer] RID of the font cache for specific variation.
*/
//go:nosplit
func (self class) FindVariation(variation_coordinates gd.Dictionary, face_index gd.Int, strength gd.Float, transform gd.Transform2D, spacing_top gd.Int, spacing_bottom gd.Int, spacing_space gd.Int, spacing_glyph gd.Int, baseline_offset gd.Float) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(variation_coordinates))
	callframe.Arg(frame, face_index)
	callframe.Arg(frame, strength)
	callframe.Arg(frame, transform)
	callframe.Arg(frame, spacing_top)
	callframe.Arg(frame, spacing_bottom)
	callframe.Arg(frame, spacing_space)
	callframe.Arg(frame, spacing_glyph)
	callframe.Arg(frame, baseline_offset)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_find_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [Array] of valid [Font] [RID]s, which can be passed to the [TextServer] methods.
*/
//go:nosplit
func (self class) GetRids(ctx gd.Lifetime) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_rids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
/*
Returns the total average font height (ascent plus descent) in pixels.
[b]Note:[/b] Real height of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the height of empty line).
*/
//go:nosplit
func (self class) GetHeight(font_size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the average font ascent (number of pixels above the baseline).
[b]Note:[/b] Real ascent of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the ascent of empty line).
*/
//go:nosplit
func (self class) GetAscent(font_size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the average font descent (number of pixels below the baseline).
[b]Note:[/b] Real descent of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the descent of empty line).
*/
//go:nosplit
func (self class) GetDescent(font_size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns average pixel offset of the underline below the baseline.
[b]Note:[/b] Real underline position of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate.
*/
//go:nosplit
func (self class) GetUnderlinePosition(font_size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns average thickness of the underline.
[b]Note:[/b] Real underline thickness of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate.
*/
//go:nosplit
func (self class) GetUnderlineThickness(font_size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns font family name.
*/
//go:nosplit
func (self class) GetFontName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_font_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns font style name.
*/
//go:nosplit
func (self class) GetFontStyleName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_font_style_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
//go:nosplit
func (self class) GetOtNameStrings(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_ot_name_strings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns font style flags, see [enum TextServer.FontStyle].
*/
//go:nosplit
func (self class) GetFontStyle() classdb.TextServerFontStyle {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerFontStyle](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_font_style, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
//go:nosplit
func (self class) GetFontWeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_font_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
//go:nosplit
func (self class) GetFontStretch() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_font_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the spacing for the given [code]type[/code] (see [enum TextServer.SpacingType]).
*/
//go:nosplit
func (self class) GetSpacing(spacing classdb.TextServerSpacingType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a set of OpenType feature tags. More info: [url=https://docs.microsoft.com/en-us/typography/opentype/spec/featuretags]OpenType feature tags[/url].
*/
//go:nosplit
func (self class) GetOpentypeFeatures(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_opentype_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets LRU cache capacity for [code]draw_*[/code] methods.
*/
//go:nosplit
func (self class) SetCacheCapacity(single_line gd.Int, multi_line gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, single_line)
	callframe.Arg(frame, multi_line)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_set_cache_capacity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the size of a bounding box of a single-line string, taking kerning, advance and subpixel positioning into account. See also [method get_multiline_string_size] and [method draw_string].
For example, to get the string size as displayed by a single-line Label, use:
[codeblocks]
[gdscript]
var string_size = $Label.get_theme_font("font").get_string_size($Label.text, HORIZONTAL_ALIGNMENT_LEFT, -1, $Label.get_theme_font_size("font_size"))
[/gdscript]
[csharp]
Label label = GetNode<Label>("Label");
Vector2 stringSize = label.GetThemeFont("font").GetStringSize(label.Text, HorizontalAlignment.Left, -1, label.GetThemeFontSize("font_size"));
[/csharp]
[/codeblocks]
[b]Note:[/b] Since kerning, advance and subpixel positioning are taken into account by [method get_string_size], using separate [method get_string_size] calls on substrings of a string then adding the results together will return a different result compared to using a single [method get_string_size] call on the full string.
[b]Note:[/b] Real height of the string is context-dependent and can be significantly different from the value returned by [method get_height].
*/
//go:nosplit
func (self class) GetStringSize(text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_string_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the size of a bounding box of a string broken into the lines, taking kerning and advance into account.
See also [method draw_multiline_string].
*/
//go:nosplit
func (self class) GetMultilineStringSize(text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, brk_flags classdb.TextServerLineBreakFlag, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_multiline_string_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draw [param text] into a canvas item using the font, at a given position, with [param modulate] color, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_string].
*/
//go:nosplit
func (self class) DrawString(canvas_item gd.RID, pos gd.Vector2, text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, modulate gd.Color, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_draw_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Breaks [param text] into lines using rules specified by [param brk_flags] and draws it into a canvas item using the font, at a given position, with [param modulate] color, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline of the first line, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_multiline_string].
*/
//go:nosplit
func (self class) DrawMultilineString(canvas_item gd.RID, pos gd.Vector2, text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, modulate gd.Color, brk_flags classdb.TextServerLineBreakFlag, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_draw_multiline_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw [param text] outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_string_outline].
*/
//go:nosplit
func (self class) DrawStringOutline(canvas_item gd.RID, pos gd.Vector2, text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, size gd.Int, modulate gd.Color, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_draw_string_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Breaks [param text] to the lines using rules specified by [param brk_flags] and draws text outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline of the first line, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_multiline_string_outline].
*/
//go:nosplit
func (self class) DrawMultilineStringOutline(canvas_item gd.RID, pos gd.Vector2, text gd.String, alignment gd.HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, size gd.Int, modulate gd.Color, brk_flags classdb.TextServerLineBreakFlag, justification_flags classdb.TextServerJustificationFlag, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_draw_multiline_string_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the size of a character. Does not take kerning into account.
[b]Note:[/b] Do not use this function to calculate width of the string character by character, use [method get_string_size] or [TextLine] instead. The height returned is the font height (see also [method get_height]) and has no relation to the glyph height.
*/
//go:nosplit
func (self class) GetCharSize(char gd.Int, font_size gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, char)
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_char_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draw a single Unicode character [param char] into a canvas item using the font, at a given position, with [param modulate] color. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
[b]Note:[/b] Do not use this function to draw strings character by character, use [method draw_string] or [TextLine] instead.
*/
//go:nosplit
func (self class) DrawChar(canvas_item gd.RID, pos gd.Vector2, char gd.Int, font_size gd.Int, modulate gd.Color) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, char)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_draw_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draw a single Unicode character [param char] outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
[b]Note:[/b] Do not use this function to draw strings character by character, use [method draw_string] or [TextLine] instead.
*/
//go:nosplit
func (self class) DrawCharOutline(canvas_item gd.RID, pos gd.Vector2, char gd.Int, font_size gd.Int, size gd.Int, modulate gd.Color) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, char)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_draw_char_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
//go:nosplit
func (self class) HasChar(char gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, char)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_has_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a string containing all the characters available in the font.
If a given character is included in more than one font data source, it appears only once in the returned string.
*/
//go:nosplit
func (self class) GetSupportedChars(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_supported_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
//go:nosplit
func (self class) IsLanguageSupported(language gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_is_language_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code], if font supports given script ([url=https://en.wikipedia.org/wiki/ISO_15924]ISO 15924[/url] code).
*/
//go:nosplit
func (self class) IsScriptSupported(script gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(script))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_is_script_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns list of OpenType features supported by font.
*/
//go:nosplit
func (self class) GetSupportedFeatureList(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_supported_feature_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns list of supported [url=https://docs.microsoft.com/en-us/typography/opentype/spec/dvaraxisreg]variation coordinates[/url], each coordinate is returned as [code]tag: Vector3i(min_value,max_value,default_value)[/code].
Font variations allow for continuous change of glyph characteristics along some given design axis, such as weight, width or slant.
To print available variation axes of a variable font:
[codeblock]
var fv = FontVariation.new()
fv.base_font = load("res://RobotoFlex.ttf")
var variation_list = fv.get_supported_variation_list()
for tag in variation_list:
    var name = TextServerManager.get_primary_interface().tag_to_name(tag)
    var values = variation_list[tag]
    print("variation axis: %s (%d)\n\tmin, max, default: %s" % [name, tag, values])
[/codeblock]
[b]Note:[/b] To set and get variation coordinates of a [FontVariation], use [member FontVariation.variation_opentype].
*/
//go:nosplit
func (self class) GetSupportedVariationList(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_supported_variation_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns number of faces in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) GetFaceCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Font.Bind_get_face_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFont() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFont() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Font", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}