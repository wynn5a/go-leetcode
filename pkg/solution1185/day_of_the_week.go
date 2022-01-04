package solution1185

import "math"

//给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。
//输入为三个整数：day、month 和year，分别表示日、月、年。
//您返回的结果必须是这几个值中的一个{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
//基姆拉尔森计算公式 或者计算与 1971/1/1（星期五） 之间的天数%7
func dayOfTheWeek(day int, month int, year int) string {
	resArray := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	if month < 3 {
		month += 12
		year -= 1
	}
	index := year + int(math.Floor(float64(year/4))) - int(math.Floor(float64(year/100))) + int(math.Floor(float64(year/400))) + 2*month + int(math.Floor(float64(3*(month+1)/5))) + day + 1
	return resArray[index%7]
}
