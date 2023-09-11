package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for data := range c {
		time.Sleep(time.Millisecond)
		fmt.Println("Worker with id - ", id, ", and with message - ", data)
	}
}

func main() {
	c := make(chan int)
	
	N := 10 // Количество воркеров 

	for i := 0; i < N; i++ {
		go worker(i, c)
	}

	for {
		data := rand.Intn(100)
		c <- data
	}
}
