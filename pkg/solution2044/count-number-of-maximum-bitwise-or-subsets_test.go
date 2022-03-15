package solution2044

import "testing"

func TestCase1(t *testing.T) {
	n := []int{3, 1}
	expected := 2
	ans := countMaxOrSubsets(n)
	if ans != expected {
		t.Errorf("Case %v: expected %v, got %v", n, expected, ans)
	}
}

func TestCase2(t *testing.T) {
	n := []int{2, 2, 2}
	expected := 7
	ans := countMaxOrSubsets(n)
	if ans != expected {
		t.Errorf("Case %v: expected %v, got %v", n, expected, ans)
	}
}
func TestCase3(t *testing.T) {
	n := []int{3, 2, 1, 5}
	expected := 6
	ans := countMaxOrSubsets(n)
	if ans != expected {
		t.Errorf("Case %v: expected %v, got %v", n, expected, ans)
	}
}
