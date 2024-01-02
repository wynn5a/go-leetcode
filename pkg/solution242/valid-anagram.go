package solution242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true
}
