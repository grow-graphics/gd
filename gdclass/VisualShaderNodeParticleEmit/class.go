package VisualShaderNodeParticleEmit

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
This node internally calls [code]emit_subparticle[/code] shader method. It will emit a particle from the configured sub-emitter and also allows to customize how its emitted. Requires a sub-emitter assigned to the particles node with this shader.

*/
type Go [1]classdb.VisualShaderNodeParticleEmit
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeParticleEmit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeParticleEmit"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Flags() classdb.VisualShaderNodeParticleEmitEmitFlags {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeParticleEmitEmitFlags(class(self).GetFlags())
}

func (self Go) SetFlags(value classdb.VisualShaderNodeParticleEmitEmitFlags) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFlags(value)
}

//go:nosplit
func (self class) SetFlags(flags classdb.VisualShaderNodeParticleEmitEmitFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleEmit.Bind_set_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlags() classdb.VisualShaderNodeParticleEmitEmitFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeParticleEmitEmitFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeParticleEmit.Bind_get_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParticleEmit() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeParticleEmit() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("VisualShaderNodeParticleEmit", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type EmitFlags = classdb.VisualShaderNodeParticleEmitEmitFlags

const (
/*If enabled, the particle starts with the position defined by this node.*/
	EmitFlagPosition EmitFlags = 1
/*If enabled, the particle starts with the rotation and scale defined by this node.*/
	EmitFlagRotScale EmitFlags = 2
/*If enabled,the particle starts with the velocity defined by this node.*/
	EmitFlagVelocity EmitFlags = 4
/*If enabled, the particle starts with the color defined by this node.*/
	EmitFlagColor EmitFlags = 8
/*If enabled, the particle starts with the [code]CUSTOM[/code] data defined by this node.*/
	EmitFlagCustom EmitFlags = 16
)