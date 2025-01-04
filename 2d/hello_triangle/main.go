package main

import (
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/Resource"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Vector2"
)

var (
	closing  bool
	viewport Resource.ID
	canvas   Resource.ID
	triangle Resource.ID
)

func onWindowInputEvent(event DisplayServer.WindowEvent) {
	switch event {
	case DisplayServer.WindowEventCloseRequest:
		closing = true
	}
}

func onResize(rect Rect2i.PositionSize) {
	size := DisplayServer.WindowGetSize()

	RenderingServer.CanvasItemClear(triangle)
	var points = []Vector2.XY{
		Vector2.New(0, 0),
		Vector2.New(size.X, 0),
		Vector2.New(size.X/2, size.Y),
	}
	var colors = []Color.RGBA{
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
	}
	RenderingServer.CanvasItemAddPolygon(triangle, points, colors)
}

func main() {
	frames := startup.Rendering()

	viewport = RenderingServer.ViewportCreate()
	defer RenderingServer.FreeRid(viewport)

	DisplayServer.WindowSetWindowEventCallback(onWindowInputEvent)
	DisplayServer.WindowSetRectChangedCallback(onResize)
	DisplayServer.WindowSetTitle("Hello Triangle")

	size := DisplayServer.WindowGetSize()

	RenderingServer.Advanced().ViewportAttachToScreen(viewport, Rect2.New(0, 0, size.X, size.Y), 0)
	RenderingServer.ViewportSetClearMode(viewport, RenderingServer.ViewportClearAlways)
	RenderingServer.ViewportSetActive(viewport, true)
	RenderingServer.ViewportSetSize(viewport, int(size.X), int(size.Y))

	canvas = RenderingServer.CanvasCreate()
	triangle = RenderingServer.CanvasItemCreate()
	defer RenderingServer.FreeRid(canvas)
	defer RenderingServer.FreeRid(triangle)

	RenderingServer.CanvasItemSetParent(triangle, canvas)
	RenderingServer.ViewportAttachCanvas(viewport, canvas)

	onResize(Rect2i.New(0, 0, size.X, size.Y))

	for range frames {
		size := DisplayServer.WindowGetSize()
		RenderingServer.ViewportSetSize(viewport, int(size.X), int(size.Y))
		RenderingServer.Advanced().ViewportAttachToScreen(viewport, Rect2.New(0, 0, size.X, size.Y), 0)
		if closing {
			break
		}
	}
}
