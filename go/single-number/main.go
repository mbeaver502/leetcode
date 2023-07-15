// https://leetcode.com/problems/single-number/
// Given a non-empty array of integers nums,
// every element appears twice except for one.
// Find that single one.
// You must implement a solution with a linear
// runtime complexity and use only constant extra space.
package main

import "log"

// 1 <= nums.length <= 3 * 104
// -3 * 104 <= nums[i] <= 3 * 104
// Each element in the array appears twice
// except for one element which appears only once.
type example struct {
	nums []int
}

func main() {
	examples := []example{
		{[]int{2, 2, 1}},
		{[]int{4, 1, 2, 1, 2}},
		{[]int{1}},
	}

	for _, ex := range examples {
		log.Printf("input:  nums = %v", ex.nums)
		log.Printf("output: %d", singleNumber(ex.nums))
	}
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// We can use the fact that a ^ a = 0 (a XOR a == 0)
	// to eliminate all the paired numbers.
	// Then we can say that the remainder r ^ 0 = r (r XOR 0 == r).
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
