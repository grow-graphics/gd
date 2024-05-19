package main

import (
	"math/rand"

	"grow.graphics/gd"
)

type Mob struct {
	gd.Class[Mob, gd.RigidBody2D] `gd:"DodgeTheCreepsMob"`

	AnimatedSprite2D gd.AnimatedSprite2D
}

func (m *Mob) Ready(godot gd.Context) {
	mobTypes := m.AnimatedSprite2D.GetSpriteFrames(godot).GetAnimationNames(godot)
	pick := int64(rand.Int() % int(mobTypes.Size()))
	m.AnimatedSprite2D.Play(mobTypes.Index(godot, pick).StringName(godot), 1, false)
}

func (m *Mob) OnVisibleOnScreenNotifier2DScreenExited() { m.Super().AsNode().QueueFree() }
