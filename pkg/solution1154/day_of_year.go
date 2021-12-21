package solution1154

import "strconv"

//给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。请你计算并返回该日期是当年的第几天。
func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	var monthDay []int
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		days[1]++
	}
	monthDay = append(monthDay, 0)
	monthDay = append(monthDay, 0)
	for i := 2; i <= 13; i++ {
		monthDay = append(monthDay, monthDay[i-1]+days[i-2])
	}

	return day + monthDay[month]
}
