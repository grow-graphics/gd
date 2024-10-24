package VisualShaderNodeVectorFunc

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeVectorBase"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A visual shader node able to perform different functions using vectors.

*/
type Go [1]classdb.VisualShaderNodeVectorFunc
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeVectorFunc
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeVectorFunc"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Function() classdb.VisualShaderNodeVectorFuncFunction {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.VisualShaderNodeVectorFuncFunction(class(self).GetFunction())
}

func (self Go) SetFunction(value classdb.VisualShaderNodeVectorFuncFunction) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFunction(value)
}

//go:nosplit
func (self class) SetFunction(fn classdb.VisualShaderNodeVectorFuncFunction)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeVectorFunc.Bind_set_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFunction() classdb.VisualShaderNodeVectorFuncFunction {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeVectorFuncFunction](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeVectorFunc.Bind_get_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeVectorFunc() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeVectorFunc() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.GD { return *((*VisualShaderNodeVectorBase.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Go { return *((*VisualShaderNodeVectorBase.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeVectorBase(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeVectorBase(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeVectorFunc", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Function = classdb.VisualShaderNodeVectorFuncFunction

const (
/*Normalizes the vector so that it has a length of [code]1[/code] but points in the same direction.*/
	FuncNormalize Function = 0
/*Clamps the value between [code]0.0[/code] and [code]1.0[/code].*/
	FuncSaturate Function = 1
/*Returns the opposite value of the parameter.*/
	FuncNegate Function = 2
/*Returns [code]1/vector[/code].*/
	FuncReciprocal Function = 3
/*Returns the absolute value of the parameter.*/
	FuncAbs Function = 4
/*Returns the arc-cosine of the parameter.*/
	FuncAcos Function = 5
/*Returns the inverse hyperbolic cosine of the parameter.*/
	FuncAcosh Function = 6
/*Returns the arc-sine of the parameter.*/
	FuncAsin Function = 7
/*Returns the inverse hyperbolic sine of the parameter.*/
	FuncAsinh Function = 8
/*Returns the arc-tangent of the parameter.*/
	FuncAtan Function = 9
/*Returns the inverse hyperbolic tangent of the parameter.*/
	FuncAtanh Function = 10
/*Finds the nearest integer that is greater than or equal to the parameter.*/
	FuncCeil Function = 11
/*Returns the cosine of the parameter.*/
	FuncCos Function = 12
/*Returns the hyperbolic cosine of the parameter.*/
	FuncCosh Function = 13
/*Converts a quantity in radians to degrees.*/
	FuncDegrees Function = 14
/*Base-e Exponential.*/
	FuncExp Function = 15
/*Base-2 Exponential.*/
	FuncExp2 Function = 16
/*Finds the nearest integer less than or equal to the parameter.*/
	FuncFloor Function = 17
/*Computes the fractional part of the argument.*/
	FuncFract Function = 18
/*Returns the inverse of the square root of the parameter.*/
	FuncInverseSqrt Function = 19
/*Natural logarithm.*/
	FuncLog Function = 20
/*Base-2 logarithm.*/
	FuncLog2 Function = 21
/*Converts a quantity in degrees to radians.*/
	FuncRadians Function = 22
/*Finds the nearest integer to the parameter.*/
	FuncRound Function = 23
/*Finds the nearest even integer to the parameter.*/
	FuncRoundeven Function = 24
/*Extracts the sign of the parameter, i.e. returns [code]-1[/code] if the parameter is negative, [code]1[/code] if it's positive and [code]0[/code] otherwise.*/
	FuncSign Function = 25
/*Returns the sine of the parameter.*/
	FuncSin Function = 26
/*Returns the hyperbolic sine of the parameter.*/
	FuncSinh Function = 27
/*Returns the square root of the parameter.*/
	FuncSqrt Function = 28
/*Returns the tangent of the parameter.*/
	FuncTan Function = 29
/*Returns the hyperbolic tangent of the parameter.*/
	FuncTanh Function = 30
/*Returns a value equal to the nearest integer to the parameter whose absolute value is not larger than the absolute value of the parameter.*/
	FuncTrunc Function = 31
/*Returns [code]1.0 - vector[/code].*/
	FuncOneminus Function = 32
/*Represents the size of the [enum Function] enum.*/
	FuncMax Function = 33
)