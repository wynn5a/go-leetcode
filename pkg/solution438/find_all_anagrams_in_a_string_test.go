package solution438

import (
	"reflect"
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		s        = "cbaebabacd"
		p        = "abc"
		expected = []int{0, 6}
	)
	actual := findAnagrams(s, p)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("findAnagrams(%s, %s) = %v; expected %v", s, p, actual, expected)
	}
}
func Test_case_2(t *testing.T) {
	var (
		s        = "abab"
		p        = "ab"
		expected = []int{0, 1, 2}
	)
	actual := findAnagrams(s, p)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("findAnagrams(%s, %s) = %v; expected %v", s, p, actual, expected)
	}
}
