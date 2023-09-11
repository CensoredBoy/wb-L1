package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seconds = 1

func readFromChanel(ch chan int) {

	for data := range ch {
		fmt.Printf("now data is %d\n", data)
	}
}

func writeToChanel(ch chan int) {
	for {
		data := rand.Intn(100)
		ch <- data
	}
}

func main() {
	channel := make(chan int)
	defer close(channel) // закрываем канал, чтобы после таймера прекратить чтение из него
	go writeToChanel(channel)
	go readFromChanel(channel)

	time.Sleep(time.Duration(seconds) * time.Second)
}