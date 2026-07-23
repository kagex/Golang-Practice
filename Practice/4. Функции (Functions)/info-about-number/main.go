package main

import (
	"fmt"
	"math"
)

func main() {
	printNumberInfo(-4)
}

func printNumberInfo(number int) {
	isPositiveNum := false
	switch {
	case number > 0:
		fmt.Printf("Число %d положительное.", number)
		isPositiveNum = true
	case number < 0:
		fmt.Printf("Число %d отрицательное.", number)
	case number == 0:
		fmt.Print("Число равно 0.")
	}

	if number%2 == 0 {
		fmt.Printf("\nЧисло %d четное.", number)
	} else {
		fmt.Printf("\nЧисло %d нечетное.", number)
	}

	if isPositiveNum {
		sqrtOfNumber := math.Sqrt(float64(number))
		if sqrtOfNumber*sqrtOfNumber == float64(number) {
			fmt.Printf("\nКвадратный корень числа %d является целым числом и равен %d.", number, int(sqrtOfNumber))
		} else {
			fmt.Printf("\nКвадратный корень числа %d не является целым числом и равен %.5f.", number, sqrtOfNumber)
		}
	}
}
