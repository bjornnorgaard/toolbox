package jqexp

import "strings"

// New prepares a string to be used as a jq expression.
// Essentially removes all newlines and tabs from the string.
func New(s string) string {
	return strings.NewReplacer("\n", "", "\t", "").Replace(s)
}
