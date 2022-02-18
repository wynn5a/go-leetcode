package solution1791

func findCenter(edges [][]int) int {
	m := make(map[int]int)
	m[edges[0][0]] = 1
	m[edges[0][1]] = 1
	for i := 1; i < len(edges); i++ {
		if _, ok := m[edges[i][0]]; ok {
			m[edges[i][0]]++
			continue
		}
		if _, ok := m[edges[i][1]]; ok {
			m[edges[i][1]]++
			continue
		}
	}
	for k, v := range m {
		if v == len(edges) {
			return k
		}
	}
	return 0
}
