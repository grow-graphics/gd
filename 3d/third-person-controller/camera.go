package main

import (
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

type CameraMode struct {
	Node3D.Extension[CameraMode]

	CameraSpeed      Float.X
	MouseSensitivity Float.X

	camera       Camera3D.Instance
	cachedCamera Camera3D.Instance
	enabled      bool
}

func NewCameraMode() *CameraMode {
	return &CameraMode{CameraSpeed: 10, MouseSensitivity: 0.01}
}

func (cm *CameraMode) Ready() {
	if OS.IsDebugBuild() {
		cm.enabled = true
	}
	cm.AsNode().SetProcess(cm.enabled)
	cm.AsNode().SetProcessInput(cm.enabled)
}

func (cm *CameraMode) Input(event InputEvent.Instance) {
	if event, ok := Object.As[InputEventKey.Instance](event); ok {
		if event.AsInputEvent().IsPressed() && !event.AsInputEvent().IsEcho() {
			if event.Keycode() == InputEventKey.KeyF10 {
				cm.toggle_camera_mode()
			}
		}
	}
}

func (cm *CameraMode) Process(dt Float.X) {
	if !cm.AsNode3D().Visible() {
		return
	}
	var movement = Vector3.Zero
	if Input.IsKeyPressed(Input.KeyW) {
		movement = Vector3.Add(movement, Vector3.Forward)
	}
	if Input.IsKeyPressed(Input.KeyA) {
		movement = Vector3.Add(movement, Vector3.Left)
	}
	if Input.IsKeyPressed(Input.KeyS) {
		movement = Vector3.Add(movement, Vector3.Back)
	}
	if Input.IsKeyPressed(Input.KeyD) {
		movement = Vector3.Add(movement, Vector3.Right)
	}
	if Input.IsKeyPressed(Input.KeyQ) {
		movement = Vector3.Add(movement, Vector3.Down)
	}
	if Input.IsKeyPressed(Input.KeyE) {
		movement = Vector3.Add(movement, Vector3.Up)
	}
	var rotation_input = Angle.Radians(-Input.GetLastMouseVelocity().X * cm.MouseSensitivity)
	var tilt_input = Angle.Radians(-Input.GetLastMouseVelocity().Y * cm.MouseSensitivity)
	var global_transform = cm.camera.AsNode3D().GlobalTransform()
	var euler_rotation = Basis.AsEulerAngles(global_transform.Basis, Angle.OrderYXZ)
	euler_rotation.X += tilt_input * Angle.Radians(dt)
	euler_rotation.X = Float.Clamp(euler_rotation.X, -Angle.Pi+0.01, Angle.Pi-0.01)
	euler_rotation.Y += rotation_input * Angle.Radians(dt)
	global_transform.Basis = Basis.Euler(euler_rotation, Angle.OrderYXZ)
	cm.camera.AsNode3D().SetGlobalTransform(global_transform)
	global_position := cm.camera.AsNode3D().GlobalPosition()
	cm.camera.AsNode3D().SetGlobalPosition(
		Vector3.Add(global_position,
			Basis.Transform(
				Vector3.MulX(movement, dt*cm.CameraSpeed),
				global_transform.Basis,
			)))
}

func (cm *CameraMode) toggle_camera_mode() {
	var tree = SceneTree.Get(cm.AsNode())
	if cm.AsNode3D().Visible() {
		tree.SetPaused(false)
		cm.cachedCamera.SetCurrent(false)
		cm.camera.AsNode().QueueFree()
		cm.AsNode3D().Hide()
		for _, node := range tree.GetNodesInGroup("camera_mode_toggle") {
			if node3d, ok := Object.As[Node3D.Instance](Node.Instance(node)); ok {
				node3d.Show()
			}
		}
	} else {
		tree.SetPaused(true)
		cm.cachedCamera = Viewport.Get(cm.AsNode()).GetCamera3d()
		cm.camera = Camera3D.New()
		cm.AsNode().AddChild(cm.camera.AsNode())
		cm.camera.SetCurrent(true)
		cm.AsNode3D().Show()
		cm.camera.SetFov(cm.cachedCamera.Fov())
		cm.AsNode3D().SetGlobalTransform(cm.cachedCamera.AsNode3D().GlobalTransform())
		for _, node := range tree.GetNodesInGroup("camera_mode_toggle") {
			if node3d, ok := Object.As[Node3D.Instance](Node.Instance(node)); ok {
				node3d.Hide()
			}
		}
	}
}
