extends Node2D

var hello: HelloWorld = HelloWorld.new()
var myclass: MyClassWithConstants = MyClassWithConstants.new()

# Called when the node enters the scene tree for the first time.
func _ready():
	hello.Print()
	
	print(hello.Arch())
	
	#var bar = hello.GetBar("Hello Bar")
	#print(bar.Message)
	
	$ExtendedNode.StringField = "Hello Property from Godot -> Go -> Godot"
	print($ExtendedNode.StringField)
	
	#hello.echo("Hello from GDScript")
	#print(hello.arch())
	#print(Engine.get_singleton("DisplayServer"))
	pass # Replace with function body.

func _notification(what):
	if what == NOTIFICATION_PREDELETE:
		hello.free()
		myclass.free()
		pass
		
# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass
