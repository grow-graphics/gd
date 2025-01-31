package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Material"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/shaders"
	"graphics.gd/startup"
	"graphics.gd/variant/Vector2"
)

func main() {
	classdb.Register[MyFirstShader]()
	startup.LoadingScene()
	sprite := startup.Pin(Sprite2D.New())

	size := DisplayServer.WindowGetSize()

	sprite.SetTexture(Resource.Load[Texture2D.Instance]("res://icon.png"))
	sprite.AsNode2D().SetPosition(Vector2.New(size.X/2, size.Y/2))
	sprite.AsNode2D().SetScale(Vector2.New(5, 5))

	var materials = []func() Material.Instance{
		func() Material.Instance {
			material := MyFirstShader2D_Texture{}
			shaders.Compile(&material)
			return material.Super().AsMaterial()
		},
		func() Material.Instance {
			material := MyFirstShader{}
			shaders.Compile(&material)
			return material.Super().AsMaterial()
		},
		func() Material.Instance {
			material := MyFirstShader2D{}
			shaders.Compile(&material)
			return material.Super().AsMaterial()
		},
		func() Material.Instance {
			material := MyFirstShader2D_UV{}
			shaders.Compile(&material)
			return material.Super().AsMaterial()
		},
	}
	var index int
	sprite.AsCanvasItem().SetMaterial(materials[index]())

	button := Button.New()
	button.SetText("Switch Shader")
	button.AsControl().SetAnchorsAndOffsetsPreset(Control.PresetCenterBottom)
	button.SetAlignment(Button.HorizontalAlignmentCenter)
	buttonSize := button.AsControl().Size()
	button.AsControl().SetPosition(Vector2.New(-buttonSize.X/2, -100))

	button.AsBaseButton().OnPressed(func() {
		index++
		index %= len(materials)
		sprite.AsCanvasItem().SetMaterial(materials[index]())
	})
	SceneTree.Add(sprite)
	SceneTree.Add(button)
	startup.Scene()
}
