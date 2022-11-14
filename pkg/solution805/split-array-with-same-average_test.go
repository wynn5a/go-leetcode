package solution805

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{3, 1}
	expected := false
	actual := splitArraySameAverage(nums)

	if actual != expected {
		t.Errorf("Test failed, expected: %v, actual: %v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := true
	actual := splitArraySameAverage(nums)

	if actual != expected {
		t.Errorf("Test failed, expected: %v, actual: %v", expected, actual)
	}
}
