package main

import (
	"fmt"

	"graphics.gd/defined"
	"graphics.gd/objects/RenderingDevice"
	"graphics.gd/objects/RenderingServer"
	"graphics.gd/objects/SceneTree"
	"graphics.gd/startup"
)

type Main struct {
	defined.Object[Main, SceneTree.Instance]
}

func (app *Main) Initialize() {
	var RD RenderingDevice.Instance = RenderingServer.GetRenderingDevice()

	fmt.Println(RD.GetDeviceName())
	fmt.Println(RD.GetDeviceVendorName())

	fmt.Println("VRAM: ", RD.GetMemoryUsage(RenderingDevice.MemoryTotal))
}

func main() {
	startup.MainLoop[Main]()
}
