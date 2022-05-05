package solution713

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{10, 5, 2, 6}
	k := 100

	expected := 8
	actual := numSubarrayProductLessThanK(nums, k)

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
