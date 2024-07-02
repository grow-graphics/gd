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

func (h *HelloName) Ready() {
	h.Button.AsObject().Connect(h.Temporary.StringName("pressed"), h.Temporary.Callable(h.OnButtonPressed), 0)
}

func (h *HelloName) OnButtonPressed() {
	h.Text.SetText(h.Temporary.String("Hello " + h.Name.GetText(h.Temporary).String()))
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gd.Register[HelloName](godot)
}
