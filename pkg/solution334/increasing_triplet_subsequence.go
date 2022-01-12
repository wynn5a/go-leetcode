package solution334

import "math"

//快速查找表，左右两个方向
//func increasingTriplet(nums []int) bool {
//	n := len(nums)
//	if n < 3 {
//		return false
//	}
//	leftMin := make([]int, n)
//	leftMin[0] = nums[0]
//	for i := 1; i < n; i++ {
//		leftMin[i] = min(leftMin[i-1], nums[i])
//	}
//	rightMax := make([]int, n)
//	rightMax[n-1] = nums[n-1]
//	for i := n - 2; i >= 0; i-- {
//		rightMax[i] = max(rightMax[i+1], nums[i])
//	}
//	for i := 1; i < n-1; i++ {
//		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
//			return true
//		}
//	}
//	return false
//}
//
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
//
//func max(a, b int) int {
//	if b > a {
//		return b
//	}
//	return a
//}

//贪心算法
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}
