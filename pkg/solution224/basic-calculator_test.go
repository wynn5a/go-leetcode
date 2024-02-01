package solution224

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"1-(+1+1)", -1},
		{"1-(+1+1)+2", 1},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("case: %d", i), func(t *testing.T) {
			result := calculate(test.input)
			if result != test.output {
				t.Errorf("calculate(%s) = %d; want %d", test.input, result, test.output)
			}
		})
	}
}
