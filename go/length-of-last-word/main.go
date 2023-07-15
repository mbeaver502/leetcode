// https://leetcode.com/problems/length-of-last-word/
// Given a string s consisting of words and spaces,
// return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.
package main

import (
	"log"
	"regexp"
)

type example struct {
	s string
}

func main() {
	examples := []example{
		{"Hello World"},
		{"   fly me   to   the moon  "},
		{"luffy is still joyboy"},
	}

	for _, ex := range examples {
		log.Printf("input: s = \"%s\"", ex.s)
		log.Printf("output: %d", lengthOfLastWord(ex.s))
	}
}

// 1 <= s.length <= 104
// s consists of only English letters and spaces ' '.
// There will be at least one word in s.
func lengthOfLastWord(s string) int {
	rex := regexp.MustCompile(`\w+`)

	// Since we have the above guarantees,
	// we're not going to bother checking idxSlice == nil.
	// This is pretty lazy. It probably would've
	// been better to iterate backwards over the
	// string and count the non-space characters
	// until hitting a space character.
	idxSlice := rex.FindAllStringIndex(s, -1)
	lastIdx := idxSlice[len(idxSlice)-1]

	return len(s[lastIdx[0]:lastIdx[1]])
}
