package solution122

import "testing"

func TestCase1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	value := maxProfit(prices)
	expected := 7

	if value != expected {
		t.Errorf("Test failed, expected: %d, actual: %d", expected, value)
	}
}

func TestCase2(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	value := maxProfit(prices)
	expected := 4

	if value != expected {
		t.Errorf("Test failed, expected: %d, actual: %d", expected, value)
	}
}

func TestCase3(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	value := maxProfit(prices)
	expected := 0

	if value != expected {
		t.Errorf("Test failed, expected: %d, actual: %d", expected, value)
	}
}
