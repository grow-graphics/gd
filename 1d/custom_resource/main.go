package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Resource"
	"graphics.gd/startup"
)

type MyResource struct {
	classdb.Extension[MyResource, Resource.Instance]

	Value string
}

type MyNode struct {
	classdb.Extension[MyNode, Node.Instance]

	Resource *MyResource
}

func main() {
	classdb.Register[MyResource]()
	classdb.Register[MyNode]()
	startup.Scene()
}
