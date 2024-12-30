package main

import (
	"fmt"

	"graphics.gd/defined"
	"graphics.gd/objects/Button"
	"graphics.gd/objects/CanvasLayer"
	"graphics.gd/objects/Label"
	"graphics.gd/objects/Node"
	"graphics.gd/objects/SceneTree"
	"graphics.gd/objects/SceneTreeTimer"
	"graphics.gd/objects/Timer"
)

type HUD struct {
	defined.Object[HUD, CanvasLayer.Instance] `gd:"DodgeTheCreepsHUD"`

	Message      Label.Instance
	MessageTimer Timer.Instance
	StartButton  Button.Instance
	ScoreLabel   Label.Instance

	StartGame chan<- struct{} `gd:"start_game"`
}

func (h *HUD) AsNode() Node.Instance { return h.Super().AsNode() }

func (h *HUD) ShowMessage(text string) {
	h.Message.SetText(text)
	h.Message.AsCanvasItem().Show()
	h.MessageTimer.Start()
}

func (h *HUD) ShowGameOver() {
	h.ShowMessage("Game Over")
	h.MessageTimer.OnTimeout(h.ShowTitle)
}

func (h *HUD) ShowTitle() {
	h.Message.SetText("Dodge the Creeps!")
	h.Message.AsCanvasItem().Show()
	var timer SceneTreeTimer.Instance = SceneTree.Instance(h.Super().AsNode().GetTree()).CreateTimer(1)
	timer.OnTimeout(h.ShowStartButton)
}

func (h *HUD) ShowStartButton() { h.StartButton.AsCanvasItem().Show() }

func (h *HUD) UpdateScore(score int) {
	h.ScoreLabel.SetText(fmt.Sprint(score))
}

func (h *HUD) OnStartButtonPressed() {
	h.StartButton.AsCanvasItem().Hide()
	h.StartGame <- struct{}{}
}

func (h *HUD) OnMessageTimerTimeout() {
	h.Message.AsCanvasItem().Hide()
}
