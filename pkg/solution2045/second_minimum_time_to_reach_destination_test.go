package solution2045

import "testing"

func Test_case_1(t *testing.T) {
	var (
		n        = 5
		edges    = [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}
		time     = 3
		change   = 5
		expected = 13
	)
	actual := secondMinimum(n, edges, time, change)
	if actual != expected {
		t.Errorf("secondMinimum(%d,%v,%d,%d) = %d; expected %d", n, edges, time, change, actual, expected)
	}
}
