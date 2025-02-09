package main

import (
	"log"

	"graphics.gd/classdb/AnimatedSprite2D"
	"graphics.gd/classdb/AtlasTexture"
	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/SpriteFrames"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

// graphics.gd equivalent to https://github.com/hajimehoshi/ebiten/blob/main/examples/animation/main.go
func main() {
	startup.LoadingScene()

	screen_size := DisplayServer.WindowGetSize()
	var img = Image.New()
	if err := img.LoadPngFromBuffer(images.Runner_png); err != nil {
		log.Fatal(err)
	}

	texture := ImageTexture.New()
	texture.SetImage(img)

	frames := SpriteFrames.New()
	frames.AddAnimation("runner")
	for i := 0; i < 8; i++ {
		frame := AtlasTexture.New()
		frame.SetAtlas(texture.AsTexture2D())
		frame.SetRegion(Rect2.New(i*32, 33, 32, 32))
		frames.AddFrame("runner", frame.AsTexture2D())
	}

	sprite := AnimatedSprite2D.New()
	sprite.AsNode2D().SetPosition(Vector2.New(screen_size.X/2, screen_size.Y/2))
	sprite.AsNode2D().SetScale(Vector2.New(4, 4))
	sprite.SetSpriteFrames(frames)
	sprite.SetSpeedScale(3)
	sprite.AsCanvasItem().SetTextureFilter(CanvasItem.TextureFilterNearest)
	sprite.SetAnimation("runner")
	sprite.Play()

	RenderingServer.SetDefaultClearColor(Color.W3C.Black)
	SceneTree.Add(sprite)

	startup.Scene()
}
