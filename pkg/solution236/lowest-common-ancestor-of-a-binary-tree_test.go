package solution236

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestCase1(t *testing.T) {
	root := &util.TreeNode{
		Val: 3,
		Left: &util.TreeNode{
			Val: 5,
			Left: &util.TreeNode{
				Val: 6,
			},
			Right: &util.TreeNode{
				Val: 2,
				Left: &util.TreeNode{
					Val: 7,
				},
				Right: &util.TreeNode{
					Val: 4,
				},
			},
		},
		Right: &util.TreeNode{
			Val: 1,
			Left: &util.TreeNode{
				Val: 0,
			},
			Right: &util.TreeNode{
				Val: 8,
			},
		},
	}
	p := &util.TreeNode{Val: 5}
	q := &util.TreeNode{Val: 1}
	result := lowestCommonAncestor(root, p, q)
	if result.Val != 3 {
		t.Errorf("lowestCommonAncestor(%v, %v, %v) = %d; want 3", root, p, q, result.Val)
	}
}

func TestCase2(t *testing.T) {
	root := &util.TreeNode{
		Val: 3,
		Left: &util.TreeNode{
			Val: 5,
			Left: &util.TreeNode{
				Val: 6,
			},
			Right: &util.TreeNode{
				Val: 2,
				Left: &util.TreeNode{
					Val: 7,
				},
				Right: &util.TreeNode{
					Val: 4,
				},
			},
		},
		Right: &util.TreeNode{
			Val: 1,
			Left: &util.TreeNode{
				Val: 0,
			},
			Right: &util.TreeNode{
				Val: 8,
			},
		},
	}
	p := &util.TreeNode{Val: 5}
	q := &util.TreeNode{Val: 4}
	result := lowestCommonAncestor(root, p, q)
	if result.Val != 5 {
		t.Errorf("lowestCommonAncestor(%v, %v, %v) = %d; want 5", root, p, q, result.Val)
	}
}

func TestCase3(t *testing.T) {
	root := &util.TreeNode{
		Val: 1,
		Left: &util.TreeNode{
			Val: 2,
		},
	}
	p := &util.TreeNode{Val: 1}
	q := &util.TreeNode{Val: 2}
	result := lowestCommonAncestor(root, p, q)
	if result.Val != 1 {
		t.Errorf("lowestCommonAncestor(%v, %v, %v) = %d; want 1", root, p, q, result.Val)
	}
}
