package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/CanvasLayer"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/SceneTreeTimer"
	"graphics.gd/classdb/Timer"
)

type HUD struct {
	classdb.Extension[HUD, CanvasLayer.Instance] `gd:"DodgeTheCreepsHUD"`

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
