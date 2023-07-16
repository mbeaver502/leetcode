// https://leetcode.com/problems/fizz-buzz/
// Given an integer n, return a string array answer (1-indexed) where:
//   - answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
//   - answer[i] == "Fizz" if i is divisible by 3.
//   - answer[i] == "Buzz" if i is divisible by 5.
//   - answer[i] == i (as a string) if none of the above conditions are true.
package main

import (
	"log"
	"strconv"
)

type example struct {
	n int
}

func main() {
	examples := []example{
		{3},
		{5},
		{15},
		{21},
	}

	for _, ex := range examples {
		log.Printf("input:  %d\n", ex.n)
		log.Printf("output: %+v\n", fizzBuzz(ex.n))
	}
}

func fizzBuzz(n int) []string {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			out[i-1] = "FizzBuzz"
		case i%5 == 0:
			out[i-1] = "Buzz"
		case i%3 == 0:
			out[i-1] = "Fizz"
		default:
			out[i-1] = strconv.Itoa(i)
		}
	}
	return out
}
