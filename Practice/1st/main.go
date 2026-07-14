package main

import "fmt"

const PI = 3.14 //Создание константы

func main() {
	age := 20 //Создание переменной age типа integer

	age = 29 //Изменение значения переменной age

	var zeroValueForInt int // Нулевое значение для переменной int, для bool - false, а для string - ""

	fmt.Println("Возраст", age)
	fmt.Println("Вывод нулевого значения для int", zeroValueForInt)
	fmt.Println("\nКстати, число Пи равно ", PI)
}
