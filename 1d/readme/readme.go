package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
)

type HelloWorld struct {
	classdb.Extension[HelloWorld, SceneTree.Instance]
}

// Initialize implements the Godot MainLoop _initialize interface (virtual function).
func (h *HelloWorld) Initialize() {
	fmt.Println("Hello World from Go!")
}

func main() {
	startup.MainLoop[HelloWorld]()
}
