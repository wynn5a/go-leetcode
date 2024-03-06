package solution226

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"reflect"
	"testing"
)

func TestCases(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{[]int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}

	for _, c := range cases {
		root := util.IntArray2TreeNode(c.input)
		result := invertTree(root)
		resultArray := util.Tree2IntArray(result)
		if !reflect.DeepEqual(resultArray, c.output) {
			t.Errorf("invertTree(%v) = %v; expect %v", c.input, resultArray, c.output)
		}
	}
}
