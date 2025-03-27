package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
)

func main() {
	classdb.Register[Box]()
	classdb.Register[Coin]()
	classdb.Register[FullScreen](NewFullScreenHandler)
	classdb.Register[DestroyedBox](NewDestroyedBox)
	startup.LoadingScene()
	SceneTree.Add(NewFullScreenHandler())
	startup.Scene()
}
