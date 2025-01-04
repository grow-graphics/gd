package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/shaders"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/vec1"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec4"
	"graphics.gd/startup"
)

type MyFirstShader struct {
	shaders.Type2D

	Offset vec2.XY `gd:"offset"`
}

type MyFirstFragment struct {
	Position vec2.XY `gd:"VERTEX"`
}

type MyFirstMaterial struct {
	Color vec4.RGBA `gd:"COLOR"`
}

func (uniform MyFirstShader) Fragment(vert shaders.CanvasItemVertexAttributes) shaders.CanvasItemVertexAttributes {
	return shaders.CanvasItemVertexAttributes{
		Position: vec2.Add(vert.Position, vec2.New(10.0, 0.1)),
	}
}

func (MyFirstShader) Material(vert shaders.CanvasItemVertexAttributes) MyFirstMaterial {
	// smooth back and forth between pink and blue
	step := vec1.Modf(vert.Global.Time, vec1.New(2.0))
	step = vec1.If(vec1.Gt(step, vec1.New(1.0)), vec1.Sub(vec1.New(2.0), step), step)
	return MyFirstMaterial{
		Color: rgba.New(step, 0.6, 0.9, 1.0),
	}
}

func (MyFirstShader) Lighting(matl MyFirstMaterial) vec4.RGBA {
	return matl.Color
}

func main() {
	classdb.Register[MyFirstShader]()
	startup.Wait()
	sprite := Sprite2D.New()
	image := Image.CreateEmpty(512, 512, false, Image.FormatRgba8)
	sprite.SetTexture(ImageTexture.Instance(ImageTexture.CreateFromImage(image)).AsTexture2D())
	material := MyFirstShader{}
	shaders.Compile(&material)
	sprite.AsCanvasItem().SetMaterial(material.Super().AsMaterial())
	SceneTree.Add(sprite)
}
