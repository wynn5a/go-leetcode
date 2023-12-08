package solution58

func lengthOfLastWord(s string) int {
	length := 0
	last := 0
	for _, v := range s {
		if v == ' ' {
			length = 0
			continue
		} else {
			length += 1
			last = length
		}

	}
	return last
}
