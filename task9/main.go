package main

import "fmt"

func double(firstChan <-chan int, secondChan chan<- int) {
	for val := range firstChan {
		secondChan <- val * 2
	}
}

func main() {
	mas := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	firstChan := make(chan int)
	secondChan := make(chan int)
	defer close(firstChan)
	defer close(secondChan)

	go double(firstChan, secondChan)

	for _, val := range mas {
		firstChan <- val
		fmt.Println(<-secondChan)
	}

}
