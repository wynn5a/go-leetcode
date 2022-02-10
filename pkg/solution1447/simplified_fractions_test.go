package solution1447

import (
	"reflect"
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		n        = 2
		expected = []string{"1/2"}
	)
	actual := simplifiedFractions(n)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("simplifiedFractions(%d) = %v; expected %v", n, actual, expected)
	}
}
func Test_case_2(t *testing.T) {
	var (
		n        = 3
		expected = []string{"1/2", "1/3", "2/3"}
	)
	actual := simplifiedFractions(n)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("simplifiedFractions(%d) = %v; expected %v", n, actual, expected)
	}
}

//func Test_case_3(t *testing.T) {
//	var (
//		n        = 4
//		expected = []string{"1/2", "1/3", "1/4", "2/3", "3/4"}
//	)
//	actual := simplifiedFractions(n)
//	if !reflect.DeepEqual(actual, expected) {
//		t.Errorf("simplifiedFractions(%d) = %v; expected %v", n, actual, expected)
//	}
//}
