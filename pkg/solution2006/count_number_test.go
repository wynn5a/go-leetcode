package solution2006

import "testing"

func Test_case_1(t *testing.T) {
	var (
		nums     = []int{1, 2, 2, 1}
		k        = 1
		expected = 4
	)
	actual := countKDifference(nums, k)
	if actual != expected {
		t.Errorf("countKDifference(%v, %d) = %d; expected %d", nums, k, actual, expected)
	}
}
func Test_case_2(t *testing.T) {
	var (
		nums     = []int{1, 3, 1}
		k        = 1
		expected = 0
	)
	actual := countKDifference(nums, k)
	if actual != expected {
		t.Errorf("countKDifference(%v, %d) = %d; expected %d", nums, k, actual, expected)
	}
}
func Test_case_3(t *testing.T) {
	var (
		nums     = []int{3, 2, 1, 5, 4}
		k        = 2
		expected = 3
	)
	actual := countKDifference(nums, k)
	if actual != expected {
		t.Errorf("countKDifference(%v, %d) = %d; expected %d", nums, k, actual, expected)
	}
}
func Test_case_4(t *testing.T) {
	var (
		nums     = []int{1, 2, 2, 1, 5, 6, 7, 8}
		k        = 3
		expected = 3
	)
	actual := countKDifference(nums, k)
	if actual != expected {
		t.Errorf("countKDifference(%v, %d) = %d; expected %d", nums, k, actual, expected)
	}
}
