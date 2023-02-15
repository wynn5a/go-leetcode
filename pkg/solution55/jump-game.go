package solution55

func canJump(nums []int) bool {
	if len(nums) == 1 && nums[0] == 0 {
		return true
	}
	l := len(nums) - 1
	max := 0
	for i, v := range nums {
		if i > max {
			break
		}
		if v+i > max {
			max = v + i
		}
		if max >= l {
			return true
		}
	}
	return false
}
