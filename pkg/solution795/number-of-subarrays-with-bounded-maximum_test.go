package solution795

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{2, 1, 4, 3}
	left, right := 2, 3
	actual := numSubarrayBoundedMax(nums, left, right)
	expect := 3

	if actual != expect {
		t.Errorf("Test case 1 failed, expected: %d, actual: %d \n", expect, actual)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{2, 9, 2, 5, 6}
	left, right := 2, 8
	actual := numSubarrayBoundedMax(nums, left, right)
	expect := 7

	if actual != expect {
		t.Errorf("Test case 2 failed, expected: %d, actual: %d \n", expect, actual)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}
	left, right := 32, 69
	actual := numSubarrayBoundedMax(nums, left, right)
	expect := 22

	if actual != expect {
		t.Errorf("Test case 3 failed, expected: %d, actual: %d \n", expect, actual)
	}
}
