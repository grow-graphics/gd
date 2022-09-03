package main

import (
	"fmt"

	"github.com/readykit/gd"
)

func main() {}

type HelloWorld struct {
	object gd.Object

	Message string
}

var NewHelloWorld = gd.Register(func(ctx gd.Context, hello *HelloWorld) gd.Object {
	hello.Message = "Hello World!"
	return gd.NewObject(ctx, &hello.object)
})

func (h HelloWorld) Print(ctx gd.Context) {
	fmt.Println(h.Message)
}

func (h HelloWorld) Echo(ctx gd.Context, s string) {
	fmt.Println(s + " from Go!")
}

type ExtendedNode struct {
	ctx gd.Context

	node gd.Node2D
}

var NewExtendedNode = gd.Register(func(ctx gd.Context, extended *ExtendedNode) gd.Node2D {
	extended.ctx = ctx
	return gd.NewNode2D(ctx, &extended.node)
})

func (e *ExtendedNode) Ready() {
	if gd.Engine.IsEditorHint() {
		fmt.Println(gd.Engine.GetLicenseText())
		return
	}

	fmt.Println(e.node.CanvasItem().Node().Object().GetClass())

	var obj = gd.NewObject(e.ctx, nil)
	fmt.Println(obj.GetClass())
	defer obj.Free(e.ctx)

	//gd.LoadSingletons()

	fmt.Println(gd.Engine.GetSingletonList())

	fmt.Println("Scene is ready!")

	fmt.Println("sin=", gd.Sin(1.5))

	fmt.Println("rotation=", e.node.GetRotation())
	e.node.SetRotation(3.14)
	fmt.Println("rotation=", e.node.GetRotation())

	pos := e.node.GetPosition()

	fmt.Println("position=", pos)

	pos.X = 100

	e.node.SetPosition(pos)
	fmt.Println("position=", pos)

	//gd.DisplayServerSingleton(gd.Engine.GetSingleton("DisplayServer"))

	//fmt.Println(godot.Engine.GetSingletonList())

	fmt.Println(gd.DisplayServer)

	gd.DisplayServer.WindowSetCurrentScreen(1, 0)
}
