package solution884

import (
	"reflect"
	"sort"
	"testing"
)

func Test_case_1(t *testing.T) {
	var (
		s1       = "this apple is sweet"
		s2       = "this apple is sour"
		expected = []string{"sweet", "sour"}
	)
	actual := uncommonFromSentences(s1, s2)
	sort.Strings(actual)
	sort.Strings(expected)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("uncommonFromSentences(%s, %s) = %v; expected %v", s1, s2, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		s1       = "apple apple"
		s2       = "banana"
		expected = []string{"banana"}
	)
	actual := uncommonFromSentences(s1, s2)
	sort.Strings(actual)
	sort.Strings(expected)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("uncommonFromSentences(%s, %s) = %v; expected %v", s1, s2, actual, expected)
	}
}
