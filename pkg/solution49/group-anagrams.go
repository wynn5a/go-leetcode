package solution49

// groupAnagrams 桶排序 + 哈希表
// 时间复杂度：O(nk)
// 空间复杂度：O(nk)
// n 为 strs 的长度，k 为 strs 中字符串的最大长度
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		count := [26]int{}
		// 将字符串中的字符出现的次数存入数组中，这样就可以将字符串转换为一个只跟字符出现次数有关的数组，从而实现字符串的分类
		for _, c := range str {
			count[c-'a']++
		}
		// 将数组转换为哈希表的键，这样就可以将字符串分类，相同的字符串就会被分到同一个桶中
		m[count] = append(m[count], str)
	}

	// 将哈希表中的值取出来，就是题目要求的结果
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
