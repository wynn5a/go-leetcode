package solution507

import "testing"

func Test_case_1(t *testing.T) {
	actual := checkPerfectNumber(6)
	if !actual {
		t.Errorf("checkPerfectNumber(6) = %v; expected %v", actual, true)
	}
}
func Test_case_2(t *testing.T) {
	actual := checkPerfectNumber(496)
	if !actual {
		t.Errorf("checkPerfectNumber(496) = %v; expected %v", actual, true)
	}
}

func Test_case_3(t *testing.T) {
	actual := checkPerfectNumber(8128)
	if !actual {
		t.Errorf("checkPerfectNumber(8128) = %v; expected %v", actual, true)
	}
}

func Test_case_4(t *testing.T) {
	actual := checkPerfectNumber(2)
	if actual {
		t.Errorf("checkPerfectNumber(2) = %v; expected %v", actual, false)
	}
}
