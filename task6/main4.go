package main

import (
	"context"
	"fmt"
	"time"
)

//Остановка горутины с помошью контекста с таймером

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Ending 1 by context time out")
				return
			default:
				fmt.Println("Normal")
			}
		}
	}()
	time.Sleep(time.Second)
}
