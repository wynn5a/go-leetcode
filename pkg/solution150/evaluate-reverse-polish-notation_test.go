package solution150

import (
	"fmt"
	"testing"
)

func TestCase(t *testing.T) {
	cases := []struct {
		input  []string
		output int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{input: []string{"4", "13", "5", "/", "+"}, output: 6},
		{input: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, output: 22},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			result := evalRPN(test.input)
			if result != test.output {
				t.Errorf("evalRPN(%v) = %d; want %d", test.input, result, test.output)
			}
		})
	}
}
