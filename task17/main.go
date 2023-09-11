package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5,2,3,4,1}
	sort.Ints(mas)
	a := sort.SearchInts(mas, 23)
	fmt.Println(mas[a])
}
