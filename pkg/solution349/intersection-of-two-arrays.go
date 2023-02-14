package solution349

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		m[v] = struct{}{}
	}

	var rst []int
	m2 := make(map[int]struct{})

	for _, v := range nums2 {
		_, ok1 := m[v]
		_, ok2 := m2[v]
		if ok1 && !ok2 {
			rst = append(rst, v)
			m2[v] = struct{}{}
		}
	}
	return rst
}
