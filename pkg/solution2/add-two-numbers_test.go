package solution2

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	l1 := util.NewListNode([]int{2, 4, 3})
	l2 := util.NewListNode([]int{5, 6, 4})

	actual := util.ListNodeToArray(addTwoNumbers(l1, l2))

	expected := []int{7, 0, 8}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

}

func TestCase2(t *testing.T) {
	l1 := util.NewListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := util.NewListNode([]int{9, 9, 9, 9})

	actual := util.ListNodeToArray(addTwoNumbers(l1, l2))

	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

}
