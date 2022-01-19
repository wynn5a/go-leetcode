package solution219

import "testing"

func Test_case_1(t *testing.T) {
	var (
		nums = []int{1, 2, 3, 1}
		k    = 3
	)
	actual := containsNearbyDuplicate(nums, k)
	if !actual {
		t.Errorf("containsNearbyDuplicate(%v, %d) = %v; expected %v", nums, k, actual, true)
	}
}

func Test_case_2(t *testing.T) {
	var (
		nums = []int{1, 0, 1, 1}
		k    = 1
	)
	actual := containsNearbyDuplicate(nums, k)
	if !actual {
		t.Errorf("containsNearbyDuplicate(%v, %d) = %v; expected %v", nums, k, actual, true)
	}
}

func Test_case_4(t *testing.T) {
	var (
		nums = []int{2, 2}
		k    = 3
	)
	actual := containsNearbyDuplicate(nums, k)
	if !actual {
		t.Errorf("containsNearbyDuplicate(%v, %d) = %v; expected %v", nums, k, actual, true)
	}
}

func Test_case_3(t *testing.T) {
	var (
		nums = []int{1, 2, 3, 1, 2, 3}
		k    = 2
	)
	actual := containsNearbyDuplicate(nums, k)
	if actual {
		t.Errorf("containsNearbyDuplicate(%v, %d) = %v; expected %v", nums, k, actual, false)
	}
}
