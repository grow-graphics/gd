package main

import (
	"fmt"

	"github.com/readykit/gd"
)

func main() {}

type HelloWorld struct {
	Object gd.Object
}

var NewHelloWorld = gd.Register(func(hello *HelloWorld) gd.Object {
	hello.Object = gd.NewObject()
	return hello.Object
})

func (h HelloWorld) Print() {
	fmt.Println("Hello World!")
}

func (h HelloWorld) Echo(s string) {
	fmt.Println(s + " from Go!")
}

type ExtendedNode struct {
	Node2D gd.Node2D
}

var NewExtendedNode = gd.Register(func(node *ExtendedNode) gd.Node2D {
	node.Node2D = gd.NewNode2D()
	return node.Node2D
})

func (e ExtendedNode) Ready() {
	if gd.Engine.IsEditorHint() {
		fmt.Println(gd.Engine.GetLicenseText())
		return
	}

	var obj = gd.NewObject()
	fmt.Println(obj.GetClass())
	defer obj.Free()

	//gd.LoadSingletons()

	fmt.Println(gd.Engine.GetSingletonList())

	fmt.Println("Scene is ready!")

	fmt.Println("sin=", gd.Sin(1.5))

	fmt.Println("rotation=", e.Node2D.GetRotation())
	e.Node2D.SetRotation(3.14)
	fmt.Println("rotation=", e.Node2D.GetRotation())

	pos := e.Node2D.GetPosition()

	fmt.Println("position=", pos)

	//fmt.Println(godot.Engine.GetSingletonList())

	//godot.DisplayServer.WindowSetCurrentScreen(1, 0)
}
