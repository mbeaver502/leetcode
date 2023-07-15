package binarytreeutils

// TreeNode defines a binary tree.
// Provided by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InOrderTraversal returns a slice of ints produced
// via an in-order traversal of the tree t.
func (t *TreeNode) InOrderTraversal() []int {
	return inorderTraversal(t)
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

// PreOrderTraversal returns a slice of ints produced
// via a pre-order traversal of the tree t.
func (t *TreeNode) PreOrderTraversal() []int {
	return preorderTraversal(t)
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	parent := []int{root.Val}
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	return append(parent, append(left, right...)...)
}

// PostOrderTraversal returns a slice of ints produced
// via a post-order traversal of the tree t.
func (t *TreeNode) PostOrderTraversal() []int {
	return postorderTraversal(t)
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	parent := []int{root.Val}
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)

	return append(left, append(right, parent...)...)
}

// MaxDepth returns the max depth of the tree t.
func (t *TreeNode) MaxDepth() int {
	return maxDepth(t)
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

// SortedArrayToBST returns a binary tree that
// represents the given sorted list nums.
func SortedArrayToBST(nums []int) *TreeNode {
	lenNums := len(nums)

	switch {
	// If we got an empty list, our tree (or leaf) is nil.
	case lenNums < 1:
		return nil

	// We found a leaf!
	case lenNums == 1:
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}

	// We want our tree to be height-balanced,
	// so start each partition at the halfway point.
	halfway := len(nums) / 2
	return &TreeNode{
		Val:   nums[halfway],
		Left:  SortedArrayToBST(nums[:halfway]),
		Right: SortedArrayToBST(nums[halfway+1:]),
	}
}

// IsSameTree determines whether the binary tree q
// is the same binary tree as t.
func (t *TreeNode) IsSameTree(q *TreeNode) bool {
	return isSameTree(t, q)
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

// IsSymmetric determines whether the binary tree t is symmetric.
func (t *TreeNode) IsSymmetric() bool {
	return isSymmetric(t)
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

// HasPathSum determines whether the binary tree t
// has a root-to-leaf path that sums to targetSum.
func (t *TreeNode) HasPathSum(targetSum int) bool {
	return hasPathSum(t, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// Make sure we reach a leaf!
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
