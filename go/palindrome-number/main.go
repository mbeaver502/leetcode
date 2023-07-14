// https://leetcode.com/problems/palindrome-number/
// Given an integer x, return true if x is a palindrome, and false otherwise.
package main

import (
	"log"
	"math"
)

type example struct {
	value int
}

func main() {
	examples := []example{
		{121},
		{-121},
		{10},
	}

	for _, example := range examples {
		log.Printf("%d => %v", example.value, isPalindrome(example.value))
	}
}

// There's probably a better, cleverer way to do this
// using just simple arithmetic, but this approach
// here is fairly intuitive.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x >= 0 && x <= 9 {
		return true
	}

	numDigits := int(math.Floor(math.Log10(float64(x)))) + 1
	digits := make([]int, numDigits)

	// Put each digit into the digits slice.
	for i := 0; i < numDigits; i++ {
		digits[i] = x % 10
		x /= 10
	}

	// Check each pair of digits for sameness.
	for i := 0; i < (numDigits / 2); i++ {
		if digits[i] != digits[numDigits-i-1] {
			return false
		}
	}

	return true
}
