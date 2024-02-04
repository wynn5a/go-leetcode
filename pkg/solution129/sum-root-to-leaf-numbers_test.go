package solution129

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		name string
		root *util.TreeNode
		want int
	}{
		{
			name: "Test Case 1",
			root: &util.TreeNode{
				Val: 1,
				Left: &util.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &util.TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: 25,
		},
		{
			name: "Test Case 2",
			root: &util.TreeNode{
				Val: 4,
				Left: &util.TreeNode{
					Val: 9,
					Left: &util.TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: &util.TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &util.TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			want: 1026,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
