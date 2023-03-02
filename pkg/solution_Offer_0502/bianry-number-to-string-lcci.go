package solution_Offer_0502

import "strings"

// 每次乘于2， 然后取整数部分的二级制表示放到结果中
func printBin(num float64) string {
	sb := &strings.Builder{}
	sb.WriteString("0.")
	for sb.Len() <= 32 && num != 0 {
		num *= 2
		digit := byte(num)
		sb.WriteByte('0' + digit)
		num -= float64(digit)
	}
	if sb.Len() <= 32 {
		return sb.String()
	}
	return "ERROR"
}
