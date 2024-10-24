package InputEventMouseButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEventMouse"
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
Stores information about mouse click events. See [method Node._input].
[b]Note:[/b] On Wear OS devices, rotary input is mapped to [constant MOUSE_BUTTON_WHEEL_UP] and [constant MOUSE_BUTTON_WHEEL_DOWN]. This can be changed to [constant MOUSE_BUTTON_WHEEL_LEFT] and [constant MOUSE_BUTTON_WHEEL_RIGHT] with the [member ProjectSettings.input_devices/pointing/android/rotary_input_scroll_axis] setting.

*/
type Go [1]classdb.InputEventMouseButton
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEventMouseButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("InputEventMouseButton"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Factor() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFactor()))
}

func (self Go) SetFactor(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFactor(gd.Float(value))
}

func (self Go) ButtonIndex() gd.MouseButton {
	gc := gd.GarbageCollector(); _ = gc
		return gd.MouseButton(class(self).GetButtonIndex())
}

func (self Go) SetButtonIndex(value gd.MouseButton) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetButtonIndex(value)
}

func (self Go) DoubleClick() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDoubleClick())
}

func (self Go) SetDoubleClick(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDoubleClick(value)
}

//go:nosplit
func (self class) SetFactor(factor gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, factor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFactor() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_get_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetButtonIndex(button_index gd.MouseButton)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, button_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_button_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetButtonIndex() gd.MouseButton {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.MouseButton](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_get_button_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCanceled(canceled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canceled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_canceled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDoubleClick(double_click bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, double_click)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_set_double_click, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDoubleClick() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouseButton.Bind_is_double_click, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventMouseButton() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventMouseButton() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEventMouse() InputEventMouse.GD { return *((*InputEventMouse.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventMouse() InputEventMouse.Go { return *((*InputEventMouse.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsInputEventMouse(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEventMouse(), name)
	}
}
func init() {classdb.Register("InputEventMouseButton", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}