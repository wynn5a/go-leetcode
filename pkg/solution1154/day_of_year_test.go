package solution1154

import "testing"

func Test_case_1(t *testing.T) {
	var (
		date     = "2019-01-09"
		expected = 9
	)
	actual := dayOfYear(date)
	if actual != expected {
		t.Errorf("dayOfYear(%v) = %d; expected %d", date, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		date     = "2019-02-10"
		expected = 41
	)
	actual := dayOfYear(date)
	if actual != expected {
		t.Errorf("dayOfYear(%v) = %d; expected %d", date, actual, expected)
	}
}

func Test_case_3(t *testing.T) {
	var (
		date     = "2003-03-01"
		expected = 60
	)
	actual := dayOfYear(date)
	if actual != expected {
		t.Errorf("dayOfYear(%v) = %d; expected %d", date, actual, expected)
	}
}
func Test_case_4(t *testing.T) {
	var (
		date     = "2004-03-01"
		expected = 61
	)
	actual := dayOfYear(date)
	if actual != expected {
		t.Errorf("dayOfYear(%v) = %d; expected %d", date, actual, expected)
	}
}
