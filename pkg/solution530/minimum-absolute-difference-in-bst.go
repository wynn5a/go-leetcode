package solution530

import "github.com/wynn5a/go-leetcode/pkg/util"

func getMinimumDifference(root *util.TreeNode) int {
	// In-order traversal
	var (
		ans, pre int
		dfs      func(*util.TreeNode)
	)
	ans = 1<<63 - 1
	pre = -1
	dfs = func(node *util.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
