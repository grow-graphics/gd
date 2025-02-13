package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Signal"
)

const MaxComboCount = 3

type Sword struct {
	classdb.Extension[Sword, Area2D.Instance] `gd:"FiniteStateMachineSword"`

	AttackFinished  Signal.Any `gd:"attack_finished()"`
	AnimationPlayer AnimationPlayer.Instance

	state                 SwordState
	attack_input_state    AttackInputState
	ready_for_next_attack bool
	combo_count           int
	attack_current        Combo
}

func (sword *Sword) AsNode() Node.Instance { return sword.Super().AsNode() }

type Combo struct {
	Damage    int
	Animation string
}

var Combos = [...]Combo{
	{
		Damage:    1,
		Animation: "attack_fast",
	},
	{
		Damage:    1,
		Animation: "attack_fast",
	},
	{
		Damage:    3,
		Animation: "attack_medium",
	},
}

type SwordState Enum.Int[struct {
	Idle   SwordState
	Attack SwordState
}]

var SwordStates = Enum.Values[SwordState]()

type AttackInputState Enum.Int[struct {
	Idle       AttackInputState
	Listening  AttackInputState
	Registered AttackInputState
}]

var AttackInputStates = Enum.Values[AttackInputState]()

func (sword *Sword) Ready() {
	sword.AnimationPlayer.AsAnimationMixer().OnAnimationFinished(sword.onAnimationFinished)
	sword.changeState(SwordStates.Idle)
}

func (sword *Sword) changeState(state SwordState) {
	switch sword.state {
	case SwordStates.Attack:
		sword.attack_input_state = AttackInputStates.Listening
		sword.ready_for_next_attack = false
	}
	switch state {
	case SwordStates.Idle:
		sword.combo_count = 0
		sword.AnimationPlayer.Stop()
		sword.Super().AsCanvasItem().SetVisible(false)
		sword.Super().SetMonitoring(false)
	case SwordStates.Attack:
		sword.attack_current = Combos[sword.combo_count-1]
		sword.AnimationPlayer.PlayNamed(sword.attack_current.Animation)
		sword.Super().AsCanvasItem().SetVisible(true)
		sword.Super().SetMonitoring(true)
	}
	sword.state = state
}

func (sword *Sword) UnhandledInput(event InputEvent.Instance) {
	if sword.state != SwordStates.Attack {
		return
	}
	if sword.attack_input_state != AttackInputStates.Listening {
		return
	}
	if event.IsActionPressed("attack") {
		sword.attack_input_state = AttackInputStates.Registered
	}
}

func (sword *Sword) PhysicsProcess(delta Float.X) {
	if sword.attack_input_state == AttackInputStates.Registered {
		sword.Attack()
	}
}

func (sword *Sword) Attack() {
	sword.combo_count %= MaxComboCount
	sword.combo_count++
	sword.changeState(SwordStates.Attack)
}

func (sword *Sword) SetAttackInputListening() {
	sword.attack_input_state = AttackInputStates.Listening
}

func (sword *Sword) SetReadyForNextAttack() {
	sword.ready_for_next_attack = true
}

func (sword *Sword) OnAttackFinished(cb func()) {
	sword.AttackFinished.Attach(Callable.New(cb))
}

func (sword *Sword) onAnimationFinished(anim_name string) {
	if sword.attack_current == (Combo{}) {
		return
	}
	if sword.attack_input_state == AttackInputStates.Registered {
		sword.Attack()
	} else {
		sword.changeState(SwordStates.Idle)
		sword.AttackFinished.Emit()
	}
}
