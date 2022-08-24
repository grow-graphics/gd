package main

import (
	"fmt"

	"github.com/readykit/gd"
	"github.com/readykit/gd/godot"
)

func main() {}

type HelloWorld struct {
	Object gd.Object
}

func NewHelloWorld(obj gd.Object) HelloWorld {
	return HelloWorld{obj}
}

func (h HelloWorld) Print() {
	fmt.Println("Hello World!")
}

func (h HelloWorld) Echo(s string) {
	fmt.Println(s + " from Go!")
}

var gdHelloWorld = gd.Register(NewHelloWorld)

type ExtendedNode struct {
	Node2D gd.Node2D
}

func NewExtendedNode(obj gd.Node2D) ExtendedNode {
	return ExtendedNode{Node2D: obj}
}

func (e ExtendedNode) Ready() {
	if godot.Engine.IsEditorHint() {

		fmt.Println(godot.Engine.GetLicenseText())

		return
	}

	fmt.Println("Scene is ready!")
	fmt.Println("rotation=", e.Node2D.GetRotation())
	e.Node2D.SetRotation(3.14)
	fmt.Println("rotation=", e.Node2D.GetRotation())
}

var gdExtendedNode = gd.Register(NewExtendedNode)
