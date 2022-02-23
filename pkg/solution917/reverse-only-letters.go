package solution917

func reverseOnlyLetters(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for !isLetter(r[i]) && i < len(s)-1 {
			i++
		}
		for !isLetter(r[j]) && j > 0 {
			j--
		}
		if i < j {
			r[i], r[j] = r[j], r[i]
		}
	}
	return string(r)
}

func isLetter(u rune) bool {
	return (u >= 'a' && u <= 'z') || (u >= 'A' && u <= 'Z')
}
