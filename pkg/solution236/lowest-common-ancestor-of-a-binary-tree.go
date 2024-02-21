package solution236

import "github.com/wynn5a/go-leetcode/pkg/util"

// lowestCommonAncestor returns the lowest common ancestor of p and q in the binary tree.
func lowestCommonAncestor(root, p, q *util.TreeNode) *util.TreeNode {
	// If the root is nil or either p or q, return the root.
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// Recursively find the lowest common ancestor in the left and right subtrees.
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// If both nodes are found in the left and right subtrees, the ancestor is the root.
	if left != nil && right != nil {
		return root
	}
	// If one of the nodes is a descendant of the other, the ancestor is the node itself.
	if left != nil {
		return left
	}
	return right
}
