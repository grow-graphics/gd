package main

import (
	"fmt"

	"grow.graphics/gd"
)

type HUD struct {
	gd.Class[HUD, gd.CanvasLayer] `gd:"DodgeTheCreepsHUD"`

	Message      gd.Label
	MessageTimer gd.Timer
	StartButton  gd.Button
	ScoreLabel   gd.Label

	StartGame gd.SignalAs[func()] `gd:"start_game"`
}

func (h *HUD) AsNode() gd.Node { return h.Super().AsNode() }

func (h *HUD) ShowMessage(text gd.String) {
	h.Message.SetText(text)
	h.Message.AsCanvasItem().Show()
	h.MessageTimer.Start(0)
}

func (h *HUD) ShowGameOver() {
	h.ShowMessage(h.Temporary.String("Game Over"))
	h.MessageTimer.AsObject().Connect(h.Temporary.StringName("timeout"), h.Temporary.Callable(h.showTitle), 0)
}

func (h *HUD) showTitle() {
	h.Message.SetText(h.Temporary.String("Dodge the Creeps!"))
	h.Message.AsCanvasItem().Show()
	timer := h.Super().AsNode().GetTree(h.Temporary).CreateTimer(h.Temporary, 1, false, false, false)
	timer.AsObject().Connect(h.Temporary.StringName("timeout"), h.Temporary.Callable(h.showStartButton), 0)
}

func (h *HUD) showStartButton() { h.StartButton.AsCanvasItem().Show() }

func (h *HUD) UpdateScore(score gd.Int) {
	h.ScoreLabel.SetText(h.Temporary.String(fmt.Sprint(score)))
}

func (h *HUD) OnStartButtonPressed() {
	h.StartButton.AsCanvasItem().Hide()
	h.StartGame.Emit()
}

func (h *HUD) OnMessageTimerTimeout() {
	h.Message.AsCanvasItem().Hide()
}
