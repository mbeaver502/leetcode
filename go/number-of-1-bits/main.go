// https://leetcode.com/problems/number-of-1-bits/
// Write a function that takes the binary representation
// of an unsigned integer and returns the number of '1' bits
// it has (also known as the Hamming weight).
package main

import (
	"log"
	"strconv"
)

type example struct {
	input uint32
}

func main() {
	examples := []example{
		{input: mustParse("00000000000000000000000000000000")},
		{input: mustParse("00000000000000000000000000001011")},
		{input: mustParse("00000000000000000000000010000000")},
		{input: mustParse("11111111111111111111111111111101")},
		{input: mustParse("11111111111111111111111111111111")},
	}

	for _, ex := range examples {
		log.Printf("input:  %32b\n", ex.input)
		log.Printf("output: %d\n", hammingWeight(ex.input))
	}
}

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}

	// Bitmask AND to determine if the last digit is 1.
	// Bitshift by 1 to shift the entire number right by
	// one position. Repeat until number is 0.
	return int((num & 1) + uint32(hammingWeight(num>>1)))
}

func mustParse(s string) uint32 {
	num, err := strconv.ParseUint(s, 2, 32)
	if err != nil {
		panic(err)
	}
	return uint32(num)
}
