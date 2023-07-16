// https://leetcode.com/problems/hamming-distance/
// The Hamming distance between two integers is
// the number of positions at which the
// corresponding bits are different.
// Given two integers x and y, return the
// Hamming distance between them.
package main

import "log"

type example struct {
	x int
	y int
}

func main() {
	examples := []example{
		{1, 4},
		{3, 1},
	}

	for _, ex := range examples {
		log.Printf("input:  x = %d, y = %d\n", ex.x, ex.y)
		log.Printf("output: %d\n", hammingDistance(ex.x, ex.y))
	}
}

func hammingDistance(x int, y int) int {
	// XOR (x^y) will give us all the differing bits.
	xor := x ^ y
	bits := 0

	// Now we can bitmask (xor&1) to isolate the final bit
	// and then bitshift the entire number by one bit right.
	for xor > 0 {
		bits += xor & 1
		xor >>= 1
	}
	return bits
}
