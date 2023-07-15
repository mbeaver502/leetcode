// https://leetcode.com/problems/symmetric-tree/
// Given the root of a binary tree, check whether
// it is a mirror of itself (i.e., symmetric around its center).
package main

import "log"

type example struct {
	tree *TreeNode
}

func main() {
	examples := []example{
		{
			nil,
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
			},
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
					Right: nil,
				},
			},
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
			},
		},
	}

	for _, ex := range examples {
		log.Printf("input:  %+v", ex.tree)
		log.Printf("%+v\n", isSymmetric(ex.tree))
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// We need to check symmetry around the given `root`.
// So let's check each side of the tree by splitting
// the tree and traversing the left and right sides.
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	// We want the outer branches to match each other
	// and for the inner branches to match each other.
	if left != nil && right != nil {
		if left.Val == right.Val {
			return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
		}
	}

	return false
}
