// https://leetcode.com/problems/symmetric-tree/
// Given the root of a binary tree, check whether
// it is a mirror of itself (i.e., symmetric around its center).
package main

import (
	"log"

	btu "github.com/mbeaver502/leetcode/go/binarytreeutils"
)

type example struct {
	tree *btu.TreeNode
}

func main() {
	examples := []example{
		{
			nil,
		},
		{
			&btu.TreeNode{
				Val: 1,
				Left: &btu.TreeNode{
					Val:   2,
					Left:  &btu.TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &btu.TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &btu.TreeNode{
					Val:   2,
					Left:  &btu.TreeNode{Val: 4, Left: nil, Right: nil},
					Right: &btu.TreeNode{Val: 3, Left: nil, Right: nil},
				},
			},
		},
		{
			&btu.TreeNode{
				Val: 1,
				Left: &btu.TreeNode{
					Val:   2,
					Left:  &btu.TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &btu.TreeNode{
					Val:   2,
					Left:  &btu.TreeNode{Val: 4, Left: nil, Right: nil},
					Right: nil,
				},
			},
		},
		{
			&btu.TreeNode{
				Val: 1,
				Left: &btu.TreeNode{
					Val:   2,
					Left:  &btu.TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &btu.TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &btu.TreeNode{
					Val:   2,
					Left:  &btu.TreeNode{Val: 4, Left: nil, Right: nil},
					Right: &btu.TreeNode{Val: 4, Left: nil, Right: nil},
				},
			},
		},
	}

	for _, ex := range examples {
		log.Printf("input:  %+v", ex.tree.InOrderTraversal())
		log.Printf("%+v\n", ex.tree.IsSymmetric())
	}
}
