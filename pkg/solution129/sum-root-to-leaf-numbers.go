package solution129

import "github.com/wynn5a/go-leetcode/pkg/util"

func sumNumbers(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	return sum(root, 0)
}

// sum returns the sum of all root-to-leaf numbers.
func sum(node *util.TreeNode, i int) int {
	if node == nil {
		return 0
	}
	// val is the value of the current node.
	val := i*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	return sum(node.Left, val) + sum(node.Right, val)
}
