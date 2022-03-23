package solution440

func getSteps(cur, n int) (steps int) {
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func findKthNumber(n, k int) int {
	current := 1
	k--
	for k > 0 {
		steps := getSteps(current, n)
		if steps <= k {
			k -= steps
			current++
		} else {
			current *= 10
			k--
		}
	}
	return current
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
