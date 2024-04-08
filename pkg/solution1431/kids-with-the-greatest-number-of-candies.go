package solution1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}
	res := make([]bool, len(candies))
	for i, candy := range candies {
		if candy+extraCandies >= maxCandies {
			res[i] = true
		}
	}
	return res
}
