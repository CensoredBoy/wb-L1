package main

import (
	"fmt"
	"sync"
)

var sum int

func square(num int) {
	sum += num * num
}


func main() {
	a := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, num := range a {
		wg.Add(1)

		go func(num int) {
			square(num)
			wg.Done()
		}(num)
	}
	wg.Wait()
	fmt.Println("sum is - ", sum)
}
