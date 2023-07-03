package solution16

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	// Test Case 1: Basic test case
	nums := []int{-1, 2, 1, -4}
	target := 1
	expected := 2
	result := threeSumClosest(nums, target)
	if result != expected {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}

	// Test Case 2: Test case with negative numbers
	nums = []int{-5, -2, 3, 6, 11}
	target = 1
	expected = -1
	result = threeSumClosest(nums, target)
	if result != expected {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}

	// Test Case 3: Test case with positive numbers
	nums = []int{1, 4, 8, 12, 16}
	target = 20
	expected = 21
	result = threeSumClosest(nums, target)
	if result != expected {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}

	// Test Case 4: Test case with duplicate numbers
	nums = []int{2, 2, 2, 2, 2}
	target = 8
	expected = 6
	result = threeSumClosest(nums, target)
	if result != expected {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}

	// Test Case 5: Test case with empty array
	nums = []int{}
	target = 5
	expected = 0
	result = threeSumClosest(nums, target)
	if result != expected {
		t.Errorf("Expected: %d, but got: %d", expected, result)
	}
}
