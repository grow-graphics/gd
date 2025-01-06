package main

import (
	"fmt"

	"graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/RenderingDevice"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/startup"
)

type Main struct {
	MainLoop.Implementation
}

func (app *Main) Initialize() {
	var RD RenderingDevice.Instance = RenderingServer.GetRenderingDevice()

	fmt.Println(RD.GetDeviceName())
	fmt.Println(RD.GetDeviceVendorName())

	fmt.Println("VRAM: ", RD.GetMemoryUsage(RenderingDevice.MemoryTotal))
}

func main() {
	startup.MainLoop(new(Main))
}
