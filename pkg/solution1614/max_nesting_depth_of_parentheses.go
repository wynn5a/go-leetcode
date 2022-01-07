package solution1614

func maxDepth(s string) int {
	res := 0
	current := 0
	for _, c := range s {
		if c == '(' {
			current++
		}
		if c == ')' {
			current--
		}

		if current > res {
			res = current
		}
	}

	return res
}
