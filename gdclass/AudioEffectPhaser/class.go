package AudioEffectPhaser

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
Combines phase-shifted signals with the original signal. The movement of the phase-shifted signals is controlled using a low-frequency oscillator.

*/
type Go [1]classdb.AudioEffectPhaser
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectPhaser
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectPhaser"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) RangeMinHz() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRangeMinHz()))
}

func (self Go) SetRangeMinHz(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRangeMinHz(gd.Float(value))
}

func (self Go) RangeMaxHz() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRangeMaxHz()))
}

func (self Go) SetRangeMaxHz(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRangeMaxHz(gd.Float(value))
}

func (self Go) RateHz() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRateHz()))
}

func (self Go) SetRateHz(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRateHz(gd.Float(value))
}

func (self Go) Feedback() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFeedback()))
}

func (self Go) SetFeedback(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFeedback(gd.Float(value))
}

func (self Go) Depth() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDepth()))
}

func (self Go) SetDepth(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepth(gd.Float(value))
}

//go:nosplit
func (self class) SetRangeMinHz(hz gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_set_range_min_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRangeMinHz() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_get_range_min_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRangeMaxHz(hz gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_set_range_max_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRangeMaxHz() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_get_range_max_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRateHz(hz gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_set_rate_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRateHz() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_get_rate_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFeedback(fbk gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fbk)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_set_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFeedback() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_get_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepth(depth gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectPhaser.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectPhaser() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectPhaser() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioEffectPhaser", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}