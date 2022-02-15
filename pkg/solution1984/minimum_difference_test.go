package solution1984

import "testing"

var TestCases = []struct {
	input    []int
	k        int
	expected int
}{
	{[]int{90}, 1, 0},
	{[]int{9, 4, 1, 7}, 2, 2},
	{[]int{1, 3, 6, 10, 15}, 2, 2},
	{[]int{87063, 61094, 44530, 21297, 95857, 93551, 9918}, 6, 74560},
}

//testing
func TestMinimumDifference(t *testing.T) {
	for _, test := range TestCases {
		if actual := minimumDifference(test.input, test.k); actual != test.expected {
			t.Errorf("minimumDifference(%v, %d) = %d, expected is %d", test.input, test.k, actual, test.expected)
		}
	}
}
