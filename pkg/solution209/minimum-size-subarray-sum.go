package solution209

import "math"

// minSubArrayLen 滑动窗口 - 同向双指针
func minSubArrayLen(target int, nums []int) int {
	left, right, sum := 0, 0, 0
	res := math.MaxInt32
	l := len(nums)

	// 当右指针还没有走到数组的末尾，让右指针持续向前走
	for right < l {
		sum += nums[right]
		for sum >= target {
			// 比较目前找到的子数组和之前找到的子数组的长度，取长度更短的那个
			if right-left+1 < res {
				res = right - left + 1
			}
			// 左指针右移，减少累计和的值，查看是否有更小长度的子数组满足条件
			sum -= nums[left]
			left++
		}
		//没有找到大于 target 的子数组，或者需要继续探索更小序列
		right++
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}
