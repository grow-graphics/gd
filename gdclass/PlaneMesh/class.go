package PlaneMesh

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
Class representing a planar [PrimitiveMesh]. This flat mesh does not have a thickness. By default, this mesh is aligned on the X and Z axes; this default rotation isn't suited for use with billboarded materials. For billboarded materials, change [member orientation] to [constant FACE_Z].
[b]Note:[/b] When using a large textured [PlaneMesh] (e.g. as a floor), you may stumble upon UV jittering issues depending on the camera angle. To solve this, increase [member subdivide_depth] and [member subdivide_width] until you no longer notice UV jittering.

*/
type Go [1]classdb.PlaneMesh
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PlaneMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PlaneMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Size() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetSize())
}

func (self Go) SetSize(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSize(value)
}

func (self Go) SubdivideWidth() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSubdivideWidth()))
}

func (self Go) SetSubdivideWidth(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSubdivideWidth(gd.Int(value))
}

func (self Go) SubdivideDepth() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSubdivideDepth()))
}

func (self Go) SetSubdivideDepth(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSubdivideDepth(gd.Int(value))
}

func (self Go) CenterOffset() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetCenterOffset())
}

func (self Go) SetCenterOffset(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCenterOffset(value)
}

func (self Go) Orientation() classdb.PlaneMeshOrientation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.PlaneMeshOrientation(class(self).GetOrientation())
}

func (self Go) SetOrientation(value classdb.PlaneMeshOrientation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOrientation(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubdivideWidth(subdivide gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, subdivide)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_set_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubdivideWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_get_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubdivideDepth(subdivide gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, subdivide)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_set_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubdivideDepth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_get_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCenterOffset(offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_set_center_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterOffset() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_get_center_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOrientation(orientation classdb.PlaneMeshOrientation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_set_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOrientation() classdb.PlaneMeshOrientation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PlaneMeshOrientation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaneMesh.Bind_get_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPlaneMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPlaneMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PlaneMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Orientation = classdb.PlaneMeshOrientation

const (
/*[PlaneMesh] will face the positive X-axis.*/
	FaceX Orientation = 0
/*[PlaneMesh] will face the positive Y-axis. This matches the behavior of the [PlaneMesh] in Godot 3.x.*/
	FaceY Orientation = 1
/*[PlaneMesh] will face the positive Z-axis. This matches the behavior of the QuadMesh in Godot 3.x.*/
	FaceZ Orientation = 2
)