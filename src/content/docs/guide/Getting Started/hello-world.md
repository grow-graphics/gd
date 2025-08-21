---
title: Hello World
slug: guide/hello-world
sidebar:
  order: 10
---

This file is all you need to start a project, save it somewhere and use `gd run` to see the result!

```go
package main

import (
	"graphics.gd/startup"

	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/SceneTree"
)

func main() {
	startup.LoadingScene() // setup the SceneTree and wait until we have access to engine functionality
	hello := Label.New()
	hello.AsControl().SetAnchorsPreset(Control.PresetFullRect) // expand the label to take up the whole screen.
	hello.SetHorizontalAlignment(GUI.HorizontalAlignmentCenter)
	hello.SetVerticalAlignment(GUI.VerticalAlignmentCenter)
	hello.SetText("Hello, World!")
	SceneTree.Add(hello)
	startup.Scene() // starts up the scene and blocks until the engine shuts down.
}
```
