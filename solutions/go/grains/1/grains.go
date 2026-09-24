package grains

import "errors"

func Square(number int) (uint64, error) {
	if number > 64 || number < 1 {
        return 0, errors.New("Wrong number, please try another")
    }
	return uint64(1) << (number - 1), nil
}

func Total() uint64 {
	return ^uint64(0)
}
