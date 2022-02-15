package solution1380

import (
	"reflect"
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		matrix   = [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
		expected = []int{15}
	)
	actual := luckyNumbers(matrix)
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("luckyNumbers(%v) = %v; expected %v", matrix, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		matrix   = [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
		expected = []int{12}
	)
	actual := luckyNumbers(matrix)
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("luckyNumbers(%v) = %v; expected %v", matrix, actual, expected)
	}
}
