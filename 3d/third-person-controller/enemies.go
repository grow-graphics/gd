package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Animation"
	"graphics.gd/classdb/AnimationNodeStateMachinePlayback"
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/AnimationTree"
	"graphics.gd/classdb/Area3D"
	"graphics.gd/classdb/AudioStreamPlayer3D"
	"graphics.gd/classdb/CollisionShape3D"
	"graphics.gd/classdb/KinematicCollision3D"
	"graphics.gd/classdb/Material"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/NavigationAgent3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/RigidBody3D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Timer"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector3"
)

type SmokePuff struct {
	classdb.Extension[SmokePuff, Node3D.Instance] `gd:"SmokePuff"`

	SmokeSounds     Node3D.Instance          `gd:"SmokeSounds"`
	AnimationPlayer AnimationPlayer.Instance `gd:"AnimationPlayer"`

	Full Signal.Void `gd:"full()"`
}

func (puff *SmokePuff) AsNode() Node.Instance {
	return puff.Super().AsNode()
}

func (puff *SmokePuff) Ready() {
	smoke_sounds := puff.SmokeSounds.AsNode().GetChildren()
	pick := smoke_sounds[Int.RandomBetween(0, len(smoke_sounds)-1)]
	Object.To[AudioStreamPlayer3D.Instance](pick).Play()

	puff.AnimationPlayer.PlayNamed("poof")
	puff.AnimationPlayer.AsAnimationMixer().OnAnimationFinished(func(anim_name string) {
		puff.Super().AsNode().QueueFree()
	})
}

func (puff *SmokePuff) SmokeAtFullDensity() {
	puff.Full.Emit()
}

var BulletScene = Resource.Load[PackedScene.Is[*Bullet]]("res://Player/Bullet.tscn")
var PuffScene = Resource.Load[PackedScene.Is[*SmokePuff]]("res://Enemies/smoke_puff/smoke_puff.tscn")

type BeeBot struct {
	classdb.Extension[BeeBot, RigidBody3D.Instance] `gd:"BeeBot"`

	ShootTimer  Float.X
	BulletSpeed Float.X
	CoinsCount  int

	ReactionAnimationPlayer AnimationPlayer.Instance     `gd:"ReactionLabel/AnimationPlayer"`
	FlyingAnimationPlayer   AnimationPlayer.Instance     `gd:"MeshRoot/AnimationPlayer"`
	DetectionArea           Area3D.Instance              `gd:"PlayerDetectionArea"`
	DeathMeshCollider       CollisionShape3D.Instance    `gd:"DeathMeshCollider"`
	Renderer                *BeeBotRenderer              `gd:"MeshRoot/bee_root"`
	DefeatSound             AudioStreamPlayer3D.Instance `gd:"DefeatSound"`

	shoot_count Float.X
	target      Node3D.ID
	alive       bool
}

func NewBeeBot() *BeeBot {
	return &BeeBot{
		ShootTimer:  1.5,
		BulletSpeed: 6,
		CoinsCount:  5,
	}
}

func (bee *BeeBot) Ready() {
	bee.alive = true
	bee.DetectionArea.OnBodyEntered(func(body Node3D.Instance) {
		if player, ok := Object.As[*DemoPlayer](body); ok {
			bee.shoot_count = 0
			bee.target = player.Super().AsNode3D().ID()
			bee.ReactionAnimationPlayer.PlayNamed("found_player")
		}
	})
	bee.DetectionArea.OnBodyExited(func(body Node3D.Instance) {
		if Object.Is[*DemoPlayer](body) {
			bee.target = 0
			bee.ReactionAnimationPlayer.PlayNamed("lost_player")
		}
	})
	bee.Renderer.PlayIdle()
}

func (bee *BeeBot) PhysicsProcess(delta Float.X) {
	if target, ok := bee.target.Instance(); ok && bee.alive {
		var target_transform = Transform3D.LookingAt(bee.Super().AsNode3D().Transform(), target.GlobalPosition(), Vector3.Up)
		bee.Super().AsNode3D().SetTransform(Transform3D.Lerp(bee.Super().AsNode3D().Transform(), target_transform, 0.1))
		bee.shoot_count += delta
		if bee.shoot_count > bee.ShootTimer {
			bee.Renderer.PlaySpitAttack()
			bee.shoot_count -= bee.ShootTimer
			var bullet = BulletScene.Instantiate()
			bullet.Shooter = bee.Super().AsNode().ID()
			origin := bee.Super().AsNode3D().GlobalPosition()
			target := Vector3.Add(target.GlobalPosition(), Vector3.Up)
			aim_direction := Vector3.Normalized(Vector3.Sub(target, bee.Super().AsNode3D().GlobalPosition()))
			bullet.Velocity = Vector3.MulX(aim_direction, bee.BulletSpeed)
			bullet.DistanceLimit = 14
			bee.Super().AsNode().GetParent().AddChild(bullet.AsNode())
			bullet.Super().SetGlobalPosition(origin)
		}
	}
}

func (bee *BeeBot) Damage(impact_point Vector3.XYZ, force Vector3.XYZ) {
	force = Vector3.LengthLimited(force, 3.0)
	RigidBody3D.Expanded(bee.Super()).ApplyImpulse(force, impact_point)
	if !bee.alive {
		return
	}
	bee.DefeatSound.Play()
	bee.alive = false
	bee.FlyingAnimationPlayer.Stop()
	// FIXME disconnect all signals ?
	bee.target = 0
	Callable.Defer(Callable.New(func() {
		bee.DeathMeshCollider.SetDisabled(false)
	}))
	bee.Super().SetGravityScale(1)
	bee.Renderer.PlayPoweroff()
	SceneTree.Get(bee.Super().AsNode()).CreateTimer(2).OnTimeout(func() {
		var puff = PuffScene.Instantiate()
		bee.Super().AsNode().GetParent().AddChild(puff.AsNode())
		puff.Super().SetGlobalPosition(bee.Super().AsNode3D().GlobalPosition())
		puff.Full.Attach(Callable.New(func() {
			for range bee.CoinsCount {
				var coin = CoinScene.Instantiate()
				bee.Super().AsNode().GetParent().AddChild(coin.AsNode())
				coin.Super().AsNode3D().SetGlobalPosition(bee.Super().AsNode3D().GlobalPosition())
				coin.Spawn()
			}
			bee.Super().AsNode().QueueFree()
		}), Signal.OneShot)
	})
}

type BeeBotRenderer struct {
	classdb.Extension[BeeBotRenderer, Node3D.Instance] `gd:"BeeBotRenderer"`

	AnimationTree AnimationTree.Instance
	Mesh          MeshInstance3D.Instance `gd:"bee_bot/Armature/Skeleton3D/bee_bot2"`
}

func (r *BeeBotRenderer) AsNode() Node.Instance {
	return r.Super().AsNode()
}

func (r *BeeBotRenderer) state_machine() AnimationNodeStateMachinePlayback.Instance {
	return Object.Get(r.AnimationTree, "parameters/StateMachine/playback").(classdb.AnimationNodeStateMachinePlayback)
}

func (r *BeeBotRenderer) Ready() {
	r.AnimationTree.AsAnimationMixer().SetActive(true)
	r.PlayIdle()
}

func (r *BeeBotRenderer) PlayIdle() {
	r.state_machine().Travel("idle")
}

func (r *BeeBotRenderer) PlaySpitAttack() {
	r.state_machine().Travel("spit_attack")
}

func (r *BeeBotRenderer) PlayPoweroff() {
	r.state_machine().Travel("power_off")
}

func (r *BeeBotRenderer) ExitTree() {
	r.Mesh.SetSurfaceOverrideMaterial(1, Material.Nil)
	r.Mesh.SetSurfaceOverrideMaterial(2, Material.Nil)
	r.Mesh.SetSurfaceOverrideMaterial(3, Material.Nil)
}

type Beetle struct {
	classdb.Extension[Beetle, RigidBody3D.Instance] `gd:"Beetle"`

	CoinsCount       int
	StoppingDistance Float.X

	ReactionAnimationPlayer AnimationPlayer.Instance     `gd:"ReactionLabel/AnimationPlayer"`
	DetectionArea           Area3D.Instance              `gd:"PlayerDetectionArea"`
	Renderer                *BeetleRenderer              `gd:"BeetlebotSkin"`
	NavigationAgent         NavigationAgent3D.Instance   `gd:"NavigationAgent"`
	DeathCollisionShape     CollisionShape3D.Instance    `gd:"DeathCollisionShape"`
	DefeatSound             AudioStreamPlayer3D.Instance `gd:"DefeatSound"`

	target Node3D.ID
	alive  bool
}

func (b *Beetle) Ready() {
	b.alive = true
	b.DetectionArea.OnBodyEntered(func(body Node3D.Instance) {
		if player, ok := Object.As[*DemoPlayer](body); ok {
			b.target = player.Super().AsNode3D().ID()
			b.ReactionAnimationPlayer.PlayNamed("found_player")
		}
	})
	b.DetectionArea.OnBodyExited(func(body Node3D.Instance) {
		if Object.Is[*DemoPlayer](body) {
			b.target = 0
			b.ReactionAnimationPlayer.PlayNamed("lost_player")
		}
	})
	b.Renderer.Idle()
}

func (b *Beetle) PhysicsProcess(delta Float.X) {
	if !b.alive {
		return
	}
	if target, ok := b.target.Instance(); ok {
		b.Renderer.Walk()
		target_look_position := target.GlobalPosition()
		target_look_position.Y = b.Super().AsNode3D().GlobalPosition().Y
		if target_look_position != Vector3.Zero {
			b.Super().AsNode3D().LookAt(target_look_position)
		}
		b.NavigationAgent.SetTargetPosition(target.GlobalPosition())
		var next_location = b.NavigationAgent.GetNextPathPosition()
		if !b.NavigationAgent.IsTargetReached() {
			var direction = Vector3.Sub(next_location, b.Super().AsNode3D().GlobalPosition())
			direction.Y = 0
			direction = Vector3.Normalized(direction)
			var collision = b.Super().AsPhysicsBody3D().MoveAndCollide(Vector3.MulX(direction, delta*3))
			if collision != KinematicCollision3D.Nil {
				var collider = collision.GetCollider()
				if player, ok := Object.As[*DemoPlayer](collider); ok {
					var impact_point = Vector3.Sub(b.Super().AsNode3D().GlobalPosition(), player.Super().AsNode3D().GlobalPosition())
					var force = Vector3.Neg(impact_point)
					force.Y = 0.5
					force = Vector3.MulX(force, 10)
					player.Damage(impact_point, force)
					b.Renderer.Attack()
				}
			}
		}
	}
}

func (b *Beetle) Damage(impact_point Vector3.XYZ, force Vector3.XYZ) {
	b.Super().SetLockRotation(false)
	force = Vector3.LengthLimited(force, 3.0)
	RigidBody3D.Expanded(b.Super()).ApplyImpulse(force, impact_point)
	if !b.alive {
		return
	}
	b.DefeatSound.Play()
	b.alive = false
	b.Renderer.Poweroff()
	// FIXME disconnect all signals ?
	b.target = 0
	Callable.Defer(Callable.New(func() {
		b.DeathCollisionShape.SetDisabled(false)
	}))
	b.Super().AsPhysicsBody3D().SetAxisLockAngularX(false)
	b.Super().AsPhysicsBody3D().SetAxisLockAngularY(false)
	b.Super().AsPhysicsBody3D().SetAxisLockAngularZ(false)
	b.Super().SetGravityScale(1)
	SceneTree.Get(b.Super().AsNode()).CreateTimer(2).OnTimeout(func() {
		var puff = PuffScene.Instantiate()
		b.Super().AsNode().GetParent().AddChild(puff.AsNode())
		puff.Super().SetGlobalPosition(b.Super().AsNode3D().GlobalPosition())
		puff.Full.Attach(Callable.New(func() {
			for range b.CoinsCount {
				var coin = CoinScene.Instantiate()
				b.Super().AsNode().GetParent().AddChild(coin.AsNode())
				coin.Super().AsNode3D().SetGlobalPosition(b.Super().AsNode3D().GlobalPosition())
				coin.Spawn()
			}
			b.Super().AsNode().QueueFree()
		}), Signal.OneShot)
	})
}

type BeetleRenderer struct {
	classdb.Extension[BeetleRenderer, Node3D.Instance] `gd:"BeetleRenderer"`

	ForceLoop            []string
	AnimationTree        AnimationTree.Instance
	AnimationPlayer      AnimationPlayer.Instance `gd:"beetle_bot/AnimationPlayer"`
	SecondaryActionTimer Timer.Instance
}

func (r *BeetleRenderer) AsNode() Node.Instance {
	return r.Super().AsNode()
}

func (r *BeetleRenderer) main_state_machine() AnimationNodeStateMachinePlayback.Instance {
	return Object.Get(r.AnimationTree, "parameters/StateMachine/playback").(classdb.AnimationNodeStateMachinePlayback)
}

func (r *BeetleRenderer) Ready() {
	r.AnimationTree.AsAnimationMixer().SetActive(true)
	for _, name := range r.ForceLoop {
		var anim = r.AnimationPlayer.AsAnimationMixer().GetAnimation(name)
		anim.SetLoopMode(Animation.LoopLinear)
	}
}

func (r *BeetleRenderer) OnSecondaryActionTimerTimeout() {
	if r.main_state_machine().GetCurrentNode() == "Idle" {
		r.Shake()
	}
	Timer.Expanded(r.SecondaryActionTimer).Start(Float.RandomBetween(3, 8))
}

func (r *BeetleRenderer) Idle()   { r.main_state_machine().Travel("Idle") }
func (r *BeetleRenderer) Walk()   { r.main_state_machine().Travel("Walk") }
func (r *BeetleRenderer) Shake()  { r.main_state_machine().Travel("Shake") }
func (r *BeetleRenderer) Attack() { r.main_state_machine().Travel("Attack") }
func (r *BeetleRenderer) Poweroff() {
	r.main_state_machine().Travel("PowerOff")
	r.SecondaryActionTimer.Stop()
}
