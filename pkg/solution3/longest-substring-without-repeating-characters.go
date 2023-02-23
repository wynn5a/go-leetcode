package solution3

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	rst := 0
	start := 0
	for i, r := range s {
		if v, exists := m[r]; exists && v >= start {
			start = v + 1
		}
		m[r] = i

		if i-start+1 > rst {
			rst = i - start + 1
		}
	}

	return rst
}
