// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along
// the longest path from the root node down to the farthest leaf node.
package main

import "log"

// The number of nodes in the tree is in the range [0, 10^4].
// -100 <= Node.val <= 100
type example struct {
	tree *TreeNode
}

func main() {
	examples := []example{
		{
			&TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		{
			nil,
		},
		{
			&TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		{
			&TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
						Left: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   18,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   25,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 40,
					Left: &TreeNode{
						Val:   35,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 50,
						Left: &TreeNode{
							Val:   45,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   60,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
		},
	}

	for _, ex := range examples {
		log.Printf("input:  %+v", ex.tree)
		log.Printf("%+v\n", maxDepth(ex.tree))
	}
}

// TreeNode defines a binary tree.
// Provided by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthLeft := 1 + maxDepth(root.Left)
	depthRight := 1 + maxDepth(root.Right)

	if depthLeft > depthRight {
		return depthLeft
	}
	return depthRight
}
