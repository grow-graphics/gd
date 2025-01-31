package main

import (
	"graphics.gd/shaders"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/vec4"
)

type MyFirstShader2D struct {
	shaders.Type2D
}

func (uniform MyFirstShader2D) Fragment(vertex shaders.Vertex2D) shaders.Fragment2D {
	return shaders.Fragment2D{
		Position: vertex.Position,
	}
}

func (uniform MyFirstShader2D) Material(vertex shaders.Fragment2D) shaders.Material2D {
	return shaders.Material2D{
		Color: rgba.New(0.4, 0.6, 0.9, 1.0),
	}
}

func (uniform MyFirstShader2D) Lighting(material shaders.Material2D) vec4.RGBA { return material.Color }

type MyFirstShader2D_UV struct {
	shaders.Type2D
}

func (uniform MyFirstShader2D_UV) Fragment(vertex shaders.Vertex2D) shaders.Fragment2D {
	return shaders.Fragment2D{
		Position: vertex.Position,
		UV:       vertex.UV,
	}
}

func (uniform MyFirstShader2D_UV) Material(fragment shaders.Fragment2D) shaders.Material2D {
	return shaders.Material2D{
		Color: rgba.New(fragment.UV.X, fragment.UV.Y, 0.5, 1.0),
	}
}

func (uniform MyFirstShader2D_UV) Lighting(material shaders.Material2D) vec4.RGBA {
	return material.Color
}

type MyFirstShader2D_Texture struct {
	shaders.Type2D
}

func (uniform MyFirstShader2D_Texture) Fragment(vertex shaders.Vertex2D) shaders.Fragment2D {
	return shaders.Fragment2D{
		Position: vertex.Position,
		UV:       vertex.UV,
	}
}

func (uniform MyFirstShader2D_Texture) Material(fragment shaders.Fragment2D) shaders.Material2D {
	color := fragment.Texture.Texture(fragment.UV)
	color.B = float.New(1.0)
	return shaders.Material2D{
		Color: color,
	}
}

func (uniform MyFirstShader2D_Texture) Lighting(material shaders.Material2D) vec4.RGBA {
	return material.Color
}
