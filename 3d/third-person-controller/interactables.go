package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Area3D"
	"graphics.gd/classdb/AudioStreamPlayer3D"
	"graphics.gd/classdb/CollisionShape3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/RigidBody3D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/SceneTreeTimer"
	"graphics.gd/classdb/Tween"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

const CoinsCount = 5

var CoinScene = Resource.Load[PackedScene.Is[*Coin]]("res://Player/Coin/Coin.tscn")
var DestroyedBoxScene = Resource.Load[PackedScene.Is[Node3D.Instance]]("res://Box/DestroyedBox.tscn")

type CoinCollector interface {
	Node3D.Any
	CollectCoin()
}

type DestroyedBox struct {
	classdb.Extension[DestroyedBox, Node3D.Instance] `gd:"DestroyedBox"`

	Model Node3D.Instance

	pieces_idx Array.Contains[int]
}

func NewDestroyedBox() *DestroyedBox {
	return &DestroyedBox{pieces_idx: Array.New(0, 1, 2, 3, 4, 5)}
}

func (box *DestroyedBox) Ready() {
	const (
		DestroyedBoxFlyingPieces  = 3
		DestroyedBoxThrowStrength = 500
	)
	Array.Shuffle(box.pieces_idx)
	for i := range DestroyedBoxFlyingPieces {
		var piece_idx = box.pieces_idx.Index(i)
		var piece = Object.To[RigidBody3D.Instance](Node.Instance(box.Model.AsNode().GetChild(piece_idx)))
		piece.AsNode3D().Show()
		piece.SetFreeze(false)
		piece.SetSleeping(false)
		piece.AsCollisionObject3D().SetCollisionMaskValue(1, true)
		var rand_vector = Vector3.Sub(Vector3.MulX(Vector3.One, 0.5), Vector3.Random())
		RigidBody3D.Advanced(piece).ApplyForce(Vector3.MulX(rand_vector, DestroyedBoxThrowStrength), rand_vector)
	}
}

type Box struct {
	classdb.Extension[Box, RigidBody3D.Instance] `gd:"Box"`

	DestroySound     AudioStreamPlayer3D.Instance
	CollisionShape3D CollisionShape3D.Instance
}

func (box *Box) Damage(_, _ Vector3.XYZ) {
	for range CoinsCount {
		var coin *Coin = CoinScene.Instantiate()
		Node.Instance.AddChild(box.Super().AsNode().GetParent(), coin.Super().AsNode())
		coin.Super().AsNode3D().SetGlobalPosition(box.Super().AsNode3D().GlobalPosition())
		coin.Spawn()
	}
	var destroyed_box = DestroyedBoxScene.Instantiate()
	box.Super().AsNode().AddChild(destroyed_box.AsNode())
	destroyed_box.SetGlobalPosition(box.Super().AsNode3D().GlobalPosition())
	Callable.Defer(Callable.New(func() {
		box.CollisionShape3D.SetDisabled(true)
	}))
	box.DestroySound.Play()
	box.DestroySound.OnFinished(func() {
		box.Super().AsNode().QueueFree()
	})
}

const (
	CoinMinLaunchRange      = 2.0
	CoinMaxLaunchRange      = 4.0
	CoinMinLaunchHeight     = 1.0
	CoinMaxLaunchHeight     = 3.0
	CoinSpawnTweenDuration  = 1.0
	CoinFollowTweenDuration = 0.5
)

type Coin struct {
	classdb.Extension[Coin, RigidBody3D.Instance] `gd:"Coin"`

	CollectAudio        AudioStreamPlayer3D.Instance
	PlayerDetectionArea Area3D.Instance

	initial_tween_position Vector3.XYZ
	target                 CoinCollector
}

func (coin *Coin) AsNode() Node.Instance { return coin.Super().AsNode() }

func (coin *Coin) Spawn() { coin.SpawnWithDelay(0.5) }

func (coin *Coin) SpawnWithDelay(delay Float.X) {
	var height = Float.RandomBetween(CoinMinLaunchHeight, CoinMaxLaunchHeight)
	var dir = Vector3.Rotated(Vector3.Forward, Vector3.Up, Angle.Radians(Float.Random()))
	var pos = Vector3.MulX(dir, Float.RandomBetween(CoinMinLaunchRange, CoinMaxLaunchRange))
	pos.Y = height
	coin.Super().ApplyCentralImpulse(pos)
	var timer SceneTreeTimer.Instance = SceneTree.Get(coin.Super().AsNode()).CreateTimer(delay)
	timer.OnTimeout(func() {
		coin.Super().AsCollisionObject3D().SetCollisionLayer(3)
	})
	coin.PlayerDetectionArea.OnBodyEntered(coin.OnBodyEntered)
}

func (coin *Coin) OnBodyEntered(body Node3D.Instance) {
	if player, ok := Object.As[*Player](Node3D.Instance(body)); ok {
		coin.SetTarget(player)
	}
}

func (coin *Coin) SetTarget(target CoinCollector) {
	if coin.target == nil {
		coin.Super().SetSleeping(true)
		coin.Super().SetFreeze(true)
		coin.initial_tween_position = coin.Super().AsNode3D().GlobalPosition()
		coin.target = target
		var tween Tween.Instance = coin.target.AsNode3D().AsNode().CreateTween()
		tween.TweenMethod(coin.follow, 0, 1, 0.5)
		tween.TweenCallback(coin.collect)
	}
}

func (coin *Coin) follow(offset any) {
	coin.Super().AsNode3D().SetGlobalPosition(Vector3.Lerp(coin.initial_tween_position, coin.target.AsNode3D().GlobalPosition(), offset.(Float.X)))
}

func (coin *Coin) collect() {
	coin.CollectAudio.SetPitchScale(Float.RandomlyDistributed(1.0, 0.1))
	coin.CollectAudio.Play()
	coin.target.CollectCoin()
	coin.Super().AsNode3D().Hide()
	coin.CollectAudio.OnFinished(func() {
		coin.Super().AsNode().QueueFree()
	})
}
