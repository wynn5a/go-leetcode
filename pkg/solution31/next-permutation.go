package solution31

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	// 从后往前找到第一个相邻升序的位置 i
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 如果找不到这样的位置，说明整个数组已经是降序排列，直接翻转整个数组
	if i < 0 {
		reverse(nums)
		return
	}
	// 从后往前找到第一个大于 nums[i] 的位置 j
	j := n - 1
	for j >= 0 && nums[j] <= nums[i] {
		j--
	}
	// 交换 nums[i] 和 nums[j]
	nums[i], nums[j] = nums[j], nums[i]
	// 翻转 i 后面的部分
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
