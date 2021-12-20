package solution475

import "sort"

//在加热器的加热半径范围内的每个房屋都可以获得供暖。
//现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。

//暴力循环
//func findRadius(houses []int, heaters []int) int {
//	var result = 0
//	for _, house := range houses {
//		var near = 500000000
//		for _, heater := range heaters {
//			distance := abs(heater-house)
//			if  distance < near {
//				near = distance
//			}
//		}
//		if near > result{
//			result = near
//		}
//	}
//
//	return result
//}

//排好序 二分查找
//TODO 还能进一步优化停止条件和查找区间
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var result = 0
	for _, house := range houses {
		near := searchNear(house, heaters)
		if near > result {
			result = near
		}
	}
	return result
}

func searchNear(house int, heaters []int) int {
	left := 0
	right := len(heaters)
	result := 500000000
	for left < right {
		mid := left + (right-left)/2
		if house == heaters[mid] {
			return 0
		}
		if house > heaters[mid] {
			if result > house-heaters[mid] {
				result = house - heaters[mid]
			}
			left = mid + 1
			if left < len(heaters) && result > abs(house-heaters[left]) {
				result = abs(house - heaters[left])
			}
			continue
		}
		if house < heaters[mid] {
			if result > heaters[mid]-house {
				result = heaters[mid] - house
			}
			right = mid - 1
			if right >= 0 && result > abs(house-heaters[right]) {
				result = abs(house - heaters[right])
			}
			continue
		}
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
