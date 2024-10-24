package InputEventKey

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEventWithModifiers"
import "grow.graphics/gd/gdclass/InputEventFromWindow"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An input event for keys on a keyboard. Supports key presses, key releases and [member echo] events. It can also be received in [method Node._unhandled_key_input].
[b]Note:[/b] Events received from the keyboard usually have all properties set. Event mappings should have only one of the [member keycode], [member physical_keycode] or [member unicode] set.
When events are compared, properties are checked in the following priority - [member keycode], [member physical_keycode] and [member unicode]. Events with the first matching value will be considered equal.

*/
type Go [1]classdb.InputEventKey

/*
Returns the Latin keycode combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_keycode_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
func (self Go) GetKeycodeWithModifiers() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(class(self).GetKeycodeWithModifiers())
}

/*
Returns the physical keycode combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_physical_keycode_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
func (self Go) GetPhysicalKeycodeWithModifiers() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(class(self).GetPhysicalKeycodeWithModifiers())
}

/*
Returns the localized key label combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_key_label_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
func (self Go) GetKeyLabelWithModifiers() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(class(self).GetKeyLabelWithModifiers())
}

/*
Returns a [String] representation of the event's [member keycode] and modifiers.
*/
func (self Go) AsTextKeycode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).AsTextKeycode(gc).String())
}

/*
Returns a [String] representation of the event's [member physical_keycode] and modifiers.
*/
func (self Go) AsTextPhysicalKeycode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).AsTextPhysicalKeycode(gc).String())
}

/*
Returns a [String] representation of the event's [member key_label] and modifiers.
*/
func (self Go) AsTextKeyLabel() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).AsTextKeyLabel(gc).String())
}

/*
Returns a [String] representation of the event's [member location]. This will be a blank string if the event is not specific to a location.
*/
func (self Go) AsTextLocation() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).AsTextLocation(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEventKey
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("InputEventKey"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Keycode() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Key(class(self).GetKeycode())
}

func (self Go) SetKeycode(value gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetKeycode(value)
}

func (self Go) PhysicalKeycode() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Key(class(self).GetPhysicalKeycode())
}

func (self Go) SetPhysicalKeycode(value gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPhysicalKeycode(value)
}

func (self Go) KeyLabel() gd.Key {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Key(class(self).GetKeyLabel())
}

func (self Go) SetKeyLabel(value gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetKeyLabel(value)
}

func (self Go) Unicode() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetUnicode()))
}

func (self Go) SetUnicode(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUnicode(gd.Int(value))
}

func (self Go) Location() gd.KeyLocation {
	gc := gd.GarbageCollector(); _ = gc
		return gd.KeyLocation(class(self).GetLocation())
}

func (self Go) SetLocation(value gd.KeyLocation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLocation(value)
}

//go:nosplit
func (self class) SetPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetKeycode(keycode gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeycode() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicalKeycode(physical_keycode gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, physical_keycode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_physical_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicalKeycode() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_physical_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeyLabel(key_label gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, key_label)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_key_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeyLabel() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_key_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUnicode(unicode gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, unicode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_unicode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUnicode() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_unicode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLocation(location gd.KeyLocation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, location)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_location, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLocation() gd.KeyLocation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.KeyLocation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_location, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEcho(echo bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, echo)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_set_echo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the Latin keycode combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_keycode_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
//go:nosplit
func (self class) GetKeycodeWithModifiers() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_keycode_with_modifiers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the physical keycode combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_physical_keycode_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
//go:nosplit
func (self class) GetPhysicalKeycodeWithModifiers() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_physical_keycode_with_modifiers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the localized key label combined with modifier keys such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See also [InputEventWithModifiers].
To get a human-readable representation of the [InputEventKey] with modifiers, use [code]OS.get_keycode_string(event.get_key_label_with_modifiers())[/code] where [code]event[/code] is the [InputEventKey].
*/
//go:nosplit
func (self class) GetKeyLabelWithModifiers() gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_get_key_label_with_modifiers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event's [member keycode] and modifiers.
*/
//go:nosplit
func (self class) AsTextKeycode(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_as_text_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event's [member physical_keycode] and modifiers.
*/
//go:nosplit
func (self class) AsTextPhysicalKeycode(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_as_text_physical_keycode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event's [member key_label] and modifiers.
*/
//go:nosplit
func (self class) AsTextKeyLabel(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_as_text_key_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [String] representation of the event's [member location]. This will be a blank string if the event is not specific to a location.
*/
//go:nosplit
func (self class) AsTextLocation(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventKey.Bind_as_text_location, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsInputEventKey() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventKey() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEventWithModifiers() InputEventWithModifiers.GD { return *((*InputEventWithModifiers.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventWithModifiers() InputEventWithModifiers.Go { return *((*InputEventWithModifiers.Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEventFromWindow() InputEventFromWindow.GD { return *((*InputEventFromWindow.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventFromWindow() InputEventFromWindow.Go { return *((*InputEventFromWindow.Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEvent() InputEvent.GD { return *((*InputEvent.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEvent() InputEvent.Go { return *((*InputEvent.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEventWithModifiers(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEventWithModifiers(), name)
	}
}
func init() {classdb.Register("InputEventKey", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}