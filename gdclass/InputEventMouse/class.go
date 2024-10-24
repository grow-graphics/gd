package InputEventMouse

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
Stores general information about mouse events.

*/
type Go [1]classdb.InputEventMouse
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEventMouse
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("InputEventMouse"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ButtonMask() gd.MouseButtonMask {
	gc := gd.GarbageCollector(); _ = gc
		return gd.MouseButtonMask(class(self).GetButtonMask())
}

func (self Go) SetButtonMask(value gd.MouseButtonMask) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetButtonMask(value)
}

func (self Go) Position() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetPosition())
}

func (self Go) SetPosition(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPosition(value)
}

func (self Go) GlobalPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetGlobalPosition())
}

func (self Go) SetGlobalPosition(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlobalPosition(value)
}

//go:nosplit
func (self class) SetButtonMask(button_mask gd.MouseButtonMask)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, button_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouse.Bind_set_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetButtonMask() gd.MouseButtonMask {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.MouseButtonMask](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouse.Bind_get_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPosition(position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouse.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouse.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalPosition(global_position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, global_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouse.Bind_set_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMouse.Bind_get_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventMouse() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventMouse() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("InputEventMouse", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}