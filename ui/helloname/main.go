package main

import (
	"graphics.gd/classdb"
	"graphics.gd/startup"

	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/TextEdit"
)

type HelloName struct {
	classdb.Extension[HelloName, Node2D.Instance]

	Name   TextEdit.Instance
	Text   Label.Instance
	Button Button.Instance
}

func (h *HelloName) Ready() {
	h.Button.AsBaseButton().OnPressed(h.OnButtonPressed)
}

func (h *HelloName) OnButtonPressed() {
	h.Text.SetText("Hello " + h.Name.Text())
}

func main() {
	classdb.Register[HelloName]()
	startup.Engine()
}
