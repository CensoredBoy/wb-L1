package main

import (
	"fmt"
	"time"
)

//остановка горутины после выполнения всех инструкций

func someGoroutine() {
	fmt.Println("do something")
	time.Sleep(time.Second)
	fmt.Println("done")
}

func main() {
	go someGoroutine()
	time.Sleep(time.Second*2)
}
