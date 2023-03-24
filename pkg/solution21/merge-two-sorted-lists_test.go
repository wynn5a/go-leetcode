package solution21

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := util.NewListNodes([]int{1, 2, 4})
	l2 := util.NewListNodes([]int{1, 3, 4})

	actual := util.ListNodeToArray(mergeTwoLists(l1, l2))
	expected := []int{1, 1, 2, 3, 4, 4}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestMergeTwoLists2(t *testing.T) {
	l1 := util.NewListNodes([]int{1, 3, 7})
	l2 := util.NewListNodes([]int{1, 3, 4})

	actual := util.ListNodeToArray(mergeTwoLists(l1, l2))
	expected := []int{1, 1, 3, 3, 4, 7}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestMergeTwoLists3(t *testing.T) {
	l1 := util.NewListNodes([]int{1, 3, 7, 9})
	l2 := util.NewListNodes([]int{1, 3, 4})

	actual := util.ListNodeToArray(mergeTwoLists(l1, l2))
	expected := []int{1, 1, 3, 3, 4, 7, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
func TestMergeTwoLists4(t *testing.T) {
	l1 := util.NewListNodes([]int{1, 3, 7, 9})
	l2 := util.NewListNodes([]int{1, 2, 10})

	actual := util.ListNodeToArray(mergeTwoLists(l1, l2))
	expected := []int{1, 1, 2, 3, 7, 9, 10}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestMergeTwoLists5(t *testing.T) {
	l1 := util.NewListNodes([]int{1})
	l2 := util.NewListNodes([]int{})

	actual := util.ListNodeToArray(mergeTwoLists(l1, l2))
	expected := []int{1}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
