package solution53

func maxSubArray(nums []int) int {
	current := 0
	max := nums[0]

	for _, v := range nums {
		if current < 0 {
			current = v
		} else if current >= 0 {
			current += v
		}
		if current > max {
			max = current
		}
	}

	return max
}
