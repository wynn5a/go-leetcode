package solution1653

// 字符串全为 “a”；
// 字符串全为 “b”；
// 字符串既有 “a” 也有 “b”，且所有 “a” 都在所有 “b”左侧。
// 第一、二种情况，是第三种情况的特殊情况
func minimumDeletions(s string) int {
	leftb := 0
	righta := 0
	for _, c := range s {
		if c == 'a' {
			righta++
		}
	}
	res := righta
	for _, c := range s {
		if c == 'a' {
			righta--
		} else {
			leftb++
		}
		res = min(res, leftb+righta)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
