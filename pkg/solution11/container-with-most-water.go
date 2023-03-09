package solution11

// two end pointer
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	i, j := 0, len(height)-1

	rst := 0

	for i < j {
		area := min(height[i], height[j]) * (j - i)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}

		if area > rst {
			rst = area
		}
	}

	return rst
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
