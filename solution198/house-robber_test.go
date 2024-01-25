package solution198

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
		{[]int{}, 0},
		{[]int{1}, 1},
	}

	for _, test := range tests {
		result := rob(test.input)
		if result != test.output {
			t.Errorf("rob(%v) = %d; want %d", test.input, result, test.output)
		}
	}
}
