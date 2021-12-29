package solution1995

import "testing"

func Test_case_1(t *testing.T) {
	var (
		nums     = []int{1, 2, 3, 6}
		expected = 1
	)
	actual := countQuadruplets(nums)
	if actual != expected {
		t.Errorf("countQuadruplets(%v) = %d; expected %d", nums, actual, expected)
	}
}
func Test_case_2(t *testing.T) {
	var (
		nums     = []int{3, 3, 6, 4, 5}
		expected = 0
	)
	actual := countQuadruplets(nums)
	if actual != expected {
		t.Errorf("countQuadruplets(%v) = %d; expected %d", nums, actual, expected)
	}
}
func Test_case_3(t *testing.T) {
	var (
		nums     = []int{1, 1, 1, 3, 5}
		expected = 4
	)
	actual := countQuadruplets(nums)
	if actual != expected {
		t.Errorf("countQuadruplets(%v) = %d; expected %d", nums, actual, expected)
	}
}
