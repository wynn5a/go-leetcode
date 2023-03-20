package solution1012

import "testing"

func TestCase1(t *testing.T) {
	n := 20
	actual := numDupDigitsAtMostN(n)
	expected := 1
	if expected != actual {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	n := 100
	actual := numDupDigitsAtMostN(n)
	expected := 10
	if expected != actual {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	n := 1000
	actual := numDupDigitsAtMostN(n)
	expected := 262
	if expected != actual {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
