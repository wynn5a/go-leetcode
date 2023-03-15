package solution7

import "math"

func reverse(x int) int {
	// 定义变量记录反转后的结果和符号
	var res, sign int

	// 判断输入整数的符号，并将其转化为正数进行处理
	if x < 0 {
		sign = -1
		x = -x
	} else {
		sign = 1
	}

	// 反转整数
	for x > 0 {
		// 获取 x 的最后一位数字
		digit := x % 10

		// 判断反转后的整数是否会溢出
		if res > (math.MaxInt32-digit)/10 {
			return 0
		}

		// 更新反转后的整数
		res = res*10 + digit

		// 移除 x 的最后一位数字
		x /= 10
	}

	// 根据符号返回最终结果
	return res * sign
}
