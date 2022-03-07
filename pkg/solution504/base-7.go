package solution504

import "strconv"

func convertToBase7(num int) string {
	res := ""
	if num == 0 {
		return "0"
	}
	flag := false
	if num < 0 {
		flag = true
		num = -num
	}
	for num > 0 {
		res = strconv.Itoa(num%7) + res
		num /= 7
	}
	if flag {
		res = "-" + res
	}
	return res
}
