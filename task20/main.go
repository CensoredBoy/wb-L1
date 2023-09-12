package main

import (
	"fmt"
	"strings"
)

//то же самое, что и в 19 задаче, только для жлементов слайса

func reverse(str string) string {
	words := strings.Split(str, " ")
	var result string
	for i, j := 0, len(words)-1; i<j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	for _, word := range(words){
		result+= word+" "
	}
	return result
}

func main() {
	str := "snow dog is god"
	str2 := reverse(str)
	fmt.Println(str2)
}
