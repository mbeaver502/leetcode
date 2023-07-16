// https://leetcode.com/problems/reverse-string/
// Write a function that reverses a string.
// The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.
package main

import "log"

type example struct {
	s []byte
}

func main() {
	examples := []example{
		{[]byte{'h', 'e', 'l', 'l', 'o'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}},
	}

	for _, ex := range examples {
		log.Printf("input:  %+v\n", ex.s)
		reverseString(ex.s)
		log.Printf("output: %+v\n", ex.s)
	}
}

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
