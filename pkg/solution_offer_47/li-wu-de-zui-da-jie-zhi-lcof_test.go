package solution_offer_47

import "testing"

func TestCase1(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	value := maxValue(grid)
	expected := 12

	if value != expected {
		t.Errorf("expected: %d, actual: %d", expected, value)
	}
}
