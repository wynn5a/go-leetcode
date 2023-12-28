package solution205

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 使用两个数组来存储字符的映射关系，避免 map 的额外开销
	mapping1 := make([]int, 128) // 用于映射 s 中的字符
	mapping2 := make([]int, 128) // 用于映射 t 中的字符
	for i := 0; i < 128; i++ {
		mapping1[i] = -1
		mapping2[i] = -1
	}

	for i := 0; i < len(s); i++ {
		c1, c2 := s[i], t[i]
		index1, index2 := mapping1[c1], mapping2[c2]

		if index1 != index2 {
			// 映射关系不一致，返回 false
			return false
		}

		if index1 == -1 {
			// 字符第一次出现，建立映射关系
			mapping1[c1] = i
			mapping2[c2] = i
		}
	}

	return true
}
