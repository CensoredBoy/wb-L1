package main

import (
	"fmt"
	"sync"
)

func square(c int) {
	fmt.Println(c * c)
}

func main() {
	a := [4]int{3, 4, 6, 7}
	wg := sync.WaitGroup{}
	for _, v := range a {
		wg.Add(1) //  добавляем +1 к счетчику группы
		go func() {
			square(v)
			wg.Done() // делаем декремент счетчика группы
		}()
		wg.Wait() // ждем, пока счетчик группы не будет равен нулю
	}
	fmt.Println("Done")

}