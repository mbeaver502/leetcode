// https://leetcode.com/problems/sum-of-left-leaves/
// Given the root of a binary tree, return the sum of all left leaves.
// A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.
package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: &TreeNode{Val: 5, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil},
	}

	fmt.Printf("input:  %+v\n", tree)
	fmt.Printf("output: %d\n", sumOfLeftLeaves(tree))
}

// TreeNode defines a binary tree.
// Provided by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// The tree is just a root, so there's no leaves.
	if root.Left == nil && root.Right == nil {
		return 0
	}

	return doSum(root.Left, false) + doSum(root.Right, true)
}

func doSum(root *TreeNode, isRight bool) int {
	if root == nil {
		return 0
	}

	// We found a leaf, and it's on the left.
	if root.Left == nil && root.Right == nil && !isRight {
		return root.Val
	}

	return doSum(root.Left, false) + doSum(root.Right, true)
}
