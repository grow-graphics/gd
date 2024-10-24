package InputEventScreenTouch

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEventFromWindow"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Stores information about multi-touch press/release input events. Supports touch press, touch release and [member index] for multi-touch count and order.

*/
type Go [1]classdb.InputEventScreenTouch
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEventScreenTouch
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("InputEventScreenTouch"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Index() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetIndex()))
}

func (self Go) SetIndex(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIndex(gd.Int(value))
}

func (self Go) Position() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetPosition())
}

func (self Go) SetPosition(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPosition(value)
}

func (self Go) DoubleTap() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDoubleTap())
}

func (self Go) SetDoubleTap(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDoubleTap(value)
}

//go:nosplit
func (self class) SetIndex(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventScreenTouch.Bind_set_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventScreenTouch.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventScreenTouch.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventScreenTouch.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventScreenTouch.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCanceled(canceled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canceled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventScreenTouch.Bind_set_canceled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDoubleTap(double_tap bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, double_tap)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventScreenTouch.Bind_set_double_tap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDoubleTap() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventScreenTouch.Bind_is_double_tap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventScreenTouch() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventScreenTouch() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsInputEventFromWindow(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEventFromWindow(), name)
	}
}
func init() {classdb.Register("InputEventScreenTouch", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}