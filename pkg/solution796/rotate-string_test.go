package solution796

import "testing"

func TestCase1(t *testing.T) {
	input := "abcde"
	goal := "cdeab"
	actual := rotateString(input, goal)
	if !actual {
		t.Errorf("Expected: %v, Actual: %v", true, actual)
	}
}
func TestCase2(t *testing.T) {
	input := "abcde"
	goal := "abced"
	actual := rotateString(input, goal)
	if actual {
		t.Errorf("Expected: %v, Actual: %v", false, actual)
	}
}
