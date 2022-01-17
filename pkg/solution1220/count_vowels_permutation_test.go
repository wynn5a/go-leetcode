package solution1220

import "testing"

func Test_case_1(t *testing.T) {
	var (
		n        = 1
		expected = 5
	)
	actual := countVowelPermutation(n)
	if actual != expected {
		t.Errorf("countVowelPermutation(%d) = %d; expected %d", n, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		n        = 2
		expected = 10
	)
	actual := countVowelPermutation(n)
	if actual != expected {
		t.Errorf("countVowelPermutation(%d) = %d; expected %d", n, actual, expected)
	}
}

func Test_case_3(t *testing.T) {
	var (
		n        = 5
		expected = 68
	)
	actual := countVowelPermutation(n)
	if actual != expected {
		t.Errorf("countVowelPermutation(%d) = %d; expected %d", n, actual, expected)
	}
}
