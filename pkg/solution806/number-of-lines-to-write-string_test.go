package solution806

import "testing"

func TestCase1(t *testing.T) {
	lines := numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz")
	if lines[0] != 3 || lines[1] != 60 {
		t.Errorf("Expected [3,60], got %v", lines)
	}
}

func TestCase2(t *testing.T) {
	lines := numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa")
	if lines[0] != 2 || lines[1] != 4 {
		t.Errorf("Expected [2,4], got %v", lines)
	}
}
