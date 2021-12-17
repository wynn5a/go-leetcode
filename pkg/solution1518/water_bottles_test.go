package solution1518

import (
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		numBottles   = 9
		numExchanges = 3
		expected     = 13
	)
	actual := numWaterBottles(numBottles, numExchanges)
	if actual != expected {
		t.Errorf("numWaterBottles(%d, %d) = %d; expected %d", numBottles, numExchanges, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		numBottles   = 15
		numExchanges = 4
		expected     = 19
	)
	actual := numWaterBottles(numBottles, numExchanges)
	if actual != expected {
		t.Errorf("numWaterBottles(%d, %d) = %d; expected %d", numBottles, numExchanges, actual, expected)
	}
}

func Test_case_3(t *testing.T) {
	var (
		numBottles   = 2
		numExchanges = 3
		expected     = 2
	)
	actual := numWaterBottles(numBottles, numExchanges)
	if actual != expected {
		t.Errorf("numWaterBottles(%d, %d) = %d; expected %d", numBottles, numExchanges, actual, expected)
	}
}

func Test_case_4(t *testing.T) {
	var (
		numBottles   = 15
		numExchanges = 8
		expected     = 17
	)
	actual := numWaterBottles(numBottles, numExchanges)
	if actual != expected {
		t.Errorf("numWaterBottles(%d, %d) = %d; expected %d", numBottles, numExchanges, actual, expected)
	}
}
