package main

import (
	"grow.graphics/gd"
)

type Player struct {
	gd.Class[Player, gd.Area2D] `gd:"DodgeTheCreepsPlayer"`

	Speed      gd.Int     // How fast the player will move (pixels/sec).
	ScreenSize gd.Vector2 // Size of the game window.

	Hit gd.SignalAs[func()] `gd:"hit"`

	AnimatedSprite2D gd.AnimatedSprite2D
	CollisionShape3D gd.CollisionShape3D
}

func (p *Player) AsNode() gd.Node { return p.Super().AsNode() }

var PlayerControls struct {
	MoveRight,
	MoveLeft,
	MoveDown,
	MoveUp gd.StringName
}

var PlayerAnimations struct {
	Walk,
	Up gd.StringName
}

func (p *Player) Ready() {
	p.Speed = 400
	p.ScreenSize = p.Super().AsCanvasItem().GetViewportRect().Size
	p.Super().AsCanvasItem().Hide()
}

func (p *Player) Process(delta gd.Float) {
	Input := gd.Input(p.Temporary)

	var velocity gd.Vector2
	if Input.IsActionPressed(PlayerControls.MoveRight, false) {
		velocity[gd.X] += 1
	}
	if Input.IsActionPressed(PlayerControls.MoveLeft, false) {
		velocity[gd.X] -= 1
	}
	if Input.IsActionPressed(PlayerControls.MoveDown, false) {
		velocity[gd.Y] += 1
	}
	if Input.IsActionPressed(PlayerControls.MoveUp, false) {
		velocity[gd.Y] -= 1
	}
	if velocity.Length() > 0 {
		velocity = velocity.Normalized().Mulf(gd.Float(p.Speed))
		p.AnimatedSprite2D.Play(gd.StringName{}, 1, false)
	} else {
		p.AnimatedSprite2D.Stop()
	}
	position := p.Super().AsNode2D().GetPosition()
	position = position.Add(velocity.Mulf(delta))
	position.Clamp(gd.Vector2{}, p.ScreenSize)
	p.Super().AsNode2D().SetPosition(position)
	if velocity[gd.X] != 0 {
		p.AnimatedSprite2D.SetAnimation(PlayerAnimations.Walk)
		p.AnimatedSprite2D.SetFlipV(false)
		p.AnimatedSprite2D.SetFlipH(velocity[gd.X] < 0)
	} else if velocity[gd.Y] != 0 {
		p.AnimatedSprite2D.SetAnimation(PlayerAnimations.Up)
		p.AnimatedSprite2D.SetFlipV(velocity[gd.Y] > 0)
	}
}

func (p *Player) Start(pos gd.Vector2) {
	p.Super().AsNode2D().SetPosition(pos)
	p.Super().AsCanvasItem().Show()
	p.CollisionShape3D.SetDisabled(false)
}

func (p *Player) OnPlayerBodyEntered(body gd.Node) {
	p.Super().AsCanvasItem().Hide()
	p.CollisionShape3D.SetDisabled(true)
	p.Hit.Emit()
	p.CollisionShape3D.AsObject().SetDeferred(p.Temporary.StringName("disabled"), p.Temporary.Variant(true))
}
