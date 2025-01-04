package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/Resource"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Vector2"
)

type HelloTriangle struct {
	classdb.Extension[HelloTriangle, MainLoop.Instance]

	viewport Resource.ID
	closing  bool

	canvas   Resource.ID
	triangle Resource.ID
}

func (app *HelloTriangle) Initialize() {
	app.viewport = RenderingServer.ViewportCreate()

	DisplayServer.WindowSetWindowEventCallback(app.OnWindowInputEvent)
	DisplayServer.WindowSetRectChangedCallback(app.OnResize)
	DisplayServer.WindowSetTitle("Hello Triangle")

	size := DisplayServer.WindowGetSize()

	RenderingServer.Advanced().ViewportAttachToScreen(app.viewport, Rect2.New(0, 0, size.X, size.Y), 0)
	RenderingServer.ViewportSetClearMode(app.viewport, RenderingServer.ViewportClearAlways)
	RenderingServer.ViewportSetActive(app.viewport, true)
	RenderingServer.ViewportSetSize(app.viewport, int(size.X), int(size.Y))

	app.canvas = RenderingServer.CanvasCreate()
	app.triangle = RenderingServer.CanvasItemCreate()

	RenderingServer.CanvasItemSetParent(app.triangle, app.canvas)
	RenderingServer.ViewportAttachCanvas(app.viewport, app.canvas)

}

func (app *HelloTriangle) OnWindowInputEvent(event DisplayServer.WindowEvent) {
	switch event {
	case DisplayServer.WindowEventCloseRequest:
		app.closing = true
	}
}

func (app *HelloTriangle) OnResize(rect Rect2i.PositionSize) {
	size := DisplayServer.WindowGetSize()

	RenderingServer.CanvasItemClear(app.triangle)
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
	RenderingServer.CanvasItemAddPolygon(app.triangle, points, colors)
}

func (app *HelloTriangle) Process(dt Float.X) bool {
	size := DisplayServer.WindowGetSize()
	RenderingServer.ViewportSetSize(app.viewport, int(size.X), int(size.Y))
	RenderingServer.Advanced().ViewportAttachToScreen(app.viewport, Rect2.New(0, 0, size.X, size.Y), 0)
	return app.closing
}

func (app *HelloTriangle) PhysicsProcess(dt Float.X) bool { return app.closing }

func (app *HelloTriangle) Finalize() {
	RenderingServer.FreeRid(app.viewport)
	RenderingServer.FreeRid(app.canvas)
	RenderingServer.FreeRid(app.triangle)
}

func main() {
	startup.MainLoop[HelloTriangle]()
}
