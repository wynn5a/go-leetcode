package solution134

import "testing"

func TestCase1(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	actual := canCompleteCircuit(gas, cost)
	expected := 3
	if actual != expected {
		t.Errorf("Test failed, expected: %d, actual: %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	actual := canCompleteCircuit(gas, cost)
	expected := -1
	if actual != expected {
		t.Errorf("Test failed, expected: %d, actual: %d", expected, actual)
	}
}
