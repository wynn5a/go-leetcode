package solution1610

import (
	"reflect"
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		points   = [][]int{{2, 1}, {2, 2}, {3, 3}}
		angle    = 90
		location = []int{1, 1}
		expected = 3
	)
	actual := visiblePoints(points, angle, location)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("visiblePoints(%d, %d, %d) = %d; expected %d", points, angle, location, actual, expected)
	}
}
