package solution

// MaxInArray maxInArray find max value from array
func MaxInArray(a []int) int {
	m := 0
	for i, e := range a {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}