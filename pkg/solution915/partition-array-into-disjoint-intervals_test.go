package solution915

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{5, 0, 3, 8, 6}
	expected := 3
	actual := partitionDisjoint(nums)
	if expected != actual {
		t.Errorf("Test failed with actual: %d, expected: %d", actual, expected)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 1, 1, 0, 6, 12}
	expected := 4
	actual := partitionDisjoint(nums)
	if expected != actual {
		t.Errorf("Test failed with actual: %d, expected: %d", actual, expected)
	}
}
