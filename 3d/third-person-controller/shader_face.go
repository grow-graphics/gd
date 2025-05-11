package main

import (
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/Spatial"
	"graphics.gd/shaders/rgb"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/swizzle"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

type FaceShader struct {
	Spatial.Shader[FaceShader]

	FaceTexture       texture.Sampler2D[vec4.XYZW] `gd:",filter_nearest,repeat_disable"`
	Intensity         float.X
	ScreenColor       vec4.RGBA `gd:",source_color"`
	ScreenRedOffset   vec2.XY
	ScreenGreenOffset vec2.XY
	ScreenBlueOffset  vec2.XY
	PixelSize         float.X
}

func NewFaceShader() *FaceShader {
	return &FaceShader{
		Intensity: float.New(1.5),
		PixelSize: float.New(20.0),
	}
}

func (s *FaceShader) Material(frag Spatial.Fragment) Spatial.Material {
	var scaled_uv = vec2.Div(vec2.Mul(vec2.Round(frag.UV), s.PixelSize), s.PixelSize)
	var face_red = float.Sub(1.0, s.FaceTexture.Sample(vec2.Add(s.ScreenRedOffset, scaled_uv)).X)
	var face_green = float.Sub(1.0, s.FaceTexture.Sample(vec2.Add(s.ScreenGreenOffset, scaled_uv)).Y)
	var face_blue = float.Sub(1.0, s.FaceTexture.Sample(vec2.Add(s.ScreenBlueOffset, scaled_uv)).Z)
	var grid_2d = vec2.Sin(vec2.Mul(vec2.Fract(vec2.Add(vec2.Mul(frag.UV, s.PixelSize), 0.5)), 3.14))
	var grid = float.Mul(grid_2d.X, grid_2d.Y)
	var iris = vec3.RGB{
		R: float.Div(float.Mul(s.ScreenColor.R, face_red), 3.0),
		G: float.Div(float.Mul(s.ScreenColor.G, face_green), 3.0),
		B: float.Div(float.Mul(s.ScreenColor.B, face_blue), 3.0),
	}
	iris = rgb.Mul(iris, float.Sub(1.0, s.ScreenColor.A))
	var fill = rgb.Mul(swizzle.RGB(s.ScreenColor), s.ScreenColor.A)
	return Spatial.Material{
		Albedo:    rgba.New(0.0, 0.0, 0.0, 1.0),
		Specular:  float.New(0.25),
		Roughness: float.New(0.45),
		Emission:  rgb.Mul(rgb.Mul(rgb.Add(iris, fill), grid), s.Intensity),
	}
}
