package solution55

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	actual := canJump(nums)
	if !actual {
		t.Errorf("expected: true, actual: false")
	}
}

func TestCase2(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	actual := canJump(nums)
	if actual {
		t.Errorf("expected: false, actual: true")
	}
}

func TestCase3(t *testing.T) {
	nums := []int{2, 0, 0}
	actual := canJump(nums)
	if !actual {
		t.Errorf("expected: true, actual: false")
	}
}
