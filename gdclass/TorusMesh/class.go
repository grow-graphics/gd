package TorusMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PrimitiveMesh"
import "grow.graphics/gd/gdclass/Mesh"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Class representing a torus [PrimitiveMesh].

*/
type Go [1]classdb.TorusMesh
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TorusMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TorusMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) InnerRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetInnerRadius()))
}

func (self Go) SetInnerRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInnerRadius(gd.Float(value))
}

func (self Go) OuterRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetOuterRadius()))
}

func (self Go) SetOuterRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOuterRadius(gd.Float(value))
}

func (self Go) Rings() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetRings()))
}

func (self Go) SetRings(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRings(gd.Int(value))
}

func (self Go) RingSegments() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetRingSegments()))
}

func (self Go) SetRingSegments(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRingSegments(gd.Int(value))
}

//go:nosplit
func (self class) SetInnerRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TorusMesh.Bind_set_inner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInnerRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TorusMesh.Bind_get_inner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOuterRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TorusMesh.Bind_set_outer_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOuterRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TorusMesh.Bind_get_outer_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRings(rings gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rings)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TorusMesh.Bind_set_rings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRings() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TorusMesh.Bind_get_rings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRingSegments(rings gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rings)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TorusMesh.Bind_set_ring_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRingSegments() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TorusMesh.Bind_get_ring_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTorusMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTorusMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPrimitiveMesh() PrimitiveMesh.GD { return *((*PrimitiveMesh.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPrimitiveMesh() PrimitiveMesh.Go { return *((*PrimitiveMesh.Go)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.GD { return *((*Mesh.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMesh() Mesh.Go { return *((*Mesh.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}
func init() {classdb.Register("TorusMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}