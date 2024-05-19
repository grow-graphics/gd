package main

import (
	"fmt"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type HelloWorld struct {
	gd.Class[HelloWorld, gd.SceneTree]
}

// Initialize implements the Godot MainLoop _initialize interface (virtual function).
func (h *HelloWorld) Initialize(godot gd.Context) {
	fmt.Println("Hello World from Go!")
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gd.Register[HelloWorld](godot)
}
