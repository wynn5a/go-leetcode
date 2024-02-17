package solution202

import "testing"

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		input  int
		output bool
	}{
		{19, true},
		{2, false},
		{1, true},
	}

	for _, test := range tests {
		result := isHappy(test.input)
		if result != test.output {
			t.Errorf("isHappy(%d) = %t; want %t", test.input, result, test.output)
		}
	}
}
