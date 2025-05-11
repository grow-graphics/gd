package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Material"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Vector2"
)

func main() {
	classdb.Register[MyFirstShader]()
	classdb.Register[MyFirstShader2D]()
	classdb.Register[MyFirstShader2D_Texture]()
	classdb.Register[MyFirstShader2D_UV]()
	startup.LoadingScene()
	sprite := startup.Pin(Sprite2D.New())

	size := DisplayServer.WindowGetSize(0)

	sprite.SetTexture(Resource.Load[Texture2D.Instance]("res://icon.png"))
	sprite.AsNode2D().SetPosition(Vector2.New(size.X/2, size.Y/2))
	sprite.AsNode2D().SetScale(Vector2.New(5, 5))

	var materials = []Material.Any{
		new(MyFirstShader2D_Texture),
		new(MyFirstShader),
		new(MyFirstShader2D),
		new(MyFirstShader2D_UV),
	}
	var index int
	sprite.AsCanvasItem().SetMaterial(materials[index].AsMaterial())

	button := Button.New()
	button.SetText("Switch Shader")
	button.AsControl().SetAnchorsAndOffsetsPreset(Control.PresetCenterBottom)
	button.SetAlignment(GUI.HorizontalAlignmentCenter)
	buttonSize := button.AsControl().Size()
	button.AsControl().SetPosition(Vector2.New(-buttonSize.X/2, -100))

	button.AsBaseButton().OnPressed(func() {
		index++
		index %= len(materials)
		sprite.AsCanvasItem().SetMaterial(materials[index].AsMaterial())
	})
	SceneTree.Add(sprite)
	SceneTree.Add(button)
	startup.Scene()
}
