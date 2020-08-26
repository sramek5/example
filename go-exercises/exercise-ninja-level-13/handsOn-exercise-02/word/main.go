// Package word provides custom functions for working with strings and words.
package word

import "strings"

// UseCount returns the number of times each word is used in a string.
   // no need to write an example for this one
   // writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string.
func Count(t string) int {
	// my written code for this func:
	xt := strings.Fields(t)
	return len(xt)
}