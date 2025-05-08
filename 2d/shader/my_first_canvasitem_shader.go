package main

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/CanvasItem"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/vec2"
)

type MyFirstShader struct {
	CanvasItem.Shader

	Offset vec2.XY `gd:"offset"`
}

func (MyFirstShader) Material(vert CanvasItem.Fragment) CanvasItem.Material {
	// smooth back and forth between pink and blue
	step := float.Mod(CanvasItem.Time, 2.0)
	step = bool.Mix(step, float.Sub(2.0, step), float.Gt(step, 1.0))
	return CanvasItem.Material{
		Color: rgba.New(step, 0.6, 0.9, 1.0),
	}
}

type MyFirstShader2D struct {
	CanvasItem.Shader
}

func (uniform MyFirstShader2D) Material(vertex CanvasItem.Fragment) CanvasItem.Material {
	return CanvasItem.Material{
		Color: rgba.New(0.4, 0.6, 0.9, 1.0),
	}
}

type MyFirstShader2D_UV struct {
	CanvasItem.Shader
}

func (uniform MyFirstShader2D_UV) Material(fragment CanvasItem.Fragment) CanvasItem.Material {
	return CanvasItem.Material{
		Color: rgba.New(fragment.UV.X, fragment.UV.Y, 0.5, 1.0),
	}
}

type MyFirstShader2D_Texture struct {
	CanvasItem.Shader
}

func (uniform MyFirstShader2D_Texture) Material(fragment CanvasItem.Fragment) CanvasItem.Material {
	color := fragment.Texture().Texture(fragment.UV)
	color.B = float.New(1.0)
	return CanvasItem.Material{
		Color: color,
	}
}
