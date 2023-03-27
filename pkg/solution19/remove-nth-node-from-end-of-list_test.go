package solution19

import (
	"github.com/wynn5a/go-leetcode/pkg/util"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	head := util.NewListNodes([]int{1, 2, 3, 4, 5})
	head = removeNthFromEnd(head, 2)
	actual := util.ListNodeToArray(head)
	expected := []int{1, 2, 3, 5}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
func TestCase2(t *testing.T) {
	head := util.NewListNodes([]int{1})
	head = removeNthFromEnd(head, 1)
	actual := util.ListNodeToArray(head)
	var expected []int
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
