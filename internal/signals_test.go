//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/Node2D"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/Vector2"
)

type CustomSignal struct {
	Node2D.Extension[CustomSignal]

	HealthChanged chan<- func() (old, new int)

	GenericSignal Signal.Solo[int]

	Health int
}

func (c *CustomSignal) Ready() {
	if c.Health == 0 {
		c.Health = 10
	}
}

func (c *CustomSignal) TakeDamage(amount int) {
	oldHealth := c.Health
	c.Health -= amount
	c.HealthChanged <- func() (int, int) { return oldHealth, c.Health }
}

func TestSignals(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		custom := new(CustomSignal)
		custom.HealthChanged = make(chan func() (int, int), 1)
		custom.TakeDamage(10)
	})
}

type CustomStringSignals struct {
	Node2D.Extension[CustomStringSignals]

	StringSignal Signal.Solo[string] `gd:"on_string(s)"`
}

// TestSignalDisconnect tests that a callable created with Callable.New can be
// disconnected from a Godot signal after being connected, using the same
// Callable.Function value for both operations. This is the bug reported in
// https://github.com/quaadgras/graphics.gd/discussions/263
func TestSignalDisconnect(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		custom := new(CustomStringSignals)
		signal := gd.NewSignalOf(custom.AsObject(), gd.NewStringName("on_string"))

		var callCount int
		do := func(s string) {
			callCount++
		}
		mycb := Callable.New(do)

		// Connect via InternalCallable (this is what SignalProxy.Attach does)
		c1 := gd.InternalCallable(mycb)
		if err := signal.Connect(c1, 0); err != 0 {
			t.Fatalf("Failed to connect signal: %d", err)
		}

		// Emit and verify the callable fires
		signal.Emit(gd.NewVariant("test"))
		if callCount != 1 {
			t.Fatalf("Expected 1 call after connect, got %d", callCount)
		}

		// Disconnect via InternalCallable (this is what SignalProxy.Remove does)
		c2 := gd.InternalCallable(mycb)
		signal.Disconnect(c2)

		// Emit again and verify the callable no longer fires
		signal.Emit(gd.NewVariant("test2"))
		if callCount != 1 {
			t.Fatalf("Expected callCount to remain 1 after disconnect, got %d", callCount)
		}
	})
}

func TestSignalString(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		custom := new(CustomStringSignals)
		signal := gd.NewSignalOf(custom.AsObject(), gd.NewStringName("on_string"))
		var triggered int
		if err := signal.Connect(gd.NewCallable(func(s string) {
			if s != "Hello World" {
				t.Fail()
			}
			triggered++
		}), 0); err != 0 {
			t.Fatalf("Failed to connect signal: %d", err)
		}
		custom.StringSignal.Attach(Callable.New(func(s string) {
			if s != "Hello World" {
				t.Fail()
			}
			triggered++
		}))
		Signal.Via(gd.SignalProxy{}, pointers.Pack(signal)).Emit(variant.New("Hello World"))
		signal.Emit(gd.NewVariant("Hello World"))
		if triggered != 4 {
			t.Fatalf("Expected 4 triggers, got %d", triggered)
		}
	})
}

type SignalPayload struct {
	Position Vector2.XY
	Speed    Vector2.XY `gd:"speed"`

	hidden int
}

type SignalBullet struct {
	Area2D.Extension[SignalBullet]

	Position Vector2.XY
	Speed    Vector2.XY
}

type SignalShooter struct {
	Node2D.Extension[SignalShooter]

	Fired   Signal.Solo[SignalPayload]
	Shot    Signal.Solo[SignalBullet]
	Spawned Signal.Solo[*SignalBullet]
	Scored  Signal.Pair[variant.Any, any]
}

func init() {
	classdb.Register[SignalBullet]()
	classdb.Register[SignalShooter]()
}

// TestSignalStructs tests that a Go function connected to the signal of an
// extension class receives the values the signal was emitted with, as the
// types the function declares. The engine carries a struct as a Dictionary
// and an extension class as an Object, which used to reach the function as
// a map[any]any, failing the call. This is the bug reported in
// https://github.com/quaadgras/graphics.gd/issues/337
func TestSignalStructs(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		shooter := new(SignalShooter)
		shooter.AsObject()
		position, speed := Vector2.XY{X: 1, Y: 2}, Vector2.XY{X: -30, Y: -400}

		var fired []SignalPayload
		shooter.Fired.Call(func(payload SignalPayload) { fired = append(fired, payload) })
		shooter.Fired.Emit(SignalPayload{Position: position, Speed: speed, hidden: 1})
		if len(fired) != 1 || fired[0] != (SignalPayload{Position: position, Speed: speed}) {
			t.Fatalf("struct: got %+v", fired)
		}

		// a struct that embeds a class extension is carried by its fields
		// like any other struct, the copy has an engine object of its own.
		var shot []Vector2.XY
		shooter.Shot.Call(func(bullet SignalBullet) { shot = append(shot, bullet.Position, bullet.Speed) })
		shooter.Shot.Emit(SignalBullet{Position: position, Speed: speed})
		if len(shot) != 2 || shot[0] != position || shot[1] != speed {
			t.Fatalf("struct embedding an extension: got %+v", shot)
		}

		var spawned []*SignalBullet
		shooter.Spawned.Attach(Callable.New(func(bullet *SignalBullet) { spawned = append(spawned, bullet) }))
		bullet := &SignalBullet{Position: position, Speed: speed}
		shooter.Spawned.Emit(bullet)
		if len(spawned) != 1 || spawned[0] != bullet {
			t.Fatalf("extension pointer: got %v, want %p", spawned, bullet)
		}

		var scored []any
		shooter.Scored.Call(func(a variant.Any, b any) { scored = append(scored, a.Interface(), b) })
		shooter.Scored.Emit(variant.New(5), "points")
		shooter.Scored.Emit(variant.Nil, nil)
		if len(scored) != 4 || scored[0] != 5 || scored[1] != "points" || scored[2] != nil || scored[3] != nil {
			t.Fatalf("variants: got %#v", scored)
		}
	})
}

type SignalEarly struct {
	Area2D.Extension[SignalEarly]

	Removed Signal.Void
	Scored  Signal.Solo[int]

	score int
}

func NewSignalEarly() *SignalEarly {
	early := new(SignalEarly)
	early.Scored.Call(func(points int) { early.score += points })
	return early
}

func init() { classdb.Register[SignalEarly](NewSignalEarly) }

// TestSignalCallBeforeObject tests that a function connected to the signal
// field of an extension class, before the value has an engine object, is
// called when the signal is emitted later. The connection used to be made
// on a copy of the field and was lost, so the signal was emitted to nobody.
// This is the bug reported in https://github.com/quaadgras/graphics.gd/issues/343
func TestSignalCallBeforeObject(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		bullet := new(SignalEarly)
		var removed int
		bullet.Removed.Call(func() { removed++ })
		before := bullet.Removed // a copy made before the value has an object.
		bullet.Removed.Emit()
		if removed != 1 {
			t.Fatalf("before the object: got %d calls, want 1", removed)
		}
		cast, ok := Object.As[*SignalEarly](bullet.AsArea2D())
		if !ok || cast != bullet {
			t.Fatalf("Object.As: got %p %v, want %p", cast, ok, bullet)
		}
		cast.Removed.Emit()
		if removed != 2 {
			t.Fatalf("after the object: got %d calls, want 2", removed)
		}
		before.Emit()
		if removed != 3 {
			t.Fatalf("through the earlier copy: got %d calls, want 3", removed)
		}
		gd.NewSignalOf(bullet.AsObject(), gd.NewStringName("removed")).Emit()
		if removed != 4 {
			t.Fatalf("from the engine: got %d calls, want 4", removed)
		}

		// connections made by a registered constructor.
		made := NewSignalEarly()
		made.AsObject()
		made.Scored.Emit(3)
		if made.score != 3 {
			t.Fatalf("constructor: got %d, want 3", made.score)
		}
	})
}
