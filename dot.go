package gd

import "reflect"

type Class interface {
	~uintptr

	class() string
	virtual(reflect.Type, string) (reflect.Method, bool)
}

func ClassNameOf[T Class](class T) string {
	return class.class()
}

func VirtualOf[T Class](class T, rtype reflect.Type, virtual string) (reflect.Method, bool) {
	return class.virtual(rtype, virtual)
}

type Vector2 struct {
	X, Y float32
}

type Array uintptr

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

type Callable uintptr

type Plane struct {
	Normal, D float32
}

type Vector3i struct {
	X, Y, Z int32
}

type Basis struct {
	X, Y, Z Vector3
}

type Name string

type String uintptr

type PhysicsServer3DExtensionRayResult struct{}
type PhysicsServer3DExtensionShapeResult struct{}
type PhysicsServer3DExtensionShapeRestInfo struct{}
type PhysicsServer3DExtensionMotionResult struct{}
type ScriptLanguageExtensionProfilingInfo struct{}
type Glyph struct{}
type CaretInfo struct{}
