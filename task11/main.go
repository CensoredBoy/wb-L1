package main

import "fmt"

//реализация с вложенным циклом за О(n^2)

func main() {
	set1 := []int{1, 2, 3, 4}
	set2 := []int{1, 2, 5, 6, 7, 8, 9}
	result := []int{}

	for _, val1 := range set1 {
		for _, val2 := range set2 {
			if val1 == val2 {
				result = append(result, val1)
			}
		}
	}
	fmt.Println(result)
}
