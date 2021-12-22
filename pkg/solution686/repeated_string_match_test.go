package solution686

import "testing"

func Test_case_1(t *testing.T) {
	var (
		a        = "abcd"
		b        = "cdabcdab"
		expected = 3
	)
	actual := repeatedStringMatch(a, b)
	if actual != expected {
		t.Errorf("repeatedStringMatch(%v, %v) = %d; expected %d", a, b, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		a        = "a"
		b        = "aa"
		expected = 2
	)
	actual := repeatedStringMatch(a, b)
	if actual != expected {
		t.Errorf("repeatedStringMatch(%v, %v) = %d; expected %d", a, b, actual, expected)
	}
}

func Test_case_3(t *testing.T) {
	var (
		a        = "a"
		b        = "a"
		expected = 1
	)
	actual := repeatedStringMatch(a, b)
	if actual != expected {
		t.Errorf("repeatedStringMatch(%v, %v) = %d; expected %d", a, b, actual, expected)
	}
}

func Test_case_4(t *testing.T) {
	var (
		a        = "abc"
		b        = "xyz"
		expected = -1
	)
	actual := repeatedStringMatch(a, b)
	if actual != expected {
		t.Errorf("repeatedStringMatch(%v, %v) = %d; expected %d", a, b, actual, expected)
	}
}

func Test_case_5(t *testing.T) {
	var (
		a        = "aaaaaaaaaaaaaaaaaaaaaaab"
		b        = "ba"
		expected = 2
	)
	actual := repeatedStringMatch(a, b)
	if actual != expected {
		t.Errorf("repeatedStringMatch(%v, %v) = %d; expected %d", a, b, actual, expected)
	}
}
