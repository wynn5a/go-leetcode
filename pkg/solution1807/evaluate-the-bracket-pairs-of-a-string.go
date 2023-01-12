package solution1807

import (
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	var k = make(map[string]string, len(knowledge))
	for _, h := range knowledge {
		k[h[0]] = h[1]
	}
	start := -1
	ans := &strings.Builder{}
	for i, c := range s {
		if c == '(' {
			start = i
			i++
		} else if c == ')' {
			if value, exists := k[s[start+1:i]]; exists {
				ans.WriteString(value)
			} else {
				ans.WriteByte('?')
			}
			start = -1
		} else if start < 0 {
			ans.WriteRune(c)
		}
	}
	return ans.String()
}
