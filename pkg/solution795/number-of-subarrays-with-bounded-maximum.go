package solution795

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	last2, last1, res := -1, -1, 0
	for i, x := range nums {
		if left <= x && x <= right {
			last1 = i
		} else if x > right {
			last2 = i
			last1 = -1
		}
		if last1 != -1 {
			res += last1 - last2
		}
	}
	return res
}
