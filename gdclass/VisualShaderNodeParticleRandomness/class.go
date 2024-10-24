package VisualShaderNodeParticleRandomness

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Randomness node will output pseudo-random values of the given type based on the specified minimum and maximum values.

*/
type Go [1]classdb.VisualShaderNodeParticleRandomness
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeParticleRandomness
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeParticleRandomness"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) OpType() classdb.VisualShaderNodeParticleRandomnessOpType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeParticleRandomnessOpType(class(self).GetOpType())
}

func (self Go) SetOpType(value classdb.VisualShaderNodeParticleRandomnessOpType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOpType(value)
}

//go:nosplit
func (self class) SetOpType(atype classdb.VisualShaderNodeParticleRandomnessOpType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleRandomness.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOpType() classdb.VisualShaderNodeParticleRandomnessOpType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeParticleRandomnessOpType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleRandomness.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParticleRandomness() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeParticleRandomness() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeParticleRandomness", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type OpType = classdb.VisualShaderNodeParticleRandomnessOpType

const (
/*A floating-point scalar.*/
	OpTypeScalar OpType = 0
/*A 2D vector type.*/
	OpTypeVector2d OpType = 1
/*A 3D vector type.*/
	OpTypeVector3d OpType = 2
/*A 4D vector type.*/
	OpTypeVector4d OpType = 3
/*Represents the size of the [enum OpType] enum.*/
	OpTypeMax OpType = 4
)