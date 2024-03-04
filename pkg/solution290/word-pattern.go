package solution290

import "strings"

func wordPattern(pattern string, s string) bool {
	if len(pattern) == 0 && len(s) == 0 {
		return true
	}

	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	word2ch := make(map[string]byte)
	ch2word := make(map[byte]string)
	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		word := words[i]
		if _, ok := word2ch[word]; !ok {
			if _, ok := ch2word[ch]; ok {
				return false
			}
			word2ch[word] = ch
			ch2word[ch] = word
		} else {
			if word2ch[word] != ch || ch2word[ch] != word {
				return false
			}
		}
	}
	return true
}
