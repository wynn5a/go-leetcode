package solution70

import "testing"

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}

	for _, test := range tests {
		result := climbStairs(test.input)
		if result != test.output {
			t.Errorf("climbStairs(%d) = %d; want %d", test.input, result, test.output)
		}
	}
}
