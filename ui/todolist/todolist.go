package main

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type TodoList struct {
	gd.Class[TodoList, gd.Control]

	Task   gd.TextEdit
	Button gd.Button

	List gd.VBoxContainer
}

func (h *TodoList) Ready() {
	h.Button.AsObject().Connect(h.Temporary.StringName("pressed"), h.Temporary.Callable(h.OnButtonPressed), 0)
}

func (h *TodoList) OnButtonPressed() {
	label := gd.Create(h.Temporary, new(gd.Label))
	label.SetText(h.Task.GetText(h.Temporary))
	h.List.AsNode().AddChild(label.AsNode(), false, 0)
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gd.Register[TodoList](godot)
}
