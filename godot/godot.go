package godot

import (
	"reflect"

	"github.com/readykit/gd"
	"github.com/readykit/gd/internal/gdc"
)

var Engine gd.Engine

var DisplayServer gd.DisplayServer

func init() {
	gdc.OnLoad(func() {
		Engine = gd.Engine(gdc.GetSingleton("Engine\000"))
		//DisplayServer = gd.DisplayServer(Engine.GetSingleton("DisplayServer\000"))
	})
}

type Extension[Type any, Extends gd.Class] struct {
	rtype reflect.Type
	ptype reflect.Type // pointer type.
	otype reflect.Type // owning class type.

	alloc func(Extends) Type

	class string           // base class name, cached on register.
	owner string           // owner class name, cached on register.
	slice []Instance[Type] // slice of instantiated instances.
}

func Register[Type any, Extends gd.Class](fn func(Extends) Type) Extension[Type, Extends] {
	var zero Type
	var parent Extends

	var rtype = reflect.TypeOf(zero)

	var extension = Extension[Type, Extends]{
		alloc: fn,
		rtype: rtype,
		ptype: reflect.PtrTo(rtype),
		class: rtype.Name() + "\000",
		owner: gd.ClassNameOf(parent),
	}

	gdc.OnLoad(extension.register)

	return extension
}

type Instance[Type any] struct {
	alive bool
	value Type
}

func (ext *Extension[Type, Extends]) Name() string {
	return ext.class
}

func (ext *Extension[Type, Extends]) Godot() string {
	return ext.owner
}

func (ext *Extension[Type, Extends]) GetVirtual(name string) (reflect.Method, bool) {
	var zero Type
	var core Extends
	return gd.VirtualOf(core, reflect.TypeOf(zero), name)
}

func (ext *Extension[Type, Extends]) Create(class gdc.ClassID) gdc.Object {
	obj := gdc.New(ext.owner)

	var instance Instance[Type]
	instance.alive = true
	instance.value = ext.alloc(Extends(obj))

	// reuse existing slot if possible.
	var id gdc.InstanceID
	for i, slot := range ext.slice {
		if !slot.alive {
			ext.slice[i] = instance
			id = gdc.InstanceID(i + 1)
			break
		}
	}

	if id == 0 {
		ext.slice = append(ext.slice, instance)
		id = gdc.InstanceID(len(ext.slice))
	}

	obj.SetInstance(class, id)

	return obj
}

func (ext *Extension[Type, Extends]) Delete(id gdc.InstanceID) {
	ext.slice[id-1].alive = false
}

func (ext *Extension[Type, Extends]) Lookup(id gdc.InstanceID) reflect.Value {
	return reflect.ValueOf(&ext.slice[id-1].value)
}

func (ext *Extension[Type, Extends]) Methods() []reflect.Method {
	var methods []reflect.Method
	for i := 0; i < ext.ptype.NumMethod(); i++ {
		methods = append(methods, ext.ptype.Method(i))
	}
	return methods
}

func (ext *Extension[Type, Extends]) register() {
	gdc.RegisterClass(ext)
}
