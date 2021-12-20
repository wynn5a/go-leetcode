package solution475

import "testing"

func Test_case_1(t *testing.T) {
	var (
		houses   = []int{1, 2, 3}
		heaters  = []int{2}
		expected = 1
	)
	actual := findRadius(houses, heaters)
	if actual != expected {
		t.Errorf("findRadius(%v, %v) = %d; expected %d", houses, heaters, actual, expected)
	}
}
func Test_case_2(t *testing.T) {
	var (
		houses   = []int{1, 2, 3, 4}
		heaters  = []int{1, 4}
		expected = 1
	)
	actual := findRadius(houses, heaters)
	if actual != expected {
		t.Errorf("findRadius(%v, %v) = %d; expected %d", houses, heaters, actual, expected)
	}
}
func Test_case_3(t *testing.T) {
	var (
		houses   = []int{1, 5}
		heaters  = []int{2}
		expected = 3
	)
	actual := findRadius(houses, heaters)
	if actual != expected {
		t.Errorf("findRadius(%v, %v) = %d; expected %d", houses, heaters, actual, expected)
	}
}
