---
title: Registering New Classes
slug: guide/classdb/registration
sidebar:
  order: 30
---

In order to register a new class that will be available within the editor, create a Go struct that embeds the
`T.Extension` type of the object you would like to extend.

All struct fields beginning with an uppercase letter will be exported as properties in the editor, and all methods
beginning with an uppercase letter will be exported for use in scripts.

```go
package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	"graphics.gd/startup"
)

type MyClass struct {
	Node.Extension[MyClass]

	MyProperty string // properties will be exported as snake_case
}

// Implements the Ready method of the [Node.Interface]
func (*MyClass) Ready() {
	fmt.Println("MyClass is ready!")
}

// methods will be available to scripts in snake_case, ie. my_script_method
func (*MyClass) MyScriptMethod() {
	fmt.Println("MyClass script method called!")
}

func main() {
	classdb.Register[MyClass]()
	startup.Scene()
}

```
