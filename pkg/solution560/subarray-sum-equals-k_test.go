package solution560

import "testing"

func TestSubarraySum(t *testing.T) {
	// 创建一个结构体切片进行测试
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{-1, -1, 1}, 1, 1},
	}

	// 循环切片进行测试
	for _, tt := range tests {
		if got := subarraySum(tt.nums, tt.k); got != tt.want {
			t.Errorf("subarraySum(%v, %v): got %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
