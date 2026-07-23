package main

import "fmt"

func main() {
	var val any = 42.013

	switch val.(type) {
	case int:
		fmt.Println("В переменной val находится тип int.")
	case string:
		fmt.Println("В переменной val находится тип string.")
	case float64:
		fmt.Println("В переменной val находится тип float64.")
	case bool:
		fmt.Println("В переменной val находится тип bool.")
	default:
		fmt.Println("В переменной val находится неизвестный тип данных.")
	}
}
