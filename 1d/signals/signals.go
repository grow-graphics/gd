package main

import (
	"time"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type Signals struct {
	gd.Class[Signals, gd.Object]

	Something chan<- struct{} `gd:"something"`
}

func (s *Signals) DoSomething() {
	go func() {
		time.Sleep(time.Second)
		s.Something <- struct{}{}
	}()
}

func main() {
	godot, ok := gdextension.Link()
	if ok {
		gd.Register[Signals](godot)
	}
}
