package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimationNodeOneShot"
	"graphics.gd/classdb/AnimationNodeStateMachinePlayback"
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/AnimationTree"
	"graphics.gd/classdb/AudioStreamPlayer3D"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/CharacterBody3D"
	"graphics.gd/classdb/ColorRect"
	"graphics.gd/classdb/HBoxContainer"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/InputMap"
	"graphics.gd/classdb/Marker3D"
	"graphics.gd/classdb/Mesh"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/PhysicsServer3D"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/RayCast3D"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/ShapeCast3D"
	"graphics.gd/classdb/SpringArm3D"
	"graphics.gd/classdb/SurfaceTool"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
)

type WeaponType Enum.Int[struct {
	Default WeaponType
	Grenade WeaponType
}]

var WeaponTypes = Enum.Values[WeaponType]()

type DemoPlayer struct {
	CharacterBody3D.Extension[DemoPlayer] `gd:"DemoPlayer"`

	WeaponSwitched Signal.Solo[string] `gd:"weapon_switched(weapon_name)"`

	MoveSpeed Float.X `
		is the maximum run speed on the ground.`
	BulletSpeed         Float.X
	AttackImpulse       Float.X
	Acceleration        Float.X
	JumpInitialImpulse  Float.X
	JumpAdditionalForce Float.X
	RotationSpeed       Float.X
	StoppingSpeed       Float.X
	MaxThrowbackForce   Float.X
	ShootCooldown       Float.X
	GrenadeCooldown     Float.X

	RotationRoot          Node3D.Instance          `gd:"CharacterRotationRoot"`
	CameraController      *CameraController        `gd:"CameraController"`
	AttackAnimationPlayer AnimationPlayer.Instance `gd:"CharacterRotationRoot/MeleeAnchor/AnimationPlayer"`
	GroundShapecast       ShapeCast3D.Instance     `gd:"GroundShapeCast"`
	GrenadeAimController  *GrenadeLauncher         `gd:"GrenadeLauncher"`
	CharacterSkin         *CharacterSkin           `gd:"CharacterRotationRoot/CharacterSkin"`

	UI struct {
		AimReticle     ColorRect.Instance     `gd:"%AimRecticle"`
		CoinsContainer HBoxContainer.Instance `gd:"%CoinsContainer"`
	}

	StepSound    AudioStreamPlayer3D.Instance `gd:"StepSound"`
	LandingSound AudioStreamPlayer3D.Instance `gd:"LandingSound"`

	equippedWeapon      WeaponType
	moveDirection       Vector3.XYZ
	lastStrongDirection Vector3.XYZ
	gravity             Float.X
	groundHeight        Float.X
	startPosition       Vector3.XYZ
	coins               int
	isOnFloorBuffer     bool

	shootCooldownTick   Float.X
	grenadeCooldownTick Float.X
}

func NewDemoPlayer() *DemoPlayer {
	return &DemoPlayer{
		MoveSpeed:           8,
		BulletSpeed:         10,
		AttackImpulse:       10,
		Acceleration:        4,
		JumpInitialImpulse:  12,
		JumpAdditionalForce: 4.5,
		RotationSpeed:       12,
		StoppingSpeed:       1,
		MaxThrowbackForce:   15,
		ShootCooldown:       0.5,
		GrenadeCooldown:     0.5,

		lastStrongDirection: Vector3.Forward,
		gravity:             -30,
	}
}

func (p *DemoPlayer) Ready() {
	p.startPosition = p.AsNode3D().GlobalTransform().Origin
	p.shootCooldownTick = p.ShootCooldown
	p.grenadeCooldownTick = p.GrenadeCooldown
	Input.SetMouseMode(Input.MouseModeCaptured)
	p.CameraController.Setup(p)
	p.GrenadeAimController.AsNode3D().SetVisible(false)
	p.WeaponSwitched.Emit(p.equippedWeapon.String())
	if !InputMap.HasAction("move_left") {
		p.register_input_actions()
	}
	p.CharacterSkin.Stepped.Attach(Callable.New(func() {
		p.StepSound.SetPitchScale(Float.RandomlyDistributed(1.2, 0.2))
		p.StepSound.Play()
	}))
}

func (p *DemoPlayer) PhysicsProcess(delta Float.X) {
	if p.GroundShapecast.GetCollisionCount() > 0 {
		for _, collision_result := range p.GroundShapecast.CollisionResult() {
			p.groundHeight = max(p.groundHeight, collision_result.Point.Y)
		}
	} else {
		p.groundHeight = p.AsNode3D().GlobalPosition().Y + p.GroundShapecast.TargetPosition().Y
	}
	if p.AsNode3D().GlobalPosition().Y < p.groundHeight {
		p.groundHeight = p.AsNode3D().GlobalPosition().Y
	}
	if Input.IsActionJustPressed("swap_weapons", false) {
		if p.equippedWeapon == WeaponTypes.Default {
			p.equippedWeapon = WeaponTypes.Grenade
			p.GrenadeAimController.AsNode3D().SetVisible(true)
		} else {
			p.equippedWeapon = WeaponTypes.Default
		}
		p.WeaponSwitched.Emit(p.equippedWeapon.String())
	}
	var is_attacking = Input.IsActionPressed("attack", false) && !p.AttackAnimationPlayer.IsPlaying()
	var is_just_attacking = Input.IsActionJustPressed("attack", false)
	var is_just_jumping = Input.IsActionJustPressed("jump", false)
	var is_aiming = Input.IsActionPressed("aim", false) && p.AsCharacterBody3D().IsOnFloor()
	var is_air_boosting = Input.IsActionPressed("jump", false) && !p.AsCharacterBody3D().IsOnFloor() && p.AsCharacterBody3D().Velocity().Y > 0
	var is_just_on_floor = p.AsCharacterBody3D().IsOnFloor() && !p.isOnFloorBuffer
	p.isOnFloorBuffer = p.AsCharacterBody3D().IsOnFloor()
	p.moveDirection = p.get_camera_oriented_input()
	if Vector3.Length(p.moveDirection) > 0.2 {
		p.lastStrongDirection = Vector3.Normalized(p.moveDirection)
	}
	if is_aiming {
		p.lastStrongDirection = Vector3.Normalized(Basis.Transform(Vector3.Back, p.CameraController.AsNode3D().GlobalTransform().Basis))
	}
	p.orient_character_to_direction(p.lastStrongDirection, delta)
	var y_velocity = p.AsCharacterBody3D().Velocity().Y
	var velocity = p.AsCharacterBody3D().Velocity()
	velocity.Y = 0
	velocity = Vector3.Lerp(velocity, Vector3.MulX(p.moveDirection, p.MoveSpeed), p.Acceleration*delta)
	if Vector3.Length(p.moveDirection) == 0 && Vector3.Length(velocity) < p.StoppingSpeed {
		velocity = Vector3.Zero
	}
	velocity.Y = y_velocity
	velocity.Y += p.gravity * delta
	if is_just_jumping {
		velocity.Y += p.JumpInitialImpulse
	} else if is_air_boosting {
		velocity.Y += p.JumpAdditionalForce * delta
	}
	p.AsCharacterBody3D().SetVelocity(velocity)
	if is_aiming {
		p.CameraController.SetPivot(CameraPivots.OverShoulder)
		p.GrenadeAimController.ThrowDirection = Quaternion.Rotate(Vector3.Forward, p.CameraController.Camera.AsNode3D().Quaternion())
		p.GrenadeAimController.FromLookPosition = p.CameraController.Camera.AsNode3D().GlobalPosition()
		p.UI.AimReticle.AsCanvasItem().SetVisible(true)
	} else {
		p.CameraController.SetPivot(CameraPivots.ThirdPerson)
		p.GrenadeAimController.ThrowDirection = p.lastStrongDirection
		p.GrenadeAimController.FromLookPosition = p.AsNode3D().GlobalPosition()
		p.UI.AimReticle.AsCanvasItem().SetVisible(false)
	}
	p.shootCooldownTick += delta
	p.grenadeCooldownTick += delta
	if is_attacking {
		switch p.equippedWeapon {
		case WeaponTypes.Default:
			if is_aiming && p.AsCharacterBody3D().IsOnFloor() {
				if p.shootCooldownTick > p.ShootCooldown {
					p.shootCooldownTick = 0
					p.shoot()
				}
			} else if is_just_attacking {
				p.attack()
			}
		case WeaponTypes.Grenade:
			if p.grenadeCooldownTick > p.GrenadeCooldown {
				p.grenadeCooldownTick = 0
				p.GrenadeAimController.ThrowGrenade()
			}
		}
	}
	if is_just_jumping {
		p.CharacterSkin.Jump()
	} else if is_just_on_floor && velocity.Y < 0 {
		p.CharacterSkin.Fall()
	} else if p.AsCharacterBody3D().IsOnFloor() {
		var xz_velocity = Vector3.New(velocity.X, 0, velocity.Z)
		if Vector3.Length(xz_velocity) > p.StoppingSpeed {
			p.CharacterSkin.SetMoving(true)
			p.CharacterSkin.SetMovingSpeed(Float.InverseLerp(0, p.MoveSpeed, Vector3.Length(xz_velocity)))
		} else {
			p.CharacterSkin.SetMoving(false)
		}
	}
	if is_just_on_floor {
		p.LandingSound.Play()
	}
	var position_before = p.AsNode3D().GlobalPosition()
	p.AsCharacterBody3D().MoveAndSlide()
	var position_after = p.AsNode3D().GlobalPosition()
	var delta_position = Vector3.Sub(position_after, position_before)
	if Vector3.Length(delta_position) < Float.Epsilon && Vector3.Length(velocity) > Float.Epsilon {
		position := p.AsNode3D().GlobalPosition()
		p.AsNode3D().SetGlobalPosition(Vector3.Add(position, Vector3.MulX(p.AsCharacterBody3D().GetWallNormal(), 0.1)))
	}
}

func (p *DemoPlayer) attack() {
	p.AttackAnimationPlayer.PlayNamed("Attack")
	p.CharacterSkin.Punch()
	p.AsCharacterBody3D().SetVelocity(Basis.Transform(Vector3.MulX(Vector3.Back, p.AttackImpulse), p.CameraController.AsNode3D().GlobalTransform().Basis))
}

func (p *DemoPlayer) shoot() {
	var bullet = BulletScene.Instantiate()
	bullet.Shooter = p.AsNode().ID()
	var origin = Vector3.Add(p.AsNode3D().GlobalPosition(), Vector3.Up)
	var aim_target = p.CameraController.GetAimTarget()
	var aim_direction = Vector3.Normalized(Vector3.Sub(aim_target, origin))
	bullet.Velocity = Vector3.MulX(aim_direction, p.BulletSpeed)
	bullet.DistanceLimit = 14
	p.AsNode().AddChild(bullet.AsNode())
	bullet.AsNode3D().SetGlobalPosition(origin)
}

func (p *DemoPlayer) ResetPosition() {
	tansform := p.AsNode3D().Transform()
	tansform.Origin = p.startPosition
	p.AsNode3D().SetTransform(tansform)
}

func (p *DemoPlayer) Body3D() RID.Body3D {
	return p.AsCollisionObject3D().GetRid()
}

func (p *DemoPlayer) CollectCoin() {
	p.coins++
	Object.Call(p.UI.CoinsContainer, "update_coins_amount", p.coins)
}

func (p *DemoPlayer) loseCoins() {
	var lost_coins = min(5, p.coins)
	p.coins -= lost_coins
	for range lost_coins {
		var coin = CoinScene.Instantiate()
		p.AsNode().AddChild(coin.AsNode())
		coin.AsNode3D().SetGlobalPosition(p.AsNode3D().GlobalPosition())
		coin.SpawnWithDelay(1.5)
	}
	Object.Call(p.UI.CoinsContainer, "update_coins_amount", p.coins)
}

func (p *DemoPlayer) get_camera_oriented_input() Vector3.XYZ {
	if p.AttackAnimationPlayer.IsPlaying() {
		return Vector3.Zero
	}
	var raw_input = Input.GetVector("move_left", "move_right", "move_up", "move_down")
	var input = Vector3.Zero
	input.X = -raw_input.X * Float.Sign(1.0-raw_input.Y*raw_input.Y/2)
	input.Z = -raw_input.Y * Float.Sign(1.0-raw_input.X*raw_input.X/2)
	input = Basis.Transform(input, p.CameraController.AsNode3D().GlobalTransform().Basis)
	input.Y = 0
	return input
}

func (p *DemoPlayer) play_foot_step_sound() {
	p.StepSound.SetPitchScale(Float.RandomlyDistributed(1.2, 0.2))
	p.StepSound.Play()
}

func (p *DemoPlayer) Damage(impact_point Vector3.XYZ, force Vector3.XYZ) {
	force.Y = Float.Abs(force.Y)
	p.AsCharacterBody3D().SetVelocity(Vector3.LengthLimited(force, p.MaxThrowbackForce))
	p.loseCoins()
}

func (p *DemoPlayer) orient_character_to_direction(direction Vector3.XYZ, delta Float.X) {
	var left_axis = Vector3.Cross(Vector3.Up, direction)
	var rotation_basis = Basis.AsQuaternion(Basis.XYZ{X: left_axis, Y: Vector3.Up, Z: direction})
	var model_scale = Basis.Scale(p.RotationRoot.Transform().Basis)
	var transform = p.RotationRoot.Transform()
	transform.Basis = Basis.Scaled(Quaternion.AsBasis(Quaternion.Slerp(Basis.AsQuaternion(transform.Basis), rotation_basis, delta*p.RotationSpeed)), model_scale)
	p.RotationRoot.SetTransform(transform)
}

func (p *DemoPlayer) register_input_actions() {
	var InputActions = map[string]any{
		"move_left":    Input.KeyA,
		"move_right":   Input.KeyD,
		"move_up":      Input.KeyW,
		"move_down":    Input.KeyS,
		"jump":         Input.KeySpace,
		"attack":       Input.MouseButtonLeft,
		"aim":          Input.MouseButtonRight,
		"swap_weapons": Input.KeyTab,
		"pause":        Input.KeyEscape,
		"camera_left":  Input.KeyQ,
		"camera_right": Input.KeyE,
		"camera_up":    Input.KeyR,
		"camera_down":  Input.KeyF,
	}
	for action, control := range InputActions {
		if InputMap.HasAction(action) {
			continue
		}
		InputMap.AddAction(action)
		switch control := control.(type) {
		case Input.Key:
			var input_key = InputEventKey.New()
			input_key.SetKeycode(control)
			InputMap.ActionAddEvent(action, input_key.AsInputEvent())
		case Input.MouseButton:
			var input_mouse = InputEventMouseButton.New()
			input_mouse.SetButtonIndex(control)
			InputMap.ActionAddEvent(action, input_mouse.AsInputEvent())
		}
	}
}

type CameraPivot Enum.Int[struct {
	OverShoulder CameraPivot
	ThirdPerson  CameraPivot
}]

var CameraPivots = Enum.Values[CameraPivot]()

type CameraController struct {
	Node3D.Extension[CameraController] `gd:"CameraController"`

	PlayerPath          Path.ToNode
	InvertMouseY        bool
	MouseSensitivity    Float.X `range:"0.1,10"`
	JoystickSensitivity Float.X `range:"0,8.0"`
	TiltUpperLimit      Angle.Radians
	TiltLowerLimit      Angle.Radians

	Camera            Camera3D.Instance    `gd:"PlayerCamera"`
	OverShoulderPivot Marker3D.Instance    `gd:"CameraOverShoulderPivot"`
	CameraSpringArm   SpringArm3D.Instance `gd:"CameraSpringArm"`
	ThirdPersonPivot  Marker3D.Instance    `gd:"CameraSpringArm/CameraThirdPersonPivot"`
	CameraRaycast     RayCast3D.Instance   `gd:"PlayerCamera/CameraRayCast"`

	aim_target         Vector3.XYZ
	aim_collider       Node.ID
	pivot              Node3D.ID
	current_pivot_type CameraPivot
	rotation_input     Float.X
	tilt_input         Float.X
	mouse_input        bool
	offset             Vector3.XYZ
	anchor             *DemoPlayer
	euler_rotation     Euler.Radians
}

func NewCameraController() *CameraController {
	return &CameraController{
		MouseSensitivity:    0.25,
		JoystickSensitivity: 2,
		TiltUpperLimit:      Angle.InRadians(-60),
		TiltLowerLimit:      Angle.InRadians(60),
	}
}

func (p *CameraController) UnhandledInput(event InputEvent.Instance) {
	p.mouse_input = Object.Is[InputEventMouseMotion.Instance](event) && Input.MouseMode() == Input.MouseModeCaptured
	if p.mouse_input {
		event := Object.To[InputEventMouseMotion.Instance](event)
		p.rotation_input = -event.AsInputEventMouseMotion().Relative().X * p.MouseSensitivity
		p.tilt_input = -event.AsInputEventMouseMotion().Relative().Y * p.MouseSensitivity
	}
}

func (p *CameraController) Process(delta Float.X) {
	if p.anchor == nil {
		return
	}
	p.rotation_input += Input.GetActionRawStrength("camera_left", false) - Input.GetActionRawStrength("camera_right", false)
	p.tilt_input += Input.GetActionRawStrength("camera_up", false) - Input.GetActionRawStrength("camera_down", false)
	if p.InvertMouseY {
		p.tilt_input *= -1
	}
	if p.CameraRaycast.IsColliding() {
		p.aim_target = p.CameraRaycast.GetCollisionPoint()
		collider := p.CameraRaycast.GetCollider()
		if collider != Object.Nil {
			p.aim_collider = Node.ID(collider.ID())
		}
	} else {
		p.aim_target = Transform3D.Transform(p.CameraRaycast.TargetPosition(), p.CameraRaycast.AsNode3D().GlobalTransform())
		p.aim_collider = 0
	}
	var target_position = Vector3.Add(p.anchor.AsNode3D().GlobalPosition(), p.offset)
	target_position.Y = Float.Lerp(p.anchor.AsNode3D().GlobalPosition().Y, p.anchor.groundHeight, 0.1)
	p.AsNode3D().SetGlobalPosition(target_position)
	p.euler_rotation.X += Angle.Radians(p.tilt_input * delta)
	p.euler_rotation.X = Float.Clamp(p.euler_rotation.X, p.TiltLowerLimit, p.TiltUpperLimit)
	p.euler_rotation.Y += Angle.Radians(p.rotation_input * delta)
	transform := p.AsNode3D().Transform()
	transform.Basis = Basis.FromEuler(p.euler_rotation, Angle.OrderYXZ)
	p.AsNode3D().SetTransform(transform)
	if pivot, ok := p.pivot.Instance(); ok {
		p.Camera.AsNode3D().SetGlobalTransform(pivot.GlobalTransform())
	}
	rotation := p.Camera.AsNode3D().Rotation()
	rotation.Z = 0
	p.Camera.AsNode3D().SetRotation(rotation)
	p.rotation_input = 0
	p.tilt_input = 0
}

func (p *CameraController) Setup(player *DemoPlayer) {
	p.anchor = player
	p.AsNode3D().SetGlobalTransform(p.anchor.AsNode3D().AsNode3D().GlobalTransform())
	p.offset = Vector3.Sub(p.AsNode3D().GlobalTransform().Origin, player.AsNode3D().GlobalTransform().Origin)
	p.SetPivot(CameraPivots.ThirdPerson)
	if pivot, ok := p.pivot.Instance(); ok {
		p.Camera.AsNode3D().SetGlobalTransform(Transform3D.Lerp(p.Camera.AsNode3D().GlobalTransform(), pivot.GlobalTransform(), 0.1))
	}
	p.CameraSpringArm.AddExcludedObject(player.AsCollisionObject3D().GetRid())
	p.CameraRaycast.AddExceptionRid(player.AsCollisionObject3D().GetRid())
}

func (p *CameraController) SetPivot(pivot CameraPivot) {
	if p.current_pivot_type == pivot {
		return
	}
	p.current_pivot_type = pivot
	switch pivot {
	case CameraPivots.OverShoulder:
		p.OverShoulderPivot.AsNode3D().LookAt(p.aim_target)
		p.pivot = p.OverShoulderPivot.AsNode3D().ID()
	case CameraPivots.ThirdPerson:
		p.pivot = p.ThirdPersonPivot.AsNode3D().ID()
	}
	p.current_pivot_type = pivot
}

func (p *CameraController) GetAimTarget() Vector3.XYZ {
	return p.aim_target
}

func (p *CameraController) GetAimCollider() Node.Instance {
	node, _ := p.aim_collider.Instance()
	return node
}

var GrenadeScene = Resource.Load[PackedScene.Is[CharacterBody3D.Instance]]("res://Player/Grenade.tscn")

type GrenadeLauncher struct {
	Node3D.Extension[GrenadeLauncher] `gd:"GrenadeLauncher"`

	MinThrowDistance Float.X
	MaxThrowDistance Float.X
	Gravity          Float.X

	FromLookPosition Vector3.XYZ
	ThrowDirection   Vector3.XYZ

	SnapMesh          Node3D.Instance         `gd:"%SnapMesh"`
	Raycast           ShapeCast3D.Instance    `gd:"%ShapeCast3D"`
	LaunchPoint       Marker3D.Instance       `gd:"%LaunchPoint"`
	TrailMeshInstance MeshInstance3D.Instance `gd:"%TrailMeshInstance"`

	throwVelocity Vector3.XYZ
	timeToLand    Float.X
}

func NewGrenadeLauncher() *GrenadeLauncher {
	return &GrenadeLauncher{
		MinThrowDistance: 7,
		MaxThrowDistance: 16,
		Gravity:          Float.X(ProjectSettings.GetSetting("physics/3d/default_gravity", Float.X(0)).(float64)),
	}
}

func (weapon *GrenadeLauncher) PhysicsProcess(delta Float.X) {
	if weapon.AsNode3D().Visible() {
		weapon.update_throw_velocity()
		weapon.draw_throw_path()
	}
}

func (weapon *GrenadeLauncher) ThrowGrenade() bool {
	if !weapon.AsNode3D().Visible() {
		return false
	}
	var grenade = GrenadeScene.Instantiate()
	weapon.AsNode().GetParent().AddChild(grenade.AsNode())
	grenade.AsNode3D().SetGlobalPosition(weapon.LaunchPoint.AsNode3D().GlobalPosition())
	Object.Call(grenade, "throw", weapon.throwVelocity)
	parent := Object.To[CharacterBody3D.Instance](grenade.AsNode().GetParent())
	PhysicsServer3D.BodyAddCollisionException(parent.AsCollisionObject3D().GetRid(), grenade.AsCollisionObject3D().GetRid())
	return true
}

func (weapon *GrenadeLauncher) update_throw_velocity() {
	var camera = Viewport.Get(weapon.AsNode()).GetCamera3d()
	var up_ratio = Float.Clamp(max(camera.AsNode3D().Rotation().X+0.5, -0.4)*2, 0, 1)
	var base_throw_distance = Float.Lerp(weapon.MinThrowDistance, weapon.MaxThrowDistance, Float.X(up_ratio))
	var throw_distance = base_throw_distance
	var global_camera_look_position = Vector3.Add(weapon.FromLookPosition, Vector3.MulX(weapon.ThrowDirection, throw_distance))
	weapon.Raycast.SetTargetPosition(Vector3.Sub(global_camera_look_position, weapon.Raycast.AsNode3D().GlobalPosition()))
	var to_target = weapon.Raycast.TargetPosition()
	if weapon.Raycast.GetCollisionCount() != 0 {
		if node, ok := Object.As[Node3D.Instance](weapon.Raycast.GetCollider(0)); ok {
			var has_target = node != Node3D.Nil && node.AsNode().IsInGroup("targeteables")
			weapon.SnapMesh.SetVisible(has_target)
			if has_target {
				to_target = Vector3.Sub(node.GlobalPosition(), weapon.LaunchPoint.AsNode3D().GlobalPosition())
				weapon.SnapMesh.SetGlobalPosition(Vector3.Add(weapon.LaunchPoint.AsNode3D().GlobalPosition(), to_target))
				weapon.SnapMesh.LookAt(weapon.LaunchPoint.AsNode3D().GlobalPosition())
			}
		}
	} else {
		weapon.SnapMesh.SetVisible(false)
	}
	var peak_height = max(to_target.Y+0.25, weapon.LaunchPoint.AsNode3D().Position().Y+0.25)
	var motion_up = peak_height
	var time_going_up = Float.Sqrt(2 * motion_up / weapon.Gravity)
	var motion_down = to_target.Y - peak_height
	var time_going_down = Float.Sqrt(2 * motion_down / weapon.Gravity)
	weapon.timeToLand = time_going_up + time_going_down
	var target_position_xz_plane = Vector3.New(to_target.X, 0, to_target.Z)
	var start_position_xz_plane = Vector3.New(weapon.LaunchPoint.AsNode3D().Position().X, 0, weapon.LaunchPoint.AsNode3D().Position().Z)
	var forward_velocity = Vector3.DivX(Vector3.Sub(target_position_xz_plane, start_position_xz_plane), weapon.timeToLand)
	var velocity_up = Float.Sqrt(2 * weapon.Gravity * motion_up)
	weapon.throwVelocity = Vector3.Add(Vector3.MulX(Vector3.Up, velocity_up), forward_velocity)
}

func (weapon *GrenadeLauncher) draw_throw_path() {
	const TimeStep = 0.05
	const TrailWidth = 0.25
	var forward_direction = Vector3.Normalized(Vector3.New(weapon.ThrowDirection.X, 0, weapon.ThrowDirection.Z))
	var left_direction = Vector3.Cross(Vector3.Up, forward_direction)
	var offset_left = Vector3.MulX(left_direction, TrailWidth/2)
	var offset_right = Vector3.MulX(Vector3.Neg(left_direction), -TrailWidth/2)
	var st = SurfaceTool.New()
	st.Begin(Mesh.PrimitiveTriangles)
	var end_time = weapon.timeToLand + 0.5
	var point_previous Vector3.XYZ
	var time_current Float.X
	for time_current < end_time {
		time_current += TimeStep
		var point_current = Vector3.Add(Vector3.MulX(weapon.throwVelocity, time_current), Vector3.MulX(Vector3.Down, weapon.Gravity*0.5*time_current*time_current))
		var trail_point_left_end = Vector3.Add(point_current, offset_left)
		var trail_point_right_end = Vector3.Add(point_current, offset_right)
		var trail_point_left_start = Vector3.Add(point_previous, offset_left)
		var trail_point_right_start = Vector3.Add(point_previous, offset_right)
		var uv_progress_end = time_current / end_time
		var uv_progress_start = uv_progress_end - (TimeStep / end_time)
		var uv_value_right_start = Vector2.MulX(Vector2.Right, uv_progress_start)
		var uv_value_right_end = Vector2.MulX(Vector2.Right, uv_progress_end)
		var uv_value_left_start = Vector2.Add(Vector2.Down, uv_value_right_start)
		var uv_value_left_end = Vector2.Add(Vector2.Down, uv_value_right_end)
		point_previous = point_current
		st.SetUv(uv_value_right_end)
		st.AddVertex(trail_point_right_end)
		st.SetUv(uv_value_left_start)
		st.AddVertex(trail_point_left_start)
		st.SetUv(uv_value_left_end)
		st.AddVertex(trail_point_left_end)
		st.SetUv(uv_value_right_start)
		st.AddVertex(trail_point_right_start)
		st.SetUv(uv_value_left_start)
		st.AddVertex(trail_point_left_start)
		st.SetUv(uv_value_right_end)
		st.AddVertex(trail_point_right_end)
	}
	st.GenerateNormals()
	weapon.TrailMeshInstance.SetMesh(st.Commit().AsMesh())
}

type CharacterSkin struct {
	Node3D.Extension[CharacterSkin] `gd:"CharacterSkin"`

	Stepped Signal.Void `gd:"stepped()"`

	MainAnimationPlayer AnimationPlayer.ID
	MovingBlendPath     string
	AnimationTree       AnimationTree.Instance

	moveSpeed Float.X
	moving    bool
}

func NewCharacterSkin() *CharacterSkin {
	return &CharacterSkin{
		MovingBlendPath: "parameters/StateMachine/move/blend_position",
	}
}

func (skin *CharacterSkin) state_machine() AnimationNodeStateMachinePlayback.Instance {
	return Object.Get(skin.AnimationTree, "parameters/StateMachine/playback").(classdb.AnimationNodeStateMachinePlayback)
}

func (skin *CharacterSkin) Moving() bool {
	return skin.moving
}
func (skin *CharacterSkin) SetMoving(moving bool) {
	skin.moving = moving
	if moving {
		skin.state_machine().Travel("move")
	} else {
		skin.state_machine().Travel("idle")
	}
}

func (skin *CharacterSkin) MoveSpeed() Float.X { return skin.moveSpeed }
func (skin *CharacterSkin) SetMovingSpeed(speed Float.X) {
	skin.moveSpeed = Float.Clamp(speed, 0, 1)
	Object.Set(skin.AnimationTree, skin.MovingBlendPath, skin.moveSpeed)
}

func (skin *CharacterSkin) Ready() {
	skin.AnimationTree.AsAnimationMixer().SetActive(true)
	if player, ok := skin.MainAnimationPlayer.Instance(); ok {
		player.SetPlaybackDefaultBlendTime(0.1)
	}
}

func (skin *CharacterSkin) Jump() { skin.state_machine().Travel("jump") }
func (skin *CharacterSkin) Fall() { skin.state_machine().Travel("fall") }
func (skin *CharacterSkin) Punch() {
	Object.Set(skin.AnimationTree, "parameters/PunchOneShot/request", AnimationNodeOneShot.OneShotRequestFire)
}
func (skin *CharacterSkin) Step() { skin.Stepped.Emit() }
