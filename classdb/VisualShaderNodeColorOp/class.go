// Code generated by the generate package DO NOT EDIT

// Package VisualShaderNodeColorOp provides methods for working with VisualShaderNodeColorOp object instances.
package VisualShaderNodeColorOp

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Angle"
import "graphics.gd/variant/Euler"
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

type _ gdclass.Node

var _ gd.Object
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
var _ Angle.Radians
var _ Euler.Radians
var _ = slices.Delete[[]struct{}, struct{}]

/*
ID is a typed object ID (reference) to an instance of this class, use it to store references to objects with
unknown lifetimes, as an ID will not panic on use if the underlying object has been destroyed.
*/
type ID Object.ID

func (id ID) Instance() (Instance, bool) { return Object.As[Instance](Object.ID(id).Instance()) }

/*
Extension can be embedded in a new struct to create an extension of this class.
T should be the type that is embedding this [Extension]
*/
type Extension[T gdclass.Interface] struct{ gdclass.Extension[T, Instance] }

/*
Applies [member operator] to two color inputs.
*/
type Instance [1]gdclass.VisualShaderNodeColorOp

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeColorOp() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeColorOp

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeColorOp"))
	casted := Instance{*(*gdclass.VisualShaderNodeColorOp)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Operator() Operator {
	return Operator(class(self).GetOperator())
}

func (self Instance) SetOperator(value Operator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op Operator) { //gd:VisualShaderNodeColorOp.set_operator
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() Operator { //gd:VisualShaderNodeColorOp.get_operator
	var frame = callframe.New()
	var r_ret = callframe.Ret[Operator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeColorOp() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeColorOp() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsVisualShaderNodeColorOp() Instance {
	return self.Super().AsVisualShaderNodeColorOp()
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsVisualShaderNode() VisualShaderNode.Instance {
	return self.Super().AsVisualShaderNode()
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsResource() Resource.Instance { return self.Super().AsResource() }
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsRefCounted() [1]gd.RefCounted { return self.Super().AsRefCounted() }
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
	gdclass.Register("VisualShaderNodeColorOp", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeColorOp{*(*gdclass.VisualShaderNodeColorOp)(unsafe.Pointer(&ptr))}
	})
}

type Operator int //gd:VisualShaderNodeColorOp.Operator

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
