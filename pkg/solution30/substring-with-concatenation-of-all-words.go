package solution30

func findSubstring(s string, words []string) []int {
	wordNum := len(words)
	ans := []int{}
	if wordNum == 0 {
		return ans
	}
	wordLen := len(words[0])

	if len(s) < (wordNum * wordLen) {
		return ans
	}

	// add to map for fast query
	allWords := make(map[string]int)
	for _, w := range words {
		allWords[w]++
	}
	for i := 0; i < len(s)-wordNum*wordLen+1; i++ {
		hasWords := make(map[string]int)
		num := 0
		for num < wordNum {
			word := s[i+num*wordLen : i+(num+1)*wordLen]
			if _, ok := allWords[word]; ok {
				hasWords[word]++
				if hasWords[word] > allWords[word] {
					break
				}
			} else {
				break
			}
			num++
		}
		if num == wordNum {
			ans = append(ans, i)
		}
	}
	return ans
}
