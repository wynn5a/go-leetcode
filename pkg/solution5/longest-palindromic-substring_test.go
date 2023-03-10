package solution5

import "testing"

func TestCase1(t *testing.T) {
	s := "babad"
	expected := "bab"
	expected2 := "aba"

	actual := longestPalindrome(s)
	if actual != expected && actual != expected2 {
		t.Errorf("exected: %s, or %s, actual: %s", expected, expected2, actual)
	}
}

func TestCase2(t *testing.T) {
	s := "cbbd"
	expected := "bb"

	actual := longestPalindrome(s)
	if actual != expected {
		t.Errorf("exected: %s,  actual: %s", expected, actual)
	}
}
