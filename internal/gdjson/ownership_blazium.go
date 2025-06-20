package gdjson

func init() {
	ClassMethodOwnership["FoldableContainer"] = map[string]map[string]OwnershipSemantics{
		"add_title_bar_control": {
			"control": OwnershipTransferred,
		},
		"remove_title_bar_control": {
			"control": ReversesTheOwnership,
		},
	}
	ClassMethodOwnership["FoldableGroup"] = map[string]map[string]OwnershipSemantics{
		"get_expanded_container": {
			"return value": MustAssertInstanceID,
		},
	}
}
