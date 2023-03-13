package solution15

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	actual := threeSum(nums)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, acutal: %v", expected, actual)
	}
}
