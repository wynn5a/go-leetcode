package solution224

func calculate(s string) int {
	// 1. 用栈来保存括号前的结果和括号前的符号
	var stack []int
	// 2. 用 sign 来保存当前的符号
	var sign = 1
	// 3. 用 result 来保存当前的结果
	var result = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			// 保存括号前的结果和括号前的符号
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0
			sign = 1
		case ')':
			// 计算括号内的结果
			result = stack[len(stack)-1]*result + stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		case ' ':
			continue
		default:
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			i--
			result += sign * num
		}
	}
	return result
}
