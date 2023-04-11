package solution42

func trap(height []int) int {

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	// 使用双指针遍历数组
	for left < right {
		if height[left] < height[right] {
			// 如果左侧的高度小于右侧，则计算左侧的积水量
			if height[left] >= leftMax {
				// 如果当前高度大于左侧最大高度，则更新左侧最大高度
				leftMax = height[left]
			} else {
				// 否则，计算积水量
				result += leftMax - height[left]
			}
			// 左指针向右移动
			left++
		} else {
			// 如果右侧的高度小于等于左侧，则计算右侧的积水量
			if height[right] >= rightMax {
				// 如果当前高度大于右侧最大高度，则更新右侧最大高度
				rightMax = height[right]
			} else {
				// 否则，计算积水量
				result += rightMax - height[right]
			}
			// 右指针向左移动
			right--
		}
	}
	return result

}
