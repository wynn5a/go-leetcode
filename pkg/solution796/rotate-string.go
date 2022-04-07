package solution796

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[i:]+s[:i] == goal {
			return true
		}
	}
	return false
}
