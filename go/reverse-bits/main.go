// https://leetcode.com/problems/reverse-bits/
// Reverse bits of a given 32 bits unsigned integer.
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
		{input: mustParse("00000010100101000001111010011100")},
		{input: mustParse("11111111111111111111111111111101")},
		{input: mustParse("11111111111111111111111111111111")},
		{input: mustParse("01010101010101010101010101010101")},
	}

	for _, ex := range examples {
		log.Printf("input:  %032b\n", ex.input)
		log.Printf("output: %032b\n", reverseBits(ex.input))
	}
}

// Basically we're going to isolate each bit
// in num by shifting it right until it hits zero.
// We get the bit by masking it (num & 1).
// Then we put that masked bit into our result
// by shifting it left by the necessary number
// of bits and bitwise OR'ing it into the result.
func reverseBits(num uint32) uint32 {
	out := uint32(0)
	i := uint32(31)
	for num > 0 {
		out |= (num & 1) << i
		num >>= 1
		i--
	}
	return out
}

func mustParse(s string) uint32 {
	num, err := strconv.ParseUint(s, 2, 32)
	if err != nil {
		panic(err)
	}
	return uint32(num)
}
