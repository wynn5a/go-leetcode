package solution539

import (
	"math"
	"sort"
)

//排序后，答案就出在首尾的差，或者前后的差
func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	ans := math.MaxInt32
	t0Minutes := getMinutes(timePoints[0])
	preMinutes := t0Minutes
	for _, t := range timePoints[1:] {
		minutes := getMinutes(t)
		ans = min(ans, minutes-preMinutes) // 相邻时间的时间差
		preMinutes = minutes
	}
	ans = min(ans, t0Minutes+1440-preMinutes) // 首尾时间的时间差
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinutes(t string) int {
	return (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
}
