package solution543

import "testing"

func TestCase1(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	actual := diameterOfBinaryTree(root)
	if actual != 3 {
		t.Errorf("Case %v, expect %v, but got %d", 1, 3, actual)
	}
}

func TestCase2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 7}},
			Right: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 5,
					Right: &TreeNode{Val: 6}},
			},
		},
	}

	actual := diameterOfBinaryTree(root)
	if actual != 5 {
		t.Errorf("Case %v, expect %v, but got %d", 2, 5, actual)
	}
}
