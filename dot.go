package gd

import (
	"reflect"

	"github.com/readykit/gd/gdnative"
)

type Vector2 struct {
	X, Y float32
}

type Array []any

type Vector3 struct {
	X, Y, Z float32
}

type NodePath string

type Quaternion struct {
	X, Y, Z, W float32
}

type Dictionary map[string]any

type Transform2D struct {
	X, Y, Rotation float32
}

type Transform3D struct {
	Basis, Origin Vector3
}

type AABB struct {
	Min, Max Vector3
}

type Rect2 struct {
	Position, Size Vector2
}

type Rect2i struct {
	Position, Size Vector2i
}

type AudioFrame struct{}

type Color struct {
	R, G, B, A float32
}

type RID int64

type Vector2i struct {
	X, Y int32
}

type Callable struct {
}

type Plane struct {
	Normal, D float32
}

type Vector3i struct {
	X, Y, Z int32
}

type Basis struct {
	X, Y, Z Vector3
}

type PhysicsServer3DExtensionRayResult struct{}
type PhysicsServer3DExtensionShapeResult struct{}
type PhysicsServer3DExtensionShapeRestInfo struct{}
type PhysicsServer3DExtensionMotionResult struct{}
type ScriptLanguageExtensionProfilingInfo struct{}
type Glyph struct{}
type CaretInfo struct{}

type Class interface {
	~uintptr

	class() string
}

type Extension[Type any, Extends Class] struct {
	rtype reflect.Type
	ptype reflect.Type // pointer type.
	otype reflect.Type // owning class type.

	alloc func(Extends) Type

	class string           // base class name, cached on register.
	owner string           // owner class name, cached on register.
	slice []Instance[Type] // slice of instantiated instances.
}

func Register[Type any, Extends Class](fn func(Extends) Type) Extension[Type, Extends] {
	var zero Type
	var parent Extends

	var rtype = reflect.TypeOf(zero)

	var extension = Extension[Type, Extends]{
		alloc: fn,
		rtype: rtype,
		ptype: reflect.PtrTo(rtype),
		class: rtype.Name(),
		owner: parent.class(),
	}

	gdnative.OnLoad(extension.register)

	return extension
}

type Instance[Type any] struct {
	alive bool
	value Type
}

func (ext *Extension[Type, Extends]) Create() gdnative.Object {
	obj := gdnative.New(ext.owner)

	var instance Instance[Type]
	instance.alive = true
	instance.value = ext.alloc(Extends(obj))

	// reuse existing slot if possible.
	var id gdnative.InstanceID
	for i, slot := range ext.slice {
		if !slot.alive {
			ext.slice[i] = instance
			id = gdnative.InstanceID(i + 1)
			break
		}
	}

	if id == 0 {
		ext.slice = append(ext.slice, instance)
		id = gdnative.InstanceID(len(ext.slice))
	}

	obj.SetInstance(ext.class, id)

	return obj
}

func (ext *Extension[Type, Extends]) Delete(id gdnative.InstanceID) {
	ext.slice[id-1].alive = false
}

func (ext *Extension[Type, Extends]) Lookup(id gdnative.InstanceID) reflect.Value {
	return reflect.ValueOf(&ext.slice[id-1].value)
}

func (ext *Extension[Type, Extends]) register() {
	gdnative.RegisterClass(ext.class, ext.owner, ext)
	for i := 0; i < ext.ptype.NumMethod(); i++ {
		method := ext.ptype.Method(i)
		gdnative.RegisterMethod(ext.class, method, ext.Lookup)
	}
}

type extension interface {
	register()
}
