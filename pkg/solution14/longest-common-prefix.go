package solution14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		// 从头开始比较，如果不相等，就截取prefix
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
		// 如果strs[i]比prefix短，就截取prefix
		if len(prefix) > len(strs[i]) {
			prefix = prefix[:len(strs[i])]
		}
	}
	return prefix
}
