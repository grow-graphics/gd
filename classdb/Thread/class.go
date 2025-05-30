// Code generated by the generate package DO NOT EDIT

// Package Thread provides methods for working with Thread object instances.
package Thread

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
A unit of execution in a process. Can run methods on [Object]s simultaneously. The use of synchronization via [Mutex] or [Semaphore] is advised if working with shared objects.
[b]Warning:[/b]
To ensure proper cleanup without crashes or deadlocks, when a [Thread]'s reference count reaches zero and it is therefore destroyed, the following conditions must be met:
- It must not have any [Mutex] objects locked.
- It must not be waiting on any [Semaphore] objects.
- [method wait_to_finish] should have been called on it.
*/
type Instance [1]gdclass.Thread

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.Thread

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsThread() Instance
}

/*
Starts a new [Thread] that calls [param callable].
If the method takes some arguments, you can pass them using [method Callable.bind].
The [param priority] of the [Thread] can be changed by passing a value from the [enum Priority] enum.
Returns [constant OK] on success, or [constant ERR_CANT_CREATE] on failure.
*/
func (self Instance) Start(callable func()) error { //gd:Thread.start
	return error(gd.ToError(Advanced(self).Start(Callable.New(callable), 1)))
}

/*
Starts a new [Thread] that calls [param callable].
If the method takes some arguments, you can pass them using [method Callable.bind].
The [param priority] of the [Thread] can be changed by passing a value from the [enum Priority] enum.
Returns [constant OK] on success, or [constant ERR_CANT_CREATE] on failure.
*/
func (self Expanded) Start(callable func(), priority Priority) error { //gd:Thread.start
	return error(gd.ToError(Advanced(self).Start(Callable.New(callable), priority)))
}

/*
Returns the current [Thread]'s ID, uniquely identifying it among all threads. If the [Thread] has not started running or if [method wait_to_finish] has been called, this returns an empty string.
*/
func (self Instance) GetId() string { //gd:Thread.get_id
	return string(Advanced(self).GetId().String())
}

/*
Returns [code]true[/code] if this [Thread] has been started. Once started, this will return [code]true[/code] until it is joined using [method wait_to_finish]. For checking if a [Thread] is still executing its task, use [method is_alive].
*/
func (self Instance) IsStarted() bool { //gd:Thread.is_started
	return bool(Advanced(self).IsStarted())
}

/*
Returns [code]true[/code] if this [Thread] is currently running the provided function. This is useful for determining if [method wait_to_finish] can be called without blocking the calling thread.
To check if a [Thread] is joinable, use [method is_started].
*/
func (self Instance) IsAlive() bool { //gd:Thread.is_alive
	return bool(Advanced(self).IsAlive())
}

/*
Joins the [Thread] and waits for it to finish. Returns the output of the [Callable] passed to [method start].
Should either be used when you want to retrieve the value returned from the method called by the [Thread] or before freeing the instance that contains the [Thread].
To determine if this can be called without blocking the calling thread, check if [method is_alive] is [code]false[/code].
*/
func (self Instance) WaitToFinish() any { //gd:Thread.wait_to_finish
	return any(Advanced(self).WaitToFinish().Interface())
}

/*
Sets whether the thread safety checks the engine normally performs in methods of certain classes (e.g., [Node]) should happen [b]on the current thread[/b].
The default, for every thread, is that they are enabled (as if called with [param enabled] being [code]true[/code]).
Those checks are conservative. That means that they will only succeed in considering a call thread-safe (and therefore allow it to happen) if the engine can guarantee such safety.
Because of that, there may be cases where the user may want to disable them ([param enabled] being [code]false[/code]) to make certain operations allowed again. By doing so, it becomes the user's responsibility to ensure thread safety (e.g., by using [Mutex]) for those objects that are otherwise protected by the engine.
[b]Note:[/b] This is an advanced usage of the engine. You are advised to use it only if you know what you are doing and there is no safer way.
[b]Note:[/b] This is useful for scripts running on either arbitrary [Thread] objects or tasks submitted to the [WorkerThreadPool]. It doesn't apply to code running during [Node] group processing, where the checks will be always performed.
[b]Note:[/b] Even in the case of having disabled the checks in a [WorkerThreadPool] task, there's no need to re-enable them at the end. The engine will do so.
*/
func SetThreadSafetyChecksEnabled(enabled bool) { //gd:Thread.set_thread_safety_checks_enabled
	self := Instance{}
	Advanced(self).SetThreadSafetyChecksEnabled(enabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Thread

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Thread"))
	casted := Instance{*(*gdclass.Thread)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Starts a new [Thread] that calls [param callable].
If the method takes some arguments, you can pass them using [method Callable.bind].
The [param priority] of the [Thread] can be changed by passing a value from the [enum Priority] enum.
Returns [constant OK] on success, or [constant ERR_CANT_CREATE] on failure.
*/
//go:nosplit
func (self class) Start(callable Callable.Function, priority Priority) Error.Code { //gd:Thread.start
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	callframe.Arg(frame, priority)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current [Thread]'s ID, uniquely identifying it among all threads. If the [Thread] has not started running or if [method wait_to_finish] has been called, this returns an empty string.
*/
//go:nosplit
func (self class) GetId() String.Readable { //gd:Thread.get_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this [Thread] has been started. Once started, this will return [code]true[/code] until it is joined using [method wait_to_finish]. For checking if a [Thread] is still executing its task, use [method is_alive].
*/
//go:nosplit
func (self class) IsStarted() bool { //gd:Thread.is_started
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_is_started, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this [Thread] is currently running the provided function. This is useful for determining if [method wait_to_finish] can be called without blocking the calling thread.
To check if a [Thread] is joinable, use [method is_started].
*/
//go:nosplit
func (self class) IsAlive() bool { //gd:Thread.is_alive
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_is_alive, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Joins the [Thread] and waits for it to finish. Returns the output of the [Callable] passed to [method start].
Should either be used when you want to retrieve the value returned from the method called by the [Thread] or before freeing the instance that contains the [Thread].
To determine if this can be called without blocking the calling thread, check if [method is_alive] is [code]false[/code].
*/
//go:nosplit
func (self class) WaitToFinish() variant.Any { //gd:Thread.wait_to_finish
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Thread.Bind_wait_to_finish, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets whether the thread safety checks the engine normally performs in methods of certain classes (e.g., [Node]) should happen [b]on the current thread[/b].
The default, for every thread, is that they are enabled (as if called with [param enabled] being [code]true[/code]).
Those checks are conservative. That means that they will only succeed in considering a call thread-safe (and therefore allow it to happen) if the engine can guarantee such safety.
Because of that, there may be cases where the user may want to disable them ([param enabled] being [code]false[/code]) to make certain operations allowed again. By doing so, it becomes the user's responsibility to ensure thread safety (e.g., by using [Mutex]) for those objects that are otherwise protected by the engine.
[b]Note:[/b] This is an advanced usage of the engine. You are advised to use it only if you know what you are doing and there is no safer way.
[b]Note:[/b] This is useful for scripts running on either arbitrary [Thread] objects or tasks submitted to the [WorkerThreadPool]. It doesn't apply to code running during [Node] group processing, where the checks will be always performed.
[b]Note:[/b] Even in the case of having disabled the checks in a [WorkerThreadPool] task, there's no need to re-enable them at the end. The engine will do so.
*/
//go:nosplit
func (self class) SetThreadSafetyChecksEnabled(enabled bool) { //gd:Thread.set_thread_safety_checks_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCallStatic(gd.Global.Methods.Thread.Bind_set_thread_safety_checks_enabled, frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsThread() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsThread() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsThread() Instance { return self.Super().AsThread() }
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
	gdclass.Register("Thread", func(ptr gd.Object) any { return [1]gdclass.Thread{*(*gdclass.Thread)(unsafe.Pointer(&ptr))} })
}

type Priority int //gd:Thread.Priority

const (
	/*A thread running with lower priority than normally.*/
	PriorityLow Priority = 0
	/*A thread with a standard priority.*/
	PriorityNormal Priority = 1
	/*A thread running with higher priority than normally.*/
	PriorityHigh Priority = 2
)
