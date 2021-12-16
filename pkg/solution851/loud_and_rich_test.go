package solution851

import (
	"reflect"
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		richer   = [][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}
		quiet    = []int{3, 2, 5, 4, 6, 1, 7, 0}
		expected = []int{5, 5, 2, 5, 4, 5, 6, 7}
	)
	actual := loudAndRich(richer, quiet)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("loudAndRich(%d, %d) = %d; expected %d", richer, quiet, actual, expected)
	}
}
