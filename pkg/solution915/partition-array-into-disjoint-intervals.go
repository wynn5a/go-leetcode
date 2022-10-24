package solution915

//左边最大值，小于等于右边最小值
func partitionDisjoint(nums []int) int {
	n := len(nums)
	minRight := make([]int, n)
	minRight[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		minRight[i] = min(nums[i], minRight[i+1])
	}

	maxLeft := nums[0]
	for i := 0; i < n-1; i++ {
		if maxLeft <= minRight[i] {
			return i
		}
		if nums[i] >= maxLeft {
			maxLeft = nums[i]
		}
	}

	return 0
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
