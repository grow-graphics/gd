package CSGMesh3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/CSGPrimitive3D"
import "grow.graphics/gd/gdclass/CSGShape3D"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This CSG node allows you to use any mesh resource as a CSG shape, provided it is closed, does not self-intersect, does not contain internal faces and has no edges that connect to more than two faces. See also [CSGPolygon3D] for drawing 2D extruded polygons to be used as CSG nodes.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.

*/
type Go [1]classdb.CSGMesh3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CSGMesh3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CSGMesh3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Mesh() gdclass.Mesh {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Mesh(class(self).GetMesh(gc))
}

func (self Go) SetMesh(value gdclass.Mesh) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMesh(value)
}

func (self Go) Material() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetMaterial(gc))
}

func (self Go) SetMaterial(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterial(value)
}

//go:nosplit
func (self class) SetMesh(mesh gdclass.Mesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGMesh3D.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMesh(ctx gd.Lifetime) gdclass.Mesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGMesh3D.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Mesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaterial(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGMesh3D.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGMesh3D.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsCSGMesh3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCSGMesh3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsCSGPrimitive3D() CSGPrimitive3D.GD { return *((*CSGPrimitive3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCSGPrimitive3D() CSGPrimitive3D.Go { return *((*CSGPrimitive3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCSGShape3D() CSGShape3D.GD { return *((*CSGShape3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCSGShape3D() CSGShape3D.Go { return *((*CSGShape3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.GD { return *((*GeometryInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() GeometryInstance3D.Go { return *((*GeometryInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCSGPrimitive3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCSGPrimitive3D(), name)
	}
}
func init() {classdb.Register("CSGMesh3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}