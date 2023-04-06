package solution18

import "testing"

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	expected := [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}
	actual := fourSum(nums, target)
	if !isEqual(expected, actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

func isEqual(arr1, arr2 [][]int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if !isSame(arr1[i], arr2[i]) {
			return false
		}
	}
	return true
}

func isSame(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestFourSumEmpty(t *testing.T) {
	var nums []int
	target := 0
	var expected [][]int
	actual := fourSum(nums, target)
	if !isEqual(expected, actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

func TestFourSumLengthLessThanFour(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 6
	var expected [][]int
	actual := fourSum(nums, target)
	if !isEqual(expected, actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}
