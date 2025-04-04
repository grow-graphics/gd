// Package CallbackTweener provides methods for working with CallbackTweener object instances.
package CallbackTweener

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Tweener"
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
[CallbackTweener] is used to call a method in a tweening sequence. See [method Tween.tween_callback] for more usage information.
The tweener will finish automatically if the callback's target object is freed.
[b]Note:[/b] [method Tween.tween_callback] is the only correct way to create [CallbackTweener]. Any [CallbackTweener] created manually will not function correctly.
*/
type Instance [1]gdclass.CallbackTweener

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCallbackTweener() Instance
}

/*
Makes the callback call delayed by given time in seconds.
[b]Example:[/b] Call [method Node.queue_free] after 2 seconds:
[codeblock]
var tween = get_tree().create_tween()
tween.tween_callback(queue_free).set_delay(2)
[/codeblock]
*/
func (self Instance) SetDelay(delay Float.X) [1]gdclass.CallbackTweener { //gd:CallbackTweener.set_delay
	return [1]gdclass.CallbackTweener(Advanced(self).SetDelay(float64(delay)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CallbackTweener

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CallbackTweener"))
	casted := Instance{*(*gdclass.CallbackTweener)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Makes the callback call delayed by given time in seconds.
[b]Example:[/b] Call [method Node.queue_free] after 2 seconds:
[codeblock]
var tween = get_tree().create_tween()
tween.tween_callback(queue_free).set_delay(2)
[/codeblock]
*/
//go:nosplit
func (self class) SetDelay(delay float64) [1]gdclass.CallbackTweener { //gd:CallbackTweener.set_delay
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CallbackTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CallbackTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.CallbackTweener](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsCallbackTweener() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCallbackTweener() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTweener() Tweener.Advanced    { return *((*Tweener.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTweener() Tweener.Instance {
	return *((*Tweener.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Tweener.Advanced(self.AsTweener()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Tweener.Instance(self.AsTweener()), name)
	}
}
func init() {
	gdclass.Register("CallbackTweener", func(ptr gd.Object) any {
		return [1]gdclass.CallbackTweener{*(*gdclass.CallbackTweener)(unsafe.Pointer(&ptr))}
	})
}
