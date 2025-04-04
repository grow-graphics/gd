// Package JavaClassWrapper provides methods for working with JavaClassWrapper object instances.
package JavaClassWrapper

import "unsafe"
import "sync"
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
The JavaClassWrapper singleton provides a way for the Godot application to send and receive data through the [url=https://developer.android.com/training/articles/perf-jni]Java Native Interface[/url] (JNI).
[b]Note:[/b] This singleton is only available in Android builds.
[codeblock]
var LocalDateTime = JavaClassWrapper.wrap("java.time.LocalDateTime")
var DateTimeFormatter = JavaClassWrapper.wrap("java.time.format.DateTimeFormatter")

var datetime = LocalDateTime.now()
var formatter = DateTimeFormatter.ofPattern("dd-MM-yyyy HH:mm:ss")

print(datetime.format(formatter))
[/codeblock]
[b]Warning:[/b] When calling Java methods, be sure to check [method JavaClassWrapper.get_exception] to check if the method threw an exception.
*/
var self [1]gdclass.JavaClassWrapper
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.JavaClassWrapper)
	self = *(*[1]gdclass.JavaClassWrapper)(unsafe.Pointer(&obj))
}

/*
Wraps a class defined in Java, and returns it as a [JavaClass] [Object] type that Godot can interact with.
[b]Note:[/b] This method only works on Android. On every other platform, this method does nothing and returns an empty [JavaClass].
*/
func Wrap(name string) [1]gdclass.JavaClass { //gd:JavaClassWrapper.wrap
	once.Do(singleton)
	return [1]gdclass.JavaClass(Advanced().Wrap(String.New(name)))
}

/*
Returns the Java exception from the last call into a Java class. If there was no exception, it will return [code]null[/code].
[b]Note:[/b] This method only works on Android. On every other platform, this method will always return [code]null[/code].
*/
func GetException() [1]gdclass.JavaObject { //gd:JavaClassWrapper.get_exception
	once.Do(singleton)
	return [1]gdclass.JavaObject(Advanced().GetException())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.JavaClassWrapper

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Wraps a class defined in Java, and returns it as a [JavaClass] [Object] type that Godot can interact with.
[b]Note:[/b] This method only works on Android. On every other platform, this method does nothing and returns an empty [JavaClass].
*/
//go:nosplit
func (self class) Wrap(name String.Readable) [1]gdclass.JavaClass { //gd:JavaClassWrapper.wrap
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaClassWrapper.Bind_wrap, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.JavaClass{gd.PointerWithOwnershipTransferredToGo[gdclass.JavaClass](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the Java exception from the last call into a Java class. If there was no exception, it will return [code]null[/code].
[b]Note:[/b] This method only works on Android. On every other platform, this method will always return [code]null[/code].
*/
//go:nosplit
func (self class) GetException() [1]gdclass.JavaObject { //gd:JavaClassWrapper.get_exception
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaClassWrapper.Bind_get_exception, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.JavaObject{gd.PointerWithOwnershipTransferredToGo[gdclass.JavaObject](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("JavaClassWrapper", func(ptr gd.Object) any {
		return [1]gdclass.JavaClassWrapper{*(*gdclass.JavaClassWrapper)(unsafe.Pointer(&ptr))}
	})
}
