extends Node2D

var hello: HelloWorld

# Called when the node enters the scene tree for the first time.
func _ready():
	hello = HelloWorld.new()
	hello.Print()
	hello.free()
	pass # Replace with function body.


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass
