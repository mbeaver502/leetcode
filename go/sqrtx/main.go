// https://leetcode.com/problems/sqrtx/
// Given a non-negative integer x, return the square root of x
// rounded down to the nearest integer.
// The returned integer should be non-negative as well.
// You must not use any built-in exponent function or operator.
package main

import "log"

type example struct {
	value int
}

func main() {
	examples := []example{
		{2},
		{4},
		{8},
		{27},
	}

	for _, ex := range examples {
		log.Printf("input: %d", ex.value)
		log.Printf("output: %d", mySqrt(ex.value))
	}
}

// 0 <= x <= (2^31) - 1
// Let's just use the binary search implementation from Wikipedia.
// https://en.wikipedia.org/wiki/Integer_square_root#Algorithm_using_binary_search
func mySqrt(x int) int {
	l := uint64(0)
	m := uint64(0)
	r := uint64(x + 1)

	for l != r-1 {
		m = (l + r) / 2

		if m*m <= uint64(x) {
			l = m
		} else {
			r = m
		}
	}

	return int(l)
}
