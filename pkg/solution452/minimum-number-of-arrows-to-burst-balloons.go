package solution452

import "sort"

// 贪心算法
// 先按照右边界排序，然后从左到右遍历，如果下一个气球的左边界大于当前气球的右边界，那么就需要一支箭，否则不需要
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// Sort the points by the end coordinate
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	xEnd := points[0][1]

	// Iterate over the points and if the current point's start coordinate is greater than the end coordinate
	// of the previous point, then we need another arrow
	for _, point := range points {
		if point[0] > xEnd {
			arrows++
			xEnd = point[1]
		}
	}

	return arrows
}
