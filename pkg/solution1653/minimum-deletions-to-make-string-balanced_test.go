package solution1653

import (
	"testing"
)

func TestCase1(t *testing.T) {
	s := "aababbab"
	expected := 2
	actual := minimumDeletions(s)

	if expected != actual {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	s := "bbaaaaabb"
	expected := 2
	actual := minimumDeletions(s)

	if expected != actual {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
