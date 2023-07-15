// https://leetcode.com/problems/remove-element/
// Given an integer array nums and an integer val,
// remove all occurrences of val in nums in-place.
// The order of the elements may be changed.
// Then return the number of elements in nums which are not equal to val.
// Consider the number of elements in nums which are not equal to val be k,
// to get accepted, you need to do the following things:
//   - Change the array nums such that the first k elements of nums
//     contain the elements which are not equal to val.
//     The remaining elements of nums are not important
//     as well as the size of nums.
//   - Return k.
package main

import "log"

type example struct {
	values []int
	target int
}

func main() {
	examples := []example{
		{[]int{3, 2, 2, 3}, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
		{[]int{0, 0, 0, 3, 2, 5, 4, 3, 2, 1, 0, 0, 0, 5, 6, 7}, 0},
	}

	for _, ex := range examples {
		log.Printf("input: nums = %v, val = %d", ex.values, ex.target)
		k := removeElement(ex.values, ex.target)
		log.Printf("output: k = %d, nums = %v", k, ex.values)
	}
}

func removeElement(nums []int, val int) int {
	// We're guaranteed that:
	// 0 <= nums.length <= 100
	// 0 <= nums[i] <= 50
	// 0 <= val <= 100

	// Worst case, we end up copying the entire
	// contents of nums into kVals if no num[i]
	// equals val.
	kVals := []int{}

	// Let's take a single pass over nums
	// and mark any instances of val as -1.
	// We'll use -1 since 0 <= nums[i] <= 50.
	for i, num := range nums {
		if num == val {
			nums[i] = -1
		} else {
			kVals = append(kVals, num)
		}
	}

	// Overwrite the first k values with
	// our saved non-val values.
	copy(nums[:len(kVals)], kVals)

	return len(kVals)
}
