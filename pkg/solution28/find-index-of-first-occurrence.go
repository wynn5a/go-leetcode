package solution28

// find the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if len(haystack)-i < len(needle) {
				return -1
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}
