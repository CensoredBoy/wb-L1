package main

import (
	"fmt"
)

//два счетчика, i и j для начала и конца строки, в каждой итерации меняем местами символы строки

func reverse(str string) string {
	sliceRune := []rune(str)
	for i, j := 0, len(sliceRune)-1; i < j; i, j = i+1, j-1 {
		sliceRune[i], sliceRune[j] = sliceRune[j], sliceRune[i]
	}
	return string(sliceRune)
}

func main() {
	str := "Danila"
	newString := reverse(str)
	fmt.Println(newString)
}
