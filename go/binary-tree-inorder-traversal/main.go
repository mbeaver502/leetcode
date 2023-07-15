// https://leetcode.com/problems/binary-tree-inorder-traversal/
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
package main

import (
	"log"
)

// The number of nodes in the tree is in the range [0, 100].
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
		traversal := inorderTraversal(ex.tree)
		log.Printf("%+v\n", traversal)
	}
}

// TreeNode defines a binary tree.
// Provided by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
