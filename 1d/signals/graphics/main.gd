extends Node

var a = Signals.new()

func _ready():
	a.connect("something", handle_sig)
	a.connect("generic", handle_generic)
	a.do_something()

func handle_sig():
	print("something happened")

func handle_generic(val):
	print(val)

func _exit_tree():
	a.free()
