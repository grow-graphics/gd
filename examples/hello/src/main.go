package main

import (
	"fmt"

	"github.com/readykit/gd"
)

type HelloWorld struct {
	gd.Object
}

func (h HelloWorld) Print() {
	fmt.Println("Hello World!")
}

var HelloWorlds gd.Extension[HelloWorld]

func main() {}
func init() {
	gd.Register(&HelloWorlds)
}
