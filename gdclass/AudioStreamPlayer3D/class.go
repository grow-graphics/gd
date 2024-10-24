package AudioStreamPlayer3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Plays audio with positional sound effects, based on the relative position of the audio listener. Positional effects include distance attenuation, directionality, and the Doppler effect. For greater realism, a low-pass filter is applied to distant sounds. This can be disabled by setting [member attenuation_filter_cutoff_hz] to [code]20500[/code].
By default, audio is heard from the camera position. This can be changed by adding an [AudioListener3D] node to the scene and enabling it by calling [method AudioListener3D.make_current] on it.
See also [AudioStreamPlayer] to play a sound non-positionally.
[b]Note:[/b] Hiding an [AudioStreamPlayer3D] node does not disable its audio output. To temporarily disable an [AudioStreamPlayer3D]'s audio output, set [member volume_db] to a very low value like [code]-100[/code] (which isn't audible to human hearing).

*/
type Go [1]classdb.AudioStreamPlayer3D

/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
func (self Go) Play() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Play(gd.Float(0.0))
}

/*
Sets the position from which audio will be played, in seconds.
*/
func (self Go) SeekTo(to_position float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SeekTo(gd.Float(to_position))
}

/*
Stops the audio.
*/
func (self Go) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Stop()
}

/*
Returns the position in the [AudioStream].
*/
func (self Go) GetPlaybackPosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetPlaybackPosition()))
}

/*
Returns whether the [AudioStreamPlayer] can return the [AudioStreamPlayback] object or not.
*/
func (self Go) HasStreamPlayback() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasStreamPlayback())
}

/*
Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer3D].
*/
func (self Go) GetStreamPlayback() gdclass.AudioStreamPlayback {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioStreamPlayback(class(self).GetStreamPlayback(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamPlayer3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamPlayer3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Stream() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetStream(gc))
}

func (self Go) SetStream(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStream(value)
}

func (self Go) AttenuationModel() classdb.AudioStreamPlayer3DAttenuationModel {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioStreamPlayer3DAttenuationModel(class(self).GetAttenuationModel())
}

func (self Go) SetAttenuationModel(value classdb.AudioStreamPlayer3DAttenuationModel) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAttenuationModel(value)
}

func (self Go) VolumeDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumeDb()))
}

func (self Go) SetVolumeDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumeDb(gd.Float(value))
}

func (self Go) UnitSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetUnitSize()))
}

func (self Go) SetUnitSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUnitSize(gd.Float(value))
}

func (self Go) MaxDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMaxDb()))
}

func (self Go) SetMaxDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxDb(gd.Float(value))
}

func (self Go) PitchScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPitchScale()))
}

func (self Go) SetPitchScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPitchScale(gd.Float(value))
}

func (self Go) Playing() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPlaying())
}

func (self Go) Autoplay() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAutoplayEnabled())
}

func (self Go) SetAutoplay(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoplay(value)
}

func (self Go) StreamPaused() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetStreamPaused())
}

func (self Go) SetStreamPaused(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStreamPaused(value)
}

func (self Go) MaxDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMaxDistance()))
}

func (self Go) SetMaxDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxDistance(gd.Float(value))
}

func (self Go) MaxPolyphony() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxPolyphony()))
}

func (self Go) SetMaxPolyphony(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxPolyphony(gd.Int(value))
}

func (self Go) PanningStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPanningStrength()))
}

func (self Go) SetPanningStrength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPanningStrength(gd.Float(value))
}

func (self Go) Bus() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetBus(gc).String())
}

func (self Go) SetBus(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBus(gc.StringName(value))
}

func (self Go) AreaMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetAreaMask()))
}

func (self Go) SetAreaMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAreaMask(gd.Int(value))
}

func (self Go) PlaybackType() classdb.AudioServerPlaybackType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioServerPlaybackType(class(self).GetPlaybackType())
}

func (self Go) SetPlaybackType(value classdb.AudioServerPlaybackType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPlaybackType(value)
}

func (self Go) EmissionAngleEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEmissionAngleEnabled())
}

func (self Go) SetEmissionAngleEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionAngleEnabled(value)
}

func (self Go) EmissionAngleDegrees() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEmissionAngle()))
}

func (self Go) SetEmissionAngleDegrees(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionAngle(gd.Float(value))
}

func (self Go) EmissionAngleFilterAttenuationDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEmissionAngleFilterAttenuationDb()))
}

func (self Go) SetEmissionAngleFilterAttenuationDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmissionAngleFilterAttenuationDb(gd.Float(value))
}

func (self Go) AttenuationFilterCutoffHz() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAttenuationFilterCutoffHz()))
}

func (self Go) SetAttenuationFilterCutoffHz(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAttenuationFilterCutoffHz(gd.Float(value))
}

func (self Go) AttenuationFilterDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAttenuationFilterDb()))
}

func (self Go) SetAttenuationFilterDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAttenuationFilterDb(gd.Float(value))
}

func (self Go) DopplerTracking() classdb.AudioStreamPlayer3DDopplerTracking {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioStreamPlayer3DDopplerTracking(class(self).GetDopplerTracking())
}

func (self Go) SetDopplerTracking(value classdb.AudioStreamPlayer3DDopplerTracking) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDopplerTracking(value)
}

//go:nosplit
func (self class) SetStream(stream gdclass.AudioStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStream(ctx gd.Lifetime) gdclass.AudioStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStream
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumeDb(volume_db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, volume_db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumeDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUnitSize(unit_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, unit_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_unit_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUnitSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_unit_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxDb(max_db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_max_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_max_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPitchScale(pitch_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pitch_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPitchScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
//go:nosplit
func (self class) Play(from_position gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_play, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the position from which audio will be played, in seconds.
*/
//go:nosplit
func (self class) SeekTo(to_position gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the audio.
*/
//go:nosplit
func (self class) Stop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPlaying() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position in the [AudioStream].
*/
//go:nosplit
func (self class) GetPlaybackPosition() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_playback_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBus(bus gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bus))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBus(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoplay(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoplayEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_is_autoplay_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxDistance(meters gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, meters)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAreaMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_area_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAreaMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_area_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionAngle(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_emission_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_emission_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionAngleEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_emission_angle_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEmissionAngleEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_is_emission_angle_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmissionAngleFilterAttenuationDb(db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_emission_angle_filter_attenuation_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmissionAngleFilterAttenuationDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_emission_angle_filter_attenuation_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAttenuationFilterCutoffHz(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_attenuation_filter_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAttenuationFilterCutoffHz() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_attenuation_filter_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAttenuationFilterDb(db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_attenuation_filter_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAttenuationFilterDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_attenuation_filter_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAttenuationModel(model classdb.AudioStreamPlayer3DAttenuationModel)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, model)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_attenuation_model, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAttenuationModel() classdb.AudioStreamPlayer3DAttenuationModel {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioStreamPlayer3DAttenuationModel](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_attenuation_model, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDopplerTracking(mode classdb.AudioStreamPlayer3DDopplerTracking)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDopplerTracking() classdb.AudioStreamPlayer3DDopplerTracking {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioStreamPlayer3DDopplerTracking](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStreamPaused(pause bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pause)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_stream_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamPaused() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_stream_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxPolyphony(max_polyphony gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxPolyphony() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPanningStrength(panning_strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, panning_strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_panning_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPanningStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_panning_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether the [AudioStreamPlayer] can return the [AudioStreamPlayback] object or not.
*/
//go:nosplit
func (self class) HasStreamPlayback() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_has_stream_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer3D].
*/
//go:nosplit
func (self class) GetStreamPlayback(ctx gd.Lifetime) gdclass.AudioStreamPlayback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_stream_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStreamPlayback
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlaybackType(playback_type classdb.AudioServerPlaybackType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, playback_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_set_playback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlaybackType() classdb.AudioServerPlaybackType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioServerPlaybackType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer3D.Bind_get_playback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnFinished(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("finished"), gc.Callable(cb), 0)
}


func (self class) AsAudioStreamPlayer3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlayer3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("AudioStreamPlayer3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type AttenuationModel = classdb.AudioStreamPlayer3DAttenuationModel

const (
/*Attenuation of loudness according to linear distance.*/
	AttenuationInverseDistance AttenuationModel = 0
/*Attenuation of loudness according to squared distance.*/
	AttenuationInverseSquareDistance AttenuationModel = 1
/*Attenuation of loudness according to logarithmic distance.*/
	AttenuationLogarithmic AttenuationModel = 2
/*No attenuation of loudness according to distance. The sound will still be heard positionally, unlike an [AudioStreamPlayer]. [constant ATTENUATION_DISABLED] can be combined with a [member max_distance] value greater than [code]0.0[/code] to achieve linear attenuation clamped to a sphere of a defined size.*/
	AttenuationDisabled AttenuationModel = 3
)
type DopplerTracking = classdb.AudioStreamPlayer3DDopplerTracking

const (
/*Disables doppler tracking.*/
	DopplerTrackingDisabled DopplerTracking = 0
/*Executes doppler tracking during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	DopplerTrackingIdleStep DopplerTracking = 1
/*Executes doppler tracking during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	DopplerTrackingPhysicsStep DopplerTracking = 2
)