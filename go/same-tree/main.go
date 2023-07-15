// https://leetcode.com/problems/same-tree/
// Given the roots of two binary trees p and q,
// write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical,
// and the nodes have the same value.
package main

import (
	"log"

	btu "github.com/mbeaver502/leetcode/go/binarytreeutils"
)

type example struct {
	p *btu.TreeNode
	q *btu.TreeNode
}

func main() {
	examples := []example{
		{
			p: nil, q: nil,
		},
		{
			p: nil, q: &btu.TreeNode{Val: 1, Left: nil, Right: nil},
		},
		{
			p: &btu.TreeNode{Val: 1, Left: nil, Right: nil}, q: nil,
		},
		{
			p: &btu.TreeNode{Val: 1, Left: &btu.TreeNode{Val: 2, Left: nil, Right: nil}, Right: &btu.TreeNode{Val: 3, Left: nil, Right: nil}},
			q: &btu.TreeNode{Val: 1, Left: &btu.TreeNode{Val: 2, Left: nil, Right: nil}, Right: &btu.TreeNode{Val: 3, Left: nil, Right: nil}},
		},
		{
			p: &btu.TreeNode{Val: 1, Left: &btu.TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil},
			q: &btu.TreeNode{Val: 1, Left: nil, Right: &btu.TreeNode{Val: 2, Left: nil, Right: nil}},
		},
		{
			p: &btu.TreeNode{Val: 1, Left: &btu.TreeNode{Val: 2, Left: nil, Right: nil}, Right: &btu.TreeNode{Val: 1, Left: nil, Right: nil}},
			q: &btu.TreeNode{Val: 1, Left: &btu.TreeNode{Val: 1, Left: nil, Right: nil}, Right: &btu.TreeNode{Val: 2, Left: nil, Right: nil}},
		},
		{
			p: &btu.TreeNode{Val: 1, Left: &btu.TreeNode{Val: 2, Left: &btu.TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}, Right: &btu.TreeNode{Val: 3, Left: nil, Right: &btu.TreeNode{Val: 5, Left: nil, Right: nil}}},
			q: &btu.TreeNode{Val: 1, Left: &btu.TreeNode{Val: 2, Left: &btu.TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}, Right: &btu.TreeNode{Val: 3, Left: nil, Right: &btu.TreeNode{Val: 5, Left: nil, Right: nil}}},
		},
	}

	for _, ex := range examples {
		log.Printf("input:  p = %+v, q = %+v", ex.p.InOrderTraversal(), ex.q.InOrderTraversal())
		log.Printf("output: %v", ex.p.IsSameTree(ex.q))
	}
}
