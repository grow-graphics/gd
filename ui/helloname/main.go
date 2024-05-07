package main

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type HelloName struct {
	gd.Class[HelloName, gd.Node2D]

	Name gd.TextEdit
	Text gd.Label

	Button gd.Button
}

func (h *HelloName) Ready(godot gd.Context) {
	h.Button.AsObject().Connect(godot.StringName("pressed"), godot.Callable(h.OnButtonPressed), 0)
}

func (h *HelloName) OnButtonPressed(godot gd.Context) {
	h.Text.SetText(godot.String("Hello " + h.Name.GetText(godot).String()))
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gd.Register[HelloName](godot)
}
