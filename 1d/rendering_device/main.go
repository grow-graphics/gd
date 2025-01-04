package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/RenderingDevice"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
)

type Main struct {
	classdb.Extension[Main, SceneTree.Instance]
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
