package solution114

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	root := &util.TreeNode{
		Val: 1,
		Left: &util.TreeNode{
			Val:   2,
			Left:  &util.TreeNode{Val: 3},
			Right: &util.TreeNode{Val: 4},
		},
		Right: &util.TreeNode{
			Val:   5,
			Right: &util.TreeNode{Val: 6},
		},
	}
	expected := []int{1, 2, 3, 4, 5, 6}
	flatten(root)
	actual := make([]int, 0)
	for root != nil {
		actual = append(actual, root.Val)
		root = root.Right
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
