package main

import "fmt"

func main() {
	num := 5 // 0101 в двоичной системе

	fmt.Println(num << 2) // Сдвиг влево на 2 позиции: 010100 = 20
	fmt.Println(num >> 1) // Сдвиг вправо на 1 позицию: 0010 = 2
	fmt.Println(num & 3)  // Побитовое И: 0101 & 0011 = 0001 = 1
	fmt.Println(num | 2)  // Побитовое ИЛИ: 0101 | 0010 = 0111 = 7
	fmt.Println(num ^ 2)  // Побитовое XOR: 0101 ^ 0010 = 0111 = 7
	fmt.Println(^num)     // Побитовое НЕ (инверсия): ~0101 = 1111 1010 = -6
}
