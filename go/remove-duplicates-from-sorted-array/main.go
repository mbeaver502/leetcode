// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Given an integer array nums sorted in non-decreasing order,
// remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same.
// Then return the number of unique elements in nums.
// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
//   - Change the array nums such that the first k elements of nums
//     contain the unique elements in the order they were present in nums initially.
//     The remaining elements of nums are not important as well as the size of nums.
//   - Return k.
package main

import "log"

type example struct {
	values []int
}

func main() {
	examples := []example{
		{[]int{1, 1, 2}},                      // 2, nums = [1,2,_]
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, // 5, nums = [0,1,2,3,4,_,_,_,_,_]
	}

	for _, example := range examples {
		log.Printf("input = %v", example.values)
		out := removeDuplicates(example.values)
		log.Printf("output = %d, %v", out, example.values)
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	nextVal := 0
	currVal := nums[0]
	numVals := 1

	for _, num := range nums {
		if num != currVal {
			numVals++
			nextVal++
			nums[nextVal] = num
			currVal = num
		}
	}

	return numVals
}
