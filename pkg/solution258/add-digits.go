package solution258

func addDigits(num int) int {
	ans := num
	for num > 9 {
		ans = 0
		for num > 0 {
			ans += num % 10
			num /= 10
		}
		num = ans
	}
	return ans
}
