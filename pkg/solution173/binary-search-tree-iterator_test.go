package solution173

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestBSTIterator(t *testing.T) {
	i := Constructor(&util.TreeNode{
		Val: 7,
		Left: &util.TreeNode{
			Val: 3,
		},
		Right: &util.TreeNode{
			Val: 15,
			Left: &util.TreeNode{
				Val: 9,
			},
			Right: &util.TreeNode{
				Val: 20,
			},
		},
	})
	if i.Next() != 3 {
		t.Errorf("i.Next() = %d; want 3", i.Next())
	}
	if i.Next() != 7 {
		t.Errorf("i.Next() = %d; want 7", i.Next())
	}
	if i.HasNext() != true {
		t.Errorf("i.HasNext() = %t; want %t", i.HasNext(), true)
	}
	if i.Next() != 9 {
		t.Errorf("i.Next() = %d; want 9", i.Next())
	}
	if i.HasNext() != true {
		t.Errorf("i.HasNext() = %t; want %t", i.HasNext(), true)
	}
	if i.Next() != 15 {
		t.Errorf("i.Next() = %d; want 15", i.Next())
	}
	if i.HasNext() != true {
		t.Errorf("i.HasNext() = %t; want %t", i.HasNext(), true)
	}
	if i.Next() != 20 {
		t.Errorf("i.Next() = %d; want 20", i.Next())
	}
	if i.HasNext() != false {
		t.Errorf("i.HasNext() = %t; want %t", i.HasNext(), false)
	}
}
