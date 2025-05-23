//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/variant/Signal"
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
	classdb.Register[CustomSignal]()

	custom := new(CustomSignal)
	custom.HealthChanged = make(chan func() (int, int), 1)
	custom.TakeDamage(10)

}
