package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1000, 100, 10, 1, 0}
	sort.Ints(numbers)
	fmt.Println(numbers)
}
