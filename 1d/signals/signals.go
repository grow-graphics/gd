package main

import (
	"time"

	"graphics.gd/classdb"
	"graphics.gd/startup"
	"graphics.gd/variant/Object"
)

type Signals struct {
	classdb.Extension[Signals, Object.Instance]

	Something chan<- struct{} `gd:"something"`
}

func (s *Signals) DoSomething() {
	go func() {
		time.Sleep(time.Second)
		s.Something <- struct{}{}
	}()
}

func main() {
	classdb.Register[Signals]()
	startup.Engine()
}
