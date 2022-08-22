package main

import (
	"github.com/readykit/gd"
)

type HelloWorld struct {
	gd.Object
}

var HelloWorlds gd.Extension[HelloWorld]

func main() {}
func init() {
	gd.Register(&HelloWorlds)
}
