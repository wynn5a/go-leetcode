package solution686

import "strings"

//给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1
func repeatedStringMatch(a string, b string) int {
	m, n, exist := len(a), len(b), make([]bool, 26)
	for _, ch := range a {
		exist[ch-'a'] = true
	}
	for _, ch := range b {
		if !(exist[ch-'a']) {
			//b contains char that does not exist in a
			return -1
		}
	}

	ret := n / m
	str := strings.Repeat(a, ret)
	for i := 0; i < 3; i++ {
		if strings.Contains(str, b) {
			return ret + i
		}
		str += a
	}
	return -1
}
