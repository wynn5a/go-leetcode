package solution2006

func countKDifference(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
		ans += cnt[num-k] + cnt[num+k]
	}
	return
}

//func countKDifference(nums []int, k int) int {
//	res := 0
//	for j, x := range nums {
//		for _, y := range nums[:j] {
//			if math.Abs(float64(x-y)) == float64(k) {
//				res++
//			}
//		}
//	}
//	return res
//}
