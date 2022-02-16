package solution1719

import "testing"

func Test_case_1(t *testing.T) {
	var (
		pair     = [][]int{{1, 2}, {2, 3}}
		expected = 1
	)
	actual := checkWays(pair)
	if actual != expected {
		t.Errorf("checkWays(%v) = %d; expected %d", pair, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		pair     = [][]int{{1, 2}, {2, 3}, {1, 3}}
		expected = 2
	)
	actual := checkWays(pair)
	if actual != expected {
		t.Errorf("checkWays(%v) = %d; expected %d", pair, actual, expected)
	}
}
func Test_case_3(t *testing.T) {
	var (
		pair     = [][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}
		expected = 0
	)
	actual := checkWays(pair)
	if actual != expected {
		t.Errorf("checkWays(%v) = %d; expected %d", pair, actual, expected)
	}
}
