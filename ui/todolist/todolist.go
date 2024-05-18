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

func (h *TodoList) Ready(godot gd.Context) {
	h.Button.AsObject().Connect(godot.StringName("pressed"), godot.Callable(h.OnButtonPressed), 0)
}

func (h *TodoList) OnButtonPressed(godot gd.Context) {
	label := gd.Create(godot, new(gd.Label))
	label.SetText(h.Task.GetText(godot))
	h.List.AsNode().AddChild(label.AsNode(), false, 0)
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gd.Register[TodoList](godot)
}
