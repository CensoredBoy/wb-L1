package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 4}
	set2 := []int{1, 2, 5, 6, 7, 8, 9}
	checkMap := make(map[int]bool)
	result := []int{}
	for _, v := range set1 {
		checkMap[v] = true
	}
	for _, v := range set2 {
		_, exists := checkMap[v]
		if exists {
			result = append(result, v)
		}
	}
	fmt.Println(result)
}
