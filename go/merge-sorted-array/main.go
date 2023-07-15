// https://leetcode.com/problems/merge-sorted-array/
// You are given two integer arrays nums1 and nums2,
// sorted in non-decreasing order, and two integers m and n,
// representing the number of elements in nums1 and nums2 respectively.
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function,
// but instead be stored inside the array nums1.
// To accommodate this, nums1 has a length of m + n, where the first m elements
// denote the elements that should be merged, and the last n elements are
// set to 0 and should be ignored. nums2 has a length of n.
package main

import "log"

type example struct {
	nums1 []int
	nums2 []int
	m     int
	n     int
}

func main() {
	examples := []example{
		{[]int{3, 0, 0}, []int{1, 2}, 1, 2},
		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3},
		{[]int{1}, []int{}, 1, 0},
		{[]int{0}, []int{1}, 0, 1},
		{[]int{1, 2, 3, 4, 5, 6, 0, 0, 0}, []int{9, 9, 9}, 6, 3},
		{[]int{9, 9, 9, 0, 0, 0, 0, 0, 0}, []int{1, 2, 3, 4, 5, 6}, 3, 6},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 10},
		{[]int{0, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, 10},
		{[]int{0, 2, 4, 6, 8, 0, 0, 0, 0, 0}, []int{1, 3, 5, 7, 9}, 5, 5},
		{[]int{1, 3, 5, 7, 9, 0, 0, 0, 0, 0}, []int{0, 2, 4, 6, 8}, 5, 5},
	}

	for _, ex := range examples {
		log.Printf("input:  nums1 = %v, m = %d, nums2 = %v, n = %d", ex.nums1, ex.m, ex.nums2, ex.n)
		merge(ex.nums1, ex.m, ex.nums2, ex.n)
		log.Printf("output: %v", ex.nums1)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// No nums2 was given, so we're done.
	if n == 0 {
		return
	}

	// No nums1 was given, so nums2 is default.
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	// Let's start by getting the first
	// m digits of nums1 out of the way.
	copy(nums1[n:], nums1[:m])

	pos := 0
	lPtr := n
	rPtr := 0

	// Let's merge elements from nums2 into nums1,
	// giving preference to the elements in nums1.
	// Whenever we run out of nums1 elements,
	// copy over the rest of nums2 and exit.
	for pos < m+n {
		if lPtr < m+n {
			if rPtr >= n {
				break
			}

			if nums1[lPtr] <= nums2[rPtr] {
				nums1[pos] = nums1[lPtr]
				lPtr++
			} else if nums1[lPtr] > nums2[rPtr] {
				nums1[pos] = nums2[rPtr]
				rPtr++
			}
		} else {
			copy(nums1[pos:], nums2[rPtr:])
			break
		}

		pos++
	}
}
