package main

import (
	"graphics.gd/defined"
	"graphics.gd/objects/Button"
	"graphics.gd/objects/Control"
	"graphics.gd/objects/Label"
	"graphics.gd/objects/TextEdit"
	"graphics.gd/objects/VBoxContainer"

	_ "graphics.gd/startup"
)

type TodoList struct {
	defined.Object[TodoList, Control.Instance]

	Task   TextEdit.Instance
	Button Button.Instance

	List VBoxContainer.Instance
}

func (h *TodoList) Ready() {
	h.Button.AsBaseButton().OnPressed(h.OnButtonPressed)
}

func (h *TodoList) OnButtonPressed() {
	label := Label.New()
	label.SetText(h.Task.Text())
	h.List.AsNode().AddChild(label.AsNode())
}

func main() {
	defined.InEditor[TodoList]()
}
