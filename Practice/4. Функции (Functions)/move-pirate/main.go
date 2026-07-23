package main

import "fmt"

func main() {

}

var pirateHP int = 3
var countRuns int = 0

func movePirate(isTrap bool) {

	countRuns++

	if pirateHP <= 0 {
		return
	}

	fmt.Println("Пират переместился на плиту", countRuns)

	if isTrap {
		pirateHP--
		if pirateHP > 0 {
			fmt.Println("Пират ранен")
		} else {
			fmt.Println("Пират убит")
			return
		}
	}

	if countRuns >= 10 {
		fmt.Println("Пират преодолел все ловушки")
	}
}
