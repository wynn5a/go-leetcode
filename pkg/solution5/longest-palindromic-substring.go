package solution5

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < length; i++ {
		//以 i 为中心的奇数长度的回文串
		l, r := i, i
		for l >= 0 && r < length && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > end-start {
			start = l + 1
			end = r - 1
		}

		// 检查以 i 和 i+1 为中心的偶数长度回文串
		l, r = i, i+1
		for l >= 0 && r < length && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > end-start {
			start = l + 1
			end = r - 1
		}
	}

	return s[start : end+1]
}
