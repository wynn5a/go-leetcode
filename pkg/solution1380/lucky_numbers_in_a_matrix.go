package solution1380

func luckyNumbers(matrix [][]int) (ans []int) {
	for _, row := range matrix {
	next:
		for j, x := range row {
			for _, y := range row {
				if y < x {
					continue next
				}
			}
			for _, r := range matrix {
				if r[j] > x {
					continue next
				}
			}
			ans = append(ans, x)
		}
	}
	return
}

//func luckyNumbers(matrix [][]int) []int {
//	var res []int
//	m := make(map[int]int, len(matrix))
//	for i := 0; i < len(matrix); i++ {
//		minj := 0
//		min := matrix[i][0]
//		for j := 1; j < len(matrix[i]); j++ {
//			if matrix[i][j] < min {
//				min = matrix[i][j]
//				minj = j
//			}
//		}
//		m[i] = minj
//	}
//	for i := 0; i < len(matrix[0]); i++ {
//		maxj := 0
//		max := matrix[0][i]
//		for j := 1; j < len(matrix); j++ {
//			if matrix[j][i] > max {
//				max = matrix[j][i]
//				maxj = j
//			}
//		}
//		if m[maxj] == i {
//			res = append(res, max)
//		}
//	}
//	return res
//}
