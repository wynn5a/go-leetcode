package solution687

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"testing"
)

func TestCase1(t *testing.T) {
	input := util.Array2TreeNode("[5,4,5,1,1,5]")

	actual := longestUnivaluePath(input)
	if actual != 2 {
		t.Errorf("Case %v, expect %v, but got %d", 1, 2, actual)
	}
}

func TestCase2(t *testing.T) {
	input := util.Array2TreeNode("[1,4,5,4,4,5]")

	actual := longestUnivaluePath(input)
	if actual != 2 {
		t.Errorf("Case %v, expect %v, but got %d", 2, 2, actual)
	}
}

func TestCase3(t *testing.T) {
	input := util.Array2TreeNode("[1]")

	actual := longestUnivaluePath(input)
	if actual != 0 {
		t.Errorf("Case %v, expect %v, but got %d", 3, 0, actual)
	}
}

func TestCase4(t *testing.T) {
	input := util.Array2TreeNode("[1,null,1,1,1,1,1,1]")

	actual := longestUnivaluePath(input)
	if actual != 4 {
		t.Errorf("Case %v, expect %v, but got %d", 4, 4, actual)
	}
}
