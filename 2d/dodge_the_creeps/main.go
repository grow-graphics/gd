package main

import (
	"fmt"
	"math/rand"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type Main struct {
	gd.Class[Main, gd.Node] `gd:"DodgeTheCreeps"`

	MobScene gd.PackedScene

	MobTimer      gd.Timer
	ScoreTimer    gd.Timer
	StartTimer    gd.Timer
	Player        *Player
	StartPosition gd.Marker2D

	Music      gd.AudioStreamPlayer
	DeathSound gd.AudioStreamPlayer

	HUD *HUD

	MobPath struct {
		gd.Path2D

		MobSpawnLocation gd.PathFollow2D
	}

	score gd.Int
}

func (m *Main) GameOver(godot gd.Context) {
	m.ScoreTimer.Stop()
	m.MobTimer.Stop()
	m.HUD.ShowGameOver(godot)
	m.Music.Stop()
	m.DeathSound.Play(0)
}

func (m *Main) NewGame(godot gd.Context) {
	m.score = 0
	m.Player.Start(godot, m.StartPosition.AsNode2D().GetPosition())
	m.StartTimer.Start(0)

	m.HUD.UpdateScore(godot, m.score)
	m.HUD.ShowMessage(godot, godot.String("Get Ready!"))

	//m.Super().GetTree(godot).CallGroup(godot.StringName("mobs"), godot.StringName("queue_free"))
	m.Music.Play(0)
}

func (m *Main) OnScoreTimerTimeout(godot gd.Context) {
	m.score++
	m.HUD.UpdateScore(godot, m.score)
}

func (m *Main) OnStartTimerTimeout(godot gd.Context) {
	m.MobTimer.Start(0)
	m.ScoreTimer.Start(0)
}

func (m *Main) OnMobTimerTimeout(godot gd.Context) {
	// Create a new instance of the Mob scene.
	mob, ok := gd.As[gd.RigidBody2D](godot, m.MobScene.Instantiate(godot, 0))
	if !ok {
		fmt.Println("failed to cast!")
		return
	}

	// Choose a random location on Path2D.
	m.MobPath.MobSpawnLocation.SetProgressRatio(rand.Float64())

	// Set the mob's direction perpendicular to the path direction.
	direction := m.MobPath.MobSpawnLocation.AsNode2D().GetRotation() + gd.Pi/2

	// Set the mob's position to a random location.
	mob.AsNode2D().SetPosition(m.MobPath.MobSpawnLocation.AsNode2D().GetPosition())

	// Add some randomness to the direction.
	direction += godot.RandfRange(-gd.Pi/4, gd.Pi/4)
	mob.AsNode2D().SetRotation(direction)

	// Choose the velocity.
	var velocity = gd.NewVector2(godot.RandfRange(150, 250), 0)
	mob.SetLinearVelocity(velocity.Rotated(gd.Radians(direction)))

	m.Super().AddChild(mob.AsNode(), true, 0)
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}

	PlayerControls.MoveRight = godot.StringName("move_right")
	PlayerControls.MoveLeft = godot.StringName("move_left")
	PlayerControls.MoveDown = godot.StringName("move_down")
	PlayerControls.MoveUp = godot.StringName("move_up")
	PlayerAnimations.Walk = godot.StringName("right")
	PlayerAnimations.Up = godot.StringName("up")

	gd.Register[HUD](godot)
	gd.Register[Player](godot)
	gd.Register[Mob](godot)
	gd.Register[Main](godot)
}
