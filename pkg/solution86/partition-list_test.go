package solution86

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	head := util.NewListNodes([]int{1, 4, 3, 2, 5, 2})
	x := 3
	expected := []int{1, 2, 2, 4, 3, 5}
	actual := partition(head, x)
	result := util.ListNodeToArray(actual)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCase2(t *testing.T) {
	head := util.NewListNodes([]int{2, 1})
	x := 2
	expected := []int{1, 2}
	actual := partition(head, x)
	result := util.ListNodeToArray(actual)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
