package main

import (
	"fmt"

	"github.com/readykit/gd"
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

type Node2d uintptr

func (gdClass Node2d) virtual(val any, name string) any {
	switch name {
	case "_onready":
		i, ok := val.(interface{ Ready() })
		if ok {
			return i.Ready
		}
	}
	return gdClass.Parent().virtual(val, name)
}

type ExtendedNode struct {
	Node2D gd.Node2D
}

func NewExtendedNode(obj gd.Node2D) ExtendedNode {
	return ExtendedNode{Node2D: obj}
}

func (e ExtendedNode) Ready() {
	fmt.Println("Scene is ready!")
}

var gdExtendedNode = gd.Register(NewExtendedNode)
