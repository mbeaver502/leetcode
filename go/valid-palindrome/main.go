// https://leetcode.com/problems/valid-palindrome/
// A phrase is a palindrome if, after converting all
// uppercase letters into lowercase letters and removing
// all non-alphanumeric characters, it reads the same
// forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.
package main

import (
	"log"
	"regexp"
	"strings"
)

type example struct {
	s string
}

func main() {
	examples := []example{
		{"A man, a plan, a canal: Panama"},
		{"race a car"},
		{"racecar"},
		{"saippuakivikauppias"},
		{"Madam I'm Adam"},
		{" "},
		{"0P"},
	}

	for _, ex := range examples {
		log.Printf("input:  \"%s\"", ex.s)
		log.Printf("output: %v", isPalindrome(ex.s))
	}
}

func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	rex := regexp.MustCompile(`[^a-z\d]`)
	str = rex.ReplaceAllString(str, "")

	end := len(str) - 1
	halfway := len(str) / 2

	for i := 0; i < halfway; i++ {
		if str[i] != str[end-i] {
			return false
		}
	}

	return true
}
