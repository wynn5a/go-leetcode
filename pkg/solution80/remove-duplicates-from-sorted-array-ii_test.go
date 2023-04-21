package solution80

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{}, 0},
	}

	for _, test := range tests {
		got := removeDuplicates(test.input)
		if got != test.output {
			t.Errorf("removeDuplicates(%v) = %d; want %d", test.input, got, test.output)
		}
	}
}
