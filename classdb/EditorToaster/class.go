// Code generated by the generate package DO NOT EDIT

// Package EditorToaster provides methods for working with EditorToaster object instances.
package EditorToaster

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
import "graphics.gd/classdb/BoxContainer"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/HBoxContainer"
import "graphics.gd/classdb/Node"
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
This object manages the functionality and display of toast notifications within the editor, ensuring timely and informative alerts are presented to users.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_editor_toaster].
*/
type Instance [1]gdclass.EditorToaster

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.EditorToaster

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorToaster() Instance
}

/*
Pushes a toast notification to the editor for display.
*/
func (self Instance) PushToast(message string) { //gd:EditorToaster.push_toast
	Advanced(self).PushToast(String.New(message), 0, String.New(""))
}

/*
Pushes a toast notification to the editor for display.
*/
func (self Expanded) PushToast(message string, severity Severity, tooltip string) { //gd:EditorToaster.push_toast
	Advanced(self).PushToast(String.New(message), severity, String.New(tooltip))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorToaster

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorToaster"))
	casted := Instance{*(*gdclass.EditorToaster)(unsafe.Pointer(&object))}
	return casted
}

/*
Pushes a toast notification to the editor for display.
*/
//go:nosplit
func (self class) PushToast(message String.Readable, severity Severity, tooltip String.Readable) { //gd:EditorToaster.push_toast
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(message)))
	callframe.Arg(frame, severity)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorToaster.Bind_push_toast, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsEditorToaster() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorToaster() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsEditorToaster() Instance { return self.Super().AsEditorToaster() }
func (self class) AsHBoxContainer() HBoxContainer.Advanced {
	return *((*HBoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsHBoxContainer() HBoxContainer.Instance {
	return self.Super().AsHBoxContainer()
}
func (self Instance) AsHBoxContainer() HBoxContainer.Instance {
	return *((*HBoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsBoxContainer() BoxContainer.Advanced {
	return *((*BoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsBoxContainer() BoxContainer.Instance {
	return self.Super().AsBoxContainer()
}
func (self Instance) AsBoxContainer() BoxContainer.Instance {
	return *((*BoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsContainer() Container.Instance { return self.Super().AsContainer() }
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced         { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsControl() Control.Instance { return self.Super().AsControl() }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsCanvasItem() CanvasItem.Instance { return self.Super().AsCanvasItem() }
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced         { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance      { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(HBoxContainer.Advanced(self.AsHBoxContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(HBoxContainer.Instance(self.AsHBoxContainer()), name)
	}
}
func init() {
	gdclass.Register("EditorToaster", func(ptr gd.Object) any {
		return [1]gdclass.EditorToaster{*(*gdclass.EditorToaster)(unsafe.Pointer(&ptr))}
	})
}

type Severity int //gd:EditorToaster.Severity

const (
	/*Toast will display with an INFO severity.*/
	SeverityInfo Severity = 0
	/*Toast will display with a WARNING severity and have a corresponding color.*/
	SeverityWarning Severity = 1
	/*Toast will display with an ERROR severity and have a corresponding color.*/
	SeverityError Severity = 2
)
