package main

import (
	"graphics.gd/defined"
	"graphics.gd/objects"
	"graphics.gd/objects/Area2D"
	"graphics.gd/objects/Input"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"

	_ "graphics.gd/startup"
)

const DefaultBallSpeed = 500

type PongBall struct {
	defined.Object[PongBall, Area2D.Instance] `gd:"PongBall"`

	Direction Vector2.XY

	speed           Float.X
	initialPosition Vector2.XY
}

func (b *PongBall) Ready() {
	b.speed = DefaultBallSpeed
	b.Direction = Vector2.Left
	b.initialPosition = b.Super().AsNode2D().Position()
}

func (b *PongBall) Process(delta Float.X) {
	node2d := b.Super().AsNode2D()
	b.speed += delta * 2
	node2d.SetPosition(Vector2.Add(node2d.Position(), Vector2.MulX(b.Direction, b.speed*delta)))
}

func (b *PongBall) Reset() {
	node2d := b.Super().AsNode2D()
	b.Direction = Vector2.Left
	node2d.SetPosition(b.initialPosition)
	b.speed = DefaultBallSpeed
}

type PongCeilingFloor struct {
	defined.Object[PongCeilingFloor, Area2D.Instance] `gd:"PongCeilingFloor"`

	BounceDirection int
}

func (cf *PongCeilingFloor) OnAreaEntered(area Area2D.Instance) {
	if ball, ok := objects.As[*PongBall](area); ok {
		ball.Direction = Vector2.Normalized(Vector2.Add(ball.Direction, Vector2.XY{0, Float.X(cf.BounceDirection)}))
	}
}

const PaddleMoveSpeed = 200

type PongPaddle struct {
	defined.Object[PongPaddle, Area2D.Instance] `gd:"PongPaddle"`

	BallDirection Float.X
	up, down      string

	screenSizeY Float.X
}

func (p *PongPaddle) Ready() {
	p.screenSizeY = p.Super().AsCanvasItem().GetViewportRect().Size.Y
	var n = p.Super().AsNode().Name()
	p.up = n + "_move_up"
	p.down = n + "_move_down"
}

func (p *PongPaddle) Process(delta Float.X) {
	node2d := p.Super().AsNode2D()
	var input = Input.GetActionStrength(p.down) - Input.GetActionStrength(p.up)
	var position = node2d.Position()
	position.Y = Float.Clamp(position.Y+input*PaddleMoveSpeed*delta, 16, p.screenSizeY-16)
	node2d.SetPosition(position)
}

func (p *PongPaddle) OnAreaEntered(area Area2D.Instance) {
	if ball, ok := objects.As[*PongBall](area); ok {
		ball.Direction = Vector2.Normalized(Vector2.New(p.BallDirection, Float.RandomBetween(-1, 1)))
	}
}

type PongWall struct {
	defined.Object[PongWall, Area2D.Instance] `gd:"PongWall"`
}

func (w *PongWall) OnAreaEntered(area Area2D.Instance) {
	if ball, ok := objects.As[*PongBall](area); ok {
		ball.Reset()
	}
}

func main() {
	defined.InEditor[PongBall]()
	defined.InEditor[PongCeilingFloor]()
	defined.InEditor[PongPaddle]()
	defined.InEditor[PongWall]()
}
