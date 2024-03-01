package solution228

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	var result []string
	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
			end = nums[i]
		} else {
			if start == end {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			start, end = nums[i], nums[i]
		}
	}
	if start == end {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
	return result
}
