extends Node

var a = Signals.new()

func _ready():
	a.connect("something", handle_sig)
	a.DoSomething()

func handle_sig():
	print("yay")


func _exit_tree():
	a.free()
