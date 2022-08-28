extends Node2D

var hello: HelloWorld

# Called when the node enters the scene tree for the first time.
func _ready():
	print(Engine.get_singleton("DisplayServer"))
	pass # Replace with function body.


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass
