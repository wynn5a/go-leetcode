package solution373

import (
	"reflect"
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		nums1    = []int{1, 7, 11}
		nums2    = []int{2, 4, 6}
		k        = 3
		excepted = [][]int{{1, 2}, {1, 4}, {1, 6}}
	)
	actual := kSmallestPairs(nums1, nums2, k)
	if !reflect.DeepEqual(actual, excepted) {
		t.Errorf("kSmallestPairs(%v, %v, %d) = %v; expected %v", nums1, nums2, k, actual, false)
	}
}

func Test_case_2(t *testing.T) {
	var (
		nums1    = []int{1, 1, 2}
		nums2    = []int{1, 2, 3}
		k        = 2
		excepted = [][]int{{1, 1}, {1, 1}}
	)
	actual := kSmallestPairs(nums1, nums2, k)
	if !reflect.DeepEqual(actual, excepted) {
		t.Errorf("kSmallestPairs(%v, %v, %d) = %v; expected %v", nums1, nums2, k, actual, false)
	}
}

func Test_case_3(t *testing.T) {
	var (
		nums1    = []int{1, 2}
		nums2    = []int{3}
		k        = 3
		excepted = [][]int{{1, 3}, {2, 3}}
	)
	actual := kSmallestPairs(nums1, nums2, k)
	if !reflect.DeepEqual(actual, excepted) {
		t.Errorf("kSmallestPairs(%v, %v, %d) = %v; expected %v", nums1, nums2, k, actual, false)
	}
}
