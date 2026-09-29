package squareroot

import "errors"

func SquareRoot(number int) (int, error) {
    if number < 0 {
        return 0, errors.New("wrong number")
    }
    
	if number == 0 {
		return 0, nil
	}

	square := 1
	for square <= number/square {
		if square*square == number {
			return square, nil
		}
		square++
	}

	return 0, errors.New("not a perfect square")
}
