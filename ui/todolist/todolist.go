package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/TextEdit"
	"graphics.gd/classdb/VBoxContainer"
	"graphics.gd/startup"
)

type TodoList struct {
	classdb.Extension[TodoList, Control.Instance]

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
	classdb.Register[TodoList]()
	startup.Engine()
}
