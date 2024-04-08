package solution1431

import "testing"

func TestCases(t *testing.T) {
	cases := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, c := range cases {
		got := kidsWithCandies(c.candies, c.extraCandies)
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("candies: %v, extraCandies: %d, expected: %v, but got: %v", c.candies, c.extraCandies, c.want, got)
			}
		}
	}
}
