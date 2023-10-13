package solution805

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{3, 1}
	actual := splitArraySameAverage(nums)

	if actual {
		t.Errorf("Test failed, expected: %v, actual: %v", false, actual)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	actual := splitArraySameAverage(nums)

	if !actual {
		t.Errorf("Test failed, expected: %v, actual: %v", true, actual)
	}
}
