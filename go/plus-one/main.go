// https://leetcode.com/problems/plus-one/
// You are given a large integer represented as an integer array digits,
// where each digits[i] is the ith digit of the integer.
// The digits are ordered from most significant to least significant
// in left-to-right order. The large integer does not contain any leading 0's.
// Increment the large integer by one and return the resulting array of digits.
package main

import "log"

type example struct {
	digits []int
}

func main() {
	examples := []example{
		{[]int{0}},
		{[]int{1}},
		{[]int{2}},
		{[]int{3}},
		{[]int{4}},
		{[]int{5}},
		{[]int{6}},
		{[]int{7}},
		{[]int{8}},
		{[]int{1, 2, 3}},
		{[]int{4, 3, 2, 1}},
		{[]int{9, 9, 9, 8}},
		{[]int{9}},
		{[]int{9, 9}},
		{[]int{9, 9, 9}},
		{[]int{9, 9, 9, 9}},
		{[]int{9, 9, 9, 9, 9}},
		{[]int{9, 9, 9, 9, 9, 9}},
	}

	for _, ex := range examples {
		log.Printf("input:  %v", ex.digits)
		log.Printf("output: %v", plusOne(ex.digits))
	}
}

// 1 <= digits.length <= 100
// 0 <= digits[i] <= 9
// digits does not contain any leading zeroes.
func plusOne(digits []int) []int {
	idx := len(digits) - 1
	carry := false

	// Increment the last digit.
	d := digits[idx] + 1
	d, carry = getCarry(d)
	digits[idx] = d

	// Carry forward any tens.
	for carry {
		carry = false
		idx--

		if idx >= 0 {
			d = digits[idx] + 1
		} else {
			d = 1
			idx = -1
		}

		d, carry = getCarry(d)

		if idx < 0 {
			digits = append([]int{0}, digits...)
			idx = 0
		}

		digits[idx] = d
	}

	return digits
}

func getCarry(d int) (int, bool) {
	if d >= 10 {
		return d % 10, true
	}
	return d, false
}
