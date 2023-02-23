package solution3

import "testing"

func TestCase1(t *testing.T) {
	s := "abcabcbb"
	expected := 3

	actual := lengthOfLongestSubstring(s)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	s := "bbbbb"
	expected := 1

	actual := lengthOfLongestSubstring(s)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	s := "pwwkew"
	expected := 3

	actual := lengthOfLongestSubstring(s)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
func TestCase4(t *testing.T) {
	s := "abba"
	expected := 2

	actual := lengthOfLongestSubstring(s)
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
