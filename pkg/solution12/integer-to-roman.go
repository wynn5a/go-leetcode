package solution12

import "strings"

func intToRoman(num int) string {
	if num > 3999 || num < 1 {
		return ""
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	i := 0
	for num > 0 {
		if num >= values[i] {
			result.WriteString(symbols[i])
			num -= values[i]
		} else {
			i++
		}
	}
	return result.String()

}
