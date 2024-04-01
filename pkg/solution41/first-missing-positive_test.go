package solution41

import "testing"

func TestCases(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
	}

	for _, c := range cases {
		got := firstMissingPositive(c.nums)
		if got != c.want {
			t.Errorf("nums: %v, expected: %v, got: %v", c.nums, c.want, got)
		}
	}
}
