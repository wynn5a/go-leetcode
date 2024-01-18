package solution114

import "github.com/wynn5a/go-leetcode/pkg/util"

func flatten(root *util.TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		// find the rightmost node of the left subtree
		rightmost := root.Left
		for rightmost.Right != nil {
			rightmost = rightmost.Right
		}
		// connect the rightmost node to the right subtree
		rightmost.Right = root.Right
		// connect the left subtree to the right subtree
		root.Right = root.Left
		// set the left subtree to nil
		root.Left = nil
	}
}
