package main

import (
	"fmt"
	"time"
)

//остановка горутины с помощью поступления сообщения из канала

func someGoroutine(quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Printf("not work\n")
			return
		default:
			fmt.Printf("work\n")
		}
	}
}

func main() {
	quit := make(chan bool)
	go someGoroutine(quit)

	time.Sleep(time.Second)
	quit <- true
}
