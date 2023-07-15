package main

import (
	"log"

	btu "github.com/mbeaver502/leetcode/go/binarytreeutils"
)

func main() {
	tree := &btu.TreeNode{
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
	}

	log.Println(tree.InOrderTraversal())
}

func hasPathSum(root *TreeNode, targetSum int) bool {

}
