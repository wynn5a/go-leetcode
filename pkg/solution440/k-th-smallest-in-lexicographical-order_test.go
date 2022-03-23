package solution440

import "testing"

func TestCase1(t *testing.T) {
	n := 13
	k := 2
	expected := 10
	ans := findKthNumber(n, k)
	if ans != expected {
		t.Errorf("Case 1: expected %d, got %d", expected, ans)
	}
}

func TestCase2(t *testing.T) {
	n := 1
	k := 1
	expected := 1
	ans := findKthNumber(n, k)
	if ans != expected {
		t.Errorf("Case 2: expected %d, got %d", expected, ans)
	}
}
