package solution687

import "github.com/wynn5a/go-leetcode/pkg/util"

func longestUnivaluePath(root *util.TreeNode) int {
	var max *int
	max = new(int)
	*max = 0
	longest(root, max)
	return *max
}

func longest(root *util.TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := longest(root.Left, max)
	right := longest(root.Right, max)

	eqLeft := false
	if root.Left != nil && root.Val == root.Left.Val {
		left += 1
		eqLeft = true
	} else {
		left = 0
	}
	eqRight := false
	if root.Right != nil && root.Val == root.Right.Val {
		right += 1
		eqRight = true
	} else {
		right = 0
	}

	result := maxOf(left, right)
	if eqRight && eqLeft && *max < left+right {
		*max = left + right
	}
	if *max < result {
		*max = result
	}

	return result
}

func maxOf(left int, right int) int {
	if left < right {
		return right
	}
	return left
}
