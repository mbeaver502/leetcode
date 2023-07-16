// https://leetcode.com/problems/ransom-note/
// Given two strings ransomNote and magazine,
// return true if ransomNote can be constructed
// by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.
package main

import "log"

type example struct {
	note     string
	magazine string
}

func main() {
	examples := []example{
		{"a", "b"},
		{"aa", "ab"},
		{"aa", "aab"},
	}

	for _, ex := range examples {
		log.Printf("input:  note = %s, magazine = %s\n", ex.note, ex.magazine)
		log.Printf("output: %+v\n", canConstruct(ex.note, ex.magazine))
	}
}

func canConstruct(ransomNote string, magazine string) bool {
	noteMap := make(map[rune]int)
	for _, c := range ransomNote {
		noteMap[c]++
	}

	magMap := make(map[rune]int)
	for _, c := range magazine {
		magMap[c]++
	}

	for k := range noteMap {
		// If a ransom letter doesn't exist in the
		// magazine, we can't construct the ransom note.
		if _, found := magMap[k]; !found {
			return false
		}

		// If a letter appears more in the ransom note
		// than in the magazine, then that means not
		// enough letters are available in the magazine.
		if noteMap[k] > magMap[k] {
			return false
		}
	}

	return true
}
