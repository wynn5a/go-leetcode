package solution1994

import "testing"

func TestCase1(t *testing.T) {
	var (
		nums     = []int{1, 2, 3, 4}
		expected = 6
	)

	actual := numberOfGoodSubsets(nums)
	if actual != expected {
		t.Errorf("numberOfGoodSubsets(%v) = %v, expected:  %v", nums, actual, expected)
	}
}
