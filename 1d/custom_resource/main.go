package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Resource"
	"graphics.gd/startup"
)

type MyResource struct {
	Resource.Extension[MyResource]

	Value string
}

type MyNode struct {
	Node.Extension[MyNode]

	Resource *MyResource
}

func main() {
	classdb.Register[MyResource]()
	classdb.Register[MyNode]()
	startup.Scene()
}
