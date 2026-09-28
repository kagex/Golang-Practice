package nthprime

import "errors"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	if n%3 == 0 {
		return n == 3
	}

	// Using formula 6k ± 1: 5, 7, 11, 13, ...
	for i := 5; i <= n/i; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
        return 0, errors.New("prime number can't be computed, N can`t be lesser than 1")
    }
    
	if n == 1 {
		return 2, nil
	}

	counter := 1
	number := 1

	for counter < n {
		number += 2
		if isPrime(number) {
			counter++
		}
	}

	return number, nil
}
