package solution902

import "testing"

func TestCase1(t *testing.T) {
	digits := []string{"1", "3", "5", "7"}
	n := 100
	expected := 20
	actual := atMostNGivenDigitSet(digits, n)

	if actual != expected {
		t.Errorf("Test case failed, expected: %d, actual: %d", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	digits := []string{"1", "4", "9"}
	n := 1000000000
	expected := 29523
	actual := atMostNGivenDigitSet(digits, n)

	if actual != expected {
		t.Errorf("Test case failed, expected: %d, actual: %d", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	digits := []string{"7"}
	n := 8
	expected := 1
	actual := atMostNGivenDigitSet(digits, n)

	if actual != expected {
		t.Errorf("Test case failed, expected: %d, actual: %d", expected, actual)
	}
}
