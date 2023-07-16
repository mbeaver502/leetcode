// https://leetcode.com/problems/isomorphic-strings/
// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving
// the order of characters. No two characters may map to the same character, but a character may map to itself.
package main

import "log"

// 1 <= s.length <= 5 * 10^4
// t.length == s.length
// s and t consist of any valid ASCII character.
type example struct {
	s string
	t string
}

func main() {
	examples := []example{
		{"egg", "add"},
		{"foo", "bar"},
		{"paper", "title"},
		{"badc", "baba"},
	}

	for _, ex := range examples {
		log.Printf("input:  s = %s, t = %s\n", ex.s, ex.t)
		log.Printf("output: %+v\n", isIsomorphic(ex.s, ex.t))
	}
}

func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, found := sMap[s[i]]; found {
			if _, found := tMap[v]; found {
				if v != t[i] {
					return false
				}
			}
		}

		if v, found := tMap[t[i]]; found {
			if _, found := sMap[v]; found {
				if v != s[i] {
					return false
				}
			}
		}

		sMap[s[i]] = t[i]
		tMap[t[i]] = s[i]
	}
	return true
}
