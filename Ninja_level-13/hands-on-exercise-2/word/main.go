// Package word provides functions to work with strings and words
package word

import "strings"

// UseCount returns the number of times each word is used in the string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count returns the number of words in the string
func Count(s string) int {
	// write the code for this func
	a := strings.Fields(s)
	return len(a)
}
