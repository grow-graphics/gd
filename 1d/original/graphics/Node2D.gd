extends Node2D

var hello: HelloWorld = HelloWorld.new()

# Called when the node enters the scene tree for the first time.
func _ready():
	hello.print()

	print(hello.arch())

	var bar = hello.get_bar("Hello Bar")
	print(bar.message)

	$ExtendedNode.string_field = "Hello Property from Godot -> Go -> Godot"
	print($ExtendedNode.string_field)
	pass

func _notification(what):
	if what == NOTIFICATION_PREDELETE:
		hello.free()
		pass

# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass
