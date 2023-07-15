// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Given an integer array nums where the elements are sorted in ascending order,
// convert it to a height-balanced binary search tree.
package main

import (
	"log"

	btu "github.com/mbeaver502/leetcode/go/binarytreeutils"
)

type example struct {
	input []int
}

func main() {
	examples := []example{
		{[]int{1}},
		{[]int{1, 3}},
		{[]int{-10, -3, 0, 5, 9}},
		{[]int{-20, -19, -18, -17, -16, -15, -14, -13, -12, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
	}

	for _, ex := range examples {
		log.Printf("input:  %+v", ex.input)
		log.Printf("output: %+v", btu.SortedArrayToBST(ex.input).InOrderTraversal())
	}
}
