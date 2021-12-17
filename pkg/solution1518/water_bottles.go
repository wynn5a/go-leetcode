package solution1518

import "fmt"

//小区便利店正在促销，用 numExchange 个空酒瓶可以兑换一瓶新酒。你购入了 numBottles 瓶酒。
//如果喝掉了酒瓶中的酒，那么酒瓶就会变成空的。
//请你计算 最多 能喝到多少瓶酒。

// 递归解法
func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + change(numExchange, numBottles)
}

//循环解法
//func numWaterBottles(numBottles int, numExchange int) int {
//	result := numBottles
//
//	for numBottles >= numExchange {
//		result += numBottles / numExchange
//		numBottles = numBottles%numExchange + numBottles/numExchange
//		fmt.Printf("left bottles: %d \n", numBottles)
//	}
//
//	return result
//}

func change(exchange int, total int) int {
	if total < exchange {
		return 0
	}
	exchanged := total / exchange
	left := total%exchange + exchanged
	fmt.Printf("exchanged: %d, left:%d \n", exchanged, left)
	return exchanged + change(exchange, left)
}
