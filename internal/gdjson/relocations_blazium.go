package gdjson

func init() {
	Relocations["FoldableContainer.set_foldable_group"] = "FoldableGroup.set"
	Relocations["FoldableContainer.get_foldable_group"] = "FoldableGroup.get"
	Relocations["ColorButton.set_color"] = "ColorButton.set_color_emit_signal"
}
