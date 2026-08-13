package main

import "fmt"

type Event struct {
	Title    string
	Date     string
	Location string
}

func main() {
	fmt.Println(createGoEvent())
}

func createGoEvent() Event {
	goBirthday := Event{
		Title:    "День рождения Golang",
		Date:     "10 ноября 2009",
		Location: "GoogleLand",
	}
	return goBirthday
}
