package main

import (
	"math/rand"

	"grow.graphics/gd"
)

type Mob struct {
	gd.Class[Mob, gd.RigidBody2D] `gd:"DodgeTheCreepsMob"`

	AnimatedSprite2D gd.AnimatedSprite2D
}

func (m *Mob) Ready() {
	mobTypes := m.AnimatedSprite2D.GetSpriteFrames(m.Temporary).GetAnimationNames(m.Temporary)
	pick := int64(rand.Int() % int(mobTypes.Size()))
	m.AnimatedSprite2D.Play(mobTypes.Index(m.Temporary, pick).StringName(m.Temporary), 1, false)
}

func (m *Mob) OnVisibleOnScreenNotifier2DScreenExited() { m.Super().AsNode().QueueFree() }
