// https://leetcode.com/problems/perfect-number/
// A perfect number is a positive integer that
// is equal to the sum of its positive divisors,
// excluding the number itself.
// A divisor of an integer x is an integer that can divide x evenly.
// Given an integer n, return true if n is a perfect number,
// otherwise return false.
package main

import (
	"log"
	"math"
)

type example struct {
	n int
}

func main() {
	examples := []example{
		{28},
		{1},
		{35},
		{6},
		{496},
	}

	for _, ex := range examples {
		log.Printf("input:  %d\n", ex.n)
		log.Printf("output: %+v\n", checkPerfectNumber(ex.n))
	}
}

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}

	sum := 1
	bound := int(math.Floor(math.Sqrt(float64(num)))) + 1
	for i := 2; i < bound; i++ {
		if num%i == 0 {
			sum += (num / i) + i
		}
	}
	return sum == num
}
