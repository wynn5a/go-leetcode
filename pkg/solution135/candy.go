package solution135

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	// Greedy
	// 1. Give one candy to every child
	// 2. If the child on the right has a higher rating, give him one more candy than the child on the left
	// 3. If the child on the left has a higher rating, give him one more candy than the child on the right
	// 4. Repeat step 2 and 3 until all children have been scanned
	// 5. Sum up all candies

	// 1. Give one candy to every child
	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			candies[i] = candies[i-1] + 1
			continue
		}
		candies[i] = 1
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	sum := 0
	for _, candy := range candies {
		sum += candy
	}
	return sum
}
