package solution17

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	digits := "23"
	excepted := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	actual := letterCombinations(digits)

	if !reflect.DeepEqual(excepted, actual) {
		t.Errorf("expected: %v, actual: %v", excepted, actual)
	}
}

func TestCase2(t *testing.T) {
	digits := "2"
	excepted := []string{"a", "b", "c"}
	actual := letterCombinations(digits)

	if !reflect.DeepEqual(excepted, actual) {
		t.Errorf("expected: %v, actual: %v", excepted, actual)
	}
}

func TestCase3(t *testing.T) {
	digits := ""
	var excepted []string
	actual := letterCombinations(digits)

	if !reflect.DeepEqual(excepted, actual) {
		t.Errorf("expected: %v, actual: %v", excepted, actual)
	}
}
