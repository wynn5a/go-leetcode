package solution1044

import "testing"

func Test_case_1(t *testing.T) {
	var (
		a        = "banana"
		expected = "ana"
	)
	actual := longestDupSubstring(a)
	if actual != expected {
		t.Errorf("longestDupSubstring(%v) = %v; expected %v", a, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		a        = "abcd"
		expected = ""
	)
	actual := longestDupSubstring(a)
	if actual != expected {
		t.Errorf("longestDupSubstring(%v) = %v; expected %v", a, actual, expected)
	}
}
