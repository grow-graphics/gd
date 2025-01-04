package main

import (
	"fmt"
	"math/rand"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/Marker2D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Path2D"
	"graphics.gd/classdb/PathFollow2D"
	"graphics.gd/classdb/RigidBody2D"
	"graphics.gd/classdb/Timer"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

type Main struct {
	classdb.Extension[Main, Node.Instance] `gd:"DodgeTheCreeps"`

	MobTimer      Timer.Instance
	ScoreTimer    Timer.Instance
	StartTimer    Timer.Instance
	Player        *Player
	StartPosition Marker2D.Instance

	Music      AudioStreamPlayer.Instance
	DeathSound AudioStreamPlayer.Instance

	HUD *HUD

	MobPath struct {
		Path2D.Instance

		MobSpawnLocation PathFollow2D.Instance
	}

	MobScene PackedScene.Instance

	score int
}

func (m *Main) GameOver() {
	m.ScoreTimer.Stop()
	m.MobTimer.Stop()
	m.HUD.ShowGameOver()
	m.Music.Stop()
	m.DeathSound.Play()
}

func (m *Main) NewGame() {
	m.score = 0
	m.Player.Start(m.StartPosition.AsNode2D().Position())
	m.StartTimer.Start()

	m.HUD.UpdateScore(m.score)
	m.HUD.ShowMessage("Get Ready!")

	//m.Super().GetTree(godot).CallGroup(godot.StringName("mobs"), godot.StringName("queue_free"))
	m.Music.Play()
}

func (m *Main) OnScoreTimerTimeout() {
	m.score++
	m.HUD.UpdateScore(m.score)
}

func (m *Main) OnStartTimerTimeout() {
	m.MobTimer.Start()
	m.ScoreTimer.Start()
}

func (m *Main) OnMobTimerTimeout() {
	// Create a new instance of the Mob scene.
	mob, ok := classdb.As[RigidBody2D.Instance](Node.Instance(m.MobScene.Instantiate()))
	if !ok {
		fmt.Println("failed to cast!")
		return
	}

	// Choose a random location on Path2D.
	m.MobPath.MobSpawnLocation.SetProgressRatio(Float.X(rand.Float64()))

	// Set the mob's direction perpendicular to the path direction.
	direction := Angle.Radians(m.MobPath.MobSpawnLocation.AsNode2D().Rotation()) + Angle.Pi/2

	// Set the mob's position to a random location.
	mob.AsNode2D().SetPosition(m.MobPath.MobSpawnLocation.AsNode2D().Position())

	// Add some randomness to the direction.
	direction += Angle.Radians(Float.RandomBetween(-Float.X(Angle.Pi/4), Float.X(Angle.Pi/4)))
	mob.AsNode2D().SetRotation(Float.X(direction))

	// Choose the velocity.
	var velocity = Vector2.New(Float.RandomBetween(150, 250), 0)
	mob.SetLinearVelocity(Vector2.Rotated(velocity, direction))

	m.Super().AddChild(mob.AsNode())
}

func main() {
	PlayerControls.MoveRight = ("move_right")
	PlayerControls.MoveLeft = ("move_left")
	PlayerControls.MoveDown = ("move_down")
	PlayerControls.MoveUp = ("move_up")
	PlayerAnimations.Walk = ("right")
	PlayerAnimations.Up = ("up")

	classdb.Register[HUD]()
	classdb.Register[Player]()
	classdb.Register[Mob]()
	classdb.Register[Main]()
	startup.Engine()
}
