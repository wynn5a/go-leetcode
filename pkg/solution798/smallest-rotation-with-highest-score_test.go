package solution798

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{2, 3, 1, 4, 0}
	expected := 3
	actual := bestRotation(nums)
	if actual != expected {
		t.Errorf("Case 1: expected: %v, actual: %v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 3, 0, 2, 4}
	expected := 0
	actual := bestRotation(nums)
	if actual != expected {
		t.Errorf("Case 2: expected: %v, actual: %v", expected, actual)
	}
}
