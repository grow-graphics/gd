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
	label.SetText(h.Task.GetText(godot))
	h.List.AsNode().AddChild(label.AsNode(), false, 0)
	// we have given ownership of the label to Godot, we we need to end the
	// lifetime within Go (so that it doesn't get freed automatically).
	mmm.End(label.AsPointer())
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gd.Register[TodoList](godot)
}
