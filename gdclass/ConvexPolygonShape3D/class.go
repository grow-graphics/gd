package ConvexPolygonShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Shape3D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 3D convex polyhedron shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D].
[ConvexPolygonShape3D] is [i]solid[/i], which means it detects collisions from objects that are fully inside it, unlike [ConcavePolygonShape3D] which is hollow. This makes it more suitable for both detection and physics.
[b]Convex decomposition:[/b] A concave polyhedron can be split up into several convex polyhedra. This allows dynamic physics bodies to have complex concave collisions (at a performance cost) and can be achieved by using several [ConvexPolygonShape3D] nodes. To generate a convex decomposition from a mesh, select the [MeshInstance3D] node, go to the [b]Mesh[/b] menu that appears above the viewport, and choose [b]Create Multiple Convex Collision Siblings[/b]. Alternatively, [method MeshInstance3D.create_multiple_convex_collisions] can be called in a script to perform this decomposition at run-time.
[b]Performance:[/b] [ConvexPolygonShape3D] is faster to check collisions against compared to [ConcavePolygonShape3D], but it is slower than primitive collision shapes such as [SphereShape3D] and [BoxShape3D]. Its use should generally be limited to medium-sized objects that cannot have their collision accurately represented by primitive shapes.

*/
type Go [1]classdb.ConvexPolygonShape3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ConvexPolygonShape3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ConvexPolygonShape3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Points() []gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Vector3(class(self).GetPoints(gc).AsSlice())
}

func (self Go) SetPoints(value []gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPoints(gc.PackedVector3Slice(value))
}

//go:nosplit
func (self class) SetPoints(points gd.PackedVector3Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConvexPolygonShape3D.Bind_set_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPoints(ctx gd.Lifetime) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConvexPolygonShape3D.Bind_get_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsConvexPolygonShape3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsConvexPolygonShape3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.GD { return *((*Shape3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsShape3D() Shape3D.Go { return *((*Shape3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape3D(), name)
	}
}
func init() {classdb.Register("ConvexPolygonShape3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}