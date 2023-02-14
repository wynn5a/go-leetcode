package solution53

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	actual := maxSubArray(nums)
	expected := 6
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1}
	actual := maxSubArray(nums)
	expected := 1
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	actual := maxSubArray(nums)
	expected := 23
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
