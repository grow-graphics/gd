//go:build !generate

package gd

import (
	"reflect"
	"unsafe"

	"runtime.link/mmm"
)

type Object struct {
	_   [0]*Object
	ptr Pointer
}

type NotificationType int32

func PointerWithOwnershipTransferredToGo(godot Lifetime, ptr uintptr) Pointer {
	return mmm.New[Pointer](godot.Lifetime, godot.API, [2]uintptr{ptr})
}

func PointerWithOwnershipTransferredToGodot(godot Lifetime, ptr Pointer) uintptr {
	raw := mmm.End(ptr)
	if raw[1] != 0 {
		panic("illegal transfer of ownership from Go -> Godot")
	}
	return raw[0]
}

func PointerMustAssertInstanceID(godot Lifetime, ptr uintptr) Pointer {
	var obj Object
	obj.SetPointer(mmm.Let[Pointer](godot.Lifetime, godot.API, [2]uintptr{ptr, 0}))
	return mmm.Let[Pointer](godot.Lifetime, godot.API, [2]uintptr{ptr, uintptr(godot.API.Object.GetInstanceID(obj))})
}

func PointerLifetimeBoundTo(godot Lifetime, obj Object, ptr uintptr) Pointer {
	return mmm.Let[Pointer](mmm.Life(obj.AsPointer()), godot.API, [2]uintptr{ptr, 0})
}

//go:nosplit
func (o Object) AsPointer() Pointer { return o.ptr }

func (o *Object) SetPointer(ptr Pointer) {
	o.ptr = ptr
}

func (self Object) AsObject() Object {
	return self
}

func (self Object) Free() { self.ptr.Free() }

type Class[T, S IsClass] struct {
	_     [0]*T
	super S

	KeepAlive Lifetime
	Temporary Lifetime
}

type ExtensionClass interface {
	PointerToClass
	SetKeepAlive(Lifetime)
	GetKeepAlive() Lifetime
	SetTemporary(Lifetime)
}

func (class *Class[T, S]) SetKeepAlive(godot Lifetime) { class.KeepAlive = godot }
func (class *Class[T, S]) SetTemporary(godot Lifetime) { class.Temporary = godot }
func (class *Class[T, S]) GetKeepAlive() Lifetime      { return class.KeepAlive }

func (class *Class[T, S]) SetPointer(ptr Pointer) {
	any(&class.super).(PointerToClass).SetPointer(ptr)
}

func (class Class[T, S]) AsPointer() Pointer {
	return class.super.AsPointer()
}

func (class Class[T, S]) AsObject() Object {
	var obj Object
	obj.SetPointer(class.super.AsPointer())
	return obj
}

func (class Class[T, S]) Pin() Lifetime {
	return Lifetime{
		Lifetime: mmm.Life(class.AsPointer()),
		API:      mmm.API(class.AsPointer()),
	}
}

func (class Class[T, S]) class() S { return class.super } //lint:ignore U1000 false positive.

func (class *Class[T, S]) Super() *S { return &class.super }

func (class Class[T, S]) Virtual(s string) reflect.Value {
	return class.super.Virtual(s)
}

func VirtualByName(class IsClass, name string) reflect.Value {
	return class.Virtual(name)
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T IsClass](godot Lifetime, class IsClass) (T, bool) {
	if ref, ok := godot.API.Instances[mmm.Get(class.AsPointer())[0]].(T); ok {
		return ref, true
	}
	var rtype = reflect.TypeOf([0]T{}).Elem()
	if rtype.Kind() == reflect.Pointer {
		rtype = rtype.Elem()
	}
	tmp := NewLifetime(godot.API)
	defer tmp.End()
	var classtag = godot.API.ClassDB.GetClassTag(tmp.StringName(rtype.Name()))
	casted := godot.API.Object.CastTo(class.AsObject(), classtag)
	if casted != (Object{}) && mmm.Get(casted.AsPointer()) != ([2]uintptr{}) {
		return (*(*T)(unsafe.Pointer(&casted))), true
	}
	var zero T
	return zero, false
}

type Singleton interface {
	IsSingleton()
}

type Pointer mmm.Pointer[API, Pointer, [2]uintptr]

func (ptr Pointer) Pointer() [2]uintptr {
	return mmm.Get(ptr)
}

func (ptr Pointer) Free() {
	if ptr.Pointer()[0] == 0 {
		mmm.End(ptr)
		return
	}

	var obj Object
	obj.ptr = ptr

	API := mmm.API(ptr)

	ctx := newContext(API)
	defer ctx.End()

	// Important that we don't destroy RefCounted objects, instead
	// they should be unreferenced instead.
	ref := API.Object.CastTo(obj, API.refCountedClassTag)
	if mmm.Get(ref.AsPointer()) != ([2]uintptr{}) {
		if (*(*RefCounted)(unsafe.Pointer(&ref))).Unreference() {
			API.Object.Destroy(obj)
		}
	} else {
		API.Object.Destroy(obj)
	}
	mmm.End(obj.AsPointer())
}

func (ptr Pointer) Virtual(string) reflect.Value {
	return reflect.Value{}
}

func (ptr Pointer) AsObject() Object {
	var obj Object
	obj.SetPointer(ptr)
	return obj
}

func (ptr Pointer) Pin() Lifetime {
	return Lifetime{
		Lifetime: mmm.Life(ptr),
		API:      mmm.API(ptr),
	}
}

func (ptr Pointer) AsPointer() Pointer     { return ptr }
func (ptr *Pointer) SetPointer(to Pointer) { *ptr = to }

type Extends[T IsClass] interface {
	IsClass
	class() T
}

type PointerToClass interface {
	IsClass
	SetPointer(Pointer)
}

type IsClass interface {
	AsPointer() Pointer
	Virtual(string) reflect.Value
	AsObject() Object
}

type IsPointer interface {
	AsPointer() Pointer
}
