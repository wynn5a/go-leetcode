package solution76

func minWindow(s string, t string) string {
	// 记录最短子串的开始位置和长度
	start, minLen := 0, len(s)+1
	left, right := 0, 0

	// 记录 window 中已经有多少字符符合要求了
	match := 0

	// needs 和 window 相当于计数器
	needs := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}

	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的一系列更新
		if needs[c] > 0 {
			window[c]++
			if window[c] == needs[c] {
				match++
			}
		}

		// 判断左侧窗口是否要收缩
		for match == len(needs) {
			// 在这里更新最小覆盖子串
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			if needs[d] > 0 {
				if window[d] == needs[d] {
					match--
				}
				window[d]--
			}
		}
	}

	// 返回最小覆盖子串
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
