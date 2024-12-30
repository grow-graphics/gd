package main

import (
	"graphics.gd/defined"

	"graphics.gd/objects/Button"
	"graphics.gd/objects/Label"
	"graphics.gd/objects/Node2D"
	"graphics.gd/objects/TextEdit"

	_ "graphics.gd/startup"
)

type HelloName struct {
	defined.Object[HelloName, Node2D.Instance]

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
	defined.InEditor[HelloName]()
}
