// https://leetcode.com/problems/add-binary/
// Given two binary strings a and b, return their sum as a binary string.
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type example struct {
	a string
	b string
}

func main() {
	examples := []example{
		{
			a: "11",
			b: "1",
		},
		{
			a: "1010",
			b: "1011",
		},
		{
			a: "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			b: "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
		},
	}

	for _, ex := range examples {
		log.Printf("input: a = %s, b = %s", ex.a, ex.b)
		log.Printf("output: %s", addBinary(ex.a, ex.b))
	}
}

// 1 <= a.length, b.length <= 10,000
// a and b consist only of '0' or '1' characters.
// Each string does not contain leading zeros except for the zero itself.
// Unfortunately, 10,000 bits is too large
// parse into a 64-bit uint. So we'll add manually.
func addBinary(a string, b string) string {
	var sb strings.Builder

	a, b = padZero(a, b)
	carry := false

	for i := len(a) - 1; i >= 0; i-- {
		n, _ := strconv.ParseUint(string(a[i]), 2, 32)
		m, _ := strconv.ParseUint(string(b[i]), 2, 32)

		s := n + m
		if carry {
			s++
			carry = false
		}

		if s >= 2 {
			s %= 2
			carry = true
		}

		sb.WriteString(strconv.FormatUint(s, 2))
	}

	if carry {
		sb.WriteString("1")
	}

	s := sb.String()
	sb.Reset()
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}

	return sb.String()
}

func padZero(a string, b string) (string, string) {
	aOut := a
	bOut := b

	lenA := len(a)
	lenB := len(b)
	lenDiff := lenA - lenB

	switch {
	case lenDiff == 0:
		return a, b
	case lenDiff > 0:
		bOut = fmt.Sprintf("%s%s", strings.Repeat("0", lenDiff), bOut)
	case lenDiff < 0:
		lenDiff *= (-1)
		aOut = fmt.Sprintf("%s%s", strings.Repeat("0", lenDiff), aOut)
	}

	return aOut, bOut
}
