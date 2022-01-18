package solution539

import (
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		timePoints = []string{"23:59", "00:00"}
		excepted   = 1
	)
	actual := findMinDifference(timePoints)
	if actual != excepted {
		t.Errorf("findMinDifference(%v) = %d; expected %d", timePoints, actual, excepted)
	}
}
func Test_case_2(t *testing.T) {
	var (
		timePoints = []string{"00:00", "23:59", "00:00"}
		excepted   = 0
	)
	actual := findMinDifference(timePoints)
	if actual != excepted {
		t.Errorf("findMinDifference(%v) = %d; expected %d", timePoints, actual, excepted)
	}
}
func Test_case_3(t *testing.T) {
	var (
		timePoints = []string{"01:00", "00:00", "01:01"}
		excepted   = 1
	)
	actual := findMinDifference(timePoints)
	if actual != excepted {
		t.Errorf("findMinDifference(%v) = %d; expected %d", timePoints, actual, excepted)
	}
}
