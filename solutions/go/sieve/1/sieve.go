package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	isPrime := make([]bool, limit+1)
	
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= limit; p++ {
		if isPrime[p] {
			for multiple := p * p; multiple <= limit; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	var result []int
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			result = append(result, i)
		}
	}

	return result
}