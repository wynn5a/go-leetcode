package solution172

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	rst := 0
	for n > 0 {
		n = n / 5
		rst += n
	}

	return rst
}
