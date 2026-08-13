package main

import (
	"errors"
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	res, err := intersectSlices(slice1, slice2)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(res)
}

func intersectSlices(slice1, slice2 []int) (result []int, err error) {
	slice1Length := len(slice1)
	slice2Length := len(slice2)

	if (slice1Length == 0 && slice2Length != 0) || (slice1Length != 0 && slice2Length == 0) {
		err = errors.New("slices cannot be nil")
		return result, err
	}

	i, j := 0, 0
	for i < slice1Length && j < slice2Length {
		switch {
		case slice1[i] < slice2[j]:
			i = i + 1
		case slice1[i] > slice2[j]:
			j = j + 1
		default:
			result = append(result, slice1[i])
			i = i + 1
			j = j + 1
		}
	}
	return
}
