package main

import (
	"fmt"

	"graphics.gd/defined"
	"graphics.gd/objects/SceneTree"
	"graphics.gd/startup"
)

type HelloWorld struct {
	defined.Object[HelloWorld, SceneTree.Instance]
}

// Initialize implements the Godot MainLoop _initialize interface (virtual function).
func (h *HelloWorld) Initialize() {
	fmt.Println("Hello World from Go!")
}

func main() {
	startup.MainLoop[HelloWorld]()
}
