package solution530

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestCases(t *testing.T) {
	root := &util.TreeNode{Val: 1, Right: &util.TreeNode{Val: 3, Left: &util.TreeNode{Val: 2}}}
	if getMinimumDifference(root) != 1 {
		t.Errorf("Expected 1, but got %d", getMinimumDifference(root))
	}
	root = &util.TreeNode{Val: 5, Left: &util.TreeNode{Val: 4}, Right: &util.TreeNode{Val: 7}}
	if getMinimumDifference(root) != 1 {
		t.Errorf("Expected 1, but got %d", getMinimumDifference(root))
	}

	root = &util.TreeNode{Val: 236, Left: &util.TreeNode{Val: 104, Right: &util.TreeNode{Val: 227}}, Right: &util.TreeNode{Val: 701, Right: &util.TreeNode{Val: 911}}}
	if getMinimumDifference(root) != 9 {
		t.Errorf("Expected 9, but got %d", getMinimumDifference(root))
	}

	root = &util.TreeNode{Val: 1, Right: &util.TreeNode{Val: 5, Left: &util.TreeNode{Val: 3}}}
	if getMinimumDifference(root) != 2 {
		t.Errorf("Expected 2, but got %d", getMinimumDifference(root))
	}
}
