package solution1576

func modifyString(s string) string {
	m := make(map[int]uint8, len(s))
	for i := 0; i < len(s); i++ {
		if '?' != s[i] {
			m[i] = s[i]
		}
	}

	var res []byte
	for i := 0; i < len(s); i++ {
		res = append(res, s[i])
		if s[i] == '?' {
			res[i] = 'a'
			for res[i] == m[i+1] || res[i] == m[i-1] {
				res[i]++
			}
			m[i] = res[i]
		}
	}

	return string(res)
}
