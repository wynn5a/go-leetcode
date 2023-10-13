package solution_The_Sieve_of_Eratosthenes

func sieve(n int) []int {
	var result []int
	isPrime := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i; i*j <= n; j++ {
				isPrime[i*j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			result = append(result, i)
		}
	}

	return result
}
