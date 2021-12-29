package solution1995

import "fmt"

//给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：
//nums[a] + nums[b] + nums[c] == nums[d] ，且
//a < b < c < d
//双向查找，利用 b 和 c 指针来找到 d - c 以及 a + b，利用 map 来存放存在的 a+b 的组合
func countQuadruplets(nums []int) (res int) {
	sum := map[int]int{}
	for i := 1; i < len(nums)-2; i++ {
		fmt.Printf("i: %d \n", i)
		for j := 0; j < i; j++ {
			fmt.Printf(" -- j: %d \n", j)
			sum[nums[i]+nums[j]]++
		}
		for j := i + 2; j < len(nums); j++ {
			fmt.Printf(" --- j: %d \n", j)
			res += sum[nums[j]-nums[i+1]]
		}
	}
	return
}

//暴力枚举找找感觉
//func countQuadruplets(nums []int) int {
//	res := 0
//	if len(nums) < 4 {
//		return res
//	}
//
//	for a, x := range nums {
//		for b := a + 1; b < len(nums); b++ {
//			for c := b + 1; c < len(nums); c++ {
//				for _, y := range nums[c+1:] {
//					if x+nums[b]+nums[c] == y {
//						res++
//					}
//				}
//			}
//		}
//	}
//
//	return res
//}
