package main

import (
	"strconv"

	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/CircleShape2D"
	"graphics.gd/classdb/CollisionShape2D"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/KinematicCollision2D"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Shape2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

type StateMachine[T any, S State[T]] struct {
	Stack Array.Contains[State[T]]
}

func (machine *StateMachine[T, S]) Change(obj *T, to State[T]) {
	machine.Current().Exit(obj)
	if machine.Stack.Len() == 0 {
		machine.Stack.PushFront(to)
	} else {
		machine.Stack.SetIndex(0, to)
	}
	machine.Current().Enter(obj)
}

func (machine *StateMachine[T, S]) ChangeBackToPrevious(obj *T) {
	if machine.Stack.Len() == 0 {
		return
	}
	machine.Current().Exit(obj)
	Array.PopFront(machine.Stack)
}

func (machine *StateMachine[T, S]) Current() State[T] {
	if machine.Stack.Len() == 0 {
		return [1]S{}[0]
	}
	return Array.First(machine.Stack)
}

func StateIs[S State[T], T any](state State[T]) bool {
	_, ok := state.(S)
	return ok
}

type State[T any] interface {
	Enter(*T)
	Exit(*T)
	HandleInput(*T, InputEvent.Instance)
	Update(*T, Float.X)
	OnAnimationFinished(*T, string)
}

type BasicStateFor[T any] struct{}

func (BasicStateFor[T]) Enter(*T)                            {}
func (BasicStateFor[T]) Exit(*T)                             {}
func (BasicStateFor[T]) HandleInput(*T, InputEvent.Instance) {}
func (BasicStateFor[T]) Update(*T, Float.X)                  {}
func (BasicStateFor[T]) OnAnimationFinished(*T, string)      {}

type Bullet struct {
	CharacterBody2D.Extension[Bullet] `gd:"FiniteStateMachineBullet"`

	CollisionShape2D CollisionShape2D.Instance

	direction Vector2.XY
	speed     Float.X
}

func NewBullet() *Bullet {
	return &Bullet{
		speed: 1000,
	}
}

func (b *Bullet) Ready() {
	b.AsCanvasItem().SetTopLevel(true)
}

func (b *Bullet) PhysicsProcess(delta Float.X) {
	root := SceneTree.Get(b.AsNode()).Root()
	if !Rect2.HasPoint(root.AsViewport().GetVisibleRect(), b.AsNode2D().Position()) {
		b.AsNode().QueueFree()
		return
	}
	var motion = Vector2.MulX(b.direction, b.speed*delta)
	var collision = b.AsPhysicsBody2D().MoveAndCollide(motion)
	if collision != KinematicCollision2D.Nil {
		b.AsNode().QueueFree()
	}
}

func (b *Bullet) Draw() {
	circle, _ := classdb.As[CircleShape2D.Instance](Shape2D.Instance(b.CollisionShape2D.Shape()))
	b.AsCanvasItem().DrawCircle(Vector2.Zero, circle.Radius(), Color.X11.White)
}

type StatesStackDisplayer struct {
	Panel.Extension[StatesStackDisplayer] `gd:"StatesStackDisplayer"`

	Player  *Player        `gd:"../../Player"`
	States  Label.Instance `gd:"%States"`
	Numbers Label.Instance `gd:"%Numbers"`
}

func (s StatesStackDisplayer) Process(Float.X) {
	var states_names = ""
	var numbers = ""
	for index, state := range s.Player.StatesStack() {
		states_names += state + "\n"
		numbers += strconv.Itoa(index) + "\n"
	}
	s.States.SetText(states_names)
	s.Numbers.SetText(numbers)
}

func main() {
	classdb.Register[Player](NewPlayer)
	classdb.Register[Sword]()
	classdb.Register[Bullet](NewBullet)
	classdb.Register[StatesStackDisplayer]()
	startup.Scene()
}
