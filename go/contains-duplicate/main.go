// https://leetcode.com/problems/contains-duplicate/
// Given an integer array nums, return true if
// any value appears at least twice in the array,
// and return false if every element is distinct.
package main

import "log"

type example struct {
	nums []int
}

func main() {
	examples := []example{
		{[]int{1, 2, 3, 1}},
		{[]int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
	}

	for _, ex := range examples {
		log.Printf("input:  %+v", ex.nums)
		log.Printf("output: %+v", containsDuplicate(ex.nums))
	}
}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if _, found := seen[num]; found {
			return true
		}
		seen[num] = true
	}
	return (len(seen) != len(nums))
}
