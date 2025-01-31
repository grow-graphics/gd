package gd

import (
	"strings"
	"unicode"
)

// whether fields in structs will be converted from PascalCase to snake_case when being exported to Godot.
var DoFieldCaseConversion bool = false

// converts PascalCase name to snake_case
// !! returns the input string if input is not in PascalCase
func PascalToSnakeCase(name string) string {
	nameRunes := []rune(name)
	current_letter := nameRunes[0]
	words := make([]string, 0)
	current_word := make([]rune, 0)
	isWordStart := func(l rune) bool { return unicode.IsUpper(l) || unicode.IsNumber(l) }
	appendLetter := func() { current_word = append(current_word, current_letter) }
	appendWord := func() {
		words = append(words, strings.ToLower(string(current_word)))
		current_word = make([]rune, 0)
	}
	for _, next_letter := range nameRunes[1:] {
		if !(unicode.IsNumber(next_letter) || unicode.IsLetter(next_letter)) {
			return name
		}
		if unicode.IsLower(current_letter) && isWordStart(next_letter) {
			appendLetter()
			appendWord()
		} else if isWordStart(current_letter) && unicode.IsLower(next_letter) {
			if len(current_word) > 0 {
				appendWord()
			}
			appendLetter()
		} else {
			appendLetter()
		}

		current_letter = next_letter
	}
	appendLetter()
	appendWord()
	return strings.Join(words, "_")
}
