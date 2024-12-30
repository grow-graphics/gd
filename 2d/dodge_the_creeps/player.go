package main

import (
	gd "graphics.gd"
	"graphics.gd/defined"
	"graphics.gd/objects/AnimatedSprite2D"
	"graphics.gd/objects/Area2D"
	"graphics.gd/objects/CollisionShape3D"
	"graphics.gd/objects/Input"
	"graphics.gd/objects/Node"
	"graphics.gd/variant"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/StringName"
	"graphics.gd/variant/Vector2"
)

type Player struct {
	defined.Object[Player, Area2D.Instance] `gd:"DodgeTheCreepsPlayer"`

	Speed      gd.Int     // How fast the player will move (pixels/sec).
	ScreenSize gd.Vector2 // Size of the game window.

	Hit chan<- struct{} `gd:"hit"`

	AnimatedSprite2D AnimatedSprite2D.Instance
	CollisionShape3D CollisionShape3D.Instance
}

func (p *Player) AsNode() Node.Instance { return p.Super().AsNode() }

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
	p.ScreenSize = p.Super().AsCanvasItem().GetViewportRect().Size
	p.Super().AsCanvasItem().Hide()
}

func (p *Player) Process(delta Float.X) {
	var velocity gd.Vector2
	if Input.IsActionPressed(PlayerControls.MoveRight) {
		velocity.X += 1
	}
	if Input.IsActionPressed(PlayerControls.MoveLeft) {
		velocity.X -= 1
	}
	if Input.IsActionPressed(PlayerControls.MoveDown) {
		velocity.Y += 1
	}
	if Input.IsActionPressed(PlayerControls.MoveUp) {
		velocity.Y -= 1
	}
	if Vector2.Length(velocity) > 0 {
		velocity = Vector2.MulX(Vector2.Normalized(velocity), p.Speed)
		p.AnimatedSprite2D.Play()
	} else {
		p.AnimatedSprite2D.Stop()
	}
	position := p.Super().AsNode2D().Position()
	position = Vector2.Add(position, Vector2.MulX(velocity, delta))
	Vector2.Clamp(position, Vector2.Zero, p.ScreenSize)
	p.Super().AsNode2D().SetPosition(position)
	if velocity.X != 0 {
		p.AnimatedSprite2D.SetAnimation(PlayerAnimations.Walk)
		p.AnimatedSprite2D.SetFlipV(false)
		p.AnimatedSprite2D.SetFlipH(velocity.X < 0)
	} else if velocity.Y != 0 {
		p.AnimatedSprite2D.SetAnimation(PlayerAnimations.Up)
		p.AnimatedSprite2D.SetFlipV(velocity.Y > 0)
	}
}

func (p *Player) Start(pos gd.Vector2) {
	p.Super().AsNode2D().SetPosition(pos)
	p.Super().AsCanvasItem().Show()
	p.CollisionShape3D.SetDisabled(false)
}

func (p *Player) OnPlayerBodyEntered(body Node.Instance) {
	p.Super().AsCanvasItem().Hide()
	p.CollisionShape3D.SetDisabled(true)
	p.Hit <- struct{}{}
	p.CollisionShape3D.AsObject().SetDeferred(StringName.New("disabled"), variant.New(true))
}
