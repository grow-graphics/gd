package main

import (
	"fmt"
	"runtime"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

/*
HelloWorld is a simple extension to demonstrate how to export
Go methods so that they can be used in scripts.
*/
type HelloWorld struct {
	gd.Class[HelloWorld, gd.Object]
}

// Print prints "Hello World"
func (h *HelloWorld) Print() {
	fmt.Println("Hello World")
}

// Echo prints the given string, signalling that it
// was printed by Go code.
func (h *HelloWorld) Echo(s gd.String) {
	fmt.Println(s.String() + " from Go!")
}

// Arch returns the current GOARCH value.
func (h *HelloWorld) Arch() gd.String {
	return h.Temporary.String(runtime.GOARCH)
}

func (h *HelloWorld) GetBar(message gd.String) gd.Object {
	var result = gd.Create(h.Temporary, new(Bar))
	result.Message = message.Copy(result.KeepAlive)
	return result.AsObject()
}

type Rotator struct {
	gd.Class[Rotator, gd.Sprite2D]
}

func (r *Rotator) Process(delta gd.Float) {
	node2D := r.Super().AsNode2D()
	node2D.SetRotation(node2D.GetRotation() + delta)
}

type StartedSignalEmitter struct {
	gd.Class[StartedSignalEmitter, gd.Node]

	started gd.SignalAs[func()]
}

func (r *StartedSignalEmitter) Ready() {
	if r.started.Emit != nil {
		r.started.Emit()
	}
}

type MyClassWithConstants struct {
	gd.Class[MyClassWithConstants, gd.Node2D]
}

func (*MyClassWithConstants) OnRegister(godot gd.Context) {
	godot.Register(gd.Enum[MyClassWithConstants, int]{
		Name: "MyEnum",
		Values: map[string]int{
			"Value1": 1,
			"Value2": 2,
		},
	})
}

/*
ExtendedNode demonstrates how to call the methods of builtin objects.
*/
type ExtendedNode struct {
	gd.Class[ExtendedNode, gd.Node2D]

	StringField gd.String
}

func (e *ExtendedNode) Ready() {
	fmt.Println("Ready!")

	node := e.Super()

	fmt.Println("class:", node.AsObject().GetClass(e.Temporary).String())

	var obj = gd.Create(e.Temporary, new(gd.Object))
	fmt.Println(obj.GetClass(e.Temporary).String())

	fmt.Println(gd.Engine(e.Temporary).GetSingletonList(e.Temporary))
	fmt.Println("Scene is ready!")

	fmt.Println("sin=", e.Temporary.Sin(1.5))

	fmt.Println("rotation=", node.GetRotation())
	node.SetRotation(3.14)
	fmt.Println("rotation=", node.GetRotation())

	pos := node.GetPosition()

	fmt.Println("position=", pos)

	pos[0] = 100

	node.SetPosition(pos)
	fmt.Println("position=", pos)

	variant := e.Temporary.Variant(node)
	result, err := variant.Call(e.Temporary, e.Temporary.StringName("get_position"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result.Interface(e.Temporary))
	}
}

type Bar struct {
	gd.Class[Bar, gd.Object]

	Message gd.String
}

// main init function, where the extensions are exported so that
// they are available to the engine.
func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	fmt.Println("Godot Version is: ", godot.Version())
	fmt.Println("Extension: ", godot.GetLibraryPath())
	gd.Register[HelloWorld](godot)
	gd.Register[Bar](godot)
	gd.Register[ExtendedNode](godot)
	gd.Register[Rotator](godot)
	gd.Register[StartedSignalEmitter](godot)
	gd.Register[MyClassWithConstants](godot)
}
