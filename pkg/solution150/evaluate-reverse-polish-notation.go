package solution150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		default:
			stack = append(stack, parseInt(token))
		}
	}
	return stack[0]
}

func parseInt(token string) int {
	num, err := strconv.Atoi(token)
	if err != nil {
		return 0
	}
	return num
}
