package main

import (
	"fmt"
	"time"
	"context"
)

//Остановка горутины с помощью отмены контекста в ходе программы

func someProcess(ctx context.Context){
	for{
		select {
			case <-ctx.Done():
				fmt.Println("Done")
				return
			default:
				fmt.Println("Do something")
		}
	}
}

func main(){
	c, cancel := context.WithCancel(context.Background())
	go someProcess(c)
	time.Sleep(time.Millisecond)
	cancel()
	time.Sleep(time.Second)
}
