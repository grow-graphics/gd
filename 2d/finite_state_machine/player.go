package main

import (
	"reflect"
	"strings"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Marker2D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Timer"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

var bullet PackedScene.Instance = Resource.Load[PackedScene.Instance]("res://player/weapon/Bullet.tscn")

type Player struct {
	classdb.Extension[Player, CharacterBody2D.Instance] `gd:"FiniteStateMachinePlayer"`

	AnimationPlayer    AnimationPlayer.Instance
	StateNameDisplayer Label.Instance
	BodyPivot          struct {
		Marker2D.Instance

		WeaponPivot struct {
			Marker2D.Instance

			Offset struct {
				Marker2D.Instance

				Sword *Sword
			}
		}
		BulletSpawn struct {
			Node2D.Instance

			CooldownTimer Timer.Instance
		}
	}

	state StateMachine[Player, playerIdle]

	look_direction            Vector2.XY
	max_walk_speed            Float.X
	max_run_speed             Float.X
	base_max_horizontal_speed Float.X
	air_acceleration          Float.X
	air_deceleration          Float.X
	air_steering_power        Float.X
	gravity                   Float.X
	enter_jump_velocity       Vector2.XY
	max_horizontal_speed      Float.X
	horizontal_speed          Float.X
	horizontal_velocity       Vector2.XY
	vertical_speed            Float.X
	jump_height               Float.X
	speed                     Float.X
	z_index_start             int
}

type (
	playerAttacking   struct{ BasicStateFor[Player] }
	playerHurting     struct{ BasicStateFor[Player] }
	playerJumping     struct{ BasicStateFor[Player] }
	playerIdle        struct{ BasicStateFor[Player] }
	playerMoving      struct{ BasicStateFor[Player] }
	playerOnTheGround struct{ BasicStateFor[Player] }
)

func NewPlayer() *Player {
	return &Player{
		base_max_horizontal_speed: 400,
		air_acceleration:          1000,
		air_deceleration:          2000,
		air_steering_power:        50,
		gravity:                   1600,
		max_walk_speed:            450,
		max_run_speed:             700,
		look_direction:            Vector2.Right,
	}
}

func (p *Player) Ready() {
	p.BodyPivot.WeaponPivot.Offset.Sword.OnAttackFinished(p.OnSwordAttackFinished)
	p.AnimationPlayer.AsAnimationMixer().OnAnimationFinished(p.OnAnimationFinished)
	p.z_index_start = p.BodyPivot.WeaponPivot.AsNode2D().AsCanvasItem().ZIndex()
}

func (p *Player) Process(delta Float.X) {
	p.state.Current().Update(p, delta)
}

func (p *Player) changeState(state State[Player]) {
	switch state.(type) {
	case playerHurting, playerJumping, playerAttacking:
		p.state.Stack.PushFront(state)
	}
	p.state.Change(p, state)
	p.StateNameDisplayer.SetText(strings.TrimPrefix(reflect.TypeOf(state).Name(), "player"))
}

func (p *Player) changeBack() {
	p.state.ChangeBackToPrevious(p)
	p.StateNameDisplayer.SetText(strings.TrimPrefix(reflect.TypeOf(p.state.Current()).Name(), "player"))
}

func (p *Player) StatesStack() []string {
	var stack []string
	for _, state := range p.state.Stack.Iter() {
		stack = append(stack, strings.TrimPrefix(reflect.TypeOf(state).Name(), "player"))
	}
	if len(stack) == 0 {
		stack = append(stack, "Idle")
	}
	return stack
}

func (p *Player) UnhandledInput(event InputEvent.Instance) {
	if event.IsActionPressed("fire") {
		gun := p.BodyPivot.BulletSpawn
		if !gun.CooldownTimer.IsStopped() {
			return
		}
		gun.CooldownTimer.Start()
		b, _ := classdb.As[*Bullet](Node.Instance(bullet.Instantiate()))
		b.Super().AsNode2D().SetPosition(gun.AsNode2D().GlobalPosition())
		b.direction = p.look_direction
		gun.AsNode().AddChild(b.Super().AsNode())
	}
	if event.IsActionPressed("attack") {
		switch p.state.Current().(type) {
		case playerAttacking, playerHurting:
			return
		}
		p.changeState(playerAttacking{})
	}
	p.state.Current().HandleInput(p, event)
}

func (p *Player) OnAnimationFinished(anim_name string) {
	p.state.Current().OnAnimationFinished(p, anim_name)
}

func (p *Player) get_input_direction() Vector2.XY {
	return Vector2.New(
		Input.GetAxis("move_left", "move_right"),
		Input.GetAxis("move_up", "move_down"),
	)
}

func (p *Player) update_look_direction(direction Vector2.XY) {
	if direction != Vector2.Zero && p.look_direction != direction {
		p.look_direction = direction
		p.BodyPivot.WeaponPivot.AsNode2D().SetRotation(Float.X(Vector2.AngleRadians(direction)))
		switch direction {
		case Vector2.Up:
			p.BodyPivot.WeaponPivot.AsCanvasItem().SetZIndex(p.z_index_start - 1)
		default:
			p.BodyPivot.WeaponPivot.AsCanvasItem().SetZIndex(p.z_index_start)
		}
	}
}

func (p *Player) OnSwordAttackFinished() { p.changeBack() }

func (playerAttacking) Enter(p *Player) {
	p.AnimationPlayer.PlayNamed("idle")
	p.BodyPivot.WeaponPivot.Offset.Sword.Attack()
}
func (playerHurting) Enter(p *Player)                                 { p.AnimationPlayer.PlayNamed("stagger") }
func (playerHurting) OnAnimationFinished(p *Player, anim_name string) { p.changeBack() }
func (playerJumping) Enter(p *Player) {
	if StateIs[playerMoving](p.state.Current()) {
		p.horizontal_speed = p.speed
		if p.speed > 0 {
			p.max_horizontal_speed = p.speed
		} else {
			p.max_horizontal_speed = p.base_max_horizontal_speed
		}
		p.enter_jump_velocity = p.Super().Velocity()
	}
	var input_direction = p.get_input_direction()
	p.update_look_direction(input_direction)
	if input_direction != Vector2.Zero {
		p.horizontal_velocity = p.enter_jump_velocity
	} else {
		p.horizontal_velocity = Vector2.Zero
	}
	p.vertical_speed = 600
	p.AnimationPlayer.PlayNamed("idle")
}
func (playerJumping) HandleInput(p *Player, event InputEvent.Instance) {
	if event.IsActionPressed("simulate_damage") {
		p.changeState(playerHurting{})
	}
}
func (playerJumping) Update(p *Player, delta Float.X) {
	var direction = p.get_input_direction()
	p.update_look_direction(direction)
	if direction != Vector2.Zero {
		p.horizontal_speed += p.air_acceleration * delta
	} else {
		p.horizontal_speed -= p.air_deceleration * delta
	}
	p.horizontal_speed = Float.Clamp(p.horizontal_speed, 0, p.max_horizontal_speed)
	var target_velocity = Vector2.MulX(Vector2.Normalized(direction), p.horizontal_speed)
	var steering_velocity = Vector2.MulX(
		Vector2.Normalized(Vector2.Sub(target_velocity, p.horizontal_velocity)), p.air_steering_power)
	p.horizontal_velocity = Vector2.Add(p.horizontal_velocity, steering_velocity)
	p.Super().SetVelocity(p.horizontal_velocity)
	p.Super().MoveAndSlide()
	p.vertical_speed -= p.gravity * delta
	p.jump_height += p.vertical_speed * delta
	p.jump_height = max(0, p.jump_height)
	pos := p.BodyPivot.AsNode2D().Position()
	pos.Y = -p.jump_height
	p.BodyPivot.AsNode2D().SetPosition(pos)
	if p.jump_height <= 0 {
		p.changeBack()
	}
}

func (playerIdle) Enter(p *Player) { p.AnimationPlayer.PlayNamed("idle") }
func (playerIdle) HandleInput(p *Player, event InputEvent.Instance) {
	playerOnTheGround{}.HandleInput(p, event)
}
func (playerIdle) Update(p *Player, delta Float.X) {
	var input_direction = p.get_input_direction()
	if input_direction != Vector2.Zero {
		p.changeState(playerMoving{})
	}
}
func (playerMoving) Enter(p *Player) {
	p.speed = 0
	p.Super().SetVelocity(Vector2.Zero)
	var input_direction = p.get_input_direction()
	p.update_look_direction(input_direction)
	p.AnimationPlayer.PlayNamed("walk")
}
func (playerMoving) HandleInput(p *Player, event InputEvent.Instance) {
	playerOnTheGround{}.HandleInput(p, event)
}
func (state playerMoving) Update(p *Player, delta Float.X) {
	var direction = p.get_input_direction()
	if Vector2.IsApproximatelyZero(direction) {
		p.changeState(playerIdle{})
		return
	}
	p.update_look_direction(direction)
	if Input.IsActionPressed("run") {
		p.speed = p.max_run_speed
	} else {
		p.speed = p.max_walk_speed
	}
	p.Super().SetVelocity(Vector2.MulX(Vector2.Normalized(direction), p.speed))
	p.Super().MoveAndSlide()
}
func (playerOnTheGround) HandleInput(p *Player, event InputEvent.Instance) {
	if event.IsActionPressed("jump") {
		p.changeState(playerJumping{})
	}
	if event.IsActionPressed("simulate_damage") {
		p.changeState(playerHurting{})
	}
}
