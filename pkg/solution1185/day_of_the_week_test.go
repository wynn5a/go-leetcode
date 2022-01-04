package solution1185

import "testing"

func Test_case_1(t *testing.T) {
	var (
		year     = 2019
		month    = 8
		day      = 31
		expected = "Saturday"
	)
	actual := dayOfTheWeek(day, month, year)
	if actual != expected {
		t.Errorf("dayOfTheWeek(%d, %d, %d) = %v; expected %v", day, month, year, actual, expected)
	}
}
