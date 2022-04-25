package solution398

import "math/rand"

type Solution []int

func Constructor(nums []int) Solution {
	return nums
}

// Pick 水塘抽样
//遍历 nums，当我们第 i 次遇到值为 target 的元素时，随机选择区间 [0,i)内的一个整数，如果其等于 0，则将返回值置为该元素的下标，否则返回值不变
func (nums Solution) Pick(target int) (ans int) {
	cnt := 0
	for i, num := range nums {
		if num == target {
			cnt++ // 第 cnt 次遇到 target
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return
}
