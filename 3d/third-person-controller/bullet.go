package main

import (
	"graphics.gd/classdb/Area3D"
	"graphics.gd/classdb/AudioStreamPlayer3D"
	"graphics.gd/classdb/Curve"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

type Bullet struct {
	Node3D.Extension[Bullet] `gd:"Bullet"`

	ScaleDecay    Curve.Instance
	DistanceLimit Float.X

	Velocity Vector3.XYZ
	Shooter  Node.ID

	Area3D  Area3D.Instance `gd:"Area3D"`
	Visuals Node3D.Instance `gd:"Bullet"`

	ProjectileSound AudioStreamPlayer3D.Instance `gd:"ProjectileSound"`

	TimeAlive  Float.X
	AliveLimit Float.X
}

func NewBullet() *Bullet {
	return &Bullet{
		DistanceLimit: 5,
	}
}

func (b *Bullet) Ready() {
	b.Area3D.OnBodyEntered(func(body Node3D.Instance) {
		if body.ID() == Node3D.ID(b.Shooter) {
			return
		}
		if body.AsNode().IsInGroup("damageables") {
			var impact_point = Vector3.Sub(b.Visuals.GlobalPosition(), body.GlobalPosition())
			Object.Call(body, "damage", impact_point, b.Velocity)
		}
		body.AsNode().QueueFree()
	})
	b.AsNode3D().LookAt(Vector3.Add(b.AsNode3D().GlobalPosition(), b.Velocity))
	b.AliveLimit = b.DistanceLimit / Vector3.Length(b.Velocity)
	b.ProjectileSound.SetPitchScale(Float.RandomlyDistributed(1.0, 0.1))
	b.ProjectileSound.Play()
}

func (b *Bullet) Process(delta Float.X) {
	b.AsNode3D().SetGlobalPosition(Vector3.Add(b.AsNode3D().GlobalPosition(), Vector3.MulX(b.Velocity, delta)))
	b.TimeAlive += delta
	b.Visuals.SetScale(Vector3.MulX(Vector3.One, b.ScaleDecay.Sample(b.TimeAlive)))
	if b.TimeAlive > b.AliveLimit {
		b.AsNode().QueueFree()
	}
}
