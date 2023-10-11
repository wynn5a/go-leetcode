package solution135

import "testing"

func TestCases(t *testing.T) {
	args := []struct {
		ratings  []int
		expected int
	}{
		{ratings: []int{1, 0, 2}, expected: 5},
		{ratings: []int{1, 2, 2}, expected: 4},
	}

	for _, arg := range args {
		actual := candy(arg.ratings)
		if actual != arg.expected {
			t.Errorf("Test failed for ratings: %v, expected: %d, actual: %d", arg.ratings, arg.expected, actual)
		}
	}
}
