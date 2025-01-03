package main

import (
	"fmt"
	"runtime"

	gd "graphics.gd"
	"graphics.gd/defined"
	"graphics.gd/objects"
	"graphics.gd/objects/Engine"
	"graphics.gd/objects/Node"
	"graphics.gd/objects/Node2D"
	"graphics.gd/objects/Sprite2D"
	"graphics.gd/startup"
	"graphics.gd/variant"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
)

/*
HelloWorld is a simple extension to demonstrate how to export
Go methods so that they can be used in scripts.
*/
type HelloWorld struct {
	defined.Object[HelloWorld, objects.Instance]
}

// Print prints "Hello World"
func (h *HelloWorld) Print() { fmt.Println("Hello World") }

// Echo prints the given string, signalling that it
// was printed by Go code.
func (h *HelloWorld) Echo(s gd.String) { fmt.Println(s.String() + " from Go!") }

// Arch returns the current GOARCH value.
func (h *HelloWorld) Arch() string { return runtime.GOARCH }

func (h *HelloWorld) GetBar(message string) *Bar {
	return &Bar{
		Message: message,
	}
}

type Rotator struct {
	defined.Object[Rotator, Sprite2D.Instance]

	Map map[string]int
}

func (r *Rotator) Process(delta Float.X) {
	node2D := r.Super().AsNode2D()
	node2D.SetRotation(node2D.Rotation() + delta)
}

type StartedSignalEmitter struct {
	defined.Object[StartedSignalEmitter, Node.Instance]

	started chan<- struct{}
}

func (r *StartedSignalEmitter) Ready() {
	select {
	case r.started <- struct{}{}:
	default:
	}
}

type MyClassWithConstants struct {
	defined.Object[MyClassWithConstants, Node2D.Instance]
}

func (*MyClassWithConstants) OnRegister() {
	/*gdextension.Register(gd.Enum[MyClassWithConstants, int]{
		Name: "MyEnum",
		Values: map[string]int{
			"Value1": 1,
			"Value2": 2,
		},
	})*/
}

/*
ExtendedNode demonstrates how to call the methods of builtin objects.
*/
type ExtendedNode struct {
	defined.Object[ExtendedNode, Node2D.Instance]

	StringField string
}

func (e *ExtendedNode) Ready() {
	fmt.Println("Ready!")

	node := e.Super()

	fmt.Println("class:", node.AsObject().GetClass().String())

	var obj = objects.New()
	fmt.Println(obj.GetClass().String())

	fmt.Println(Engine.GetSingletonList())
	fmt.Println("Scene is ready!")

	fmt.Println("sin=", Angle.Sin(1.5))

	fmt.Println("rotation=", node.Rotation())
	node.SetRotation(3.14)
	fmt.Println("rotation=", node.Rotation())

	pos := node.Position()

	fmt.Println("position=", pos)

	pos.X = 100

	node.SetPosition(pos)
	fmt.Println("position=", pos)

	v := variant.New(node)
	fmt.Printf("result: %[1]T %[1]v\n", variant.Call(v, "get_position"))
}

type Bar struct {
	Message string
}

// main init function, where the extensions are exported so that
// they are available to the engine.
func main() {
	defined.InEditor[HelloWorld]()
	defined.InEditor[ExtendedNode]()
	defined.InEditor[Rotator]()
	defined.InEditor[StartedSignalEmitter]()
	defined.InEditor[MyClassWithConstants]()

	fmt.Println("Engine Version is: ", gd.Version())
	fmt.Println("Extension: ", startup.LibraryPath())
	startup.Engine()
	fmt.Println("Shutting Down...")
}
