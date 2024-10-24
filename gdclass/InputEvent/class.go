package InputEvent

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
Abstract base class of all types of input events. See [method Node._input].

*/
type Go [1]classdb.InputEvent

/*
Returns [code]true[/code] if this input event matches a pre-defined action of any type.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func (self Go) IsAction(action string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsAction(gc.StringName(action), false))
}

/*
Returns [code]true[/code] if the given action is being pressed (and is not an echo event for [InputEventKey] events, unless [param allow_echo] is [code]true[/code]). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] Due to keyboard ghosting, [method is_action_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
func (self Go) IsActionPressed(action string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsActionPressed(gc.StringName(action), false, false))
}

/*
Returns [code]true[/code] if the given action is released (i.e. not pressed). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func (self Go) IsActionReleased(action string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsActionReleased(gc.StringName(action), false))
}

/*
Returns a value between 0.0 and 1.0 depending on the given actions' state. Useful for getting the value of events of type [InputEventJoypadMotion].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func (self Go) GetActionStrength(action string) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetActionStrength(gc.StringName(action), false)))
}

/*
Returns [code]true[/code] if this input event has been canceled.
*/
func (self Go) IsCanceled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsCanceled())
}

/*
Returns [code]true[/code] if this input event is pressed. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
[b]Note:[/b] Due to keyboard ghosting, [method is_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
func (self Go) IsPressed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPressed())
}

/*
Returns [code]true[/code] if this input event is released. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
*/
func (self Go) IsReleased() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsReleased())
}

/*
Returns [code]true[/code] if this input event is an echo event (only for events of type [InputEventKey]). An echo event is a repeated key event sent when the user is holding down the key. Any other event type returns [code]false[/code].
[b]Note:[/b] The rate at which echo events are sent is typically around 20 events per second (after holding down the key for roughly half a second). However, the key repeat delay/speed can be changed by the user or disabled entirely in the operating system settings. To ensure your project works correctly on all configurations, do not assume the user has a specific key repeat configuration in your project's behavior.
*/
func (self Go) IsEcho() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsEcho())
}

/*
Returns a [String] representation of the event.
*/
func (self Go) AsText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).AsText(gc).String())
}

/*
Returns [code]true[/code] if the specified [param event] matches this event. Only valid for action events i.e key ([InputEventKey]), button ([InputEventMouseButton] or [InputEventJoypadButton]), axis [InputEventJoypadMotion] or action ([InputEventAction]) events.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
func (self Go) IsMatch(event gdclass.InputEvent) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsMatch(event, true))
}

/*
Returns [code]true[/code] if this input event's type is one that can be assigned to an input action.
*/
func (self Go) IsActionType() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsActionType())
}

/*
Returns [code]true[/code] if the given input event and this input event can be added together (only for events of type [InputEventMouseMotion]).
The given input event's position, global position and speed will be copied. The resulting [code]relative[/code] is a sum of both events. Both events' modifiers have to be identical.
*/
func (self Go) Accumulate(with_event gdclass.InputEvent) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).Accumulate(with_event))
}

/*
Returns a copy of the given input event which has been offset by [param local_ofs] and transformed by [param xform]. Relevant for events of type [InputEventMouseButton], [InputEventMouseMotion], [InputEventScreenTouch], [InputEventScreenDrag], [InputEventMagnifyGesture] and [InputEventPanGesture].
*/
func (self Go) XformedBy(xform gd.Transform2D) gdclass.InputEvent {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.InputEvent(class(self).XformedBy(gc, xform, gd.Vector2{0, 0}))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEvent
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("InputEvent"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Device() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetDevice()))
}

func (self Go) SetDevice(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDevice(gd.Int(value))
}

//go:nosplit
func (self class) SetDevice(device gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, device)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_set_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDevice() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_get_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this input event matches a pre-defined action of any type.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) IsAction(action gd.StringName, exact_match bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given action is being pressed (and is not an echo event for [InputEventKey] events, unless [param allow_echo] is [code]true[/code]). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
[b]Note:[/b] Due to keyboard ghosting, [method is_action_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
//go:nosplit
func (self class) IsActionPressed(action gd.StringName, allow_echo bool, exact_match bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, allow_echo)
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_action_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given action is released (i.e. not pressed). Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) IsActionReleased(action gd.StringName, exact_match bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_action_released, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a value between 0.0 and 1.0 depending on the given actions' state. Useful for getting the value of events of type [InputEventJoypadMotion].
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) GetActionStrength(action gd.StringName, exact_match bool) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_get_action_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this input event has been canceled.
*/
//go:nosplit
func (self class) IsCanceled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_canceled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this input event is pressed. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
[b]Note:[/b] Due to keyboard ghosting, [method is_pressed] may return [code]false[/code] even if one of the action's keys is pressed. See [url=$DOCS_URL/tutorials/inputs/input_examples.html#keyboard-events]Input examples[/url] in the documentation for more information.
*/
//go:nosplit
func (self class) IsPressed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this input event is released. Not relevant for events of type [InputEventMouseMotion] or [InputEventScreenDrag].
*/
//go:nosplit
func (self class) IsReleased() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_released, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this input event is an echo event (only for events of type [InputEventKey]). An echo event is a repeated key event sent when the user is holding down the key. Any other event type returns [code]false[/code].
[b]Note:[/b] The rate at which echo events are sent is typically around 20 events per second (after holding down the key for roughly half a second). However, the key repeat delay/speed can be changed by the user or disabled entirely in the operating system settings. To ensure your project works correctly on all configurations, do not assume the user has a specific key repeat configuration in your project's behavior.
*/
//go:nosplit
func (self class) IsEcho() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_echo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event.
*/
//go:nosplit
func (self class) AsText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_as_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified [param event] matches this event. Only valid for action events i.e key ([InputEventKey]), button ([InputEventMouseButton] or [InputEventJoypadButton]), axis [InputEventJoypadMotion] or action ([InputEventAction]) events.
If [param exact_match] is [code]false[/code], it ignores additional input modifiers for [InputEventKey] and [InputEventMouseButton] events, and the direction for [InputEventJoypadMotion] events.
*/
//go:nosplit
func (self class) IsMatch(event gdclass.InputEvent, exact_match bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	callframe.Arg(frame, exact_match)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_match, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this input event's type is one that can be assigned to an input action.
*/
//go:nosplit
func (self class) IsActionType() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_is_action_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given input event and this input event can be added together (only for events of type [InputEventMouseMotion]).
The given input event's position, global position and speed will be copied. The resulting [code]relative[/code] is a sum of both events. Both events' modifiers have to be identical.
*/
//go:nosplit
func (self class) Accumulate(with_event gdclass.InputEvent) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(with_event[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_accumulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a copy of the given input event which has been offset by [param local_ofs] and transformed by [param xform]. Relevant for events of type [InputEventMouseButton], [InputEventMouseMotion], [InputEventScreenTouch], [InputEventScreenDrag], [InputEventMagnifyGesture] and [InputEventPanGesture].
*/
//go:nosplit
func (self class) XformedBy(ctx gd.Lifetime, xform gd.Transform2D, local_ofs gd.Vector2) gdclass.InputEvent {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	callframe.Arg(frame, local_ofs)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEvent.Bind_xformed_by, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.InputEvent
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsInputEvent() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEvent() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("InputEvent", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}