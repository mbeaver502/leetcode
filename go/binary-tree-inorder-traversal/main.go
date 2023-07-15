// https://leetcode.com/problems/binary-tree-inorder-traversal/
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
package main

import (
	"log"

	btu "github.com/mbeaver502/leetcode/go/binarytreeutils"
)

// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
type example struct {
	tree *btu.TreeNode
}

func main() {
	examples := []example{
		{
			&btu.TreeNode{
				Val:  1,
				Left: nil,
				Right: &btu.TreeNode{
					Val: 2,
					Left: &btu.TreeNode{
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
			&btu.TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		{
			&btu.TreeNode{
				Val: 30,
				Left: &btu.TreeNode{
					Val: 20,
					Left: &btu.TreeNode{
						Val: 15,
						Left: &btu.TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &btu.TreeNode{
							Val:   18,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &btu.TreeNode{
						Val:   25,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &btu.TreeNode{
					Val: 40,
					Left: &btu.TreeNode{
						Val:   35,
						Left:  nil,
						Right: nil,
					},
					Right: &btu.TreeNode{
						Val: 50,
						Left: &btu.TreeNode{
							Val:   45,
							Left:  nil,
							Right: nil,
						},
						Right: &btu.TreeNode{
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
		log.Printf("output: %+v\n", ex.tree.InOrderTraversal())
	}
}
