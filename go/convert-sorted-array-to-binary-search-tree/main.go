// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Given an integer array nums where the elements are sorted in ascending order,
// convert it to a height-balanced binary search tree.
package main

import "log"

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
		tree := sortedArrayToBST(ex.input)
		log.Printf("output: %+v", inorderTraversal(tree))
	}
}

// TreeNode defines a binary tree.
// Provided by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	lenNums := len(nums)

	switch {
	// If we got an empty list, our tree (or leaf) is nil.
	case lenNums < 1:
		return nil

	// We found a leaf!
	case lenNums == 1:
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}

	// We want our tree to be height-balanced,
	// so start each partition at the halfway point.
	halfway := len(nums) / 2
	return &TreeNode{
		Val:   nums[halfway],
		Left:  sortedArrayToBST(nums[:halfway]),
		Right: sortedArrayToBST(nums[halfway+1:]),
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	parent := []int{root.Val}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	return append(left, append(parent, right...)...)
}
