package solution57

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			result = append(result, interval)
		} else if interval[0] > newInterval[1] {
			result = append(result, newInterval)
			newInterval = interval
		} else if interval[1] >= newInterval[0] || interval[0] <= newInterval[1] {
			newInterval[0] = min(interval[0], newInterval[0])
			newInterval[1] = max(interval[1], newInterval[1])
		}
	}
	result = append(result, newInterval)
	return result
}
