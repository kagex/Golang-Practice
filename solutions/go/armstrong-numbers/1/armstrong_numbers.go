package armstrongnumbers

func IsNumber(n int) bool {
	copyN := n
	digits := make([]int, 0)

	for copyN > 0 {
		digits = append(digits, copyN%10)
		copyN /= 10
	}

	if len(digits) == 0 {
		return true
	}

	sum := 0
	power := len(digits)

	for _, v := range digits {
		sum += intPow(v, power)
	}

	return n == sum
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}