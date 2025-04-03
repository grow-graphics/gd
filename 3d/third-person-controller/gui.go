package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/GridContainer"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/variant/Enum"
)

type InstructionType Enum.Int[struct {
	Keyboard InstructionType
	Joypad   InstructionType
}]

var InstructionTypes = Enum.Values[InstructionType]()

type DemoPage struct {
	classdb.Extension[DemoPage, Node.Instance]

	DemoPageRoot          Control.Instance       `gd:"CanvasLayer/DemoPageRoot"`
	ResumeButton          Button.Instance        `gd:"CanvasLayer/DemoPageRoot/Content/MarginContainer/Buttons/Resume"`
	ExitButton            Button.Instance        `gd:"CanvasLayer/DemoPageRoot/Content/MarginContainer/Buttons/Exit"`
	KeyboardButton        Button.Instance        `gd:"%KeyboardButton"`
	JoypadButton          Button.Instance        `gd:"%JoypadButton"`
	GridContainerKeyboard GridContainer.Instance `gd:"%GridContainerKeyboard"`
	GridContainerJoypad   GridContainer.Instance `gd:"%GridContainerJoypad"`

	demoMouseMode Input.MouseModeValue
}

func (page *DemoPage) Ready() {
	var tree SceneTree.Instance = page.Super().AsNode().GetTree()
	tree.SetPaused(true)
	page.demoMouseMode = Input.MouseMode()
	//page.ResumeButton.AsBaseButton().OnPressed(page.resume_demo)
	page.ExitButton.AsBaseButton().OnPressed(tree.Quit)
}
