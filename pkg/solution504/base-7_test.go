package solution504

import "testing"

func TestCase1(t *testing.T) {
	n := 100
	expected := "202"
	actual := convertToBase7(n)
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	n := -7
	expected := "-10"
	actual := convertToBase7(n)
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
