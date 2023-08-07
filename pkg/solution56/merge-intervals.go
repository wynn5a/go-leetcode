package solution56

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 对 intervals 中的元素进行排序，依据每个子数组（区间）的左边界字面值
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片
	merged := make([][]int, 0, len(intervals))

	for _, interval := range intervals {
		// 如果结果集为空，或当前区间的左端点 > 结果集最后一个区间的右端点
		// 则不需要合并，直接将当前区间加入到结果集即可。
		if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
			merged = append(merged, interval)
		} else {
			// 否则，当前区间与结果集的最后一个区间存在重叠，进行合并操作。
			// 合并后的区间的右端点为当前区间和结果集中最后一个区间右端点的较大值。
			if interval[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = interval[1]
			}
		}
	}

	return merged
}
