// https://leetcode.com/problems/two-sum/
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

package main

import "log"

type example struct {
	nums   []int
	target int
}

func main() {
	examples := []example{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}

	for _, ex := range examples {
		log.Printf("%v, %d => %v", ex.nums, ex.target, twoSum(ex.nums, ex.target))
	}
}

func twoSum(nums []int, target int) []int {
	// We'll map each encountered number to its index.
	// If the number we want is in the map, then we're done.
	// We're looking for a + b = t, but we can simplify
	// by looking for a = t - b in our map.
	numMap := make(map[int]int)
	for i, n := range nums {
		if j, found := numMap[target-n]; found {
			return []int{i, j}
		}
		numMap[n] = i
	}
	return []int{-1, -1}
}
