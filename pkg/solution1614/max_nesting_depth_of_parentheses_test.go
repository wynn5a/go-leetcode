package solution1614

import "testing"

func Test_case_1(t *testing.T) {
	var (
		s        = "(1+(2*3)+((8)/4))+1"
		expected = 3
	)
	actual := maxDepth(s)
	if actual != expected {
		t.Errorf("maxDepth(%v) = %d; expected %d", s, actual, expected)
	}
}
func Test_case_2(t *testing.T) {
	var (
		s        = "(1)+((2))+(((3)))"
		expected = 3
	)
	actual := maxDepth(s)
	if actual != expected {
		t.Errorf("maxDepth(%v) = %d; expected %d", s, actual, expected)
	}
}
func Test_case_3(t *testing.T) {
	var (
		s        = "1+(2*3)/(2-1)"
		expected = 1
	)
	actual := maxDepth(s)
	if actual != expected {
		t.Errorf("maxDepth(%v) = %d; expected %d", s, actual, expected)
	}
}
func Test_case_4(t *testing.T) {
	var (
		s        = "1+2*3/2-1"
		expected = 0
	)
	actual := maxDepth(s)
	if actual != expected {
		t.Errorf("maxDepth(%v) = %d; expected %d", s, actual, expected)
	}
}
