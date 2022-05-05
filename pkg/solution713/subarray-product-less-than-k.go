package solution713

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	count := 0
	rst := 1
	left := 0
	right := 0
	for right < len(nums) {
		rst *= nums[right]
		for rst >= k {
			rst /= nums[left]
			left++
		}

		//每次右指针位移到一个新位置，应该加上 x 种数组组合：
		//  nums[right]
		//  nums[right-1], nums[right]
		//  nums[right-2], nums[right-1], nums[right]
		//  nums[left], ......, nums[right-2], nums[right-1], nums[right]
		//共有 right - left + 1 种
		count += right - left + 1
		right++
	}

	return count
}
