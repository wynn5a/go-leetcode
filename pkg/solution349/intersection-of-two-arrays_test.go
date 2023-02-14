package solution349

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{1, 2}
	rst := []int{1, 2}
	actual := intersection(nums1, nums2)
	if !reflect.DeepEqual(actual, rst) {
		t.Errorf("expected: %v,  actual: %v \n", rst, actual)
	}
}

func TestCase2(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	rst := []int{9, 4}
	actual := intersection(nums1, nums2)
	if !reflect.DeepEqual(actual, rst) {
		t.Errorf("expected: %v,  actual: %v \n", rst, actual)
	}
}
