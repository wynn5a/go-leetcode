package solution652

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestCase1(t *testing.T) {
	actual := findDuplicateSubtrees(util.Array2TreeNode("[1,2,3,4,null,2,4,null,null,4]"))
	//[[2,4],[4]]
	for _, node := range actual {
		if node.Val == 2 {
			if node.Left == nil || node.Left.Val != 4 {
				t.Errorf("fail")
			}
		} else if node.Val != 4 {
			t.Errorf("fail")
		}
	}
}
