package utils

import (
	"math"
	"regexp"
)

var notNumberRE = regexp.MustCompile(`\D`)

// CleanNumber remove all not number characters
func CleanNumber(value string) string {
	return notNumberRE.ReplaceAllString(value, "")
}

// Round rounds a number by decimal places
// Eg.: var x = Round(1.234, 2) // 1.23
// Eg.: var x = Round(1.235, 2) // 1.24
func Round(f float64, places int) float64 {
	if places <= 0 {
		return math.Round(f)
	}

	shift := math.Pow10(places)

	return math.Round(f*shift) / shift
}

// RoundUp rounds up a number by decimal places
// Eg.: var x = RoundUp(1.234, 2) // 1.24
// Eg.: var x = RoundUp(1.235, 2) // 1.24
func RoundUp(f float64, places int) float64 {
	if places <= 0 {
		return math.Ceil(f)
	}

	shift := math.Pow10(places)

	return math.Ceil(Round(f*shift, 1)) / shift
}

// RoundDown rounds down a number by decimal places
// Eg.: var x = RoundDown(1.234, 2) // 1.23
// Eg.: var x = RoundDown(1.235, 2) // 1.23
func RoundDown(f float64, places int) float64 {
	if places <= 0 {
		return math.Floor(f)
	}

	shift := math.Pow10(places)

	return math.Floor(Round(f*shift, 1)) / shift
}
