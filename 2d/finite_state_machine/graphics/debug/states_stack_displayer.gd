extends Panel

@onready var fsm_node: Node = get_node(^"../../Player")

func _process(_delta: float) -> void:
	var states_names := ""
	var numbers := ""
	var index := 0

	for state: String in fsm_node.states_stack():
		states_names += state + "\n"
		numbers += str(index) + "\n"
		index += 1

	%States.text = states_names
	%Numbers.text = numbers
