package main

import (
	"fmt"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
	"grow.graphics/rd"
)

type Main struct {
	gd.Class[Main, gd.SceneTree]

	RD rd.Interface
}

func (app *Main) Initialize() {
	app.RD = gd.RenderingDevice(app.KeepAlive)

	fmt.Println(app.RD.DeviceName())
	fmt.Println(app.RD.DeviceVendor())

	fmt.Println("VRAM: ", app.RD.MemoryUsage(rd.MemoryTotal))
}

func main() {
	godot, ok := gdextension.Link()
	if ok {
		gd.Register[Main](godot)
	}
}
