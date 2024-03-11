package solution77

import (
	"reflect"
	"testing"
)

// TestCombine 测试 combine 函数
func TestCombine(t *testing.T) {
	testcases := []struct {
		n, k     int
		expected [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{1, 1, [][]int{{1}}},
		{3, 4, [][]int{}},
		{5, 5, [][]int{{1, 2, 3, 4, 5}}},
		{3, 0, [][]int{{}}}, // 注意这里，k为0我们期待一个包含空切片的切片
		{6, 3, [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 2, 6}, {1, 3, 4}, {1, 3, 5}, {1, 3, 6}, {1, 4, 5}, {1, 4, 6}, {1, 5, 6}, {2, 3, 4}, {2, 3, 5}, {2, 3, 6}, {2, 4, 5}, {2, 4, 6}, {2, 5, 6}, {3, 4, 5}, {3, 4, 6}, {3, 5, 6}, {4, 5, 6}}},
	}

	for _, tc := range testcases {
		got := combine(tc.n, tc.k)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("combine(%v, %v) = %v, expected %v", tc.n, tc.k, got, tc.expected)
		}
	}
}
