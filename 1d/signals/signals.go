package main

import (
	"time"

	"graphics.gd/classdb"
	"graphics.gd/startup"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Signal"
)

type Signals struct {
	Object.Extension[Signals]

	Something chan<- struct{} `gd:"something"`

	Generic Signal.Solo[int] `gd:"generic"`
}

func (s *Signals) DoSomething() {
	go func() {
		time.Sleep(time.Second)
		s.Something <- struct{}{}
		s.Generic.Emit(22)
	}()
}

func main() {
	classdb.Register[Signals]()
	startup.Scene()
}
