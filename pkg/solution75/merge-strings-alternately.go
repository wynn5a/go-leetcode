package solution75

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	r2 := []rune(word2)
	l2 := len(word2)
	for p, r := range word1 {
		sb.WriteRune(r)
		if p < l2 {
			sb.WriteRune(r2[p])
		}
	}

	l1 := len(word1)
	if l2 > l1 {
		for i := 0; i < l2-l1; i++ {
			sb.WriteRune(r2[l1+i])
		}
	}

	return sb.String()
}
