package solution1775

import "testing"

func TestCase1(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{1, 1, 2, 2, 2, 2}

	actual := minOperations(nums1, nums2)
	expect := 3

	if actual != expect {
		t.Errorf("Test failed, expected %d, got: %d", expect, actual)
	}
}

func TestCase2(t *testing.T) {
	nums1 := []int{1, 1, 1, 1, 1, 1, 1}
	nums2 := []int{6}

	actual := minOperations(nums1, nums2)
	expect := -1

	if actual != expect {
		t.Errorf("Test failed, expected %d, got: %d", expect, actual)
	}
}
