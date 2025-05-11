package main

import (
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/GridContainer"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/TextureButton"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Float"
)

type InstructionType Enum.Int[struct {
	Keyboard InstructionType
	Joypad   InstructionType
}]

var InstructionTypes = Enum.Values[InstructionType]()

type DemoPage struct {
	Node.Extension[DemoPage] `gd:"DemoPage"`

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
	var tree = SceneTree.Get(page.AsNode())
	tree.SetPaused(true)
	page.demoMouseMode = Input.MouseMode()
	Input.SetMouseMode(Input.MouseModeVisible)
	page.ResumeButton.AsBaseButton().OnPressed(page.resume_demo)
	page.ExitButton.AsBaseButton().OnPressed(func() {
		SceneTree.Get(page.AsNode()).Quit()
	})
	page.KeyboardButton.AsBaseButton().OnPressed(func() {
		page.change_instruction(InstructionTypes.Keyboard)
	})
	page.JoypadButton.AsBaseButton().OnPressed(func() {
		page.change_instruction(InstructionTypes.Joypad)
	})
	if len(Input.GetConnectedJoypads()) > 0 {
		page.change_instruction(InstructionTypes.Joypad)
	} else {
		page.change_instruction(InstructionTypes.Keyboard)
	}
}

func (page *DemoPage) resume_demo() {
	SceneTree.Get(page.AsNode()).SetPaused(false)
	var tween = page.AsNode().CreateTween()
	tween.TweenProperty(page.DemoPageRoot.AsObject(), "modulate", Color.Transparent, 0.3)
	tween.TweenCallback(page.DemoPageRoot.AsCanvasItem().Hide)
	Input.SetMouseMode(page.demoMouseMode)
}

func (page *DemoPage) change_instruction(itype InstructionType) {
	withAlpha := func(alpha Float.X, c Color.RGBA) Color.RGBA {
		c.A = alpha
		return c
	}
	switch itype {
	case InstructionTypes.Keyboard:
		page.KeyboardButton.AsCanvasItem().SetModulate(withAlpha(1, page.KeyboardButton.AsCanvasItem().Modulate()))
		page.JoypadButton.AsCanvasItem().SetModulate(withAlpha(0.3, page.JoypadButton.AsCanvasItem().Modulate()))
		page.GridContainerKeyboard.AsCanvasItem().Show()
		page.GridContainerJoypad.AsCanvasItem().Hide()
	case InstructionTypes.Joypad:
		page.KeyboardButton.AsCanvasItem().SetModulate(withAlpha(0.3, page.KeyboardButton.AsCanvasItem().Modulate()))
		page.JoypadButton.AsCanvasItem().SetModulate(withAlpha(1, page.JoypadButton.AsCanvasItem().Modulate()))
		page.GridContainerKeyboard.AsCanvasItem().Hide()
		page.GridContainerJoypad.AsCanvasItem().Show()
	}
	page.KeyboardButton.AsControl().ReleaseFocus()
	page.JoypadButton.AsControl().ReleaseFocus()
}

func (page *DemoPage) Input(event InputEvent.Instance) {
	if event.IsActionPressed("pause") && !event.IsEcho() {
		if SceneTree.Get(page.AsNode()).Paused() {
			page.resume_demo()
		} else {
			page.pause_demo()
		}
	}
}

func (page *DemoPage) pause_demo() {
	page.demoMouseMode = Input.MouseMode()
	SceneTree.Get(page.AsNode()).SetPaused(true)
	page.DemoPageRoot.AsCanvasItem().Show()
	var tween = page.AsNode().CreateTween()
	tween.TweenProperty(page.DemoPageRoot.AsObject(), "modulate", Color.X11.White, 0.3)
	Input.SetMouseMode(Input.MouseModeVisible)
}

type DemoLinkButton struct {
	TextureButton.Extension[DemoLinkButton] `gd:"DemoLinkButton"`

	Link string
}

func (button *DemoLinkButton) Ready() {
	button.AsBaseButton().OnPressed(func() {
		OS.ShellOpen(button.Link)
	})
}
