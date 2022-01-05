package solution1576

import "testing"

func Test_case_1(t *testing.T) {
	var (
		s        = "j?qg??b"
		expected = "jaqgacb"
	)
	actual := modifyString(s)
	if actual != expected {
		t.Errorf("modifyString(%v) = %v; expected %v", s, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		s        = "?qg??b"
		expected = "aqgacb"
	)
	actual := modifyString(s)
	if actual != expected {
		t.Errorf("modifyString(%v) = %v; expected %v", s, actual, expected)
	}
}
