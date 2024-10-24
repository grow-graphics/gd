package Environment

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
Resource for environment nodes (like [WorldEnvironment]) that define multiple environment operations (such as background [Sky] or [Color], ambient light, fog, depth-of-field...). These parameters affect the final render of the scene. The order of these operations is:
- Depth of Field Blur
- Glow
- Tonemap (Auto Exposure)
- Adjustments

*/
type Go [1]classdb.Environment

/*
Sets the intensity of the glow level [param idx]. A value above [code]0.0[/code] enables the level. Each level relies on the previous level. This means that enabling higher glow levels will slow down the glow effect rendering, even if previous levels aren't enabled.
*/
func (self Go) SetGlowLevel(idx int, intensity float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowLevel(gd.Int(idx), gd.Float(intensity))
}

/*
Returns the intensity of the glow level [param idx].
*/
func (self Go) GetGlowLevel(idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetGlowLevel(gd.Int(idx))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Environment
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Environment"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BackgroundMode() classdb.EnvironmentBGMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EnvironmentBGMode(class(self).GetBackground())
}

func (self Go) SetBackgroundMode(value classdb.EnvironmentBGMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackground(value)
}

func (self Go) BackgroundColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetBgColor())
}

func (self Go) SetBackgroundColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBgColor(value)
}

func (self Go) BackgroundEnergyMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBgEnergyMultiplier()))
}

func (self Go) SetBackgroundEnergyMultiplier(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBgEnergyMultiplier(gd.Float(value))
}

func (self Go) BackgroundIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBgIntensity()))
}

func (self Go) SetBackgroundIntensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBgIntensity(gd.Float(value))
}

func (self Go) BackgroundCanvasMaxLayer() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCanvasMaxLayer()))
}

func (self Go) SetBackgroundCanvasMaxLayer(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCanvasMaxLayer(gd.Int(value))
}

func (self Go) BackgroundCameraFeedId() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCameraFeedId()))
}

func (self Go) SetBackgroundCameraFeedId(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCameraFeedId(gd.Int(value))
}

func (self Go) Sky() gdclass.Sky {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Sky(class(self).GetSky(gc))
}

func (self Go) SetSky(value gdclass.Sky) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSky(value)
}

func (self Go) SkyCustomFov() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSkyCustomFov()))
}

func (self Go) SetSkyCustomFov(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkyCustomFov(gd.Float(value))
}

func (self Go) SkyRotation() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetSkyRotation())
}

func (self Go) SetSkyRotation(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkyRotation(value)
}

func (self Go) AmbientLightSource() classdb.EnvironmentAmbientSource {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EnvironmentAmbientSource(class(self).GetAmbientSource())
}

func (self Go) SetAmbientLightSource(value classdb.EnvironmentAmbientSource) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAmbientSource(value)
}

func (self Go) AmbientLightColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetAmbientLightColor())
}

func (self Go) SetAmbientLightColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAmbientLightColor(value)
}

func (self Go) AmbientLightSkyContribution() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAmbientLightSkyContribution()))
}

func (self Go) SetAmbientLightSkyContribution(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAmbientLightSkyContribution(gd.Float(value))
}

func (self Go) AmbientLightEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAmbientLightEnergy()))
}

func (self Go) SetAmbientLightEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAmbientLightEnergy(gd.Float(value))
}

func (self Go) ReflectedLightSource() classdb.EnvironmentReflectionSource {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EnvironmentReflectionSource(class(self).GetReflectionSource())
}

func (self Go) SetReflectedLightSource(value classdb.EnvironmentReflectionSource) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReflectionSource(value)
}

func (self Go) TonemapMode() classdb.EnvironmentToneMapper {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EnvironmentToneMapper(class(self).GetTonemapper())
}

func (self Go) SetTonemapMode(value classdb.EnvironmentToneMapper) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTonemapper(value)
}

func (self Go) TonemapExposure() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTonemapExposure()))
}

func (self Go) SetTonemapExposure(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTonemapExposure(gd.Float(value))
}

func (self Go) TonemapWhite() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTonemapWhite()))
}

func (self Go) SetTonemapWhite(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTonemapWhite(gd.Float(value))
}

func (self Go) SsrEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSsrEnabled())
}

func (self Go) SetSsrEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsrEnabled(value)
}

func (self Go) SsrMaxSteps() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSsrMaxSteps()))
}

func (self Go) SetSsrMaxSteps(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsrMaxSteps(gd.Int(value))
}

func (self Go) SsrFadeIn() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsrFadeIn()))
}

func (self Go) SetSsrFadeIn(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsrFadeIn(gd.Float(value))
}

func (self Go) SsrFadeOut() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsrFadeOut()))
}

func (self Go) SetSsrFadeOut(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsrFadeOut(gd.Float(value))
}

func (self Go) SsrDepthTolerance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsrDepthTolerance()))
}

func (self Go) SetSsrDepthTolerance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsrDepthTolerance(gd.Float(value))
}

func (self Go) SsaoEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSsaoEnabled())
}

func (self Go) SetSsaoEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoEnabled(value)
}

func (self Go) SsaoRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsaoRadius()))
}

func (self Go) SetSsaoRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoRadius(gd.Float(value))
}

func (self Go) SsaoIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsaoIntensity()))
}

func (self Go) SetSsaoIntensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoIntensity(gd.Float(value))
}

func (self Go) SsaoPower() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsaoPower()))
}

func (self Go) SetSsaoPower(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoPower(gd.Float(value))
}

func (self Go) SsaoDetail() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsaoDetail()))
}

func (self Go) SetSsaoDetail(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoDetail(gd.Float(value))
}

func (self Go) SsaoHorizon() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsaoHorizon()))
}

func (self Go) SetSsaoHorizon(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoHorizon(gd.Float(value))
}

func (self Go) SsaoSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsaoSharpness()))
}

func (self Go) SetSsaoSharpness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoSharpness(gd.Float(value))
}

func (self Go) SsaoLightAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsaoDirectLightAffect()))
}

func (self Go) SetSsaoLightAffect(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoDirectLightAffect(gd.Float(value))
}

func (self Go) SsaoAoChannelAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsaoAoChannelAffect()))
}

func (self Go) SetSsaoAoChannelAffect(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsaoAoChannelAffect(gd.Float(value))
}

func (self Go) SsilEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSsilEnabled())
}

func (self Go) SetSsilEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsilEnabled(value)
}

func (self Go) SsilRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsilRadius()))
}

func (self Go) SetSsilRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsilRadius(gd.Float(value))
}

func (self Go) SsilIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsilIntensity()))
}

func (self Go) SetSsilIntensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsilIntensity(gd.Float(value))
}

func (self Go) SsilSharpness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsilSharpness()))
}

func (self Go) SetSsilSharpness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsilSharpness(gd.Float(value))
}

func (self Go) SsilNormalRejection() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSsilNormalRejection()))
}

func (self Go) SetSsilNormalRejection(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSsilNormalRejection(gd.Float(value))
}

func (self Go) SdfgiEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSdfgiEnabled())
}

func (self Go) SetSdfgiEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiEnabled(value)
}

func (self Go) SdfgiUseOcclusion() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSdfgiUsingOcclusion())
}

func (self Go) SetSdfgiUseOcclusion(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiUseOcclusion(value)
}

func (self Go) SdfgiReadSkyLight() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSdfgiReadingSkyLight())
}

func (self Go) SetSdfgiReadSkyLight(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiReadSkyLight(value)
}

func (self Go) SdfgiBounceFeedback() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSdfgiBounceFeedback()))
}

func (self Go) SetSdfgiBounceFeedback(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiBounceFeedback(gd.Float(value))
}

func (self Go) SdfgiCascades() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSdfgiCascades()))
}

func (self Go) SetSdfgiCascades(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiCascades(gd.Int(value))
}

func (self Go) SdfgiMinCellSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSdfgiMinCellSize()))
}

func (self Go) SetSdfgiMinCellSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiMinCellSize(gd.Float(value))
}

func (self Go) SdfgiCascade0Distance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSdfgiCascade0Distance()))
}

func (self Go) SetSdfgiCascade0Distance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiCascade0Distance(gd.Float(value))
}

func (self Go) SdfgiMaxDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSdfgiMaxDistance()))
}

func (self Go) SetSdfgiMaxDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiMaxDistance(gd.Float(value))
}

func (self Go) SdfgiYScale() classdb.EnvironmentSDFGIYScale {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EnvironmentSDFGIYScale(class(self).GetSdfgiYScale())
}

func (self Go) SetSdfgiYScale(value classdb.EnvironmentSDFGIYScale) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiYScale(value)
}

func (self Go) SdfgiEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSdfgiEnergy()))
}

func (self Go) SetSdfgiEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiEnergy(gd.Float(value))
}

func (self Go) SdfgiNormalBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSdfgiNormalBias()))
}

func (self Go) SetSdfgiNormalBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiNormalBias(gd.Float(value))
}

func (self Go) SdfgiProbeBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSdfgiProbeBias()))
}

func (self Go) SetSdfgiProbeBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSdfgiProbeBias(gd.Float(value))
}

func (self Go) GlowEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsGlowEnabled())
}

func (self Go) SetGlowEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowEnabled(value)
}

func (self Go) GlowNormalized() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsGlowNormalized())
}

func (self Go) SetGlowNormalized(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowNormalized(value)
}

func (self Go) GlowIntensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlowIntensity()))
}

func (self Go) SetGlowIntensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowIntensity(gd.Float(value))
}

func (self Go) GlowStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlowStrength()))
}

func (self Go) SetGlowStrength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowStrength(gd.Float(value))
}

func (self Go) GlowMix() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlowMix()))
}

func (self Go) SetGlowMix(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowMix(gd.Float(value))
}

func (self Go) GlowBloom() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlowBloom()))
}

func (self Go) SetGlowBloom(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowBloom(gd.Float(value))
}

func (self Go) GlowBlendMode() classdb.EnvironmentGlowBlendMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EnvironmentGlowBlendMode(class(self).GetGlowBlendMode())
}

func (self Go) SetGlowBlendMode(value classdb.EnvironmentGlowBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowBlendMode(value)
}

func (self Go) GlowHdrThreshold() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlowHdrBleedThreshold()))
}

func (self Go) SetGlowHdrThreshold(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowHdrBleedThreshold(gd.Float(value))
}

func (self Go) GlowHdrScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlowHdrBleedScale()))
}

func (self Go) SetGlowHdrScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowHdrBleedScale(gd.Float(value))
}

func (self Go) GlowHdrLuminanceCap() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlowHdrLuminanceCap()))
}

func (self Go) SetGlowHdrLuminanceCap(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowHdrLuminanceCap(gd.Float(value))
}

func (self Go) GlowMapStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGlowMapStrength()))
}

func (self Go) SetGlowMapStrength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowMapStrength(gd.Float(value))
}

func (self Go) GlowMap() gdclass.Texture {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture(class(self).GetGlowMap(gc))
}

func (self Go) SetGlowMap(value gdclass.Texture) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlowMap(value)
}

func (self Go) FogEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFogEnabled())
}

func (self Go) SetFogEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogEnabled(value)
}

func (self Go) FogMode() classdb.EnvironmentFogMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.EnvironmentFogMode(class(self).GetFogMode())
}

func (self Go) SetFogMode(value classdb.EnvironmentFogMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogMode(value)
}

func (self Go) FogLightColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetFogLightColor())
}

func (self Go) SetFogLightColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogLightColor(value)
}

func (self Go) FogLightEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogLightEnergy()))
}

func (self Go) SetFogLightEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogLightEnergy(gd.Float(value))
}

func (self Go) FogSunScatter() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogSunScatter()))
}

func (self Go) SetFogSunScatter(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogSunScatter(gd.Float(value))
}

func (self Go) FogDensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogDensity()))
}

func (self Go) SetFogDensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogDensity(gd.Float(value))
}

func (self Go) FogAerialPerspective() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogAerialPerspective()))
}

func (self Go) SetFogAerialPerspective(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogAerialPerspective(gd.Float(value))
}

func (self Go) FogSkyAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogSkyAffect()))
}

func (self Go) SetFogSkyAffect(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogSkyAffect(gd.Float(value))
}

func (self Go) FogHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogHeight()))
}

func (self Go) SetFogHeight(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogHeight(gd.Float(value))
}

func (self Go) FogHeightDensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogHeightDensity()))
}

func (self Go) SetFogHeightDensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogHeightDensity(gd.Float(value))
}

func (self Go) FogDepthCurve() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogDepthCurve()))
}

func (self Go) SetFogDepthCurve(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogDepthCurve(gd.Float(value))
}

func (self Go) FogDepthBegin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogDepthBegin()))
}

func (self Go) SetFogDepthBegin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogDepthBegin(gd.Float(value))
}

func (self Go) FogDepthEnd() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFogDepthEnd()))
}

func (self Go) SetFogDepthEnd(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFogDepthEnd(gd.Float(value))
}

func (self Go) VolumetricFogEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVolumetricFogEnabled())
}

func (self Go) SetVolumetricFogEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogEnabled(value)
}

func (self Go) VolumetricFogDensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogDensity()))
}

func (self Go) SetVolumetricFogDensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogDensity(gd.Float(value))
}

func (self Go) VolumetricFogAlbedo() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetVolumetricFogAlbedo())
}

func (self Go) SetVolumetricFogAlbedo(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogAlbedo(value)
}

func (self Go) VolumetricFogEmission() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetVolumetricFogEmission())
}

func (self Go) SetVolumetricFogEmission(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogEmission(value)
}

func (self Go) VolumetricFogEmissionEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogEmissionEnergy()))
}

func (self Go) SetVolumetricFogEmissionEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogEmissionEnergy(gd.Float(value))
}

func (self Go) VolumetricFogGiInject() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogGiInject()))
}

func (self Go) SetVolumetricFogGiInject(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogGiInject(gd.Float(value))
}

func (self Go) VolumetricFogAnisotropy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogAnisotropy()))
}

func (self Go) SetVolumetricFogAnisotropy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogAnisotropy(gd.Float(value))
}

func (self Go) VolumetricFogLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogLength()))
}

func (self Go) SetVolumetricFogLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogLength(gd.Float(value))
}

func (self Go) VolumetricFogDetailSpread() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogDetailSpread()))
}

func (self Go) SetVolumetricFogDetailSpread(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogDetailSpread(gd.Float(value))
}

func (self Go) VolumetricFogAmbientInject() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogAmbientInject()))
}

func (self Go) SetVolumetricFogAmbientInject(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogAmbientInject(gd.Float(value))
}

func (self Go) VolumetricFogSkyAffect() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogSkyAffect()))
}

func (self Go) SetVolumetricFogSkyAffect(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogSkyAffect(gd.Float(value))
}

func (self Go) VolumetricFogTemporalReprojectionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVolumetricFogTemporalReprojectionEnabled())
}

func (self Go) SetVolumetricFogTemporalReprojectionEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogTemporalReprojectionEnabled(value)
}

func (self Go) VolumetricFogTemporalReprojectionAmount() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumetricFogTemporalReprojectionAmount()))
}

func (self Go) SetVolumetricFogTemporalReprojectionAmount(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumetricFogTemporalReprojectionAmount(gd.Float(value))
}

func (self Go) AdjustmentEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAdjustmentEnabled())
}

func (self Go) SetAdjustmentEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdjustmentEnabled(value)
}

func (self Go) AdjustmentBrightness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAdjustmentBrightness()))
}

func (self Go) SetAdjustmentBrightness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdjustmentBrightness(gd.Float(value))
}

func (self Go) AdjustmentContrast() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAdjustmentContrast()))
}

func (self Go) SetAdjustmentContrast(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdjustmentContrast(gd.Float(value))
}

func (self Go) AdjustmentSaturation() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAdjustmentSaturation()))
}

func (self Go) SetAdjustmentSaturation(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdjustmentSaturation(gd.Float(value))
}

func (self Go) AdjustmentColorCorrection() gdclass.Texture {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture(class(self).GetAdjustmentColorCorrection(gc))
}

func (self Go) SetAdjustmentColorCorrection(value gdclass.Texture) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdjustmentColorCorrection(value)
}

//go:nosplit
func (self class) SetBackground(mode classdb.EnvironmentBGMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackground() classdb.EnvironmentBGMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentBGMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSky(sky gdclass.Sky)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(sky[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSky(ctx gd.Lifetime) gdclass.Sky {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Sky
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyCustomFov(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyCustomFov() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sky_custom_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyRotation(euler_radians gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, euler_radians)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sky_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyRotation() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sky_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBgColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBgColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBgEnergyMultiplier(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_bg_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBgEnergyMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_bg_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBgIntensity(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_bg_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBgIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_bg_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCanvasMaxLayer(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCanvasMaxLayer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_canvas_max_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCameraFeedId(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraFeedId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_camera_feed_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientLightColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ambient_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientLightColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ambient_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientSource(source classdb.EnvironmentAmbientSource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ambient_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientSource() classdb.EnvironmentAmbientSource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentAmbientSource](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ambient_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientLightEnergy(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ambient_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientLightEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ambient_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientLightSkyContribution(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ambient_light_sky_contribution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientLightSkyContribution() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ambient_light_sky_contribution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReflectionSource(source classdb.EnvironmentReflectionSource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_reflection_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReflectionSource() classdb.EnvironmentReflectionSource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentReflectionSource](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_reflection_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTonemapper(mode classdb.EnvironmentToneMapper)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_tonemapper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTonemapper() classdb.EnvironmentToneMapper {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentToneMapper](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_tonemapper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTonemapExposure(exposure gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exposure)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_tonemap_exposure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTonemapExposure() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_tonemap_exposure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTonemapWhite(white gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, white)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_tonemap_white, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTonemapWhite() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_tonemap_white, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSsrEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_ssr_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrMaxSteps(max_steps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_steps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_max_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsrMaxSteps() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssr_max_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrFadeIn(fade_in gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fade_in)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_fade_in, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsrFadeIn() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssr_fade_in, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrFadeOut(fade_out gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fade_out)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_fade_out, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsrFadeOut() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssr_fade_out, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsrDepthTolerance(depth_tolerance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, depth_tolerance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssr_depth_tolerance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsrDepthTolerance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssr_depth_tolerance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSsaoEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_ssao_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoIntensity(intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoPower(power gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, power)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_power, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoPower() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_power, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoDetail(detail gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoDetail() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoHorizon(horizon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, horizon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_horizon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoHorizon() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_horizon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoSharpness(sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoDirectLightAffect(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_direct_light_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoDirectLightAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_direct_light_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsaoAoChannelAffect(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssao_ao_channel_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsaoAoChannelAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssao_ao_channel_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSsilEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_ssil_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsilRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssil_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilIntensity(intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsilIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssil_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilSharpness(sharpness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sharpness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsilSharpness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssil_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSsilNormalRejection(normal_rejection gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, normal_rejection)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_ssil_normal_rejection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSsilNormalRejection() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_ssil_normal_rejection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSdfgiEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_sdfgi_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiCascades(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_cascades, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiCascades() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_cascades, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiMinCellSize(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_min_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiMinCellSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_min_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiMaxDistance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiMaxDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiCascade0Distance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_cascade0_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiCascade0Distance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_cascade0_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiYScale(scale classdb.EnvironmentSDFGIYScale)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_y_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiYScale() classdb.EnvironmentSDFGIYScale {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentSDFGIYScale](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_y_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiUseOcclusion(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_use_occlusion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSdfgiUsingOcclusion() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_sdfgi_using_occlusion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiBounceFeedback(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_bounce_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiBounceFeedback() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_bounce_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiReadSkyLight(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_read_sky_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSdfgiReadingSkyLight() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_sdfgi_reading_sky_light, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiEnergy(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiNormalBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiNormalBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_normal_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSdfgiProbeBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_sdfgi_probe_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSdfgiProbeBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_sdfgi_probe_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGlowEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_glow_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the intensity of the glow level [param idx]. A value above [code]0.0[/code] enables the level. Each level relies on the previous level. This means that enabling higher glow levels will slow down the glow effect rendering, even if previous levels aren't enabled.
*/
//go:nosplit
func (self class) SetGlowLevel(idx gd.Int, intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the intensity of the glow level [param idx].
*/
//go:nosplit
func (self class) GetGlowLevel(idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowNormalized(normalize bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, normalize)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGlowNormalized() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_glow_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowIntensity(intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowStrength(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowMix(mix gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mix)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowMix() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowBloom(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_bloom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowBloom() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_bloom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowBlendMode(mode classdb.EnvironmentGlowBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowBlendMode() classdb.EnvironmentGlowBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentGlowBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowHdrBleedThreshold(threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_hdr_bleed_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowHdrBleedThreshold() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_hdr_bleed_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowHdrBleedScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_hdr_bleed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowHdrBleedScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_hdr_bleed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowHdrLuminanceCap(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_hdr_luminance_cap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowHdrLuminanceCap() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_hdr_luminance_cap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowMapStrength(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_map_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowMapStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_map_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlowMap(mode gdclass.Texture)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mode[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_glow_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlowMap(ctx gd.Lifetime) gdclass.Texture {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_glow_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFogEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogMode(mode classdb.EnvironmentFogMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogMode() classdb.EnvironmentFogMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.EnvironmentFogMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogLightColor(light_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogLightColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_light_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogLightEnergy(light_energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light_energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogLightEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_light_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogSunScatter(sun_scatter gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sun_scatter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_sun_scatter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogSunScatter() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_sun_scatter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogDensity(density gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogDensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogHeight(height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogHeightDensity(height_density gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height_density)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_height_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogHeightDensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_height_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogAerialPerspective(aerial_perspective gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, aerial_perspective)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_aerial_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogAerialPerspective() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_aerial_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogSkyAffect(sky_affect gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sky_affect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogSkyAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogDepthCurve(curve gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_depth_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogDepthCurve() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_depth_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogDepthBegin(begin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_depth_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogDepthBegin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_depth_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFogDepthEnd(end gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, end)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_fog_depth_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFogDepthEnd() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_fog_depth_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVolumetricFogEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_volumetric_fog_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogEmission(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogEmission() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogAlbedo(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogAlbedo() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogDensity(density gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogDensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogEmissionEnergy(begin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_emission_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogEmissionEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_emission_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogAnisotropy(anisotropy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, anisotropy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogAnisotropy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_anisotropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogLength(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogDetailSpread(detail_spread gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail_spread)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_detail_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogDetailSpread() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_detail_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogGiInject(gi_inject gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gi_inject)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_gi_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogGiInject() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_gi_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogAmbientInject(enabled gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_ambient_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogAmbientInject() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_ambient_inject, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogSkyAffect(sky_affect gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sky_affect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogSkyAffect() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_sky_affect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogTemporalReprojectionEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_temporal_reprojection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVolumetricFogTemporalReprojectionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_volumetric_fog_temporal_reprojection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumetricFogTemporalReprojectionAmount(temporal_reprojection_amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, temporal_reprojection_amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_volumetric_fog_temporal_reprojection_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumetricFogTemporalReprojectionAmount() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_volumetric_fog_temporal_reprojection_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAdjustmentEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_is_adjustment_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentBrightness(brightness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, brightness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_brightness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdjustmentBrightness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_adjustment_brightness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentContrast(contrast gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, contrast)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_contrast, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdjustmentContrast() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_adjustment_contrast, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentSaturation(saturation gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, saturation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_saturation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdjustmentSaturation() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_adjustment_saturation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdjustmentColorCorrection(color_correction gdclass.Texture)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(color_correction[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_set_adjustment_color_correction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdjustmentColorCorrection(ctx gd.Lifetime) gdclass.Texture {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Environment.Bind_get_adjustment_color_correction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsEnvironment() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEnvironment() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("Environment", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type BGMode = classdb.EnvironmentBGMode

const (
/*Clears the background using the clear color defined in [member ProjectSettings.rendering/environment/defaults/default_clear_color].*/
	BgClearColor BGMode = 0
/*Clears the background using a custom clear color.*/
	BgColor BGMode = 1
/*Displays a user-defined sky in the background.*/
	BgSky BGMode = 2
/*Displays a [CanvasLayer] in the background.*/
	BgCanvas BGMode = 3
/*Keeps on screen every pixel drawn in the background. This is the fastest background mode, but it can only be safely used in fully-interior scenes (no visible sky or sky reflections). If enabled in a scene where the background is visible, "ghost trail" artifacts will be visible when moving the camera.*/
	BgKeep BGMode = 4
/*Displays a camera feed in the background.*/
	BgCameraFeed BGMode = 5
/*Represents the size of the [enum BGMode] enum.*/
	BgMax BGMode = 6
)
type AmbientSource = classdb.EnvironmentAmbientSource

const (
/*Gather ambient light from whichever source is specified as the background.*/
	AmbientSourceBg AmbientSource = 0
/*Disable ambient light. This provides a slight performance boost over [constant AMBIENT_SOURCE_SKY].*/
	AmbientSourceDisabled AmbientSource = 1
/*Specify a specific [Color] for ambient light. This provides a slight performance boost over [constant AMBIENT_SOURCE_SKY].*/
	AmbientSourceColor AmbientSource = 2
/*Gather ambient light from the [Sky] regardless of what the background is.*/
	AmbientSourceSky AmbientSource = 3
)
type ReflectionSource = classdb.EnvironmentReflectionSource

const (
/*Use the background for reflections.*/
	ReflectionSourceBg ReflectionSource = 0
/*Disable reflections. This provides a slight performance boost over other options.*/
	ReflectionSourceDisabled ReflectionSource = 1
/*Use the [Sky] for reflections regardless of what the background is.*/
	ReflectionSourceSky ReflectionSource = 2
)
type ToneMapper = classdb.EnvironmentToneMapper

const (
/*Linear tonemapper operator. Reads the linear data and passes it on unmodified. This can cause bright lighting to look blown out, with noticeable clipping in the output colors.*/
	ToneMapperLinear ToneMapper = 0
/*Reinhardt tonemapper operator. Performs a variation on rendered pixels' colors by this formula: [code]color = color / (1 + color)[/code]. This avoids clipping bright highlights, but the resulting image can look a bit dull.*/
	ToneMapperReinhardt ToneMapper = 1
/*Filmic tonemapper operator. This avoids clipping bright highlights, with a resulting image that usually looks more vivid than [constant TONE_MAPPER_REINHARDT].*/
	ToneMapperFilmic ToneMapper = 2
/*Use the Academy Color Encoding System tonemapper. ACES is slightly more expensive than other options, but it handles bright lighting in a more realistic fashion by desaturating it as it becomes brighter. ACES typically has a more contrasted output compared to [constant TONE_MAPPER_REINHARDT] and [constant TONE_MAPPER_FILMIC].
[b]Note:[/b] This tonemapping operator is called "ACES Fitted" in Godot 3.x.*/
	ToneMapperAces ToneMapper = 3
)
type GlowBlendMode = classdb.EnvironmentGlowBlendMode

const (
/*Additive glow blending mode. Mostly used for particles, glows (bloom), lens flare, bright sources.*/
	GlowBlendModeAdditive GlowBlendMode = 0
/*Screen glow blending mode. Increases brightness, used frequently with bloom.*/
	GlowBlendModeScreen GlowBlendMode = 1
/*Soft light glow blending mode. Modifies contrast, exposes shadows and highlights (vivid bloom).*/
	GlowBlendModeSoftlight GlowBlendMode = 2
/*Replace glow blending mode. Replaces all pixels' color by the glow value. This can be used to simulate a full-screen blur effect by tweaking the glow parameters to match the original image's brightness.*/
	GlowBlendModeReplace GlowBlendMode = 3
/*Mixes the glow with the underlying color to avoid increasing brightness as much while still maintaining a glow effect.*/
	GlowBlendModeMix GlowBlendMode = 4
)
type FogMode = classdb.EnvironmentFogMode

const (
/*Use a physically-based fog model defined primarily by fog density.*/
	FogModeExponential FogMode = 0
/*Use a simple fog model defined by start and end positions and a custom curve. While not physically accurate, this model can be useful when you need more artistic control.*/
	FogModeDepth FogMode = 1
)
type SDFGIYScale = classdb.EnvironmentSDFGIYScale

const (
/*Use 50% scale for SDFGI on the Y (vertical) axis. SDFGI cells will be twice as short as they are wide. This allows providing increased GI detail and reduced light leaking with thin floors and ceilings. This is usually the best choice for scenes that don't feature much verticality.*/
	SdfgiYScale50Percent SDFGIYScale = 0
/*Use 75% scale for SDFGI on the Y (vertical) axis. This is a balance between the 50% and 100% SDFGI Y scales.*/
	SdfgiYScale75Percent SDFGIYScale = 1
/*Use 100% scale for SDFGI on the Y (vertical) axis. SDFGI cells will be as tall as they are wide. This is usually the best choice for highly vertical scenes. The downside is that light leaking may become more noticeable with thin floors and ceilings.*/
	SdfgiYScale100Percent SDFGIYScale = 2
)