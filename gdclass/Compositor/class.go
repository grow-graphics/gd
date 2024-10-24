package Compositor

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
The compositor resource stores attributes used to customize how a [Viewport] is rendered.

*/
type Go [1]classdb.Compositor
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Compositor
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Compositor"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) CompositorEffects() gd.ArrayOf[gdclass.CompositorEffect] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gdclass.CompositorEffect](class(self).GetCompositorEffects(gc))
}

func (self Go) SetCompositorEffects(value gd.ArrayOf[gdclass.CompositorEffect]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCompositorEffects(value)
}

//go:nosplit
func (self class) SetCompositorEffects(compositor_effects gd.ArrayOf[gdclass.CompositorEffect])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(compositor_effects))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Compositor.Bind_set_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCompositorEffects(ctx gd.Lifetime) gd.ArrayOf[gdclass.CompositorEffect] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Compositor.Bind_get_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.CompositorEffect](ret)
}
func (self class) AsCompositor() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCompositor() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("Compositor", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}