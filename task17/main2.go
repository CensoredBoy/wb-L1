package main

import (
	"fmt"
)

func main() {
	test := []int{1,2,3,4,5}
	fmt.Println(bin_search(5, test))
}


func bin_search(elem int, a []int) (result int) {
	mid := len(a) / 2
    switch {
    case len(a) == 0:
        result = -1 // не нашелся
    case a[mid] > elem:
        result = bin_search(elem, a[:mid])
    case a[mid] < elem:
        result = bin_search(elem, a[mid+1:])
        if result >= 0 { // так как индексация среза с нуля, не забываем про предидущие итерации
            result += mid + 1
        }
    default:
        result = mid
    }
    return
}