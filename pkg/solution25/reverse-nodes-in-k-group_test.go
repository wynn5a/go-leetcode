package solution25

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"reflect"
	"testing"
)

func TestCases(t *testing.T) {
	cases := []struct {
		head   []int
		k      int
		expect []int
	}{
		{
			head:   []int{1, 2, 3, 4, 5},
			k:      2,
			expect: []int{2, 1, 4, 3, 5},
		},
		{
			head:   []int{1, 2, 3, 4, 5},
			k:      3,
			expect: []int{3, 2, 1, 4, 5},
		},
		{
			head:   []int{1, 2, 3, 4, 5},
			k:      1,
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			head:   []int{1},
			k:      1,
			expect: []int{1},
		},
		{
			head:   []int{1, 2},
			k:      2,
			expect: []int{2, 1},
		},
	}

	for _, c := range cases {
		head := util.NewListNodes(c.head)
		result := reverseKGroup(head, c.k)
		resultArray := util.ListNodeToArray(result)
		if !reflect.DeepEqual(resultArray, c.expect) {
			t.Errorf("reverseKGroup(%v, %d) = %v; expect %v", c.head, c.k, resultArray, c.expect)
		}
	}
}
