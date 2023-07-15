// https://leetcode.com/problems/path-sum/
// Given the root of a binary tree and an integer targetSum,
// return true if the tree has a root-to-leaf path such that
// adding up all the values along the path equals targetSum.
// A leaf is a node with no children.
package main

import (
	"log"

	btu "github.com/mbeaver502/leetcode/go/binarytreeutils"
)

type example struct {
	tree      *btu.TreeNode
	targetSum int
}

func main() {
	examples := []example{
		{
			tree:      nil,
			targetSum: 0,
		},
		{
			tree: &btu.TreeNode{
				Val: 1,
				Left: &btu.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			targetSum: 1,
		},
		{
			tree: &btu.TreeNode{
				Val: 1,
				Left: &btu.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &btu.TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			targetSum: 5,
		},
		{
			tree: &btu.TreeNode{
				Val: 5,
				Left: &btu.TreeNode{
					Val: 4,
					Left: &btu.TreeNode{
						Val: 11,
						Left: &btu.TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: &btu.TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
				Right: &btu.TreeNode{
					Val: 8,
					Left: &btu.TreeNode{
						Val:   13,
						Left:  nil,
						Right: nil,
					},
					Right: &btu.TreeNode{
						Val:  4,
						Left: nil,
						Right: &btu.TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			targetSum: 22,
		},
	}

	for _, ex := range examples {
		log.Printf("input:  tree = %+v, targetSum = %d\n", ex.tree.InOrderTraversal(), ex.targetSum)
		log.Printf("output: %v\n", ex.tree.HasPathSum(ex.targetSum))
	}
}
