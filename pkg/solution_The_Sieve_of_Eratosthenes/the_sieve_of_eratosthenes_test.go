package solution_The_Sieve_of_Eratosthenes

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected []int
	}{

		{
			name:     "test0",
			input:    2,
			expected: []int{2},
		},

		{
			name:     "test1",
			input:    10,
			expected: []int{2, 3, 5, 7},
		},
		{
			name:     "test2",
			input:    20,
			expected: []int{2, 3, 5, 7, 11, 13, 17, 19},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := sieve(c.input)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("sieve(%d) == %v, expected %v", c.input, actual, c.expected)
			}
		})
	}
}
