extends Node

var a = Signals.new()

func _ready():
	a.connect("something", handle_sig)
	a.do_something()

func handle_sig():
	print("something happened")

func _exit_tree():
	a.free()
