package solution806

func numberOfLines(widths []int, s string) []int {
	var lines, chars int
	for _, c := range s {
		if chars+widths[c-'a'] > 100 {
			lines++
			chars = 0
		}
		chars += widths[c-'a']
	}
	if chars > 0 {
		lines++
	}
	return []int{lines, chars}
}
