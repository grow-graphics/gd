package AspectRatioContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A container type that arranges its child controls in a way that preserves their proportions automatically when the container is resized. Useful when a container has a dynamic size and the child nodes must adjust their sizes accordingly without losing their aspect ratios.

*/
type Go [1]classdb.AspectRatioContainer
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AspectRatioContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AspectRatioContainer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Ratio() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRatio()))
}

func (self Go) SetRatio(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRatio(gd.Float(value))
}

func (self Go) StretchMode() classdb.AspectRatioContainerStretchMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AspectRatioContainerStretchMode(class(self).GetStretchMode())
}

func (self Go) SetStretchMode(value classdb.AspectRatioContainerStretchMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStretchMode(value)
}

func (self Go) AlignmentHorizontal() classdb.AspectRatioContainerAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AspectRatioContainerAlignmentMode(class(self).GetAlignmentHorizontal())
}

func (self Go) SetAlignmentHorizontal(value classdb.AspectRatioContainerAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlignmentHorizontal(value)
}

func (self Go) AlignmentVertical() classdb.AspectRatioContainerAlignmentMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AspectRatioContainerAlignmentMode(class(self).GetAlignmentVertical())
}

func (self Go) SetAlignmentVertical(value classdb.AspectRatioContainerAlignmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlignmentVertical(value)
}

//go:nosplit
func (self class) SetRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_set_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_get_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStretchMode(stretch_mode classdb.AspectRatioContainerStretchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stretch_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStretchMode() classdb.AspectRatioContainerStretchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AspectRatioContainerStretchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlignmentHorizontal(alignment_horizontal classdb.AspectRatioContainerAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment_horizontal)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_set_alignment_horizontal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlignmentHorizontal() classdb.AspectRatioContainerAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AspectRatioContainerAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_get_alignment_horizontal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlignmentVertical(alignment_vertical classdb.AspectRatioContainerAlignmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment_vertical)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_set_alignment_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlignmentVertical() classdb.AspectRatioContainerAlignmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AspectRatioContainerAlignmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AspectRatioContainer.Bind_get_alignment_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAspectRatioContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAspectRatioContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {classdb.Register("AspectRatioContainer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type StretchMode = classdb.AspectRatioContainerStretchMode

const (
/*The height of child controls is automatically adjusted based on the width of the container.*/
	StretchWidthControlsHeight StretchMode = 0
/*The width of child controls is automatically adjusted based on the height of the container.*/
	StretchHeightControlsWidth StretchMode = 1
/*The bounding rectangle of child controls is automatically adjusted to fit inside the container while keeping the aspect ratio.*/
	StretchFit StretchMode = 2
/*The width and height of child controls is automatically adjusted to make their bounding rectangle cover the entire area of the container while keeping the aspect ratio.
When the bounding rectangle of child controls exceed the container's size and [member Control.clip_contents] is enabled, this allows to show only the container's area restricted by its own bounding rectangle.*/
	StretchCover StretchMode = 3
)
type AlignmentMode = classdb.AspectRatioContainerAlignmentMode

const (
/*Aligns child controls with the beginning (left or top) of the container.*/
	AlignmentBegin AlignmentMode = 0
/*Aligns child controls with the center of the container.*/
	AlignmentCenter AlignmentMode = 1
/*Aligns child controls with the end (right or bottom) of the container.*/
	AlignmentEnd AlignmentMode = 2
)