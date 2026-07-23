package main

import "fmt"

type Day int

const (
	_ Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	isWeekend := isWeekend(6)
	fmt.Printf("День является выходным?\n%t", isWeekend)
}

func isWeekend(day Day) (isWeekend bool) {
	if day == Saturday || day == Sunday {
		isWeekend = true
	}
	return
}
