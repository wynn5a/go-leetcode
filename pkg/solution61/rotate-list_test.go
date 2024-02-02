package solution61

import (
	"fmt"
	"github.com/wynn5a/go-leetcode/pkg/util"
	"reflect"
	"testing"
)

func TestRotateList(t *testing.T) {
	cases := []struct {
		input  []int
		k      int
		output []int
	}{
		{input: []int{1, 2, 3, 4, 5}, k: 2, output: []int{4, 5, 1, 2, 3}},
		{input: []int{0, 1, 2}, k: 4, output: []int{2, 0, 1}},
		{input: []int{1, 2}, k: 1, output: []int{2, 1}},
		{input: []int{1}, k: 0, output: []int{1}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("case: %d", i), func(t *testing.T) {
			head := util.NewListNodes(test.input)
			result := rotateRight(head, test.k)
			if !reflect.DeepEqual(util.ListNodeToArray(result), test.output) {
				t.Errorf("rotateRight(%v, %d) = %v; want %v", test.input, test.k, util.ListNodeToArray(result), test.output)
			}
		})
	}
}
