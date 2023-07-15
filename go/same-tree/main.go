// https://leetcode.com/problems/same-tree/
// Given the roots of two binary trees p and q,
// write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical,
// and the nodes have the same value.
package main

import "log"

type example struct {
	p *TreeNode
	q *TreeNode
}

func main() {
	examples := []example{
		{
			p: nil, q: nil,
		},
		{
			p: nil, q: &TreeNode{Val: 1, Left: nil, Right: nil},
		},
		{
			p: &TreeNode{Val: 1, Left: nil, Right: nil}, q: nil,
		},
		{
			p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		},
		{
			p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil},
			q: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
		},
		{
			p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 1, Left: nil, Right: nil}},
			q: &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
		},
		{
			p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: &TreeNode{Val: 5, Left: nil, Right: nil}}},
			q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: &TreeNode{Val: 5, Left: nil, Right: nil}}},
		},
	}

	for _, ex := range examples {
		log.Printf("input:  p = %+v, q = %+v", inorderTraversal(ex.p), inorderTraversal(ex.q))
		log.Printf("output: %v", isSameTree(ex.p, ex.q))
	}

}

// TreeNode defines a binary tree.
// Provided by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		if p.Val == q.Val {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}
	return false
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
