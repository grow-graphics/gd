package main

import (
	"graphics.gd/shaders"
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/vec2"
)

type MyFirstShader struct {
	shaders.Type2D

	Offset vec2.XY `gd:"offset"`
}

func (MyFirstShader) Material(vert shaders.Fragment2D) shaders.Material2D {
	// smooth back and forth between pink and blue
	step := float.Mod(vert.Time, 2.0)
	step = bool.Mix(step, float.Sub(2.0, step), float.Gt(step, 1.0))
	return shaders.Material2D{
		Color: rgba.New(step, 0.6, 0.9, 1.0),
	}
}

type MyFirstShader2D struct {
	shaders.Type2D
}

func (uniform MyFirstShader2D) Material(vertex shaders.Fragment2D) shaders.Material2D {
	return shaders.Material2D{
		Color: rgba.New(0.4, 0.6, 0.9, 1.0),
	}
}

type MyFirstShader2D_UV struct {
	shaders.Type2D
}

func (uniform MyFirstShader2D_UV) Material(fragment shaders.Fragment2D) shaders.Material2D {
	return shaders.Material2D{
		Color: rgba.New(fragment.UV.X, fragment.UV.Y, 0.5, 1.0),
	}
}

type MyFirstShader2D_Texture struct {
	shaders.Type2D
}

func (uniform MyFirstShader2D_Texture) Material(fragment shaders.Fragment2D) shaders.Material2D {
	color := fragment.Texture.Texture(fragment.UV)
	color.B = float.New(1.0)
	return shaders.Material2D{
		Color: color,
	}
}
