package main

import "fmt"

//swap двух чисел с помощью синтаксиса go и математически

func main() {
	first, second := 1488, 20

	first, second = second, first
	fmt.Printf("first - %d, second - %d\n", first, second)

	first = first + second
	second = first - second
	first = first - second
	fmt.Printf("first - %d, second - %d", first, second)
}
