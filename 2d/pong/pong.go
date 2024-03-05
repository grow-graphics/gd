package main

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

const DefaultBallSpeed = 500

type PongBall struct {
	gd.Class[PongBall, gd.Area2D] `gd:"PongBall"`

	Direction gd.Vector2

	speed           gd.Float
	initialPosition gd.Vector2
}

func (b *PongBall) Ready(gd.Context) {
	b.speed = DefaultBallSpeed
	b.Direction = gd.Const(gd.Vector2.LEFT)
	b.initialPosition = b.Super().AsNode2D().GetPosition()
}

func (b *PongBall) Process(_ gd.Context, delta gd.Float) {
	node2d := b.Super().AsNode2D()
	b.speed += delta * 2
	node2d.SetPosition(node2d.GetPosition().Add(b.Direction.Mulf(b.speed * delta)))
}

func (b *PongBall) Reset() {
	node2d := b.Super().AsNode2D()
	b.Direction = gd.Const(gd.Vector2.LEFT)
	node2d.SetPosition(b.initialPosition)
	b.speed = DefaultBallSpeed
}

type PongCeilingFloor struct {
	gd.Class[PongCeilingFloor, gd.Area2D] `gd:"PongCeilingFloor"`

	BounceDirection gd.Int
}

func (cf *PongCeilingFloor) OnAreaEntered(godot gd.Context, area gd.Area2D) {
	if ball, ok := gd.As[*PongBall](godot, area); ok {
		ball.Direction = gd.Vector2.Add(ball.Direction, gd.NewVector2(0, gd.Float(cf.BounceDirection))).Normalized()
	}
}

const PaddleMoveSpeed = 200

type PongPaddle struct {
	gd.Class[PongPaddle, gd.Area2D] `gd:"PongPaddle"`

	BallDirection gd.Float
	up, down      gd.StringName

	screenSizeY gd.Float
}

func (p *PongPaddle) Ready(godot gd.Context) {
	p.screenSizeY = p.Super().AsCanvasItem().GetViewportRect().Size.Y()
	var n = p.Super().AsNode().GetName(godot).String()
	p.up = p.Pin().StringName(n + "_move_up")
	p.down = p.Pin().StringName(n + "_move_down")
}

func (p *PongPaddle) Process(godot gd.Context, delta gd.Float) {
	node2d := p.Super().AsNode2D()
	var input = gd.Input(godot).GetActionStrength(p.down, false) - gd.Input(godot).GetActionStrength(p.up, false)
	var position = node2d.GetPosition()
	position.SetY(gd.Float(gd.Clamp(gd.Float(position[1])+input*PaddleMoveSpeed*delta, 16, gd.Float(p.screenSizeY-16))))
	node2d.SetPosition(position)

}

func (p *PongPaddle) OnAreaEntered(godot gd.Context, area gd.Area2D) {
	if ball, ok := gd.As[*PongBall](godot, area); ok {
		ball.Direction = (gd.NewVector2(p.BallDirection, godot.Randf()*2-1)).Normalized()
	}
}

type PongWall struct {
	gd.Class[PongWall, gd.Area2D] `gd:"PongWall"`
}

func (w *PongWall) OnAreaEntered(godot gd.Context, area gd.Area2D) {
	if ball, ok := gd.As[*PongBall](godot, area); ok {
		ball.Reset()
	}
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		panic("could not link to godot")
	}
	gd.Register[PongBall](godot)
	gd.Register[PongCeilingFloor](godot)
	gd.Register[PongPaddle](godot)
	gd.Register[PongWall](godot)
}
