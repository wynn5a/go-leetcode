package solution507

import "math"

//对于一个正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。
//给定一个整数n，如果是完美数，返回 true，否则返回 false
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	m := make(map[int]struct{})
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			if _, exists := m[i]; !exists {
				m[i] = struct{}{}
				m[num/i] = struct{}{}
				sum += i
				sum += num / i
			}
		}
		if sum > num {
			return false
		}
	}

	return sum == num
}
