package solution32

import "testing"

func TestCase1(t *testing.T) {
	s := "(()"
	expected := 2
	actual := longestValidParentheses(s)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	s := ")()())"
	expected := 4
	actual := longestValidParentheses(s)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	s := "()(())"
	expected := 6
	actual := longestValidParentheses(s)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
