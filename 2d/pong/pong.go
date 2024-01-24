package main

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

const DefaultBallSpeed = 200

type PongBall struct {
	gd.Class[PongBall, gd.Area2D] `gd:"PongBall"`

	Direction gd.Vector2

	speed           gd.Float
	initialPosition gd.Vector2

	engine gd.Engine
}

func (b *PongBall) Ready(gd.Context) {
	b.speed = DefaultBallSpeed
	b.Direction = gd.Vector2{-1, 0}
	b.initialPosition = b.Super().AsNode2D().GetPosition()
}

func (b *PongBall) Process(_ gd.Context, delta gd.Float) {
	if b.engine.IsEditorHint() {
		return
	}

	node2d := b.Super().AsNode2D()

	b.speed += delta * 2
	node2d.SetPosition(node2d.GetPosition().Add(b.Direction.ScaledBy(b.speed * delta)))
}

func (b *PongBall) Reset() {
	node2d := b.Super().AsNode2D()
	b.Direction = gd.Vector2{-1, 0}
	node2d.SetPosition(b.initialPosition)
	b.speed = DefaultBallSpeed
}

type PongCeilingFloor struct {
	gd.Class[PongCeilingFloor, gd.Area2D] `gd:"PongCeilingFloor"`

	BounceDirection gd.Int
}

func (cf *PongCeilingFloor) OnAreaEntered(godot gd.Context, area gd.Area2D) {
	if ball, ok := gd.As[*PongBall](godot, area); ok {
		ball.Direction = (ball.Direction.Add(gd.Vector2{0, float32(cf.BounceDirection)})).Normalized()
	}
}

const PaddleMoveSpeed = 200

type PongPaddle struct {
	gd.Class[PongPaddle, gd.Area2D] `gd:"PongPaddle"`

	input  gd.Input
	engine gd.Engine

	BallDirection gd.Float
	up, down      gd.StringName

	screenSizeY float32
}

func (p *PongPaddle) Ready(godot gd.Context) {
	p.screenSizeY = p.Super().AsCanvasItem().GetViewportRect().Size[1]
	var n = p.Super().AsNode().GetName(godot).String()
	p.up = p.Pin().StringName(n + "_move_up")
	p.down = p.Pin().StringName(n + "_move_down")
}

func (p *PongPaddle) Process(godot gd.Context, delta gd.Float) {
	if p.engine.IsEditorHint() {
		return
	}
	node2d := p.Super().AsNode2D()
	var input = p.input.GetActionStrength(p.down, false) - p.input.GetActionStrength(p.up, false)
	var position = node2d.GetPosition()
	position[1] = gd.Clamp(position[1]+float32(input*PaddleMoveSpeed*delta), 16, p.screenSizeY-16)
	node2d.SetPosition(position)

}

func (p *PongPaddle) OnAreaEntered(godot gd.Context, area gd.Area2D) {
	if ball, ok := gd.As[*PongBall](godot, area); ok {
		ball.Direction = (gd.Vector2{float32(p.BallDirection), float32(godot.Randf()*2 - 1)}).Normalized()
	}
}

type PongWall struct {
	gd.Class[PongWall, gd.Area2D] `gd:"PongWall"`
}

func (w *PongWall) Hello(gd.Context) {}

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
