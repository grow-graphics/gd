package main

import (
	"time"

	gd "graphics.gd"
	"graphics.gd/defined"
	"graphics.gd/startup"
)

type Signals struct {
	defined.Object[Signals, gd.Object]

	Something chan<- struct{} `gd:"something"`
}

func (s *Signals) DoSomething() {
	go func() {
		time.Sleep(time.Second)
		s.Something <- struct{}{}
	}()
}

func main() {
	defined.InEditor[Signals]()
	startup.Engine()
}
