package main

import (
	"fmt"
	"strings"
)

func check(s string) bool {
	str := strings.ToLower(s)
	mapResult := make(map[rune]int)
	for _, char := range str {
		mapResult[char]++
	}
	if len(mapResult) != len(str) {
		return false
	}
	return true
}

func main() {
	str := "dflkjs"
	fmt.Println("result is - ", check(str))
}
