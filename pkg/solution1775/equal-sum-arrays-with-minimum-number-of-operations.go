package solution1775

func minOperations(nums1 []int, nums2 []int) (ans int) {
	if 6*len(nums1) < len(nums2) || 6*len(nums2) < len(nums1) {
		return -1 // 优化
	}
	d := 0 // 数组元素和的差，我们要让这个差变为 0
	for _, x := range nums2 {
		d += x
	}
	for _, x := range nums1 {
		d -= x
	}
	if d < 0 {
		d = -d
		nums1, nums2 = nums2, nums1 // 统一让 nums1 的数变大，nums2 的数变小
	}
	cnt := [6]int{} // 统计每个数的最大变化量
	for _, x := range nums1 {
		cnt[6-x]++
	} // nums1 的变成 6
	for _, x := range nums2 {
		cnt[x-1]++
	} // nums2 的变成 1
	for i := 5; ; i-- { // 从大到小枚举最大变化量 5 4 3 2 1
		if i*cnt[i] >= d { // 可以让 d 变为 0
			return ans + (d+i-1)/i
		}
		ans += cnt[i] // 需要所有最大变化量为 i 的数
		d -= i * cnt[i]
	}
}
