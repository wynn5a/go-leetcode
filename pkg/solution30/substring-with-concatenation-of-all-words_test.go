package solution30

// 假设 sort 包已导入
import (
	"reflect"
	"sort"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		s      string
		words  []string
		expect []int
	}{
		{
			s:      "barfoothefoobarman",
			words:  []string{"foo", "bar"},
			expect: []int{0, 9},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "good"},
			expect: []int{8},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "word"},
			expect: []int{},
		}, {
			s:      "barfoofoobarthefoobarman",
			words:  []string{"bar", "foo", "the"},
			expect: []int{6, 9, 12},
		},
	}

	for _, test := range tests {
		actual := findSubstring(test.s, test.words)

		// 对预期和实际结果进行排序
		sort.Ints(actual)
		sort.Ints(test.expect)

		if !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("Input: %s, %v\nOutput: %v\nExpect: %v\n", test.s, test.words, actual, test.expect)
		}
	}
}
