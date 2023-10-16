package solution238

import (
	"reflect"
	"testing"
)

func TestCases(t *testing.T) {
	tests := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			desc: "Case 1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			desc: "Case 2",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
	}

	for _, s := range tests {
		t.Run(s.desc, func(t *testing.T) {
			if got := productExceptSelf(s.nums); !reflect.DeepEqual(got, s.want) {
				t.Errorf("nums: %v, want: %v, got: %v", s.nums, s.want, got)
			}
		})
	}
}
