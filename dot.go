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
	class() string
}

type Extension[T Class] struct {
	rtype reflect.Type
	ptype reflect.Type // pointer type.
	otype reflect.Type // owning class type.

	class string        // base class name, cached on register.
	owner string        // owner class name, cached on register.
	slice []Instance[T] // slice of instantiated instances.
}

type Instance[T Class] struct {
	alive bool
	value T
}

func (ext *Extension[T]) Create() gdnative.Object {
	obj := gdnative.New(ext.owner)

	var instance Instance[T]
	instance.alive = true
	// reuse existing memory if possible.

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

	reflect.ValueOf(&ext.slice[id-1].value).Elem().Field(0).SetUint(uint64(obj))

	obj.SetInstance(ext.class, id)

	return obj
}

func (ext *Extension[T]) Delete(id gdnative.InstanceID) {
	ext.slice[id-1].alive = false
}

func (ext *Extension[T]) Lookup(id gdnative.InstanceID) reflect.Value {
	return reflect.ValueOf(&ext.slice[id-1].value)
}

func (ext *Extension[T]) register() {
	var native T

	ext.rtype = reflect.TypeOf(native)
	ext.ptype = reflect.TypeOf(&native)
	ext.otype = ext.rtype.Field(0).Type
	ext.class = ext.rtype.Name()
	ext.owner = native.class()

	gdnative.RegisterClass(ext.class, ext.owner, ext)
	for i := 0; i < ext.ptype.NumMethod(); i++ {
		method := ext.rtype.Method(i)
		gdnative.RegisterMethod(ext.class, method, ext.Lookup)
	}
}

type extension interface {
	register()
}

func Register(extensions ...extension) {
	for _, extension := range extensions {
		extension.register()
	}
}
