package main

import (
	"fmt"

	"github.com/readykit/gd"
)

func main() {}

type HelloWorld struct {
	Object gd.Object
}

func NewHelloWorld(obj gd.Object) HelloWorld {
	return HelloWorld{obj}
}

func (h HelloWorld) Print() {
	fmt.Println("Hello World!")
}

var gdHelloWorld = gd.Register(NewHelloWorld)
