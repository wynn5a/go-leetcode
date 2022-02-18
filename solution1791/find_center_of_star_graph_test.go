package solution1791

import "testing"

func Test_case_1(t *testing.T) {
	var (
		edges    = [][]int{{1, 2}, {2, 3}, {4, 2}}
		expected = 2
	)
	actual := findCenter(edges)
	if actual != expected {
		t.Errorf("findCenter(%v) = %d; expected %d", edges, actual, expected)
	}
}
