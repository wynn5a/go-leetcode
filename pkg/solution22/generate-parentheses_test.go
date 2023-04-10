package solution22

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	testCases := []struct {
		n        int
		expected []string
	}{
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
	}

	for _, tc := range testCases {
		result := generateParenthesis(tc.n)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("generateParenthesis(%d) = %v; expected %v", tc.n, result, tc.expected)
		}
	}
}
