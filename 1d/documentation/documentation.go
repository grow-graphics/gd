package main

import (
	"graphics.gd/classdb"
	"graphics.gd/startup"
	"graphics.gd/variant/Object"
)

type ClassWithDocumentation struct {
	Object.Extension[ClassWithDocumentation] `gd:"ClassWithDocumentation"
		serves as an example of how to document a class in Go code.`

	MyField int `gd:"my_field"
		can store an integer value.`
}

func (c *ClassWithDocumentation) MyMethod() {}

func main() {
	classdb.Register[ClassWithDocumentation](
		map[string]any{
			"ignore_this_method": (*ClassWithDocumentation).MyMethod,
		},
		map[string]string{
			"ignore_this_method": `
				does not do anything.`,
		},
	)
	startup.Scene()
}
