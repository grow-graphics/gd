package AudioEffectPitchShift

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioEffect"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Allows modulation of pitch independently of tempo. All frequencies can be increased/decreased with minimal effect on transients.

*/
type Go [1]classdb.AudioEffectPitchShift
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectPitchShift
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectPitchShift"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) PitchScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPitchScale()))
}

func (self Go) SetPitchScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPitchScale(gd.Float(value))
}

func (self Go) Oversampling() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetOversampling()))
}

func (self Go) SetOversampling(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOversampling(gd.Int(value))
}

func (self Go) FftSize() classdb.AudioEffectPitchShiftFFTSize {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioEffectPitchShiftFFTSize(class(self).GetFftSize())
}

func (self Go) SetFftSize(value classdb.AudioEffectPitchShiftFFTSize) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFftSize(value)
}

//go:nosplit
func (self class) SetPitchScale(rate gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPitchShift.Bind_set_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPitchScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPitchShift.Bind_get_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOversampling(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPitchShift.Bind_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOversampling() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPitchShift.Bind_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFftSize(size classdb.AudioEffectPitchShiftFFTSize)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPitchShift.Bind_set_fft_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFftSize() classdb.AudioEffectPitchShiftFFTSize {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioEffectPitchShiftFFTSize](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPitchShift.Bind_get_fft_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectPitchShift() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectPitchShift() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffect() AudioEffect.GD { return *((*AudioEffect.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffect() AudioEffect.Go { return *((*AudioEffect.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}
func init() {classdb.Register("AudioEffectPitchShift", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type FFTSize = classdb.AudioEffectPitchShiftFFTSize

const (
/*Use a buffer of 256 samples for the Fast Fourier transform. Lowest latency, but least stable over time.*/
	FftSize256 FFTSize = 0
/*Use a buffer of 512 samples for the Fast Fourier transform. Low latency, but less stable over time.*/
	FftSize512 FFTSize = 1
/*Use a buffer of 1024 samples for the Fast Fourier transform. This is a compromise between latency and stability over time.*/
	FftSize1024 FFTSize = 2
/*Use a buffer of 2048 samples for the Fast Fourier transform. High latency, but stable over time.*/
	FftSize2048 FFTSize = 3
/*Use a buffer of 4096 samples for the Fast Fourier transform. Highest latency, but most stable over time.*/
	FftSize4096 FFTSize = 4
/*Represents the size of the [enum FFTSize] enum.*/
	FftSizeMax FFTSize = 5
)