package solution226

import "github.com/wynn5a/go-leetcode/pkg/util"

func invertTree(root *util.TreeNode) *util.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
