package LightmapGI

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [LightmapGI] node is used to compute and store baked lightmaps. Lightmaps are used to provide high-quality indirect lighting with very little light leaking. [LightmapGI] can also provide rough reflections using spherical harmonics if [member directional] is enabled. Dynamic objects can receive indirect lighting thanks to [i]light probes[/i], which can be automatically placed by setting [member generate_probes_subdiv] to a value other than [constant GENERATE_PROBES_DISABLED]. Additional lightmap probes can also be added by creating [LightmapProbe] nodes. The downside is that lightmaps are fully static and cannot be baked in an exported project. Baking a [LightmapGI] node is also slower compared to [VoxelGI].
[b]Procedural generation:[/b] Lightmap baking functionality is only available in the editor. This means [LightmapGI] is not suited to procedurally generated or user-built levels. For procedurally generated or user-built levels, use [VoxelGI] or SDFGI instead (see [member Environment.sdfgi_enabled]).
[b]Performance:[/b] [LightmapGI] provides the best possible run-time performance for global illumination. It is suitable for low-end hardware including integrated graphics and mobile devices.
[b]Note:[/b] Due to how lightmaps work, most properties only have a visible effect once lightmaps are baked again.
[b]Note:[/b] Lightmap baking on [CSGShape3D]s and [PrimitiveMesh]es is not supported, as these cannot store UV2 data required for baking.
[b]Note:[/b] If no custom lightmappers are installed, [LightmapGI] can only be baked from devices that support the Forward+ or Mobile rendering backends.

*/
type Go [1]classdb.LightmapGI
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.LightmapGI
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("LightmapGI"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Quality() classdb.LightmapGIBakeQuality {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.LightmapGIBakeQuality(class(self).GetBakeQuality())
}

func (self Go) SetQuality(value classdb.LightmapGIBakeQuality) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakeQuality(value)
}

func (self Go) Bounces() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBounces()))
}

func (self Go) SetBounces(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBounces(gd.Int(value))
}

func (self Go) BounceIndirectEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBounceIndirectEnergy()))
}

func (self Go) SetBounceIndirectEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBounceIndirectEnergy(gd.Float(value))
}

func (self Go) Directional() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDirectional())
}

func (self Go) SetDirectional(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDirectional(value)
}

func (self Go) UseTextureForBounces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingTextureForBounces())
}

func (self Go) SetUseTextureForBounces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseTextureForBounces(value)
}

func (self Go) Interior() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsInterior())
}

func (self Go) SetInterior(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInterior(value)
}

func (self Go) UseDenoiser() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingDenoiser())
}

func (self Go) SetUseDenoiser(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseDenoiser(value)
}

func (self Go) DenoiserStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDenoiserStrength()))
}

func (self Go) SetDenoiserStrength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDenoiserStrength(gd.Float(value))
}

func (self Go) DenoiserRange() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetDenoiserRange()))
}

func (self Go) SetDenoiserRange(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDenoiserRange(gd.Int(value))
}

func (self Go) Bias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBias()))
}

func (self Go) SetBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBias(gd.Float(value))
}

func (self Go) TexelScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTexelScale()))
}

func (self Go) SetTexelScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTexelScale(gd.Float(value))
}

func (self Go) MaxTextureSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxTextureSize()))
}

func (self Go) SetMaxTextureSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxTextureSize(gd.Int(value))
}

func (self Go) EnvironmentMode() classdb.LightmapGIEnvironmentMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.LightmapGIEnvironmentMode(class(self).GetEnvironmentMode())
}

func (self Go) SetEnvironmentMode(value classdb.LightmapGIEnvironmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnvironmentMode(value)
}

func (self Go) EnvironmentCustomSky() gdclass.Sky {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Sky(class(self).GetEnvironmentCustomSky(gc))
}

func (self Go) SetEnvironmentCustomSky(value gdclass.Sky) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnvironmentCustomSky(value)
}

func (self Go) EnvironmentCustomColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetEnvironmentCustomColor())
}

func (self Go) SetEnvironmentCustomColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnvironmentCustomColor(value)
}

func (self Go) EnvironmentCustomEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEnvironmentCustomEnergy()))
}

func (self Go) SetEnvironmentCustomEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnvironmentCustomEnergy(gd.Float(value))
}

func (self Go) CameraAttributes() gdclass.CameraAttributes {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.CameraAttributes(class(self).GetCameraAttributes(gc))
}

func (self Go) SetCameraAttributes(value gdclass.CameraAttributes) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCameraAttributes(value)
}

func (self Go) GenerateProbesSubdiv() classdb.LightmapGIGenerateProbes {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.LightmapGIGenerateProbes(class(self).GetGenerateProbes())
}

func (self Go) SetGenerateProbesSubdiv(value classdb.LightmapGIGenerateProbes) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGenerateProbes(value)
}

func (self Go) LightData() gdclass.LightmapGIData {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.LightmapGIData(class(self).GetLightData(gc))
}

func (self Go) SetLightData(value gdclass.LightmapGIData) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLightData(value)
}

//go:nosplit
func (self class) SetLightData(data gdclass.LightmapGIData)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_light_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightData(ctx gd.Lifetime) gdclass.LightmapGIData {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_light_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.LightmapGIData
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakeQuality(bake_quality classdb.LightmapGIBakeQuality)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bake_quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_bake_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakeQuality() classdb.LightmapGIBakeQuality {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.LightmapGIBakeQuality](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_bake_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBounces(bounces gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bounces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBounces() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBounceIndirectEnergy(bounce_indirect_energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bounce_indirect_energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_bounce_indirect_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBounceIndirectEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_bounce_indirect_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGenerateProbes(subdivision classdb.LightmapGIGenerateProbes)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, subdivision)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_generate_probes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGenerateProbes() classdb.LightmapGIGenerateProbes {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.LightmapGIGenerateProbes](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_generate_probes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironmentMode(mode classdb.LightmapGIEnvironmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_environment_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironmentMode() classdb.LightmapGIEnvironmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.LightmapGIEnvironmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_environment_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironmentCustomSky(sky gdclass.Sky)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(sky[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_environment_custom_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironmentCustomSky(ctx gd.Lifetime) gdclass.Sky {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_environment_custom_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Sky
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironmentCustomColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_environment_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironmentCustomColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_environment_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironmentCustomEnergy(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_environment_custom_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironmentCustomEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_environment_custom_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTexelScale(texel_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texel_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_texel_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexelScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_texel_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxTextureSize(max_texture_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_texture_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_max_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxTextureSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_max_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseDenoiser(use_denoiser bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_denoiser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_use_denoiser, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingDenoiser() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_is_using_denoiser, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDenoiserStrength(denoiser_strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, denoiser_strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_denoiser_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDenoiserStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_denoiser_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDenoiserRange(denoiser_range gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, denoiser_range)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_denoiser_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDenoiserRange() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_denoiser_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInterior(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsInterior() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_is_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDirectional(directional bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, directional)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_directional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDirectional() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_is_directional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseTextureForBounces(use_texture_for_bounces bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_texture_for_bounces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_use_texture_for_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingTextureForBounces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_is_using_texture_for_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCameraAttributes(camera_attributes gdclass.CameraAttributes)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(camera_attributes[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraAttributes(ctx gd.Lifetime) gdclass.CameraAttributes {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.CameraAttributes
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsLightmapGI() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLightmapGI() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}
func init() {classdb.Register("LightmapGI", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type BakeQuality = classdb.LightmapGIBakeQuality

const (
/*Low bake quality (fastest bake times). The quality of this preset can be adjusted by changing [member ProjectSettings.rendering/lightmapping/bake_quality/low_quality_ray_count] and [member ProjectSettings.rendering/lightmapping/bake_quality/low_quality_probe_ray_count].*/
	BakeQualityLow BakeQuality = 0
/*Medium bake quality (fast bake times). The quality of this preset can be adjusted by changing [member ProjectSettings.rendering/lightmapping/bake_quality/medium_quality_ray_count] and [member ProjectSettings.rendering/lightmapping/bake_quality/medium_quality_probe_ray_count].*/
	BakeQualityMedium BakeQuality = 1
/*High bake quality (slow bake times). The quality of this preset can be adjusted by changing [member ProjectSettings.rendering/lightmapping/bake_quality/high_quality_ray_count] and [member ProjectSettings.rendering/lightmapping/bake_quality/high_quality_probe_ray_count].*/
	BakeQualityHigh BakeQuality = 2
/*Highest bake quality (slowest bake times). The quality of this preset can be adjusted by changing [member ProjectSettings.rendering/lightmapping/bake_quality/ultra_quality_ray_count] and [member ProjectSettings.rendering/lightmapping/bake_quality/ultra_quality_probe_ray_count].*/
	BakeQualityUltra BakeQuality = 3
)
type GenerateProbes = classdb.LightmapGIGenerateProbes

const (
/*Don't generate lightmap probes for lighting dynamic objects.*/
	GenerateProbesDisabled GenerateProbes = 0
/*Lowest level of subdivision (fastest bake times, smallest file sizes).*/
	GenerateProbesSubdiv4 GenerateProbes = 1
/*Low level of subdivision (fast bake times, small file sizes).*/
	GenerateProbesSubdiv8 GenerateProbes = 2
/*High level of subdivision (slow bake times, large file sizes).*/
	GenerateProbesSubdiv16 GenerateProbes = 3
/*Highest level of subdivision (slowest bake times, largest file sizes).*/
	GenerateProbesSubdiv32 GenerateProbes = 4
)
type BakeError = classdb.LightmapGIBakeError

const (
/*Lightmap baking was successful.*/
	BakeErrorOk BakeError = 0
/*Lightmap baking failed because the root node for the edited scene could not be accessed.*/
	BakeErrorNoSceneRoot BakeError = 1
/*Lightmap baking failed as the lightmap data resource is embedded in a foreign resource.*/
	BakeErrorForeignData BakeError = 2
/*Lightmap baking failed as there is no lightmapper available in this Godot build.*/
	BakeErrorNoLightmapper BakeError = 3
/*Lightmap baking failed as the [LightmapGIData] save path isn't configured in the resource.*/
	BakeErrorNoSavePath BakeError = 4
/*Lightmap baking failed as there are no meshes whose [member GeometryInstance3D.gi_mode] is [constant GeometryInstance3D.GI_MODE_STATIC] and with valid UV2 mapping in the current scene. You may need to select 3D scenes in the Import dock and change their global illumination mode accordingly.*/
	BakeErrorNoMeshes BakeError = 5
/*Lightmap baking failed as the lightmapper failed to analyze some of the meshes marked as static for baking.*/
	BakeErrorMeshesInvalid BakeError = 6
/*Lightmap baking failed as the resulting image couldn't be saved or imported by Godot after it was saved.*/
	BakeErrorCantCreateImage BakeError = 7
/*The user aborted the lightmap baking operation (typically by clicking the [b]Cancel[/b] button in the progress dialog).*/
	BakeErrorUserAborted BakeError = 8
/*Lightmap baking failed as the maximum texture size is too small to fit some of the meshes marked for baking.*/
	BakeErrorTextureSizeTooSmall BakeError = 9
/*Lightmap baking failed as the lightmap is too small.*/
	BakeErrorLightmapTooSmall BakeError = 10
/*Lightmap baking failed as the lightmap was unable to fit into an atlas.*/
	BakeErrorAtlasTooSmall BakeError = 11
)
type EnvironmentMode = classdb.LightmapGIEnvironmentMode

const (
/*Ignore environment lighting when baking lightmaps.*/
	EnvironmentModeDisabled EnvironmentMode = 0
/*Use the scene's environment lighting when baking lightmaps.
[b]Note:[/b] If baking lightmaps in a scene with no [WorldEnvironment] node, this will act like [constant ENVIRONMENT_MODE_DISABLED]. The editor's preview sky and sun is [i]not[/i] taken into account by [LightmapGI] when baking lightmaps.*/
	EnvironmentModeScene EnvironmentMode = 1
/*Use [member environment_custom_sky] as a source of environment lighting when baking lightmaps.*/
	EnvironmentModeCustomSky EnvironmentMode = 2
/*Use [member environment_custom_color] multiplied by [member environment_custom_energy] as a constant source of environment lighting when baking lightmaps.*/
	EnvironmentModeCustomColor EnvironmentMode = 3
)