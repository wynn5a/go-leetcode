package solution2373

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	grid := [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}
	expected := [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}
	actual := largestLocal(grid)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	grid := [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}
	expected := [][]int{{9, 9}, {8, 6}}
	actual := largestLocal(grid)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
