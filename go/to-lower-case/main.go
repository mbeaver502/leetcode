// https://leetcode.com/problems/to-lower-case/
// Given a string s, return the string after replacing
// every uppercase letter with the same lowercase letter.
package main

import (
	"log"
	"strings"
)

// 1 <= s.length <= 100
// s consists of printable ASCII characters.
type example struct {
	s string
}

func main() {
	examples := []example{
		{"Hello"},
		{"here"},
		{"LOVELY"},
	}

	for _, ex := range examples {
		log.Printf("input:  %s\n", ex.s)
		log.Printf("output: %s\n", toLowerCase(ex.s))
	}
}

func toLowerCase(s string) string {
	var sb strings.Builder

	// We're guaranteed to have ASCII characters,
	// so we can do simple manipulation here.
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c = (c - 'A') + 'a'
		}
		sb.WriteRune(c)
	}

	return sb.String()
}
