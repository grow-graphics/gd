package main

import (
	"fmt"

	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/CanvasLayer"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/SceneTreeTimer"
	"graphics.gd/classdb/Timer"
	"graphics.gd/variant/Signal"
)

type HUD struct {
	CanvasLayer.Extension[HUD] `gd:"DodgeTheCreepsHUD"`

	Message      Label.Instance
	MessageTimer Timer.Instance
	StartButton  Button.Instance
	ScoreLabel   Label.Instance

	StartGame Signal.Void `gd:"start_game"`
}

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
	var timer SceneTreeTimer.Instance = SceneTree.Get(h.AsNode()).CreateTimer(1)
	timer.OnTimeout(h.ShowStartButton)
}

func (h *HUD) ShowStartButton() { h.StartButton.AsCanvasItem().Show() }

func (h *HUD) UpdateScore(score int) {
	h.ScoreLabel.SetText(fmt.Sprint(score))
}

func (h *HUD) OnStartButtonPressed() {
	h.StartButton.AsCanvasItem().Hide()
	h.StartGame.Emit()
}

func (h *HUD) OnMessageTimerTimeout() {
	h.Message.AsCanvasItem().Hide()
}
