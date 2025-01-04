package main

import (
	"graphics.gd/startup"

	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/SceneTree"
)

func main() {
	startup.Wait()
	hello := Label.New()
	hello.AsControl().SetAnchorsPreset(Control.PresetFullRect)
	hello.SetHorizontalAlignment(Label.HorizontalAlignmentCenter)
	hello.SetVerticalAlignment(Label.VerticalAlignmentCenter)
	hello.SetText("Hello, World!")
	SceneTree.Add(hello)
}
