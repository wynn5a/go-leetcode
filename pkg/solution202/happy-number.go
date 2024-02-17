package solution202

// Floyd's Cycle detection algorithm
func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = digitSquareSum(slow)
		fast = digitSquareSum(digitSquareSum(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}

func digitSquareSum(number int) int {
	sum := 0
	for number > 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	return sum
}
