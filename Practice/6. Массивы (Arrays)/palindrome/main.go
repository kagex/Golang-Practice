package main

import (
	"fmt"
)

func main() {
	testArray := [10]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 1}
	isPalindrome(testArray)
}

func isPalindrome(array [10]int) {
	length := len(array) - 1
	for i := 0; i < length/2; i++ {
		if array[i] != array[length-i] {
			fmt.Print("Не палиндром!")
			return
		}
	}
	fmt.Print("Это палиндром!")
}
