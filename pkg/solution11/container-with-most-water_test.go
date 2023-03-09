package solution11

import "testing"

func TestCase1(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	actual := maxArea(height)
	expected := 49

	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
func TestCase2(t *testing.T) {
	height := []int{1, 1}
	actual := maxArea(height)
	expected := 1

	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
