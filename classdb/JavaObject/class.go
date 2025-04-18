// Package JavaObject provides methods for working with JavaObject object instances.
package JavaObject

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
Represents an object from the Java Native Interface. It can be returned from Java methods called on [JavaClass] or other [JavaObject]s. See [JavaClassWrapper] for an example.
[b]Note:[/b] This class only works on Android. On any other platform, this class does nothing.
[b]Note:[/b] This class is not to be confused with [JavaScriptObject].
*/
type Instance [1]gdclass.JavaObject

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsJavaObject() Instance
}

/*
Returns the [JavaClass] that this object is an instance of.
*/
func (self Instance) GetJavaClass() [1]gdclass.JavaClass { //gd:JavaObject.get_java_class
	return [1]gdclass.JavaClass(Advanced(self).GetJavaClass())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.JavaObject

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("JavaObject"))
	casted := Instance{*(*gdclass.JavaObject)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the [JavaClass] that this object is an instance of.
*/
//go:nosplit
func (self class) GetJavaClass() [1]gdclass.JavaClass { //gd:JavaObject.get_java_class
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaObject.Bind_get_java_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.JavaClass{gd.PointerWithOwnershipTransferredToGo[gdclass.JavaClass](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsJavaObject() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJavaObject() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("JavaObject", func(ptr gd.Object) any { return [1]gdclass.JavaObject{*(*gdclass.JavaObject)(unsafe.Pointer(&ptr))} })
}
