package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type threadMap struct {
	key   int
	value int
}

func AddToMap(ch chan threadMap, data map[int]int, mx *sync.Mutex) {
	c := <-ch
	mx.Lock()
	data[c.key] = c.value
	mx.Unlock()
}

func PushToAdd(ch chan threadMap) {
	key := rand.Intn(100)
	value := rand.Intn(100)
	ch <- threadMap{key, value}
}

func main() {
	ch := make(chan threadMap)
	m := make(map[int]int)
	mx := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go PushToAdd(ch)
		go AddToMap(ch, m, &mx)
	}

	time.Sleep(time.Second)
	fmt.Println(m)
}
