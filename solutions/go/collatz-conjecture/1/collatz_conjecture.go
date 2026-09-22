package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	counter := 0
    for {
        if n == 1 {break}
        if n%2 == 0 {
            n /= 2
        } else {
            n = n*3 + 1
        }
        counter++
        if counter > 500 {
            return 0, errors.New("probably inf loop")
        }
    }
    return counter, nil
}
