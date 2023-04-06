package solution8

import "math"

func myAtoi(str string) int {
	result := 0
	sign := 1
	i := 0

	// Skip white spaces
	for i < len(str) && str[i] == ' ' {
		i++
	}

	// Check for the sign
	if i < len(str) && (str[i] == '+' || str[i] == '-') {
		if str[i] == '+' {
			sign = 1
		} else {
			sign = -1
		}
		i++
	}

	// Convert the string to integer
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		digit := int(str[i] - '0')
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result*10 + digit
		i++
	}

	return result * sign
}
