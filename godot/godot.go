package godot

import (
	"github.com/readykit/gd"
	"github.com/readykit/gd/gdnative"
)

var Engine gd.Engine

func init() {
	gdnative.OnLoad(func() {
		Engine = gd.Engine(gdnative.GetSingleton("Engine\000"))
	})
}
