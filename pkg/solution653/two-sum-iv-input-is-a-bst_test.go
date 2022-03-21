package solution653

import "testing"

func TestCase1(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 7}
	got := findTarget(root, 9)
	if !got {
		t.Errorf("Case 1: expected %v, got %v", true, got)
	}
}

func TestCase2(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 7}
	got := findTarget(root, 28)
	if got {
		t.Errorf("Case 2: expected %v, got %v", false, got)
	}
}
