package solution167

func twoSum(numbers []int, target int) []int {
	// 双指针 一个指针指向数组头部，一个指针指向数组尾部
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			// 如果和小于目标值，说明左指针指向的值太小，需要右移
			left++
		} else {
			// 如果和大于目标值，说明右指针指向的值太大，需要左移
			right--
		}
	}
	return []int{}
}
