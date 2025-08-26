package main

import (
	"graphics.gd/classdb/AnimatedSprite2D"
	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/CollisionShape3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/Vector2"
)

type Player struct {
	Area2D.Extension[Player] `gd:"DodgeTheCreepsPlayer"`

	Speed      int        // How fast the player will move (pixels/sec).
	ScreenSize Vector2.XY // Size of the game window.

	Hit Signal.Void `gd:"hit"`

	AnimatedSprite2D AnimatedSprite2D.Instance
	CollisionShape3D CollisionShape3D.Instance
}

var PlayerControls struct {
	MoveRight,
	MoveLeft,
	MoveDown,
	MoveUp string
}

var PlayerAnimations struct {
	Walk,
	Up string
}

func (p *Player) Ready() {
	p.Speed = 400
	p.ScreenSize = p.AsCanvasItem().GetViewportRect().Size
	p.AsCanvasItem().Hide()
}

func (p *Player) Process(delta Float.X) {
	var velocity Vector2.XY
	if Input.IsActionPressed(PlayerControls.MoveRight, false) {
		velocity.X += 1
	}
	if Input.IsActionPressed(PlayerControls.MoveLeft, false) {
		velocity.X -= 1
	}
	if Input.IsActionPressed(PlayerControls.MoveDown, false) {
		velocity.Y += 1
	}
	if Input.IsActionPressed(PlayerControls.MoveUp, false) {
		velocity.Y -= 1
	}
	if Vector2.Length(velocity) > 0 {
		velocity = Vector2.MulX(Vector2.Normalized(velocity), p.Speed)
		p.AnimatedSprite2D.Play()
	} else {
		p.AnimatedSprite2D.Stop()
	}
	position := p.AsNode2D().Position()
	position = Vector2.Add(position, Vector2.MulX(velocity, delta))
	Vector2.Clamp(position, Vector2.Zero, p.ScreenSize)
	p.AsNode2D().SetPosition(position)
	if velocity.X != 0 {
		p.AnimatedSprite2D.SetAnimation(PlayerAnimations.Walk)
		p.AnimatedSprite2D.SetFlipV(false)
		p.AnimatedSprite2D.SetFlipH(velocity.X < 0)
	} else if velocity.Y != 0 {
		p.AnimatedSprite2D.SetAnimation(PlayerAnimations.Up)
		p.AnimatedSprite2D.SetFlipV(velocity.Y > 0)
	}
}

func (p *Player) Start(pos Vector2.XY) {
	p.AsNode2D().SetPosition(pos)
	p.AsCanvasItem().Show()
	p.CollisionShape3D.SetDisabled(false)
}

func (p *Player) OnPlayerBodyEntered(body Node.Instance) {
	p.AsCanvasItem().Hide()
	p.CollisionShape3D.SetDisabled(true)
	p.Hit.Emit()
	Callable.Defer(Callable.New(func() {
		p.CollisionShape3D.SetDisabled(true)
	}))
}
