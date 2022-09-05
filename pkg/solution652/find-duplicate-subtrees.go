package solution652

import (
	"fmt"
	"github.com/wynn5a/go-leetcode/pkg/util"
)

func findDuplicateSubtrees(root *util.TreeNode) []*util.TreeNode {
	repeat := map[*util.TreeNode]struct{}{}
	seen := map[string]*util.TreeNode{}
	var dfs func(*util.TreeNode) string
	dfs = func(node *util.TreeNode) string {
		if node == nil {
			return ""
		}
		serial := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
		if n, ok := seen[serial]; ok {
			repeat[n] = struct{}{}
		} else {
			seen[serial] = node
		}
		return serial
	}
	dfs(root)
	ans := make([]*util.TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}
