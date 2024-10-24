package TextServer

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
[TextServer] is the API backend for managing fonts and rendering text.
[b]Note:[/b] This is a low-level API, consider using [TextLine], [TextParagraph], and [Font] classes instead.
This is an abstract class, so to get the currently active [TextServer] instance, use the following code:
[codeblocks]
[gdscript]
var ts = TextServerManager.get_primary_interface()
[/gdscript]
[csharp]
var ts = TextServerManager.GetPrimaryInterface();
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.TextServer

/*
Returns [code]true[/code] if the server supports a feature.
*/
func (self Go) HasFeature(feature classdb.TextServerFeature) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasFeature(feature))
}

/*
Returns the name of the server interface.
*/
func (self Go) GetName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetName(gc).String())
}

/*
Returns text server features, see [enum Feature].
*/
func (self Go) GetFeatures() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFeatures()))
}

/*
Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
[b]Note:[/b] This function should be called before any other TextServer functions used, otherwise it won't have any effect.
*/
func (self Go) LoadSupportData(filename string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).LoadSupportData(gc.String(filename)))
}

/*
Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
*/
func (self Go) GetSupportDataFilename() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSupportDataFilename(gc).String())
}

/*
Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
*/
func (self Go) GetSupportDataInfo() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSupportDataInfo(gc).String())
}

/*
Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
[b]Note:[/b] This function is used by during project export, to include TextServer database.
*/
func (self Go) SaveSupportData(filename string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).SaveSupportData(gc.String(filename)))
}

/*
Returns [code]true[/code] if locale is right-to-left.
*/
func (self Go) IsLocaleRightToLeft(locale string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLocaleRightToLeft(gc.String(locale)))
}

/*
Converts readable feature, variation, script, or language name to OpenType tag.
*/
func (self Go) NameToTag(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).NameToTag(gc.String(name))))
}

/*
Converts OpenType tag to readable feature, variation, script, or language name.
*/
func (self Go) TagToName(tag int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).TagToName(gc, gd.Int(tag)).String())
}

/*
Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
*/
func (self Go) Has(rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).Has(rid))
}

/*
Frees an object created by this [TextServer].
*/
func (self Go) FreeRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FreeRid(rid)
}

/*
Creates a new, empty font cache entry resource. To free the resulting resource, use the [method free_rid] method.
*/
func (self Go) CreateFont() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).CreateFont())
}

/*
Creates a new variation existing font which is reusing the same glyph cache and font data. To free the resulting resource, use the [method free_rid] method.
*/
func (self Go) CreateFontLinkedVariation(font_rid gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).CreateFontLinkedVariation(font_rid))
}

/*
Sets font source data, e.g contents of the dynamic font source file.
*/
func (self Go) FontSetData(font_rid gd.RID, data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetData(font_rid, gc.PackedByteSlice(data))
}

/*
Sets an active face index in the TrueType / OpenType collection.
*/
func (self Go) FontSetFaceIndex(font_rid gd.RID, face_index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetFaceIndex(font_rid, gd.Int(face_index))
}

/*
Returns an active face index in the TrueType / OpenType collection.
*/
func (self Go) FontGetFaceIndex(font_rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetFaceIndex(font_rid)))
}

/*
Returns number of faces in the TrueType / OpenType collection.
*/
func (self Go) FontGetFaceCount(font_rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetFaceCount(font_rid)))
}

/*
Sets the font style flags, see [enum FontStyle].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], [method font_set_embolden], or [method font_set_transform] instead.
*/
func (self Go) FontSetStyle(font_rid gd.RID, style classdb.TextServerFontStyle) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetStyle(font_rid, style)
}

/*
Returns font style flags, see [enum FontStyle].
*/
func (self Go) FontGetStyle(font_rid gd.RID) classdb.TextServerFontStyle {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerFontStyle(class(self).FontGetStyle(font_rid))
}

/*
Sets the font family name.
*/
func (self Go) FontSetName(font_rid gd.RID, name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetName(font_rid, gc.String(name))
}

/*
Returns font family name.
*/
func (self Go) FontGetName(font_rid gd.RID) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).FontGetName(gc, font_rid).String())
}

/*
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
func (self Go) FontGetOtNameStrings(font_rid gd.RID) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).FontGetOtNameStrings(gc, font_rid))
}

/*
Sets the font style name.
*/
func (self Go) FontSetStyleName(font_rid gd.RID, name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetStyleName(font_rid, gc.String(name))
}

/*
Returns font style name.
*/
func (self Go) FontGetStyleName(font_rid gd.RID) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).FontGetStyleName(gc, font_rid).String())
}

/*
Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], or [method font_set_embolden] instead.
*/
func (self Go) FontSetWeight(font_rid gd.RID, weight int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetWeight(font_rid, gd.Int(weight))
}

/*
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (self Go) FontGetWeight(font_rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetWeight(font_rid)))
}

/*
Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], or [method font_set_transform] instead.
*/
func (self Go) FontSetStretch(font_rid gd.RID, weight int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetStretch(font_rid, gd.Int(weight))
}

/*
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (self Go) FontGetStretch(font_rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetStretch(font_rid)))
}

/*
Sets font anti-aliasing mode.
*/
func (self Go) FontSetAntialiasing(font_rid gd.RID, antialiasing classdb.TextServerFontAntialiasing) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetAntialiasing(font_rid, antialiasing)
}

/*
Returns font anti-aliasing mode.
*/
func (self Go) FontGetAntialiasing(font_rid gd.RID) classdb.TextServerFontAntialiasing {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerFontAntialiasing(class(self).FontGetAntialiasing(font_rid))
}

/*
If set to [code]true[/code], embedded font bitmap loading is disabled (bitmap-only and color fonts ignore this property).
*/
func (self Go) FontSetDisableEmbeddedBitmaps(font_rid gd.RID, disable_embedded_bitmaps bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetDisableEmbeddedBitmaps(font_rid, disable_embedded_bitmaps)
}

/*
Returns whether the font's embedded bitmap loading is disabled.
*/
func (self Go) FontGetDisableEmbeddedBitmaps(font_rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontGetDisableEmbeddedBitmaps(font_rid))
}

/*
If set to [code]true[/code] font texture mipmap generation is enabled.
*/
func (self Go) FontSetGenerateMipmaps(font_rid gd.RID, generate_mipmaps bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetGenerateMipmaps(font_rid, generate_mipmaps)
}

/*
Returns [code]true[/code] if font texture mipmap generation is enabled.
*/
func (self Go) FontGetGenerateMipmaps(font_rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontGetGenerateMipmaps(font_rid))
}

/*
If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
[b]Note:[/b] MSDF font rendering does not render glyphs with overlapping shapes correctly. Overlapping shapes are not valid per the OpenType standard, but are still commonly found in many font files, especially those converted by Google Fonts. To avoid issues with overlapping glyphs, consider downloading the font file directly from the type foundry instead of relying on Google Fonts.
*/
func (self Go) FontSetMultichannelSignedDistanceField(font_rid gd.RID, msdf bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetMultichannelSignedDistanceField(font_rid, msdf)
}

/*
Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
*/
func (self Go) FontIsMultichannelSignedDistanceField(font_rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontIsMultichannelSignedDistanceField(font_rid))
}

/*
Sets the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (self Go) FontSetMsdfPixelRange(font_rid gd.RID, msdf_pixel_range int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetMsdfPixelRange(font_rid, gd.Int(msdf_pixel_range))
}

/*
Returns the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (self Go) FontGetMsdfPixelRange(font_rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetMsdfPixelRange(font_rid)))
}

/*
Sets source font size used to generate MSDF textures.
*/
func (self Go) FontSetMsdfSize(font_rid gd.RID, msdf_size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetMsdfSize(font_rid, gd.Int(msdf_size))
}

/*
Returns source font size used to generate MSDF textures.
*/
func (self Go) FontGetMsdfSize(font_rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetMsdfSize(font_rid)))
}

/*
Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
*/
func (self Go) FontSetFixedSize(font_rid gd.RID, fixed_size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetFixedSize(font_rid, gd.Int(fixed_size))
}

/*
Returns bitmap font fixed size.
*/
func (self Go) FontGetFixedSize(font_rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetFixedSize(font_rid)))
}

/*
Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
*/
func (self Go) FontSetFixedSizeScaleMode(font_rid gd.RID, fixed_size_scale_mode classdb.TextServerFixedSizeScaleMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetFixedSizeScaleMode(font_rid, fixed_size_scale_mode)
}

/*
Returns bitmap font scaling mode.
*/
func (self Go) FontGetFixedSizeScaleMode(font_rid gd.RID) classdb.TextServerFixedSizeScaleMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerFixedSizeScaleMode(class(self).FontGetFixedSizeScaleMode(font_rid))
}

/*
If set to [code]true[/code], system fonts can be automatically used as fallbacks.
*/
func (self Go) FontSetAllowSystemFallback(font_rid gd.RID, allow_system_fallback bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetAllowSystemFallback(font_rid, allow_system_fallback)
}

/*
Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
*/
func (self Go) FontIsAllowSystemFallback(font_rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontIsAllowSystemFallback(font_rid))
}

/*
If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
*/
func (self Go) FontSetForceAutohinter(font_rid gd.RID, force_autohinter bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetForceAutohinter(font_rid, force_autohinter)
}

/*
Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting. Used by dynamic fonts only.
*/
func (self Go) FontIsForceAutohinter(font_rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontIsForceAutohinter(font_rid))
}

/*
Sets font hinting mode. Used by dynamic fonts only.
*/
func (self Go) FontSetHinting(font_rid gd.RID, hinting classdb.TextServerHinting) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetHinting(font_rid, hinting)
}

/*
Returns the font hinting mode. Used by dynamic fonts only.
*/
func (self Go) FontGetHinting(font_rid gd.RID) classdb.TextServerHinting {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerHinting(class(self).FontGetHinting(font_rid))
}

/*
Sets font subpixel glyph positioning mode.
*/
func (self Go) FontSetSubpixelPositioning(font_rid gd.RID, subpixel_positioning classdb.TextServerSubpixelPositioning) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetSubpixelPositioning(font_rid, subpixel_positioning)
}

/*
Returns font subpixel glyph positioning mode.
*/
func (self Go) FontGetSubpixelPositioning(font_rid gd.RID) classdb.TextServerSubpixelPositioning {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerSubpixelPositioning(class(self).FontGetSubpixelPositioning(font_rid))
}

/*
Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (self Go) FontSetEmbolden(font_rid gd.RID, strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetEmbolden(font_rid, gd.Float(strength))
}

/*
Returns font embolden strength.
*/
func (self Go) FontGetEmbolden(font_rid gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetEmbolden(font_rid)))
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
func (self Go) FontSetSpacing(font_rid gd.RID, spacing classdb.TextServerSpacingType, value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetSpacing(font_rid, spacing, gd.Int(value))
}

/*
Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
func (self Go) FontGetSpacing(font_rid gd.RID, spacing classdb.TextServerSpacingType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetSpacing(font_rid, spacing)))
}

/*
Sets extra baseline offset (as a fraction of font height).
*/
func (self Go) FontSetBaselineOffset(font_rid gd.RID, baseline_offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetBaselineOffset(font_rid, gd.Float(baseline_offset))
}

/*
Returns extra baseline offset (as a fraction of font height).
*/
func (self Go) FontGetBaselineOffset(font_rid gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetBaselineOffset(font_rid)))
}

/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
For example, to simulate italic typeface by slanting, apply the following transform [code]Transform2D(1.0, slant, 0.0, 1.0, 0.0, 0.0)[/code].
*/
func (self Go) FontSetTransform(font_rid gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetTransform(font_rid, transform)
}

/*
Returns 2D transform applied to the font outlines.
*/
func (self Go) FontGetTransform(font_rid gd.RID) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).FontGetTransform(font_rid))
}

/*
Sets variation coordinates for the specified font cache entry. See [method font_supported_variation_list] for more info.
*/
func (self Go) FontSetVariationCoordinates(font_rid gd.RID, variation_coordinates gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetVariationCoordinates(font_rid, variation_coordinates)
}

/*
Returns variation coordinates for the specified font cache entry. See [method font_supported_variation_list] for more info.
*/
func (self Go) FontGetVariationCoordinates(font_rid gd.RID) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).FontGetVariationCoordinates(gc, font_rid))
}

/*
Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (self Go) FontSetOversampling(font_rid gd.RID, oversampling float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetOversampling(font_rid, gd.Float(oversampling))
}

/*
Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (self Go) FontGetOversampling(font_rid gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetOversampling(font_rid)))
}

/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
func (self Go) FontGetSizeCacheList(font_rid gd.RID) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](class(self).FontGetSizeCacheList(gc, font_rid))
}

/*
Removes all font sizes from the cache entry.
*/
func (self Go) FontClearSizeCache(font_rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontClearSizeCache(font_rid)
}

/*
Removes specified font size from the cache entry.
*/
func (self Go) FontRemoveSizeCache(font_rid gd.RID, size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontRemoveSizeCache(font_rid, size)
}

/*
Sets the font ascent (number of pixels above the baseline).
*/
func (self Go) FontSetAscent(font_rid gd.RID, size int, ascent float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetAscent(font_rid, gd.Int(size), gd.Float(ascent))
}

/*
Returns the font ascent (number of pixels above the baseline).
*/
func (self Go) FontGetAscent(font_rid gd.RID, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetAscent(font_rid, gd.Int(size))))
}

/*
Sets the font descent (number of pixels below the baseline).
*/
func (self Go) FontSetDescent(font_rid gd.RID, size int, descent float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetDescent(font_rid, gd.Int(size), gd.Float(descent))
}

/*
Returns the font descent (number of pixels below the baseline).
*/
func (self Go) FontGetDescent(font_rid gd.RID, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetDescent(font_rid, gd.Int(size))))
}

/*
Sets pixel offset of the underline below the baseline.
*/
func (self Go) FontSetUnderlinePosition(font_rid gd.RID, size int, underline_position float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetUnderlinePosition(font_rid, gd.Int(size), gd.Float(underline_position))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Go) FontGetUnderlinePosition(font_rid gd.RID, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetUnderlinePosition(font_rid, gd.Int(size))))
}

/*
Sets thickness of the underline in pixels.
*/
func (self Go) FontSetUnderlineThickness(font_rid gd.RID, size int, underline_thickness float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetUnderlineThickness(font_rid, gd.Int(size), gd.Float(underline_thickness))
}

/*
Returns thickness of the underline in pixels.
*/
func (self Go) FontGetUnderlineThickness(font_rid gd.RID, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetUnderlineThickness(font_rid, gd.Int(size))))
}

/*
Sets scaling factor of the color bitmap font.
*/
func (self Go) FontSetScale(font_rid gd.RID, size int, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetScale(font_rid, gd.Int(size), gd.Float(scale))
}

/*
Returns scaling factor of the color bitmap font.
*/
func (self Go) FontGetScale(font_rid gd.RID, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetScale(font_rid, gd.Int(size))))
}

/*
Returns number of textures used by font cache entry.
*/
func (self Go) FontGetTextureCount(font_rid gd.RID, size gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetTextureCount(font_rid, size)))
}

/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method font_remove_glyph] to remove them manually.
*/
func (self Go) FontClearTextures(font_rid gd.RID, size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontClearTextures(font_rid, size)
}

/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, remove them manually, using [method font_remove_glyph].
*/
func (self Go) FontRemoveTexture(font_rid gd.RID, size gd.Vector2i, texture_index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontRemoveTexture(font_rid, size, gd.Int(texture_index))
}

/*
Sets font cache texture image data.
*/
func (self Go) FontSetTextureImage(font_rid gd.RID, size gd.Vector2i, texture_index int, image gdclass.Image) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetTextureImage(font_rid, size, gd.Int(texture_index), image)
}

/*
Returns font cache texture image data.
*/
func (self Go) FontGetTextureImage(font_rid gd.RID, size gd.Vector2i, texture_index int) gdclass.Image {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Image(class(self).FontGetTextureImage(gc, font_rid, size, gd.Int(texture_index)))
}

/*
Sets array containing glyph packing data.
*/
func (self Go) FontSetTextureOffsets(font_rid gd.RID, size gd.Vector2i, texture_index int, offset []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetTextureOffsets(font_rid, size, gd.Int(texture_index), gc.PackedInt32Slice(offset))
}

/*
Returns array containing glyph packing data.
*/
func (self Go) FontGetTextureOffsets(font_rid gd.RID, size gd.Vector2i, texture_index int) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).FontGetTextureOffsets(gc, font_rid, size, gd.Int(texture_index)).AsSlice())
}

/*
Returns list of rendered glyphs in the cache entry.
*/
func (self Go) FontGetGlyphList(font_rid gd.RID, size gd.Vector2i) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).FontGetGlyphList(gc, font_rid, size).AsSlice())
}

/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method font_remove_texture] to remove them manually.
*/
func (self Go) FontClearGlyphs(font_rid gd.RID, size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontClearGlyphs(font_rid, size)
}

/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method font_remove_texture] to remove them manually.
*/
func (self Go) FontRemoveGlyph(font_rid gd.RID, size gd.Vector2i, glyph int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontRemoveGlyph(font_rid, size, gd.Int(glyph))
}

/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Go) FontGetGlyphAdvance(font_rid gd.RID, size int, glyph int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).FontGetGlyphAdvance(font_rid, gd.Int(size), gd.Int(glyph)))
}

/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Go) FontSetGlyphAdvance(font_rid gd.RID, size int, glyph int, advance gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetGlyphAdvance(font_rid, gd.Int(size), gd.Int(glyph), advance)
}

/*
Returns glyph offset from the baseline.
*/
func (self Go) FontGetGlyphOffset(font_rid gd.RID, size gd.Vector2i, glyph int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).FontGetGlyphOffset(font_rid, size, gd.Int(glyph)))
}

/*
Sets glyph offset from the baseline.
*/
func (self Go) FontSetGlyphOffset(font_rid gd.RID, size gd.Vector2i, glyph int, offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetGlyphOffset(font_rid, size, gd.Int(glyph), offset)
}

/*
Returns size of the glyph.
*/
func (self Go) FontGetGlyphSize(font_rid gd.RID, size gd.Vector2i, glyph int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).FontGetGlyphSize(font_rid, size, gd.Int(glyph)))
}

/*
Sets size of the glyph.
*/
func (self Go) FontSetGlyphSize(font_rid gd.RID, size gd.Vector2i, glyph int, gl_size gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetGlyphSize(font_rid, size, gd.Int(glyph), gl_size)
}

/*
Returns rectangle in the cache texture containing the glyph.
*/
func (self Go) FontGetGlyphUvRect(font_rid gd.RID, size gd.Vector2i, glyph int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).FontGetGlyphUvRect(font_rid, size, gd.Int(glyph)))
}

/*
Sets rectangle in the cache texture containing the glyph.
*/
func (self Go) FontSetGlyphUvRect(font_rid gd.RID, size gd.Vector2i, glyph int, uv_rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetGlyphUvRect(font_rid, size, gd.Int(glyph), uv_rect)
}

/*
Returns index of the cache texture containing the glyph.
*/
func (self Go) FontGetGlyphTextureIdx(font_rid gd.RID, size gd.Vector2i, glyph int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetGlyphTextureIdx(font_rid, size, gd.Int(glyph))))
}

/*
Sets index of the cache texture containing the glyph.
*/
func (self Go) FontSetGlyphTextureIdx(font_rid gd.RID, size gd.Vector2i, glyph int, texture_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetGlyphTextureIdx(font_rid, size, gd.Int(glyph), gd.Int(texture_idx))
}

/*
Returns resource ID of the cache texture containing the glyph.
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
func (self Go) FontGetGlyphTextureRid(font_rid gd.RID, size gd.Vector2i, glyph int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).FontGetGlyphTextureRid(font_rid, size, gd.Int(glyph)))
}

/*
Returns size of the cache texture containing the glyph.
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
func (self Go) FontGetGlyphTextureSize(font_rid gd.RID, size gd.Vector2i, glyph int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).FontGetGlyphTextureSize(font_rid, size, gd.Int(glyph)))
}

/*
Returns outline contours of the glyph as a [Dictionary] with the following contents:
[code]points[/code]         - [PackedVector3Array], containing outline points. [code]x[/code] and [code]y[/code] are point coordinates. [code]z[/code] is the type of the point, using the [enum ContourPointTag] values.
[code]contours[/code]       - [PackedInt32Array], containing indices the end points of each contour.
[code]orientation[/code]    - [bool], contour orientation. If [code]true[/code], clockwise contours must be filled.
*/
func (self Go) FontGetGlyphContours(font gd.RID, size int, index int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).FontGetGlyphContours(gc, font, gd.Int(size), gd.Int(index)))
}

/*
Returns list of the kerning overrides.
*/
func (self Go) FontGetKerningList(font_rid gd.RID, size int) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](class(self).FontGetKerningList(gc, font_rid, gd.Int(size)))
}

/*
Removes all kerning overrides.
*/
func (self Go) FontClearKerningMap(font_rid gd.RID, size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontClearKerningMap(font_rid, gd.Int(size))
}

/*
Removes kerning override for the pair of glyphs.
*/
func (self Go) FontRemoveKerning(font_rid gd.RID, size int, glyph_pair gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontRemoveKerning(font_rid, gd.Int(size), glyph_pair)
}

/*
Sets kerning for the pair of glyphs.
*/
func (self Go) FontSetKerning(font_rid gd.RID, size int, glyph_pair gd.Vector2i, kerning gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetKerning(font_rid, gd.Int(size), glyph_pair, kerning)
}

/*
Returns kerning for the pair of glyphs.
*/
func (self Go) FontGetKerning(font_rid gd.RID, size int, glyph_pair gd.Vector2i) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).FontGetKerning(font_rid, gd.Int(size), glyph_pair))
}

/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector]. See [method font_get_char_from_glyph_index].
*/
func (self Go) FontGetGlyphIndex(font_rid gd.RID, size int, char int, variation_selector int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetGlyphIndex(font_rid, gd.Int(size), gd.Int(char), gd.Int(variation_selector))))
}

/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method font_get_glyph_index].
*/
func (self Go) FontGetCharFromGlyphIndex(font_rid gd.RID, size int, glyph_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FontGetCharFromGlyphIndex(font_rid, gd.Int(size), gd.Int(glyph_index))))
}

/*
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
func (self Go) FontHasChar(font_rid gd.RID, char int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontHasChar(font_rid, gd.Int(char)))
}

/*
Returns a string containing all the characters available in the font.
*/
func (self Go) FontGetSupportedChars(font_rid gd.RID) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).FontGetSupportedChars(gc, font_rid).String())
}

/*
Renders the range of characters to the font cache texture.
*/
func (self Go) FontRenderRange(font_rid gd.RID, size gd.Vector2i, start int, end int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontRenderRange(font_rid, size, gd.Int(start), gd.Int(end))
}

/*
Renders specified glyph to the font cache texture.
*/
func (self Go) FontRenderGlyph(font_rid gd.RID, size gd.Vector2i, index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontRenderGlyph(font_rid, size, gd.Int(index))
}

/*
Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
[b]Note:[/b] Glyph index is specific to the font, use glyphs indices returned by [method shaped_text_get_glyphs] or [method font_get_glyph_index].
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
func (self Go) FontDrawGlyph(font_rid gd.RID, canvas gd.RID, size int, pos gd.Vector2, index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontDrawGlyph(font_rid, canvas, gd.Int(size), pos, gd.Int(index), gd.Color{1, 1, 1, 1})
}

/*
Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
[b]Note:[/b] Glyph index is specific to the font, use glyphs indices returned by [method shaped_text_get_glyphs] or [method font_get_glyph_index].
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
func (self Go) FontDrawGlyphOutline(font_rid gd.RID, canvas gd.RID, size int, outline_size int, pos gd.Vector2, index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontDrawGlyphOutline(font_rid, canvas, gd.Int(size), gd.Int(outline_size), pos, gd.Int(index), gd.Color{1, 1, 1, 1})
}

/*
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
func (self Go) FontIsLanguageSupported(font_rid gd.RID, language string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontIsLanguageSupported(font_rid, gc.String(language)))
}

/*
Adds override for [method font_is_language_supported].
*/
func (self Go) FontSetLanguageSupportOverride(font_rid gd.RID, language string, supported bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetLanguageSupportOverride(font_rid, gc.String(language), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
func (self Go) FontGetLanguageSupportOverride(font_rid gd.RID, language string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontGetLanguageSupportOverride(font_rid, gc.String(language)))
}

/*
Remove language support override.
*/
func (self Go) FontRemoveLanguageSupportOverride(font_rid gd.RID, language string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontRemoveLanguageSupportOverride(font_rid, gc.String(language))
}

/*
Returns list of language support overrides.
*/
func (self Go) FontGetLanguageSupportOverrides(font_rid gd.RID) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).FontGetLanguageSupportOverrides(gc, font_rid).Strings(gc))
}

/*
Returns [code]true[/code], if font supports given script (ISO 15924 code).
*/
func (self Go) FontIsScriptSupported(font_rid gd.RID, script string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontIsScriptSupported(font_rid, gc.String(script)))
}

/*
Adds override for [method font_is_script_supported].
*/
func (self Go) FontSetScriptSupportOverride(font_rid gd.RID, script string, supported bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetScriptSupportOverride(font_rid, gc.String(script), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
func (self Go) FontGetScriptSupportOverride(font_rid gd.RID, script string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).FontGetScriptSupportOverride(font_rid, gc.String(script)))
}

/*
Removes script support override.
*/
func (self Go) FontRemoveScriptSupportOverride(font_rid gd.RID, script string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontRemoveScriptSupportOverride(font_rid, gc.String(script))
}

/*
Returns list of script support overrides.
*/
func (self Go) FontGetScriptSupportOverrides(font_rid gd.RID) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).FontGetScriptSupportOverrides(gc, font_rid).Strings(gc))
}

/*
Sets font OpenType feature set override.
*/
func (self Go) FontSetOpentypeFeatureOverrides(font_rid gd.RID, overrides gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetOpentypeFeatureOverrides(font_rid, overrides)
}

/*
Returns font OpenType feature set override.
*/
func (self Go) FontGetOpentypeFeatureOverrides(font_rid gd.RID) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).FontGetOpentypeFeatureOverrides(gc, font_rid))
}

/*
Returns the dictionary of the supported OpenType features.
*/
func (self Go) FontSupportedFeatureList(font_rid gd.RID) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).FontSupportedFeatureList(gc, font_rid))
}

/*
Returns the dictionary of the supported OpenType variation coordinates.
*/
func (self Go) FontSupportedVariationList(font_rid gd.RID) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).FontSupportedVariationList(gc, font_rid))
}

/*
Returns the font oversampling factor, shared by all fonts in the TextServer.
*/
func (self Go) FontGetGlobalOversampling() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).FontGetGlobalOversampling()))
}

/*
Sets oversampling factor, shared by all font in the TextServer.
[b]Note:[/b] This value can be automatically changed by display server.
*/
func (self Go) FontSetGlobalOversampling(oversampling float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FontSetGlobalOversampling(gd.Float(oversampling))
}

/*
Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
*/
func (self Go) GetHexCodeBoxSize(size int, index int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetHexCodeBoxSize(gd.Int(size), gd.Int(index)))
}

/*
Draws box displaying character hexadecimal code. Used for replacing missing characters.
*/
func (self Go) DrawHexCodeBox(canvas gd.RID, size int, pos gd.Vector2, index int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DrawHexCodeBox(canvas, gd.Int(size), pos, gd.Int(index), color)
}

/*
Creates a new buffer for complex text layout, with the given [param direction] and [param orientation]. To free the resulting buffer, use [method free_rid] method.
[b]Note:[/b] Direction is ignored if server does not support [constant FEATURE_BIDI_LAYOUT] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] Orientation is ignored if server does not support [constant FEATURE_VERTICAL_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
func (self Go) CreateShapedText() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).CreateShapedText(0, 0))
}

/*
Clears text buffer (removes text and inline objects).
*/
func (self Go) ShapedTextClear(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextClear(rid)
}

/*
Sets desired text direction. If set to [constant DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
[b]Note:[/b] Direction is ignored if server does not support [constant FEATURE_BIDI_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
func (self Go) ShapedTextSetDirection(shaped gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextSetDirection(shaped, 0)
}

/*
Returns direction of the text.
*/
func (self Go) ShapedTextGetDirection(shaped gd.RID) classdb.TextServerDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerDirection(class(self).ShapedTextGetDirection(shaped))
}

/*
Returns direction of the text, inferred by the BiDi algorithm.
*/
func (self Go) ShapedTextGetInferredDirection(shaped gd.RID) classdb.TextServerDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerDirection(class(self).ShapedTextGetInferredDirection(shaped))
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
func (self Go) ShapedTextSetBidiOverride(shaped gd.RID, override gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextSetBidiOverride(shaped, override)
}

/*
Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (self Go) ShapedTextSetCustomPunctuation(shaped gd.RID, punct string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextSetCustomPunctuation(shaped, gc.String(punct))
}

/*
Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (self Go) ShapedTextGetCustomPunctuation(shaped gd.RID) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).ShapedTextGetCustomPunctuation(gc, shaped).String())
}

/*
Sets ellipsis character used for text clipping.
*/
func (self Go) ShapedTextSetCustomEllipsis(shaped gd.RID, char int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextSetCustomEllipsis(shaped, gd.Int(char))
}

/*
Returns ellipsis character used for text clipping.
*/
func (self Go) ShapedTextGetCustomEllipsis(shaped gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextGetCustomEllipsis(shaped)))
}

/*
Sets desired text orientation.
[b]Note:[/b] Orientation is ignored if server does not support [constant FEATURE_VERTICAL_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
func (self Go) ShapedTextSetOrientation(shaped gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextSetOrientation(shaped, 0)
}

/*
Returns text orientation.
*/
func (self Go) ShapedTextGetOrientation(shaped gd.RID) classdb.TextServerOrientation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerOrientation(class(self).ShapedTextGetOrientation(shaped))
}

/*
If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
*/
func (self Go) ShapedTextSetPreserveInvalid(shaped gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextSetPreserveInvalid(shaped, enabled)
}

/*
Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
[b]Note:[/b] If set to [code]false[/code], nothing is displayed in place of invalid characters.
*/
func (self Go) ShapedTextGetPreserveInvalid(shaped gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ShapedTextGetPreserveInvalid(shaped))
}

/*
If set to [code]true[/code] text buffer will display control characters.
*/
func (self Go) ShapedTextSetPreserveControl(shaped gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextSetPreserveControl(shaped, enabled)
}

/*
Returns [code]true[/code] if text buffer is configured to display control characters.
*/
func (self Go) ShapedTextGetPreserveControl(shaped gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ShapedTextGetPreserveControl(shaped))
}

/*
Sets extra spacing added between glyphs or lines in pixels.
*/
func (self Go) ShapedTextSetSpacing(shaped gd.RID, spacing classdb.TextServerSpacingType, value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextSetSpacing(shaped, spacing, gd.Int(value))
}

/*
Returns extra spacing added between glyphs or lines in pixels.
*/
func (self Go) ShapedTextGetSpacing(shaped gd.RID, spacing classdb.TextServerSpacingType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextGetSpacing(shaped, spacing)))
}

/*
Adds text span and font to draw it to the text buffer.
*/
func (self Go) ShapedTextAddString(shaped gd.RID, text string, fonts gd.ArrayOf[gd.RID], size int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ShapedTextAddString(shaped, gc.String(text), fonts, gd.Int(size), ([1]gd.Dictionary{}[0]), gc.String(""), gc.Variant(([1]gd.Variant{}[0]))))
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (self Go) ShapedTextAddObject(shaped gd.RID, key gd.Variant, size gd.Vector2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ShapedTextAddObject(shaped, key, size, 5, gd.Int(1), gd.Float(0.0)))
}

/*
Sets new size and alignment of embedded object.
*/
func (self Go) ShapedTextResizeObject(shaped gd.RID, key gd.Variant, size gd.Vector2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ShapedTextResizeObject(shaped, key, size, 5, gd.Float(0.0)))
}

/*
Returns number of text spans added using [method shaped_text_add_string] or [method shaped_text_add_object].
*/
func (self Go) ShapedGetSpanCount(shaped gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedGetSpanCount(shaped)))
}

/*
Returns text span metadata.
*/
func (self Go) ShapedGetSpanMeta(shaped gd.RID, index int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).ShapedGetSpanMeta(gc, shaped, gd.Int(index)))
}

/*
Changes text span font, font size, and OpenType features, without changing the text.
*/
func (self Go) ShapedSetSpanUpdateFont(shaped gd.RID, index int, fonts gd.ArrayOf[gd.RID], size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedSetSpanUpdateFont(shaped, gd.Int(index), fonts, gd.Int(size), ([1]gd.Dictionary{}[0]))
}

/*
Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
*/
func (self Go) ShapedTextSubstr(shaped gd.RID, start int, length int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).ShapedTextSubstr(shaped, gd.Int(start), gd.Int(length)))
}

/*
Returns the parent buffer from which the substring originates.
*/
func (self Go) ShapedTextGetParent(shaped gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).ShapedTextGetParent(shaped))
}

/*
Adjusts text width to fit to specified width, returns new text width.
*/
func (self Go) ShapedTextFitToWidth(shaped gd.RID, width float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).ShapedTextFitToWidth(shaped, gd.Float(width), 3)))
}

/*
Aligns shaped text to the given tab-stops.
*/
func (self Go) ShapedTextTabAlign(shaped gd.RID, tab_stops []float32) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).ShapedTextTabAlign(shaped, gc.PackedFloat32Slice(tab_stops))))
}

/*
Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
[b]Note:[/b] It is not necessary to call this function manually, buffer will be shaped automatically as soon as any of its output data is requested.
*/
func (self Go) ShapedTextShape(shaped gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ShapedTextShape(shaped))
}

/*
Returns [code]true[/code] if buffer is successfully shaped.
*/
func (self Go) ShapedTextIsReady(shaped gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ShapedTextIsReady(shaped))
}

/*
Returns [code]true[/code] if text buffer contains any visible characters.
*/
func (self Go) ShapedTextHasVisibleChars(shaped gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).ShapedTextHasVisibleChars(shaped))
}

/*
Returns an array of glyphs in the visual order.
*/
func (self Go) ShapedTextGetGlyphs(shaped gd.RID) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](class(self).ShapedTextGetGlyphs(gc, shaped))
}

/*
Returns text glyphs in the logical order.
*/
func (self Go) ShapedTextSortLogical(shaped gd.RID) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](class(self).ShapedTextSortLogical(gc, shaped))
}

/*
Returns number of glyphs in the buffer.
*/
func (self Go) ShapedTextGetGlyphCount(shaped gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextGetGlyphCount(shaped)))
}

/*
Returns substring buffer character range in the parent buffer.
*/
func (self Go) ShapedTextGetRange(shaped gd.RID) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).ShapedTextGetRange(shaped))
}

/*
Breaks text to the lines and columns. Returns character ranges for each segment.
*/
func (self Go) ShapedTextGetLineBreaksAdv(shaped gd.RID, width []float32) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).ShapedTextGetLineBreaksAdv(gc, shaped, gc.PackedFloat32Slice(width), gd.Int(0), true, 3).AsSlice())
}

/*
Breaks text to the lines and returns character ranges for each line.
*/
func (self Go) ShapedTextGetLineBreaks(shaped gd.RID, width float64) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).ShapedTextGetLineBreaks(gc, shaped, gd.Float(width), gd.Int(0), 3).AsSlice())
}

/*
Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum GraphemeFlag]).
*/
func (self Go) ShapedTextGetWordBreaks(shaped gd.RID) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).ShapedTextGetWordBreaks(gc, shaped, 264, 4).AsSlice())
}

/*
Returns the position of the overrun trim.
*/
func (self Go) ShapedTextGetTrimPos(shaped gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextGetTrimPos(shaped)))
}

/*
Returns position of the ellipsis.
*/
func (self Go) ShapedTextGetEllipsisPos(shaped gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextGetEllipsisPos(shaped)))
}

/*
Returns array of the glyphs in the ellipsis.
*/
func (self Go) ShapedTextGetEllipsisGlyphs(shaped gd.RID) gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](class(self).ShapedTextGetEllipsisGlyphs(gc, shaped))
}

/*
Returns number of glyphs in the ellipsis.
*/
func (self Go) ShapedTextGetEllipsisGlyphCount(shaped gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextGetEllipsisGlyphCount(shaped)))
}

/*
Trims text if it exceeds the given width.
*/
func (self Go) ShapedTextOverrunTrimToWidth(shaped gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextOverrunTrimToWidth(shaped, gd.Float(0), 0)
}

/*
Returns array of inline objects.
*/
func (self Go) ShapedTextGetObjects(shaped gd.RID) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).ShapedTextGetObjects(gc, shaped))
}

/*
Returns bounding rectangle of the inline object.
*/
func (self Go) ShapedTextGetObjectRect(shaped gd.RID, key gd.Variant) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).ShapedTextGetObjectRect(shaped, key))
}

/*
Returns the character range of the inline object.
*/
func (self Go) ShapedTextGetObjectRange(shaped gd.RID, key gd.Variant) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).ShapedTextGetObjectRange(shaped, key))
}

/*
Returns the glyph index of the inline object.
*/
func (self Go) ShapedTextGetObjectGlyph(shaped gd.RID, key gd.Variant) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextGetObjectGlyph(shaped, key)))
}

/*
Returns size of the text.
*/
func (self Go) ShapedTextGetSize(shaped gd.RID) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).ShapedTextGetSize(shaped))
}

/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
[b]Note:[/b] Overall ascent can be higher than font ascent, if some glyphs are displaced from the baseline.
*/
func (self Go) ShapedTextGetAscent(shaped gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).ShapedTextGetAscent(shaped)))
}

/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
[b]Note:[/b] Overall descent can be higher than font descent, if some glyphs are displaced from the baseline.
*/
func (self Go) ShapedTextGetDescent(shaped gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).ShapedTextGetDescent(shaped)))
}

/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
func (self Go) ShapedTextGetWidth(shaped gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).ShapedTextGetWidth(shaped)))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Go) ShapedTextGetUnderlinePosition(shaped gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).ShapedTextGetUnderlinePosition(shaped)))
}

/*
Returns thickness of the underline.
*/
func (self Go) ShapedTextGetUnderlineThickness(shaped gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).ShapedTextGetUnderlineThickness(shaped)))
}

/*
Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
*/
func (self Go) ShapedTextGetCarets(shaped gd.RID, position int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).ShapedTextGetCarets(gc, shaped, gd.Int(position)))
}

/*
Returns selection rectangles for the specified character range.
*/
func (self Go) ShapedTextGetSelection(shaped gd.RID, start int, end int) []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return []gd.Vector2(class(self).ShapedTextGetSelection(gc, shaped, gd.Int(start), gd.Int(end)).AsSlice())
}

/*
Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
*/
func (self Go) ShapedTextHitTestGrapheme(shaped gd.RID, coords float64) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextHitTestGrapheme(shaped, gd.Float(coords))))
}

/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
func (self Go) ShapedTextHitTestPosition(shaped gd.RID, coords float64) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextHitTestPosition(shaped, gd.Float(coords))))
}

/*
Returns composite character's bounds as offsets from the start of the line.
*/
func (self Go) ShapedTextGetGraphemeBounds(shaped gd.RID, pos int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).ShapedTextGetGraphemeBounds(shaped, gd.Int(pos)))
}

/*
Returns grapheme end position closest to the [param pos].
*/
func (self Go) ShapedTextNextGraphemePos(shaped gd.RID, pos int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextNextGraphemePos(shaped, gd.Int(pos))))
}

/*
Returns grapheme start position closest to the [param pos].
*/
func (self Go) ShapedTextPrevGraphemePos(shaped gd.RID, pos int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextPrevGraphemePos(shaped, gd.Int(pos))))
}

/*
Returns array of the composite character boundaries.
*/
func (self Go) ShapedTextGetCharacterBreaks(shaped gd.RID) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).ShapedTextGetCharacterBreaks(gc, shaped).AsSlice())
}

/*
Returns composite character end position closest to the [param pos].
*/
func (self Go) ShapedTextNextCharacterPos(shaped gd.RID, pos int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextNextCharacterPos(shaped, gd.Int(pos))))
}

/*
Returns composite character start position closest to the [param pos].
*/
func (self Go) ShapedTextPrevCharacterPos(shaped gd.RID, pos int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextPrevCharacterPos(shaped, gd.Int(pos))))
}

/*
Returns composite character position closest to the [param pos].
*/
func (self Go) ShapedTextClosestCharacterPos(shaped gd.RID, pos int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).ShapedTextClosestCharacterPos(shaped, gd.Int(pos))))
}

/*
Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (self Go) ShapedTextDraw(shaped gd.RID, canvas gd.RID, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextDraw(shaped, canvas, pos, gd.Float(-1), gd.Float(-1), gd.Color{1, 1, 1, 1})
}

/*
Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (self Go) ShapedTextDrawOutline(shaped gd.RID, canvas gd.RID, pos gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShapedTextDrawOutline(shaped, canvas, pos, gd.Float(-1), gd.Float(-1), gd.Int(1), gd.Color{1, 1, 1, 1})
}

/*
Returns dominant direction of in the range of text.
*/
func (self Go) ShapedTextGetDominantDirectionInRange(shaped gd.RID, start int, end int) classdb.TextServerDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerDirection(class(self).ShapedTextGetDominantDirectionInRange(shaped, gd.Int(start), gd.Int(end)))
}

/*
Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
If [param language] is omitted, the active locale will be used.
*/
func (self Go) FormatNumber(number string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).FormatNumber(gc, gc.String(number), gc.String("")).String())
}

/*
Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
*/
func (self Go) ParseNumber(number string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).ParseNumber(gc, gc.String(number), gc.String("")).String())
}

/*
Returns percent sign used in the [param language].
*/
func (self Go) PercentSign() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).PercentSign(gc, gc.String("")).String())
}

/*
Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
When [param chars_per_line] is greater than zero, line break boundaries are returned instead.
[codeblock]
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_word_breaks("The Godot Engine, 4")) # Prints [0, 3, 4, 9, 10, 16, 18, 19], which corresponds to the following substrings: "The", "Godot", "Engine", "4"
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 5)) # Prints [0, 3, 4, 9, 10, 15, 15, 19], which corresponds to the following substrings: "The", "Godot", "Engin", "e, 4"
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 10)) # Prints [0, 9, 10, 19], which corresponds to the following substrings: "The Godot", "Engine, 4"
[/codeblock]
*/
func (self Go) StringGetWordBreaks(s string) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).StringGetWordBreaks(gc, gc.String(s), gc.String(""), gd.Int(0)).AsSlice())
}

/*
Returns array of the composite character boundaries.
[codeblock]
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_word_breaks("Test ❤️‍🔥 Test")) # Prints [1, 2, 3, 4, 5, 9, 10, 11, 12, 13, 14]
[/codeblock]
*/
func (self Go) StringGetCharacterBreaks(s string) []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).StringGetCharacterBreaks(gc, gc.String(s), gc.String("")).AsSlice())
}

/*
Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
[b]Note:[/b] This method doesn't detect invisible characters, for spoof detection use it in combination with [method spoof_check].
[b]Note:[/b] Always returns [code]-1[/code] if the server does not support the [constant FEATURE_UNICODE_SECURITY] feature.
*/
func (self Go) IsConfusable(s string, dict []string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).IsConfusable(gc.String(s), gc.PackedStringSlice(dict))))
}

/*
Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
[b]Note:[/b] Always returns [code]false[/code] if the server does not support the [constant FEATURE_UNICODE_SECURITY] feature.
*/
func (self Go) SpoofCheck(s string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).SpoofCheck(gc.String(s)))
}

/*
Strips diacritics from the string.
[b]Note:[/b] The result may be longer or shorter than the original.
*/
func (self Go) StripDiacritics(s string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).StripDiacritics(gc, gc.String(s)).String())
}

/*
Returns [code]true[/code] if [param string] is a valid identifier.
If the text server supports the [constant FEATURE_UNICODE_IDENTIFIERS] feature, a valid identifier must:
- Conform to normalization form C.
- Begin with a Unicode character of class XID_Start or [code]"_"[/code].
- May contain Unicode characters of class XID_Continue in the other positions.
- Use UAX #31 recommended scripts only (mixed scripts are allowed).
If the [constant FEATURE_UNICODE_IDENTIFIERS] feature is not supported, a valid identifier must:
- Begin with a Unicode character of class XID_Start or [code]"_"[/code].
- May contain Unicode characters of class XID_Continue in the other positions.
*/
func (self Go) IsValidIdentifier(s string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsValidIdentifier(gc.String(s)))
}

/*
Returns [code]true[/code] if the given code point is a valid letter, i.e. it belongs to the Unicode category "L".
*/
func (self Go) IsValidLetter(unicode int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsValidLetter(gd.Int(unicode)))
}

/*
Returns the string converted to uppercase.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
func (self Go) StringToUpper(s string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).StringToUpper(gc, gc.String(s), gc.String("")).String())
}

/*
Returns the string converted to lowercase.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
func (self Go) StringToLower(s string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).StringToLower(gc, gc.String(s), gc.String("")).String())
}

/*
Returns the string converted to title case.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
func (self Go) StringToTitle(s string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).StringToTitle(gc, gc.String(s), gc.String("")).String())
}

/*
Default implementation of the BiDi algorithm override function. See [enum StructuredTextParser] for more info.
*/
func (self Go) ParseStructuredText(parser_type classdb.TextServerStructuredTextParser, args gd.Array, text string) gd.ArrayOf[gd.Vector3i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector3i](class(self).ParseStructuredText(gc, parser_type, args, gc.String(text)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextServer"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns [code]true[/code] if the server supports a feature.
*/
//go:nosplit
func (self class) HasFeature(feature classdb.TextServerFeature) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the server interface.
*/
//go:nosplit
func (self class) GetName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns text server features, see [enum Feature].
*/
//go:nosplit
func (self class) GetFeatures() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_get_features, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
[b]Note:[/b] This function should be called before any other TextServer functions used, otherwise it won't have any effect.
*/
//go:nosplit
func (self class) LoadSupportData(filename gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(filename))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_load_support_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
*/
//go:nosplit
func (self class) GetSupportDataFilename(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_get_support_data_filename, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
*/
//go:nosplit
func (self class) GetSupportDataInfo(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_get_support_data_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
[b]Note:[/b] This function is used by during project export, to include TextServer database.
*/
//go:nosplit
func (self class) SaveSupportData(filename gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(filename))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_save_support_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if locale is right-to-left.
*/
//go:nosplit
func (self class) IsLocaleRightToLeft(locale gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(locale))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_is_locale_right_to_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Converts readable feature, variation, script, or language name to OpenType tag.
*/
//go:nosplit
func (self class) NameToTag(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_name_to_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Converts OpenType tag to readable feature, variation, script, or language name.
*/
//go:nosplit
func (self class) TagToName(ctx gd.Lifetime, tag gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tag)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_tag_to_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
*/
//go:nosplit
func (self class) Has(rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_has, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Frees an object created by this [TextServer].
*/
//go:nosplit
func (self class) FreeRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new, empty font cache entry resource. To free the resulting resource, use the [method free_rid] method.
*/
//go:nosplit
func (self class) CreateFont() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_create_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new variation existing font which is reusing the same glyph cache and font data. To free the resulting resource, use the [method free_rid] method.
*/
//go:nosplit
func (self class) CreateFontLinkedVariation(font_rid gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_create_font_linked_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets font source data, e.g contents of the dynamic font source file.
*/
//go:nosplit
func (self class) FontSetData(font_rid gd.RID, data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) FontSetFaceIndex(font_rid gd.RID, face_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, face_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) FontGetFaceIndex(font_rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of faces in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) FontGetFaceCount(font_rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_face_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the font style flags, see [enum FontStyle].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], [method font_set_embolden], or [method font_set_transform] instead.
*/
//go:nosplit
func (self class) FontSetStyle(font_rid gd.RID, style classdb.TextServerFontStyle)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, style)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_style, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font style flags, see [enum FontStyle].
*/
//go:nosplit
func (self class) FontGetStyle(font_rid gd.RID) classdb.TextServerFontStyle {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[classdb.TextServerFontStyle](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_style, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the font family name.
*/
//go:nosplit
func (self class) FontSetName(font_rid gd.RID, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font family name.
*/
//go:nosplit
func (self class) FontGetName(ctx gd.Lifetime, font_rid gd.RID) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
//go:nosplit
func (self class) FontGetOtNameStrings(ctx gd.Lifetime, font_rid gd.RID) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_ot_name_strings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the font style name.
*/
//go:nosplit
func (self class) FontSetStyleName(font_rid gd.RID, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_style_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font style name.
*/
//go:nosplit
func (self class) FontGetStyleName(ctx gd.Lifetime, font_rid gd.RID) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_style_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], or [method font_set_embolden] instead.
*/
//go:nosplit
func (self class) FontSetWeight(font_rid gd.RID, weight gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
//go:nosplit
func (self class) FontGetWeight(font_rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], or [method font_set_transform] instead.
*/
//go:nosplit
func (self class) FontSetStretch(font_rid gd.RID, weight gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
//go:nosplit
func (self class) FontGetStretch(font_rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets font anti-aliasing mode.
*/
//go:nosplit
func (self class) FontSetAntialiasing(font_rid gd.RID, antialiasing classdb.TextServerFontAntialiasing)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, antialiasing)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font anti-aliasing mode.
*/
//go:nosplit
func (self class) FontGetAntialiasing(font_rid gd.RID) classdb.TextServerFontAntialiasing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[classdb.TextServerFontAntialiasing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If set to [code]true[/code], embedded font bitmap loading is disabled (bitmap-only and color fonts ignore this property).
*/
//go:nosplit
func (self class) FontSetDisableEmbeddedBitmaps(font_rid gd.RID, disable_embedded_bitmaps bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, disable_embedded_bitmaps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the font's embedded bitmap loading is disabled.
*/
//go:nosplit
func (self class) FontGetDisableEmbeddedBitmaps(font_rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If set to [code]true[/code] font texture mipmap generation is enabled.
*/
//go:nosplit
func (self class) FontSetGenerateMipmaps(font_rid gd.RID, generate_mipmaps bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, generate_mipmaps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if font texture mipmap generation is enabled.
*/
//go:nosplit
func (self class) FontGetGenerateMipmaps(font_rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
[b]Note:[/b] MSDF font rendering does not render glyphs with overlapping shapes correctly. Overlapping shapes are not valid per the OpenType standard, but are still commonly found in many font files, especially those converted by Google Fonts. To avoid issues with overlapping glyphs, consider downloading the font file directly from the type foundry instead of relying on Google Fonts.
*/
//go:nosplit
func (self class) FontSetMultichannelSignedDistanceField(font_rid gd.RID, msdf bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, msdf)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
*/
//go:nosplit
func (self class) FontIsMultichannelSignedDistanceField(font_rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_is_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the width of the range around the shape between the minimum and maximum representable signed distance.
*/
//go:nosplit
func (self class) FontSetMsdfPixelRange(font_rid gd.RID, msdf_pixel_range gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, msdf_pixel_range)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the width of the range around the shape between the minimum and maximum representable signed distance.
*/
//go:nosplit
func (self class) FontGetMsdfPixelRange(font_rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets source font size used to generate MSDF textures.
*/
//go:nosplit
func (self class) FontSetMsdfSize(font_rid gd.RID, msdf_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, msdf_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_msdf_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns source font size used to generate MSDF textures.
*/
//go:nosplit
func (self class) FontGetMsdfSize(font_rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_msdf_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
*/
//go:nosplit
func (self class) FontSetFixedSize(font_rid gd.RID, fixed_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, fixed_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_fixed_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns bitmap font fixed size.
*/
//go:nosplit
func (self class) FontGetFixedSize(font_rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_fixed_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
*/
//go:nosplit
func (self class) FontSetFixedSizeScaleMode(font_rid gd.RID, fixed_size_scale_mode classdb.TextServerFixedSizeScaleMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, fixed_size_scale_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns bitmap font scaling mode.
*/
//go:nosplit
func (self class) FontGetFixedSizeScaleMode(font_rid gd.RID) classdb.TextServerFixedSizeScaleMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[classdb.TextServerFixedSizeScaleMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If set to [code]true[/code], system fonts can be automatically used as fallbacks.
*/
//go:nosplit
func (self class) FontSetAllowSystemFallback(font_rid gd.RID, allow_system_fallback bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, allow_system_fallback)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
*/
//go:nosplit
func (self class) FontIsAllowSystemFallback(font_rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_is_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
*/
//go:nosplit
func (self class) FontSetForceAutohinter(font_rid gd.RID, force_autohinter bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, force_autohinter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontIsForceAutohinter(font_rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_is_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets font hinting mode. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontSetHinting(font_rid gd.RID, hinting classdb.TextServerHinting)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, hinting)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_hinting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the font hinting mode. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontGetHinting(font_rid gd.RID) classdb.TextServerHinting {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[classdb.TextServerHinting](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_hinting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets font subpixel glyph positioning mode.
*/
//go:nosplit
func (self class) FontSetSubpixelPositioning(font_rid gd.RID, subpixel_positioning classdb.TextServerSubpixelPositioning)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, subpixel_positioning)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font subpixel glyph positioning mode.
*/
//go:nosplit
func (self class) FontGetSubpixelPositioning(font_rid gd.RID) classdb.TextServerSubpixelPositioning {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[classdb.TextServerSubpixelPositioning](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) FontSetEmbolden(font_rid gd.RID, strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font embolden strength.
*/
//go:nosplit
func (self class) FontGetEmbolden(font_rid gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) FontSetSpacing(font_rid gd.RID, spacing classdb.TextServerSpacingType, value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) FontGetSpacing(font_rid gd.RID, spacing classdb.TextServerSpacingType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) FontSetBaselineOffset(font_rid gd.RID, baseline_offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, baseline_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) FontGetBaselineOffset(font_rid gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
For example, to simulate italic typeface by slanting, apply the following transform [code]Transform2D(1.0, slant, 0.0, 1.0, 0.0, 0.0)[/code].
*/
//go:nosplit
func (self class) FontSetTransform(font_rid gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns 2D transform applied to the font outlines.
*/
//go:nosplit
func (self class) FontGetTransform(font_rid gd.RID) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets variation coordinates for the specified font cache entry. See [method font_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) FontSetVariationCoordinates(font_rid gd.RID, variation_coordinates gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(variation_coordinates))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns variation coordinates for the specified font cache entry. See [method font_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) FontGetVariationCoordinates(ctx gd.Lifetime, font_rid gd.RID) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontSetOversampling(font_rid gd.RID, oversampling gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, oversampling)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontGetOversampling(font_rid gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
//go:nosplit
func (self class) FontGetSizeCacheList(ctx gd.Lifetime, font_rid gd.RID) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_size_cache_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Removes all font sizes from the cache entry.
*/
//go:nosplit
func (self class) FontClearSizeCache(font_rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_clear_size_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified font size from the cache entry.
*/
//go:nosplit
func (self class) FontRemoveSizeCache(font_rid gd.RID, size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_remove_size_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) FontSetAscent(font_rid gd.RID, size gd.Int, ascent gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, ascent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) FontGetAscent(font_rid gd.RID, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) FontSetDescent(font_rid gd.RID, size gd.Int, descent gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, descent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) FontGetDescent(font_rid gd.RID, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) FontSetUnderlinePosition(font_rid gd.RID, size gd.Int, underline_position gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) FontGetUnderlinePosition(font_rid gd.RID, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets thickness of the underline in pixels.
*/
//go:nosplit
func (self class) FontSetUnderlineThickness(font_rid gd.RID, size gd.Int, underline_thickness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_thickness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns thickness of the underline in pixels.
*/
//go:nosplit
func (self class) FontGetUnderlineThickness(font_rid gd.RID, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) FontSetScale(font_rid gd.RID, size gd.Int, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) FontGetScale(font_rid gd.RID, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of textures used by font cache entry.
*/
//go:nosplit
func (self class) FontGetTextureCount(font_rid gd.RID, size gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_texture_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method font_remove_glyph] to remove them manually.
*/
//go:nosplit
func (self class) FontClearTextures(font_rid gd.RID, size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_clear_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, remove them manually, using [method font_remove_glyph].
*/
//go:nosplit
func (self class) FontRemoveTexture(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_remove_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets font cache texture image data.
*/
//go:nosplit
func (self class) FontSetTextureImage(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, image gdclass.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_texture_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font cache texture image data.
*/
//go:nosplit
func (self class) FontGetTextureImage(ctx gd.Lifetime, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) gdclass.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_texture_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets array containing glyph packing data.
*/
//go:nosplit
func (self class) FontSetTextureOffsets(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, offset gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, mmm.Get(offset))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns array containing glyph packing data.
*/
//go:nosplit
func (self class) FontGetTextureOffsets(ctx gd.Lifetime, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns list of rendered glyphs in the cache entry.
*/
//go:nosplit
func (self class) FontGetGlyphList(ctx gd.Lifetime, font_rid gd.RID, size gd.Vector2i) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method font_remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) FontClearGlyphs(font_rid gd.RID, size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_clear_glyphs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method font_remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) FontRemoveGlyph(font_rid gd.RID, size gd.Vector2i, glyph gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_remove_glyph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) FontGetGlyphAdvance(font_rid gd.RID, size gd.Int, glyph gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) FontSetGlyphAdvance(font_rid gd.RID, size gd.Int, glyph gd.Int, advance gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, advance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns glyph offset from the baseline.
*/
//go:nosplit
func (self class) FontGetGlyphOffset(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets glyph offset from the baseline.
*/
//go:nosplit
func (self class) FontSetGlyphOffset(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns size of the glyph.
*/
//go:nosplit
func (self class) FontGetGlyphSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets size of the glyph.
*/
//go:nosplit
func (self class) FontSetGlyphSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, gl_size gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, gl_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_glyph_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) FontGetGlyphUvRect(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) FontSetGlyphUvRect(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, uv_rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, uv_rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) FontGetGlyphTextureIdx(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) FontSetGlyphTextureIdx(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, texture_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, texture_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns resource ID of the cache texture containing the glyph.
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
//go:nosplit
func (self class) FontGetGlyphTextureRid(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_texture_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns size of the cache texture containing the glyph.
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
//go:nosplit
func (self class) FontGetGlyphTextureSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns outline contours of the glyph as a [Dictionary] with the following contents:
[code]points[/code]         - [PackedVector3Array], containing outline points. [code]x[/code] and [code]y[/code] are point coordinates. [code]z[/code] is the type of the point, using the [enum ContourPointTag] values.
[code]contours[/code]       - [PackedInt32Array], containing indices the end points of each contour.
[code]orientation[/code]    - [bool], contour orientation. If [code]true[/code], clockwise contours must be filled.
*/
//go:nosplit
func (self class) FontGetGlyphContours(ctx gd.Lifetime, font gd.RID, size gd.Int, index gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font)
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_contours, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns list of the kerning overrides.
*/
//go:nosplit
func (self class) FontGetKerningList(ctx gd.Lifetime, font_rid gd.RID, size gd.Int) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_kerning_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Removes all kerning overrides.
*/
//go:nosplit
func (self class) FontClearKerningMap(font_rid gd.RID, size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_clear_kerning_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes kerning override for the pair of glyphs.
*/
//go:nosplit
func (self class) FontRemoveKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_remove_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) FontSetKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i, kerning gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	callframe.Arg(frame, kerning)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) FontGetKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector]. See [method font_get_char_from_glyph_index].
*/
//go:nosplit
func (self class) FontGetGlyphIndex(font_rid gd.RID, size gd.Int, char gd.Int, variation_selector gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, char)
	callframe.Arg(frame, variation_selector)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method font_get_glyph_index].
*/
//go:nosplit
func (self class) FontGetCharFromGlyphIndex(font_rid gd.RID, size gd.Int, glyph_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_char_from_glyph_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
//go:nosplit
func (self class) FontHasChar(font_rid gd.RID, char gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, char)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_has_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a string containing all the characters available in the font.
*/
//go:nosplit
func (self class) FontGetSupportedChars(ctx gd.Lifetime, font_rid gd.RID) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_supported_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Renders the range of characters to the font cache texture.
*/
//go:nosplit
func (self class) FontRenderRange(font_rid gd.RID, size gd.Vector2i, start gd.Int, end gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_render_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renders specified glyph to the font cache texture.
*/
//go:nosplit
func (self class) FontRenderGlyph(font_rid gd.RID, size gd.Vector2i, index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_render_glyph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
[b]Note:[/b] Glyph index is specific to the font, use glyphs indices returned by [method shaped_text_get_glyphs] or [method font_get_glyph_index].
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
//go:nosplit
func (self class) FontDrawGlyph(font_rid gd.RID, canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, size)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_draw_glyph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
[b]Note:[/b] Glyph index is specific to the font, use glyphs indices returned by [method shaped_text_get_glyphs] or [method font_get_glyph_index].
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
//go:nosplit
func (self class) FontDrawGlyphOutline(font_rid gd.RID, canvas gd.RID, size gd.Int, outline_size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, size)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_draw_glyph_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
//go:nosplit
func (self class) FontIsLanguageSupported(font_rid gd.RID, language gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_is_language_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds override for [method font_is_language_supported].
*/
//go:nosplit
func (self class) FontSetLanguageSupportOverride(font_rid gd.RID, language gd.String, supported bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(language))
	callframe.Arg(frame, supported)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
//go:nosplit
func (self class) FontGetLanguageSupportOverride(font_rid gd.RID, language gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Remove language support override.
*/
//go:nosplit
func (self class) FontRemoveLanguageSupportOverride(font_rid gd.RID, language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_remove_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns list of language support overrides.
*/
//go:nosplit
func (self class) FontGetLanguageSupportOverrides(ctx gd.Lifetime, font_rid gd.RID) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_language_support_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code], if font supports given script (ISO 15924 code).
*/
//go:nosplit
func (self class) FontIsScriptSupported(font_rid gd.RID, script gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(script))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_is_script_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds override for [method font_is_script_supported].
*/
//go:nosplit
func (self class) FontSetScriptSupportOverride(font_rid gd.RID, script gd.String, supported bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(script))
	callframe.Arg(frame, supported)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
//go:nosplit
func (self class) FontGetScriptSupportOverride(font_rid gd.RID, script gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(script))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes script support override.
*/
//go:nosplit
func (self class) FontRemoveScriptSupportOverride(font_rid gd.RID, script gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(script))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_remove_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns list of script support overrides.
*/
//go:nosplit
func (self class) FontGetScriptSupportOverrides(ctx gd.Lifetime, font_rid gd.RID) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_script_support_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets font OpenType feature set override.
*/
//go:nosplit
func (self class) FontSetOpentypeFeatureOverrides(font_rid gd.RID, overrides gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, mmm.Get(overrides))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns font OpenType feature set override.
*/
//go:nosplit
func (self class) FontGetOpentypeFeatureOverrides(ctx gd.Lifetime, font_rid gd.RID) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the dictionary of the supported OpenType features.
*/
//go:nosplit
func (self class) FontSupportedFeatureList(ctx gd.Lifetime, font_rid gd.RID) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_supported_feature_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the dictionary of the supported OpenType variation coordinates.
*/
//go:nosplit
func (self class) FontSupportedVariationList(ctx gd.Lifetime, font_rid gd.RID) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_supported_variation_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the font oversampling factor, shared by all fonts in the TextServer.
*/
//go:nosplit
func (self class) FontGetGlobalOversampling() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_get_global_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets oversampling factor, shared by all font in the TextServer.
[b]Note:[/b] This value can be automatically changed by display server.
*/
//go:nosplit
func (self class) FontSetGlobalOversampling(oversampling gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, oversampling)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_font_set_global_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
*/
//go:nosplit
func (self class) GetHexCodeBoxSize(size gd.Int, index gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_get_hex_code_box_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draws box displaying character hexadecimal code. Used for replacing missing characters.
*/
//go:nosplit
func (self class) DrawHexCodeBox(canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, size)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_draw_hex_code_box, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new buffer for complex text layout, with the given [param direction] and [param orientation]. To free the resulting buffer, use [method free_rid] method.
[b]Note:[/b] Direction is ignored if server does not support [constant FEATURE_BIDI_LAYOUT] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] Orientation is ignored if server does not support [constant FEATURE_VERTICAL_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
//go:nosplit
func (self class) CreateShapedText(direction classdb.TextServerDirection, orientation classdb.TextServerOrientation) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_create_shaped_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears text buffer (removes text and inline objects).
*/
//go:nosplit
func (self class) ShapedTextClear(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets desired text direction. If set to [constant DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
[b]Note:[/b] Direction is ignored if server does not support [constant FEATURE_BIDI_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
//go:nosplit
func (self class) ShapedTextSetDirection(shaped gd.RID, direction classdb.TextServerDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns direction of the text.
*/
//go:nosplit
func (self class) ShapedTextGetDirection(shaped gd.RID) classdb.TextServerDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[classdb.TextServerDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns direction of the text, inferred by the BiDi algorithm.
*/
//go:nosplit
func (self class) ShapedTextGetInferredDirection(shaped gd.RID) classdb.TextServerDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[classdb.TextServerDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_inferred_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
//go:nosplit
func (self class) ShapedTextSetBidiOverride(shaped gd.RID, override gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(override))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_set_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
//go:nosplit
func (self class) ShapedTextSetCustomPunctuation(shaped gd.RID, punct gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(punct))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_set_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
//go:nosplit
func (self class) ShapedTextGetCustomPunctuation(ctx gd.Lifetime, shaped gd.RID) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets ellipsis character used for text clipping.
*/
//go:nosplit
func (self class) ShapedTextSetCustomEllipsis(shaped gd.RID, char gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, char)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_set_custom_ellipsis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns ellipsis character used for text clipping.
*/
//go:nosplit
func (self class) ShapedTextGetCustomEllipsis(shaped gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_custom_ellipsis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets desired text orientation.
[b]Note:[/b] Orientation is ignored if server does not support [constant FEATURE_VERTICAL_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
//go:nosplit
func (self class) ShapedTextSetOrientation(shaped gd.RID, orientation classdb.TextServerOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_set_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns text orientation.
*/
//go:nosplit
func (self class) ShapedTextGetOrientation(shaped gd.RID) classdb.TextServerOrientation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[classdb.TextServerOrientation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
*/
//go:nosplit
func (self class) ShapedTextSetPreserveInvalid(shaped gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_set_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
[b]Note:[/b] If set to [code]false[/code], nothing is displayed in place of invalid characters.
*/
//go:nosplit
func (self class) ShapedTextGetPreserveInvalid(shaped gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If set to [code]true[/code] text buffer will display control characters.
*/
//go:nosplit
func (self class) ShapedTextSetPreserveControl(shaped gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_set_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if text buffer is configured to display control characters.
*/
//go:nosplit
func (self class) ShapedTextGetPreserveControl(shaped gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets extra spacing added between glyphs or lines in pixels.
*/
//go:nosplit
func (self class) ShapedTextSetSpacing(shaped gd.RID, spacing classdb.TextServerSpacingType, value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_set_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns extra spacing added between glyphs or lines in pixels.
*/
//go:nosplit
func (self class) ShapedTextGetSpacing(shaped gd.RID, spacing classdb.TextServerSpacingType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds text span and font to draw it to the text buffer.
*/
//go:nosplit
func (self class) ShapedTextAddString(shaped gd.RID, text gd.String, fonts gd.ArrayOf[gd.RID], size gd.Int, opentype_features gd.Dictionary, language gd.String, meta gd.Variant) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, mmm.Get(fonts))
	callframe.Arg(frame, size)
	callframe.Arg(frame, mmm.Get(opentype_features))
	callframe.Arg(frame, mmm.Get(language))
	callframe.Arg(frame, mmm.Get(meta))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_add_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
//go:nosplit
func (self class) ShapedTextAddObject(shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, length gd.Int, baseline gd.Float) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, length)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_add_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets new size and alignment of embedded object.
*/
//go:nosplit
func (self class) ShapedTextResizeObject(shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, baseline gd.Float) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_resize_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of text spans added using [method shaped_text_add_string] or [method shaped_text_add_object].
*/
//go:nosplit
func (self class) ShapedGetSpanCount(shaped gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_get_span_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns text span metadata.
*/
//go:nosplit
func (self class) ShapedGetSpanMeta(ctx gd.Lifetime, shaped gd.RID, index gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_get_span_meta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Changes text span font, font size, and OpenType features, without changing the text.
*/
//go:nosplit
func (self class) ShapedSetSpanUpdateFont(shaped gd.RID, index gd.Int, fonts gd.ArrayOf[gd.RID], size gd.Int, opentype_features gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(fonts))
	callframe.Arg(frame, size)
	callframe.Arg(frame, mmm.Get(opentype_features))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_set_span_update_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
*/
//go:nosplit
func (self class) ShapedTextSubstr(shaped gd.RID, start gd.Int, length gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, start)
	callframe.Arg(frame, length)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_substr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the parent buffer from which the substring originates.
*/
//go:nosplit
func (self class) ShapedTextGetParent(shaped gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adjusts text width to fit to specified width, returns new text width.
*/
//go:nosplit
func (self class) ShapedTextFitToWidth(shaped gd.RID, width gd.Float, justification_flags classdb.TextServerJustificationFlag) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, width)
	callframe.Arg(frame, justification_flags)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_fit_to_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Aligns shaped text to the given tab-stops.
*/
//go:nosplit
func (self class) ShapedTextTabAlign(shaped gd.RID, tab_stops gd.PackedFloat32Array) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(tab_stops))
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_tab_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
[b]Note:[/b] It is not necessary to call this function manually, buffer will be shaped automatically as soon as any of its output data is requested.
*/
//go:nosplit
func (self class) ShapedTextShape(shaped gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if buffer is successfully shaped.
*/
//go:nosplit
func (self class) ShapedTextIsReady(shaped gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_is_ready, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if text buffer contains any visible characters.
*/
//go:nosplit
func (self class) ShapedTextHasVisibleChars(shaped gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_has_visible_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of glyphs in the visual order.
*/
//go:nosplit
func (self class) ShapedTextGetGlyphs(ctx gd.Lifetime, shaped gd.RID) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_glyphs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns text glyphs in the logical order.
*/
//go:nosplit
func (self class) ShapedTextSortLogical(ctx gd.Lifetime, shaped gd.RID) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_sort_logical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns number of glyphs in the buffer.
*/
//go:nosplit
func (self class) ShapedTextGetGlyphCount(shaped gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_glyph_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns substring buffer character range in the parent buffer.
*/
//go:nosplit
func (self class) ShapedTextGetRange(shaped gd.RID) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Breaks text to the lines and columns. Returns character ranges for each segment.
*/
//go:nosplit
func (self class) ShapedTextGetLineBreaksAdv(ctx gd.Lifetime, shaped gd.RID, width gd.PackedFloat32Array, start gd.Int, once bool, break_flags classdb.TextServerLineBreakFlag) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(width))
	callframe.Arg(frame, start)
	callframe.Arg(frame, once)
	callframe.Arg(frame, break_flags)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_line_breaks_adv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Breaks text to the lines and returns character ranges for each line.
*/
//go:nosplit
func (self class) ShapedTextGetLineBreaks(ctx gd.Lifetime, shaped gd.RID, width gd.Float, start gd.Int, break_flags classdb.TextServerLineBreakFlag) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, width)
	callframe.Arg(frame, start)
	callframe.Arg(frame, break_flags)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_line_breaks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum GraphemeFlag]).
*/
//go:nosplit
func (self class) ShapedTextGetWordBreaks(ctx gd.Lifetime, shaped gd.RID, grapheme_flags classdb.TextServerGraphemeFlag, skip_grapheme_flags classdb.TextServerGraphemeFlag) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, grapheme_flags)
	callframe.Arg(frame, skip_grapheme_flags)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_word_breaks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the position of the overrun trim.
*/
//go:nosplit
func (self class) ShapedTextGetTrimPos(shaped gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_trim_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns position of the ellipsis.
*/
//go:nosplit
func (self class) ShapedTextGetEllipsisPos(shaped gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_ellipsis_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns array of the glyphs in the ellipsis.
*/
//go:nosplit
func (self class) ShapedTextGetEllipsisGlyphs(ctx gd.Lifetime, shaped gd.RID) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_ellipsis_glyphs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns number of glyphs in the ellipsis.
*/
//go:nosplit
func (self class) ShapedTextGetEllipsisGlyphCount(shaped gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_ellipsis_glyph_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Trims text if it exceeds the given width.
*/
//go:nosplit
func (self class) ShapedTextOverrunTrimToWidth(shaped gd.RID, width gd.Float, overrun_trim_flags classdb.TextServerTextOverrunFlag)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, width)
	callframe.Arg(frame, overrun_trim_flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_overrun_trim_to_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns array of inline objects.
*/
//go:nosplit
func (self class) ShapedTextGetObjects(ctx gd.Lifetime, shaped gd.RID) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns bounding rectangle of the inline object.
*/
//go:nosplit
func (self class) ShapedTextGetObjectRect(shaped gd.RID, key gd.Variant) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_object_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the character range of the inline object.
*/
//go:nosplit
func (self class) ShapedTextGetObjectRange(shaped gd.RID, key gd.Variant) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_object_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the glyph index of the inline object.
*/
//go:nosplit
func (self class) ShapedTextGetObjectGlyph(shaped gd.RID, key gd.Variant) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, mmm.Get(key))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_object_glyph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns size of the text.
*/
//go:nosplit
func (self class) ShapedTextGetSize(shaped gd.RID) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
[b]Note:[/b] Overall ascent can be higher than font ascent, if some glyphs are displaced from the baseline.
*/
//go:nosplit
func (self class) ShapedTextGetAscent(shaped gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
[b]Note:[/b] Overall descent can be higher than font descent, if some glyphs are displaced from the baseline.
*/
//go:nosplit
func (self class) ShapedTextGetDescent(shaped gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
//go:nosplit
func (self class) ShapedTextGetWidth(shaped gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) ShapedTextGetUnderlinePosition(shaped gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns thickness of the underline.
*/
//go:nosplit
func (self class) ShapedTextGetUnderlineThickness(shaped gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
*/
//go:nosplit
func (self class) ShapedTextGetCarets(ctx gd.Lifetime, shaped gd.RID, position gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns selection rectangles for the specified character range.
*/
//go:nosplit
func (self class) ShapedTextGetSelection(ctx gd.Lifetime, shaped gd.RID, start gd.Int, end gd.Int) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
*/
//go:nosplit
func (self class) ShapedTextHitTestGrapheme(shaped gd.RID, coords gd.Float) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_hit_test_grapheme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
//go:nosplit
func (self class) ShapedTextHitTestPosition(shaped gd.RID, coords gd.Float) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_hit_test_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns composite character's bounds as offsets from the start of the line.
*/
//go:nosplit
func (self class) ShapedTextGetGraphemeBounds(shaped gd.RID, pos gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_grapheme_bounds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns grapheme end position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextNextGraphemePos(shaped gd.RID, pos gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_next_grapheme_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns grapheme start position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextPrevGraphemePos(shaped gd.RID, pos gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_prev_grapheme_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns array of the composite character boundaries.
*/
//go:nosplit
func (self class) ShapedTextGetCharacterBreaks(ctx gd.Lifetime, shaped gd.RID) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_character_breaks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns composite character end position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextNextCharacterPos(shaped gd.RID, pos gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_next_character_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns composite character start position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextPrevCharacterPos(shaped gd.RID, pos gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_prev_character_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns composite character position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextClosestCharacterPos(shaped gd.RID, pos gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_closest_character_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
//go:nosplit
func (self class) ShapedTextDraw(shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, clip_l)
	callframe.Arg(frame, clip_r)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
//go:nosplit
func (self class) ShapedTextDrawOutline(shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, outline_size gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, clip_l)
	callframe.Arg(frame, clip_r)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_draw_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns dominant direction of in the range of text.
*/
//go:nosplit
func (self class) ShapedTextGetDominantDirectionInRange(shaped gd.RID, start gd.Int, end gd.Int) classdb.TextServerDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[classdb.TextServerDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_shaped_text_get_dominant_direction_in_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
If [param language] is omitted, the active locale will be used.
*/
//go:nosplit
func (self class) FormatNumber(ctx gd.Lifetime, number gd.String, language gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(number))
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_format_number, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
*/
//go:nosplit
func (self class) ParseNumber(ctx gd.Lifetime, number gd.String, language gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(number))
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_parse_number, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns percent sign used in the [param language].
*/
//go:nosplit
func (self class) PercentSign(ctx gd.Lifetime, language gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_percent_sign, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
When [param chars_per_line] is greater than zero, line break boundaries are returned instead.
[codeblock]
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_word_breaks("The Godot Engine, 4")) # Prints [0, 3, 4, 9, 10, 16, 18, 19], which corresponds to the following substrings: "The", "Godot", "Engine", "4"
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 5)) # Prints [0, 3, 4, 9, 10, 15, 15, 19], which corresponds to the following substrings: "The", "Godot", "Engin", "e, 4"
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 10)) # Prints [0, 9, 10, 19], which corresponds to the following substrings: "The Godot", "Engine, 4"
[/codeblock]
*/
//go:nosplit
func (self class) StringGetWordBreaks(ctx gd.Lifetime, s gd.String, language gd.String, chars_per_line gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	callframe.Arg(frame, mmm.Get(language))
	callframe.Arg(frame, chars_per_line)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_string_get_word_breaks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns array of the composite character boundaries.
[codeblock]
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_word_breaks("Test ❤️‍🔥 Test")) # Prints [1, 2, 3, 4, 5, 9, 10, 11, 12, 13, 14]
[/codeblock]
*/
//go:nosplit
func (self class) StringGetCharacterBreaks(ctx gd.Lifetime, s gd.String, language gd.String) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_string_get_character_breaks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
[b]Note:[/b] This method doesn't detect invisible characters, for spoof detection use it in combination with [method spoof_check].
[b]Note:[/b] Always returns [code]-1[/code] if the server does not support the [constant FEATURE_UNICODE_SECURITY] feature.
*/
//go:nosplit
func (self class) IsConfusable(s gd.String, dict gd.PackedStringArray) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	callframe.Arg(frame, mmm.Get(dict))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_is_confusable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
[b]Note:[/b] Always returns [code]false[/code] if the server does not support the [constant FEATURE_UNICODE_SECURITY] feature.
*/
//go:nosplit
func (self class) SpoofCheck(s gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_spoof_check, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Strips diacritics from the string.
[b]Note:[/b] The result may be longer or shorter than the original.
*/
//go:nosplit
func (self class) StripDiacritics(ctx gd.Lifetime, s gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_strip_diacritics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [param string] is a valid identifier.
If the text server supports the [constant FEATURE_UNICODE_IDENTIFIERS] feature, a valid identifier must:
- Conform to normalization form C.
- Begin with a Unicode character of class XID_Start or [code]"_"[/code].
- May contain Unicode characters of class XID_Continue in the other positions.
- Use UAX #31 recommended scripts only (mixed scripts are allowed).
If the [constant FEATURE_UNICODE_IDENTIFIERS] feature is not supported, a valid identifier must:
- Begin with a Unicode character of class XID_Start or [code]"_"[/code].
- May contain Unicode characters of class XID_Continue in the other positions.
*/
//go:nosplit
func (self class) IsValidIdentifier(s gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_is_valid_identifier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given code point is a valid letter, i.e. it belongs to the Unicode category "L".
*/
//go:nosplit
func (self class) IsValidLetter(unicode gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, unicode)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_is_valid_letter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the string converted to uppercase.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
//go:nosplit
func (self class) StringToUpper(ctx gd.Lifetime, s gd.String, language gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_string_to_upper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the string converted to lowercase.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
//go:nosplit
func (self class) StringToLower(ctx gd.Lifetime, s gd.String, language gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_string_to_lower, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the string converted to title case.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
//go:nosplit
func (self class) StringToTitle(ctx gd.Lifetime, s gd.String, language gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_string_to_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Default implementation of the BiDi algorithm override function. See [enum StructuredTextParser] for more info.
*/
//go:nosplit
func (self class) ParseStructuredText(ctx gd.Lifetime, parser_type classdb.TextServerStructuredTextParser, args gd.Array, text gd.String) gd.ArrayOf[gd.Vector3i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, parser_type)
	callframe.Arg(frame, mmm.Get(args))
	callframe.Arg(frame, mmm.Get(text))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServer.Bind_parse_structured_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector3i](ret)
}
func (self class) AsTextServer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextServer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("TextServer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type FontAntialiasing = classdb.TextServerFontAntialiasing

const (
/*Font glyphs are rasterized as 1-bit bitmaps.*/
	FontAntialiasingNone FontAntialiasing = 0
/*Font glyphs are rasterized as 8-bit grayscale anti-aliased bitmaps.*/
	FontAntialiasingGray FontAntialiasing = 1
/*Font glyphs are rasterized for LCD screens.
LCD subpixel layout is determined by the value of [code]gui/theme/lcd_subpixel_layout[/code] project settings.
LCD subpixel anti-aliasing mode is suitable only for rendering horizontal, unscaled text in 2D.*/
	FontAntialiasingLcd FontAntialiasing = 2
)
type FontLCDSubpixelLayout = classdb.TextServerFontLCDSubpixelLayout

const (
/*Unknown or unsupported subpixel layout, LCD subpixel antialiasing is disabled.*/
	FontLcdSubpixelLayoutNone FontLCDSubpixelLayout = 0
/*Horizontal RGB subpixel layout.*/
	FontLcdSubpixelLayoutHrgb FontLCDSubpixelLayout = 1
/*Horizontal BGR subpixel layout.*/
	FontLcdSubpixelLayoutHbgr FontLCDSubpixelLayout = 2
/*Vertical RGB subpixel layout.*/
	FontLcdSubpixelLayoutVrgb FontLCDSubpixelLayout = 3
/*Vertical BGR subpixel layout.*/
	FontLcdSubpixelLayoutVbgr FontLCDSubpixelLayout = 4
/*Represents the size of the [enum FontLCDSubpixelLayout] enum.*/
	FontLcdSubpixelLayoutMax FontLCDSubpixelLayout = 5
)
type Direction = classdb.TextServerDirection

const (
/*Text direction is determined based on contents and current locale.*/
	DirectionAuto Direction = 0
/*Text is written from left to right.*/
	DirectionLtr Direction = 1
/*Text is written from right to left.*/
	DirectionRtl Direction = 2
/*Text writing direction is the same as base string writing direction. Used for BiDi override only.*/
	DirectionInherited Direction = 3
)
type Orientation = classdb.TextServerOrientation

const (
/*Text is written horizontally.*/
	OrientationHorizontal Orientation = 0
/*Left to right text is written vertically from top to bottom.
Right to left text is written vertically from bottom to top.*/
	OrientationVertical Orientation = 1
)
type JustificationFlag = classdb.TextServerJustificationFlag

const (
/*Do not justify text.*/
	JustificationNone JustificationFlag = 0
/*Justify text by adding and removing kashidas.*/
	JustificationKashida JustificationFlag = 1
/*Justify text by changing width of the spaces between the words.*/
	JustificationWordBound JustificationFlag = 2
/*Remove trailing and leading spaces from the justified text.*/
	JustificationTrimEdgeSpaces JustificationFlag = 4
/*Only apply justification to the part of the text after the last tab.*/
	JustificationAfterLastTab JustificationFlag = 8
/*Apply justification to the trimmed line with ellipsis.*/
	JustificationConstrainEllipsis JustificationFlag = 16
/*Do not apply justification to the last line of the paragraph.*/
	JustificationSkipLastLine JustificationFlag = 32
/*Do not apply justification to the last line of the paragraph with visible characters (takes precedence over [constant JUSTIFICATION_SKIP_LAST_LINE]).*/
	JustificationSkipLastLineWithVisibleChars JustificationFlag = 64
/*Always apply justification to the paragraphs with a single line ([constant JUSTIFICATION_SKIP_LAST_LINE] and [constant JUSTIFICATION_SKIP_LAST_LINE_WITH_VISIBLE_CHARS] are ignored).*/
	JustificationDoNotSkipSingleLine JustificationFlag = 128
)
type AutowrapMode = classdb.TextServerAutowrapMode

const (
/*Autowrap is disabled.*/
	AutowrapOff AutowrapMode = 0
/*Wraps the text inside the node's bounding rectangle by allowing to break lines at arbitrary positions, which is useful when very limited space is available.*/
	AutowrapArbitrary AutowrapMode = 1
/*Wraps the text inside the node's bounding rectangle by soft-breaking between words.*/
	AutowrapWord AutowrapMode = 2
/*Behaves similarly to [constant AUTOWRAP_WORD], but force-breaks a word if that single word does not fit in one line.*/
	AutowrapWordSmart AutowrapMode = 3
)
type LineBreakFlag = classdb.TextServerLineBreakFlag

const (
/*Do not break the line.*/
	BreakNone LineBreakFlag = 0
/*Break the line at the line mandatory break characters (e.g. [code]"\n"[/code]).*/
	BreakMandatory LineBreakFlag = 1
/*Break the line between the words.*/
	BreakWordBound LineBreakFlag = 2
/*Break the line between any unconnected graphemes.*/
	BreakGraphemeBound LineBreakFlag = 4
/*Should be used only in conjunction with [constant BREAK_WORD_BOUND], break the line between any unconnected graphemes, if it's impossible to break it between the words.*/
	BreakAdaptive LineBreakFlag = 8
/*Remove edge spaces from the broken line segments.*/
	BreakTrimEdgeSpaces LineBreakFlag = 16
/*Subtract first line indentation width from all lines after the first one.*/
	BreakTrimIndent LineBreakFlag = 32
)
type VisibleCharactersBehavior = classdb.TextServerVisibleCharactersBehavior

const (
/*Trims text before the shaping. e.g, increasing [member Label.visible_characters] or [member RichTextLabel.visible_characters] value is visually identical to typing the text.*/
	VcCharsBeforeShaping VisibleCharactersBehavior = 0
/*Displays glyphs that are mapped to the first [member Label.visible_characters] or [member RichTextLabel.visible_characters] characters from the beginning of the text.*/
	VcCharsAfterShaping VisibleCharactersBehavior = 1
/*Displays [member Label.visible_ratio] or [member RichTextLabel.visible_ratio] glyphs, starting from the left or from the right, depending on [member Control.layout_direction] value.*/
	VcGlyphsAuto VisibleCharactersBehavior = 2
/*Displays [member Label.visible_ratio] or [member RichTextLabel.visible_ratio] glyphs, starting from the left.*/
	VcGlyphsLtr VisibleCharactersBehavior = 3
/*Displays [member Label.visible_ratio] or [member RichTextLabel.visible_ratio] glyphs, starting from the right.*/
	VcGlyphsRtl VisibleCharactersBehavior = 4
)
type OverrunBehavior = classdb.TextServerOverrunBehavior

const (
/*No text trimming is performed.*/
	OverrunNoTrimming OverrunBehavior = 0
/*Trims the text per character.*/
	OverrunTrimChar OverrunBehavior = 1
/*Trims the text per word.*/
	OverrunTrimWord OverrunBehavior = 2
/*Trims the text per character and adds an ellipsis to indicate that parts are hidden.*/
	OverrunTrimEllipsis OverrunBehavior = 3
/*Trims the text per word and adds an ellipsis to indicate that parts are hidden.*/
	OverrunTrimWordEllipsis OverrunBehavior = 4
)
type TextOverrunFlag = classdb.TextServerTextOverrunFlag

const (
/*No trimming is performed.*/
	OverrunNoTrim TextOverrunFlag = 0
/*Trims the text when it exceeds the given width.*/
	OverrunTrim TextOverrunFlag = 1
/*Trims the text per word instead of per grapheme.*/
	OverrunTrimWordOnly TextOverrunFlag = 2
/*Determines whether an ellipsis should be added at the end of the text.*/
	OverrunAddEllipsis TextOverrunFlag = 4
/*Determines whether the ellipsis at the end of the text is enforced and may not be hidden.*/
	OverrunEnforceEllipsis TextOverrunFlag = 8
/*Accounts for the text being justified before attempting to trim it (see [enum JustificationFlag]).*/
	OverrunJustificationAware TextOverrunFlag = 16
)
type GraphemeFlag = classdb.TextServerGraphemeFlag

const (
/*Grapheme is supported by the font, and can be drawn.*/
	GraphemeIsValid GraphemeFlag = 1
/*Grapheme is part of right-to-left or bottom-to-top run.*/
	GraphemeIsRtl GraphemeFlag = 2
/*Grapheme is not part of source text, it was added by justification process.*/
	GraphemeIsVirtual GraphemeFlag = 4
/*Grapheme is whitespace.*/
	GraphemeIsSpace GraphemeFlag = 8
/*Grapheme is mandatory break point (e.g. [code]"\n"[/code]).*/
	GraphemeIsBreakHard GraphemeFlag = 16
/*Grapheme is optional break point (e.g. space).*/
	GraphemeIsBreakSoft GraphemeFlag = 32
/*Grapheme is the tabulation character.*/
	GraphemeIsTab GraphemeFlag = 64
/*Grapheme is kashida.*/
	GraphemeIsElongation GraphemeFlag = 128
/*Grapheme is punctuation character.*/
	GraphemeIsPunctuation GraphemeFlag = 256
/*Grapheme is underscore character.*/
	GraphemeIsUnderscore GraphemeFlag = 512
/*Grapheme is connected to the previous grapheme. Breaking line before this grapheme is not safe.*/
	GraphemeIsConnected GraphemeFlag = 1024
/*It is safe to insert a U+0640 before this grapheme for elongation.*/
	GraphemeIsSafeToInsertTatweel GraphemeFlag = 2048
/*Grapheme is an object replacement character for the embedded object.*/
	GraphemeIsEmbeddedObject GraphemeFlag = 4096
/*Grapheme is a soft hyphen.*/
	GraphemeIsSoftHyphen GraphemeFlag = 8192
)
type Hinting = classdb.TextServerHinting

const (
/*Disables font hinting (smoother but less crisp).*/
	HintingNone Hinting = 0
/*Use the light font hinting mode.*/
	HintingLight Hinting = 1
/*Use the default font hinting mode (crisper but less smooth).
[b]Note:[/b] This hinting mode changes both horizontal and vertical glyph metrics. If applied to monospace font, some glyphs might have different width.*/
	HintingNormal Hinting = 2
)
type SubpixelPositioning = classdb.TextServerSubpixelPositioning

const (
/*Glyph horizontal position is rounded to the whole pixel size, each glyph is rasterized once.*/
	SubpixelPositioningDisabled SubpixelPositioning = 0
/*Glyph horizontal position is rounded based on font size.
- To one quarter of the pixel size if font size is smaller or equal to [constant SUBPIXEL_POSITIONING_ONE_QUARTER_MAX_SIZE].
- To one half of the pixel size if font size is smaller or equal to [constant SUBPIXEL_POSITIONING_ONE_HALF_MAX_SIZE].
- To the whole pixel size for larger fonts.*/
	SubpixelPositioningAuto SubpixelPositioning = 1
/*Glyph horizontal position is rounded to one half of the pixel size, each glyph is rasterized up to two times.*/
	SubpixelPositioningOneHalf SubpixelPositioning = 2
/*Glyph horizontal position is rounded to one quarter of the pixel size, each glyph is rasterized up to four times.*/
	SubpixelPositioningOneQuarter SubpixelPositioning = 3
/*Maximum font size which will use one half of the pixel subpixel positioning in [constant SUBPIXEL_POSITIONING_AUTO] mode.*/
	SubpixelPositioningOneHalfMaxSize SubpixelPositioning = 20
/*Maximum font size which will use one quarter of the pixel subpixel positioning in [constant SUBPIXEL_POSITIONING_AUTO] mode.*/
	SubpixelPositioningOneQuarterMaxSize SubpixelPositioning = 16
)
type Feature = classdb.TextServerFeature

const (
/*TextServer supports simple text layouts.*/
	FeatureSimpleLayout Feature = 1
/*TextServer supports bidirectional text layouts.*/
	FeatureBidiLayout Feature = 2
/*TextServer supports vertical layouts.*/
	FeatureVerticalLayout Feature = 4
/*TextServer supports complex text shaping.*/
	FeatureShaping Feature = 8
/*TextServer supports justification using kashidas.*/
	FeatureKashidaJustification Feature = 16
/*TextServer supports complex line/word breaking rules (e.g. dictionary based).*/
	FeatureBreakIterators Feature = 32
/*TextServer supports loading bitmap fonts.*/
	FeatureFontBitmap Feature = 64
/*TextServer supports loading dynamic (TrueType, OpeType, etc.) fonts.*/
	FeatureFontDynamic Feature = 128
/*TextServer supports multichannel signed distance field dynamic font rendering.*/
	FeatureFontMsdf Feature = 256
/*TextServer supports loading system fonts.*/
	FeatureFontSystem Feature = 512
/*TextServer supports variable fonts.*/
	FeatureFontVariable Feature = 1024
/*TextServer supports locale dependent and context sensitive case conversion.*/
	FeatureContextSensitiveCaseConversion Feature = 2048
/*TextServer require external data file for some features, see [method load_support_data].*/
	FeatureUseSupportData Feature = 4096
/*TextServer supports UAX #31 identifier validation, see [method is_valid_identifier].*/
	FeatureUnicodeIdentifiers Feature = 8192
/*TextServer supports [url=https://unicode.org/reports/tr36/]Unicode Technical Report #36[/url] and [url=https://unicode.org/reports/tr39/]Unicode Technical Standard #39[/url] based spoof detection features.*/
	FeatureUnicodeSecurity Feature = 16384
)
type ContourPointTag = classdb.TextServerContourPointTag

const (
/*Contour point is on the curve.*/
	ContourCurveTagOn ContourPointTag = 1
/*Contour point isn't on the curve, but serves as a control point for a conic (quadratic) Bézier arc.*/
	ContourCurveTagOffConic ContourPointTag = 0
/*Contour point isn't on the curve, but serves as a control point for a cubic Bézier arc.*/
	ContourCurveTagOffCubic ContourPointTag = 2
)
type SpacingType = classdb.TextServerSpacingType

const (
/*Spacing for each glyph.*/
	SpacingGlyph SpacingType = 0
/*Spacing for the space character.*/
	SpacingSpace SpacingType = 1
/*Spacing at the top of the line.*/
	SpacingTop SpacingType = 2
/*Spacing at the bottom of the line.*/
	SpacingBottom SpacingType = 3
/*Represents the size of the [enum SpacingType] enum.*/
	SpacingMax SpacingType = 4
)
type FontStyle = classdb.TextServerFontStyle

const (
/*Font is bold.*/
	FontBold FontStyle = 1
/*Font is italic or oblique.*/
	FontItalic FontStyle = 2
/*Font have fixed-width characters.*/
	FontFixedWidth FontStyle = 4
)
type StructuredTextParser = classdb.TextServerStructuredTextParser

const (
/*Use default Unicode BiDi algorithm.*/
	StructuredTextDefault StructuredTextParser = 0
/*BiDi override for URI.*/
	StructuredTextUri StructuredTextParser = 1
/*BiDi override for file path.*/
	StructuredTextFile StructuredTextParser = 2
/*BiDi override for email.*/
	StructuredTextEmail StructuredTextParser = 3
/*BiDi override for lists. Structured text options: list separator [String].*/
	StructuredTextList StructuredTextParser = 4
/*BiDi override for GDScript.*/
	StructuredTextGdscript StructuredTextParser = 5
/*User defined structured text BiDi override function.*/
	StructuredTextCustom StructuredTextParser = 6
)
type FixedSizeScaleMode = classdb.TextServerFixedSizeScaleMode

const (
/*Bitmap font is not scaled.*/
	FixedSizeScaleDisable FixedSizeScaleMode = 0
/*Bitmap font is scaled to the closest integer multiple of the font's fixed size. This is the recommended option for pixel art fonts.*/
	FixedSizeScaleIntegerOnly FixedSizeScaleMode = 1
/*Bitmap font is scaled to an arbitrary (fractional) size. This is the recommended option for non-pixel art fonts.*/
	FixedSizeScaleEnabled FixedSizeScaleMode = 2
)