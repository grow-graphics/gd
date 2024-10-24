package MethodTweener

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Tweener"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[MethodTweener] is similar to a combination of [CallbackTweener] and [PropertyTweener]. It calls a method providing an interpolated value as a parameter. See [method Tween.tween_method] for more usage information.
The tweener will finish automatically if the callback's target object is freed.
[b]Note:[/b] [method Tween.tween_method] is the only correct way to create [MethodTweener]. Any [MethodTweener] created manually will not function correctly.

*/
type Go [1]classdb.MethodTweener

/*
Sets the time in seconds after which the [MethodTweener] will start interpolating. By default there's no delay.
*/
func (self Go) SetDelay(delay float64) gdclass.MethodTweener {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.MethodTweener(class(self).SetDelay(gc, gd.Float(delay)))
}

/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
func (self Go) SetTrans(trans classdb.TweenTransitionType) gdclass.MethodTweener {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.MethodTweener(class(self).SetTrans(gc, trans))
}

/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
func (self Go) SetEase(ease classdb.TweenEaseType) gdclass.MethodTweener {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.MethodTweener(class(self).SetEase(gc, ease))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MethodTweener
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MethodTweener"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Sets the time in seconds after which the [MethodTweener] will start interpolating. By default there's no delay.
*/
//go:nosplit
func (self class) SetDelay(ctx gd.Lifetime, delay gd.Float) gdclass.MethodTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MethodTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.MethodTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetTrans(ctx gd.Lifetime, trans classdb.TweenTransitionType) gdclass.MethodTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, trans)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MethodTweener.Bind_set_trans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.MethodTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetEase(ctx gd.Lifetime, ease classdb.TweenEaseType) gdclass.MethodTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ease)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MethodTweener.Bind_set_ease, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.MethodTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsMethodTweener() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMethodTweener() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTweener() Tweener.GD { return *((*Tweener.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTweener() Tweener.Go { return *((*Tweener.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTweener(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTweener(), name)
	}
}
func init() {classdb.Register("MethodTweener", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}