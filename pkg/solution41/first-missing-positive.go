package solution41

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 因为数组下标从 0 开始，如果 nums[i] == i+1，说明 nums[i] 已经在正确的位置上
		// 交换 nums[i] 和 nums[nums[i]-1] 的位置, 使得 nums[i] 放在正确的位置上
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
