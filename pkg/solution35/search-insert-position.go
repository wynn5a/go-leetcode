package solution35

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	// binary search
	left, right := 0, len(nums)-1
	for left <= right {
		// avoid overflow and calculate mid
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
