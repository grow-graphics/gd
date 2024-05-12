package main

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
	"runtime.link/mmm"
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
	label := gd.Create(h.List.AsPointer().Pin(), new(gd.Label))
	text := h.Task.GetText(godot)
	label.SetText(text)
	h.List.AsNode().AddChild(label.AsNode(), false, 0)
	mmm.End(label.AsPointer())
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gd.Register[TodoList](godot)
}
