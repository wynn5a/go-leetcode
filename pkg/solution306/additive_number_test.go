package solution306

import "testing"

func Test_case_1(t *testing.T) {
	var (
		s = "111111111"
	)
	actual := isAdditiveNumber(s)
	if actual != false {
		t.Errorf("isAdditiveNumber(%v) = %v; expected %v", s, actual, false)
	}
}
func Test_case_2(t *testing.T) {
	var (
		s = "112358"
	)
	actual := isAdditiveNumber(s)
	if actual != true {
		t.Errorf("isAdditiveNumber(%v) = %v; expected %v", s, actual, true)
	}
}
