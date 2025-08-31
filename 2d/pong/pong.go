package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"

	_ "graphics.gd/startup"
)

const DefaultBallSpeed = 500

type PongBall struct {
	Area2D.Extension[PongBall] `gd:"PongBall"`

	Direction Vector2.XY

	speed           Float.X
	initialPosition Vector2.XY
}

func (b *PongBall) Ready() {
	b.speed = DefaultBallSpeed
	b.Direction = Vector2.Left
	b.initialPosition = b.AsNode2D().Position()
}

func (b *PongBall) Process(delta Float.X) {
	node2d := b.AsNode2D()
	b.speed += delta * 2
	node2d.SetPosition(Vector2.Add(node2d.Position(), Vector2.MulX(b.Direction, b.speed*delta)))
}

func (b *PongBall) Reset() {
	node2d := b.AsNode2D()
	b.Direction = Vector2.Left
	node2d.SetPosition(b.initialPosition)
	b.speed = DefaultBallSpeed
}

type PongCeilingFloor struct {
	Area2D.Extension[PongCeilingFloor] `gd:"PongCeilingFloor"`

	BounceDirection int
}

func (cf *PongCeilingFloor) Ready() {
	cf.AsArea2D().OnAreaEntered(cf.OnAreaEntered)
}

func (cf *PongCeilingFloor) OnAreaEntered(area Area2D.Instance) {
	if ball, ok := Object.As[*PongBall](area); ok {
		ball.Direction = Vector2.Normalized(Vector2.Add(ball.Direction, Vector2.XY{0, Float.X(cf.BounceDirection)}))
	}
}

const PaddleMoveSpeed = 200

// PongPaddle is controlled by the players.
type PongPaddle struct {
	Area2D.Extension[PongPaddle] `gd:"PongPaddle"`

	BallDirection Float.X // The direction the ball will take when it hits the paddle.
	up, down      string

	screenSizeY Float.X
}

func (p *PongPaddle) Ready() {
	p.screenSizeY = p.AsCanvasItem().GetViewportRect().Size.Y
	var n = p.AsNode().Name()
	p.up = n + "_move_up"
	p.down = n + "_move_down"
	p.AsArea2D().OnAreaEntered(p.OnAreaEntered)
}

func (p *PongPaddle) Process(delta Float.X) {
	node2d := p.AsNode2D()
	var input = Input.GetActionStrength(p.down, false) - Input.GetActionStrength(p.up, false)
	var position = node2d.Position()
	position.Y = Float.Clamp(position.Y+input*PaddleMoveSpeed*delta, 16, p.screenSizeY-16)
	node2d.SetPosition(position)
}

// OnAreaEntered should be hooked up to the "area_entered" signal of the PongPaddle node.
func (p *PongPaddle) OnAreaEntered(area Area2D.Instance) {
	if ball, ok := Object.As[*PongBall](area); ok {
		ball.Direction = Vector2.Normalized(Vector2.New(p.BallDirection, Float.RandomBetween(-1, 1)))
	}
}

type PongWall struct {
	Area2D.Extension[PongWall] `gd:"PongWall"`
}

func (w *PongWall) Ready() {
	w.AsArea2D().OnAreaEntered(w.OnAreaEntered)
}

func (w *PongWall) OnAreaEntered(area Area2D.Instance) {
	if ball, ok := Object.As[*PongBall](area); ok {
		ball.Reset()
	}
}

func main() {
	classdb.Register[PongBall]()
	classdb.Register[PongCeilingFloor]()
	classdb.Register[PongPaddle]()
	classdb.Register[PongWall]()
	startup.Scene()
}
