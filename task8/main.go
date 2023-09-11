package main

import "fmt"

func main() {
	var number int64 = 100
	var bitNumber int64 = 5
	changeToOne := false
	switch changeToOne {
	case true:
		number |= 1 << bitNumber //если нужно поменять на единицу, используем "ИЛИ"" и маску и нулями
	default:
		number &= ^(1 << bitNumber) //если нужно поменять на ноль, используем "И" и маску с единицами
	}

	fmt.Println("result is - ", number)
}
