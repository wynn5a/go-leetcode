package solution258

import "testing"

func TestCase1(t *testing.T) {
	n := 38
	expected := 2
	ans := addDigits(n)
	if ans != expected {
		t.Errorf("Case %d: expected %d, got %d", n, expected, ans)
	}
}

func TestCase2(t *testing.T) {
	n := 0
	expected := 0
	ans := addDigits(n)
	if ans != expected {
		t.Errorf("Case %d: expected %d, got %d", n, expected, ans)
	}
}
