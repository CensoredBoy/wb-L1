package main

import (
	"fmt"
	"time"
)


//остановка горутины с помощью закрытия канала

func someProcess(a chan int) {
	for{
		select{
		default:
			fmt.Println("Do something")
			data, ok := <-a
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			_ = data
		}
	}
}

func main() {
	a := make(chan int)
	go someProcess(a)
	time.Sleep(time.Second)
	close(a)
	time.Sleep(time.Second)
}