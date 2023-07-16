// https://leetcode.com/problems/counting-bits/
// Given an integer n, return an array ans of
// length n + 1 such that for each i (0 <= i <= n),
// ans[i] is the number of 1's in the binary representation of i.
package main

import "log"

type example struct {
	n int
}

func main() {
	examples := []example{
		{2},
		{5},
		{10},
		{32},
	}

	for _, ex := range examples {
		log.Printf("input:  %d\n", ex.n)
		log.Printf("output: %+v\n", countBits(ex.n))
	}
}

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = countOnes(i)
	}
	return ans
}

func countOnes(i int) int {
	count := 0
	for i > 0 {
		if i&1 == 1 {
			count++
		}
		i >>= 1
	}
	return count
}
