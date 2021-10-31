package utils

import "strings"

//LeftPad fill string with a determined character
func LeftPad(str string, lgt int, pad string) string {
	content := []rune(str)

	if len(content) >= lgt {
		return str
	}

	return strings.Repeat(pad, lgt-len(content)) + str
}
