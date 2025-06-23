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
	classdb.Register[CameraMode](NewCameraMode)
	classdb.Register[DemoPage]()
	classdb.Register[DemoLinkButton]()
	classdb.Register[BeeBotRenderer]()
	classdb.Register[BeeBot](NewBeeBot)
	classdb.Register[Bullet](NewBullet)
	classdb.Register[SmokePuff]()
	classdb.Register[Beetle]()
	classdb.Register[BeetleRenderer]()
	classdb.Register[CameraController](NewCameraController)
	classdb.Register[DemoPlayer](NewDemoPlayer)
	classdb.Register[CharacterSkin](NewCharacterSkin)
	classdb.Register[GrenadeLauncher](NewGrenadeLauncher)
	classdb.Register[FaceShader](NewFaceShader)
	startup.LoadingScene()
	SceneTree.Add(NewFullScreenHandler())
	startup.Scene()
}
