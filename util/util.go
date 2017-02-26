package util

import (
	"strings"
)

// Replace all occurrences of a space with an underscore in given string
func ReplaceSpaces(in string) string {
	return strings.Replace(in, " ", "_", -1)
}

// Replace all underscores with a space in given string
func ReplaceUnderscores(in string) string {
	return strings.Replace(in, "_", " ", -1)
}
