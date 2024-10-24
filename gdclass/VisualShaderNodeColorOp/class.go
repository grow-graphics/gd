package VisualShaderNodeColorOp

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
Applies [member operator] to two color inputs.

*/
type Go [1]classdb.VisualShaderNodeColorOp
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeColorOp
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeColorOp"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Operator() classdb.VisualShaderNodeColorOpOperator {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeColorOpOperator(class(self).GetOperator())
}

func (self Go) SetOperator(value classdb.VisualShaderNodeColorOpOperator) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op classdb.VisualShaderNodeColorOpOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeColorOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOperator() classdb.VisualShaderNodeColorOpOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeColorOpOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeColorOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeColorOp() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeColorOp() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("VisualShaderNodeColorOp", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Operator = classdb.VisualShaderNodeColorOpOperator

const (
/*Produce a screen effect with the following formula:
[codeblock]
result = vec3(1.0) - (vec3(1.0) - a) * (vec3(1.0) - b);
[/codeblock]*/
	OpScreen Operator = 0
/*Produce a difference effect with the following formula:
[codeblock]
result = abs(a - b);
[/codeblock]*/
	OpDifference Operator = 1
/*Produce a darken effect with the following formula:
[codeblock]
result = min(a, b);
[/codeblock]*/
	OpDarken Operator = 2
/*Produce a lighten effect with the following formula:
[codeblock]
result = max(a, b);
[/codeblock]*/
	OpLighten Operator = 3
/*Produce an overlay effect with the following formula:
[codeblock]
for (int i = 0; i < 3; i++) {
    float base = a[i];
    float blend = b[i];
    if (base < 0.5) {
        result[i] = 2.0 * base * blend;
    } else {
        result[i] = 1.0 - 2.0 * (1.0 - blend) * (1.0 - base);
    }
}
[/codeblock]*/
	OpOverlay Operator = 4
/*Produce a dodge effect with the following formula:
[codeblock]
result = a / (vec3(1.0) - b);
[/codeblock]*/
	OpDodge Operator = 5
/*Produce a burn effect with the following formula:
[codeblock]
result = vec3(1.0) - (vec3(1.0) - a) / b;
[/codeblock]*/
	OpBurn Operator = 6
/*Produce a soft light effect with the following formula:
[codeblock]
for (int i = 0; i < 3; i++) {
    float base = a[i];
    float blend = b[i];
    if (base < 0.5) {
        result[i] = base * (blend + 0.5);
    } else {
        result[i] = 1.0 - (1.0 - base) * (1.0 - (blend - 0.5));
    }
}
[/codeblock]*/
	OpSoftLight Operator = 7
/*Produce a hard light effect with the following formula:
[codeblock]
for (int i = 0; i < 3; i++) {
    float base = a[i];
    float blend = b[i];
    if (base < 0.5) {
        result[i] = base * (2.0 * blend);
    } else {
        result[i] = 1.0 - (1.0 - base) * (1.0 - 2.0 * (blend - 0.5));
    }
}
[/codeblock]*/
	OpHardLight Operator = 8
/*Represents the size of the [enum Operator] enum.*/
	OpMax Operator = 9
)