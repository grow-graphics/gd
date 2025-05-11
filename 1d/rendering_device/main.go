package main

import (
	"fmt"

	"graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/Rendering"
	"graphics.gd/classdb/RenderingDevice"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/startup"
)

type Main struct {
	MainLoop.Implementation
}

func (app *Main) Initialize() {
	fmt.Println(RenderingServer.GetVideoAdapterName())

	var RD RenderingDevice.Instance = RenderingServer.GetRenderingDevice()
	if RD == RenderingDevice.Nil {
		fmt.Println("no RenderingDevice available on this platform!")
		return
	}

	fmt.Println(RD.GetDeviceName())
	fmt.Println(RD.GetDeviceVendorName())

	fmt.Println("VRAM: ", RD.GetMemoryUsage(Rendering.MemoryTotal))
}

func main() {
	startup.MainLoop(new(Main))
}
