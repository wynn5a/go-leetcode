package solution34

// 时间复杂度：O(logn)
// 空间复杂度：O(1)
// 思路：二分查找
// 查找第一个大于等于给定值的元素 和最后一个小于等于给定值的元素 返回第一个大于等于给定值的元素和最后一个小于等于给定值的元素的下标
func searchRange(nums []int, target int) []int {
	left := searchFirst(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := searchLast(nums[left:], target) + left
	return []int{left, right}
}

func searchFirst(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchLast(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
