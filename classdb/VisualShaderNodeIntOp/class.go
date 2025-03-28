// Package VisualShaderNodeIntOp provides methods for working with VisualShaderNodeIntOp object instances.
package VisualShaderNodeIntOp

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Applies [member operator] to two integer inputs: [code]a[/code] and [code]b[/code].
*/
type Instance [1]gdclass.VisualShaderNodeIntOp

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeIntOp() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeIntOp

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeIntOp"))
	casted := Instance{*(*gdclass.VisualShaderNodeIntOp)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Operator() gdclass.VisualShaderNodeIntOpOperator {
	return gdclass.VisualShaderNodeIntOpOperator(class(self).GetOperator())
}

func (self Instance) SetOperator(value gdclass.VisualShaderNodeIntOpOperator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op gdclass.VisualShaderNodeIntOpOperator) { //gd:VisualShaderNodeIntOp.set_operator
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeIntOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() gdclass.VisualShaderNodeIntOpOperator { //gd:VisualShaderNodeIntOp.get_operator
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeIntOpOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeIntOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeIntOp() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeIntOp() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Advanced(self.AsVisualShaderNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Instance(self.AsVisualShaderNode()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeIntOp", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeIntOp{*(*gdclass.VisualShaderNodeIntOp)(unsafe.Pointer(&ptr))}
	})
}

type Operator = gdclass.VisualShaderNodeIntOpOperator //gd:VisualShaderNodeIntOp.Operator

const (
	/*Sums two numbers using [code]a + b[/code].*/
	OpAdd Operator = 0
	/*Subtracts two numbers using [code]a - b[/code].*/
	OpSub Operator = 1
	/*Multiplies two numbers using [code]a * b[/code].*/
	OpMul Operator = 2
	/*Divides two numbers using [code]a / b[/code].*/
	OpDiv Operator = 3
	/*Calculates the remainder of two numbers using [code]a % b[/code].*/
	OpMod Operator = 4
	/*Returns the greater of two numbers. Translates to [code]max(a, b)[/code] in the Godot Shader Language.*/
	OpMax Operator = 5
	/*Returns the lesser of two numbers. Translates to [code]max(a, b)[/code] in the Godot Shader Language.*/
	OpMin Operator = 6
	/*Returns the result of bitwise [code]AND[/code] operation on the integer. Translates to [code]a & b[/code] in the Godot Shader Language.*/
	OpBitwiseAnd Operator = 7
	/*Returns the result of bitwise [code]OR[/code] operation for two integers. Translates to [code]a | b[/code] in the Godot Shader Language.*/
	OpBitwiseOr Operator = 8
	/*Returns the result of bitwise [code]XOR[/code] operation for two integers. Translates to [code]a ^ b[/code] in the Godot Shader Language.*/
	OpBitwiseXor Operator = 9
	/*Returns the result of bitwise left shift operation on the integer. Translates to [code]a << b[/code] in the Godot Shader Language.*/
	OpBitwiseLeftShift Operator = 10
	/*Returns the result of bitwise right shift operation on the integer. Translates to [code]a >> b[/code] in the Godot Shader Language.*/
	OpBitwiseRightShift Operator = 11
	/*Represents the size of the [enum Operator] enum.*/
	OpEnumSize Operator = 12
)
