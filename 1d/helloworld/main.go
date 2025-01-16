package main

import (
	"fmt"

	"graphics.gd/startup"
)

func main() {
	startup.Loader()
	fmt.Println("Hello, World!")
	startup.Engine()
}
