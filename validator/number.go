package validator

import "regexp"

var notNumberRE = regexp.MustCompile(`\D`)

// cleanNumber remove all not number characters
func cleanNumber(value string) string {
	return notNumberRE.ReplaceAllString(value, "")
}
