package solution128

import "testing"

func TestCase(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	expected := 4
	actual := longestConsecutive(nums)
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 0, 1}
	expected := 3
	actual := longestConsecutive(nums)
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{1, 2, 0, 1, 3}
	expected := 4
	actual := longestConsecutive(nums)
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
