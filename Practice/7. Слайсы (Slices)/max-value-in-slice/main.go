package main

import (
	"errors"
	"fmt"
)

func main() {
	testSlice := []int{-1, -5, -3}
	res, err := Max(testSlice)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(res)
}

func Max(slice []int) (max int, err error) {
	if len(slice) == 0 {
		err := errors.New("slice is nil or empty")
		return 0, err
	}

	max = slice[0]
	for _, value := range slice {
		if max < value {
			max = value
		}
	}

	return max, err
}
