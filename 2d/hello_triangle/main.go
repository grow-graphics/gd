package main

import (
	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type HelloTriangle struct {
	gd.Class[HelloTriangle, gd.MainLoop]

	viewport gd.RID
	closing  gd.Bool

	canvas   gd.RID
	triangle gd.RID
}

func (app *HelloTriangle) Initialize(godot gd.Context) {
	RS := gd.RenderingServer(godot)
	app.viewport = RS.ViewportCreate()

	DS := gd.DisplayServer(godot)
	DS.WindowSetWindowEventCallback(godot.Callable(app.OnWindowInputEvent), 0)
	DS.WindowSetTitle(godot.String("Hello Triangle"), 0)

	size := DS.WindowGetSize(0)

	RS.ViewportAttachToScreen(app.viewport, gd.NewRect2(0, 0, gd.Float(size.X()), gd.Float(size.Y())), 0)
	RS.ViewportSetClearMode(app.viewport, gd.RenderingServerViewportClearAlways)
	RS.ViewportSetActive(app.viewport, true)
	RS.ViewportSetSize(app.viewport, size.X(), size.Y())

	app.canvas = RS.CanvasCreate()
	app.triangle = RS.CanvasItemCreate()

	RS.CanvasItemSetParent(app.triangle, app.canvas)

	RS.ViewportAttachCanvas(app.viewport, app.canvas)
}

func (app *HelloTriangle) OnWindowInputEvent(godot gd.Context, event gd.DisplayServerWindowEvent) {
	switch event {
	case gd.DisplayServerWindowEventCloseRequest:
		app.closing = true
	}
}

func (app *HelloTriangle) Process(godot gd.Context, delta gd.Float) gd.Bool {
	RS := gd.RenderingServer(godot)
	DS := gd.DisplayServer(godot)
	size := DS.WindowGetSize(0)
	RS.ViewportSetSize(app.viewport, size.X(), size.Y())
	RS.ViewportAttachToScreen(app.viewport, gd.NewRect2(0, 0, gd.Float(size.X()), gd.Float(size.Y())), 0)
	RS.CanvasItemClear(app.triangle)
	var points = godot.PackedVector2Array()
	points.Append(gd.NewVector2(0, 0))
	points.Append(gd.NewVector2(gd.Float(size.X()), 0))
	points.Append(gd.NewVector2(gd.Float(size.X())/2, gd.Float(size.Y())))

	var colors = godot.PackedColorArray()
	colors.Append(gd.Color{1, 0, 0, 1})
	colors.Append(gd.Color{0, 1, 0, 1})
	colors.Append(gd.Color{0, 0, 1, 1})

	RS.CanvasItemAddPolygon(app.triangle, points, colors, gd.PackedVector2Array{}, 0)
	return app.closing
}

func (app *HelloTriangle) PhysicsProcess(godot gd.Context, delta gd.Float) gd.Bool {
	return app.closing
}

func (app *HelloTriangle) Finalize(godot gd.Context) {
	RS := gd.RenderingServer(godot)
	RS.FreeRid(app.viewport)
	RS.FreeRid(app.canvas)
	RS.FreeRid(app.triangle)
}

func main() {
	godot, ok := gdextension.Link()
	if ok {
		gd.Register[HelloTriangle](godot)
	}
}
