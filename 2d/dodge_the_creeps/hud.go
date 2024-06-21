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

func (h *HUD) ShowMessage(godot gd.Context, text gd.String) {
	h.Message.SetText(text)
	h.Message.AsCanvasItem().Show()
	h.MessageTimer.Start(0)
}

func (h *HUD) ShowGameOver(godot gd.Context) {
	h.ShowMessage(godot, godot.String("Game Over"))
	h.MessageTimer.AsObject().Connect(godot.StringName("timeout"), godot.Callable(h.showTitle), 0)
}

func (h *HUD) showTitle(godot gd.Context) {
	h.Message.SetText(godot.String("Dodge the Creeps!"))
	h.Message.AsCanvasItem().Show()
	timer := h.Super().AsNode().GetTree(godot).CreateTimer(godot, 1, false, false, false)
	timer.AsObject().Connect(godot.StringName("timeout"), godot.Callable(h.showStartButton), 0)
}

func (h *HUD) showStartButton(godot gd.Context) { h.StartButton.AsCanvasItem().Show() }

func (h *HUD) UpdateScore(godot gd.Context, score gd.Int) {
	h.ScoreLabel.SetText(godot.String(fmt.Sprint(score)))
}

func (h *HUD) OnStartButtonPressed(gd.Context) {
	h.StartButton.AsCanvasItem().Hide()
	h.StartGame.Emit()
}

func (h *HUD) OnMessageTimerTimeout(godot gd.Context) {
	h.Message.AsCanvasItem().Hide()
}
