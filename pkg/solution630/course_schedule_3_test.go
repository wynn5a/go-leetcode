package solution630

import "testing"

func Test_case_1(t *testing.T) {
	var (
		in       = [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}
		expected = 3
	)
	actual := scheduleCourse(in)
	if actual != expected {
		t.Errorf("scheduleCourse(%d) = %d; expected %d", in, actual, expected)
	}
}


func Test_case_2(t *testing.T) {
	var (
		in       = [][]int{{3,2}, {4,3}}
		expected = 0
	)
	actual := scheduleCourse(in)
	if actual != expected {
		t.Errorf("scheduleCourse(%d) = %d; expected %d", in, actual, expected)
	}
}

func Test_case_3(t *testing.T) {
	var (
		in       = [][]int{{1,2}}
		expected = 1
	)
	actual := scheduleCourse(in)
	if actual != expected {
		t.Errorf("scheduleCourse(%d) = %d; expected %d", in, actual, expected)
	}
}