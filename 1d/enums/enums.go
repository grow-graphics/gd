package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	"graphics.gd/startup"
	"graphics.gd/variant/Enum"
)

type MyNode struct {
	Node.Extension[MyNode]

	Animal Animal `gd:"animal"`
}

type Animal Enum.Int[struct {
	Cat Animal
	Dog Animal
}]

var Animals = Enum.Values[Animal]()

func main() {
	classdb.Register[MyNode]()
	startup.Scene()
}
