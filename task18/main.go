package main

import (
	"fmt"
	"sync"
)

type counter struct {
	i int
	mx sync.RWMutex
}

func (c *counter) increment(wg *sync.WaitGroup) {
	c.mx.Lock()
	c.i++
	c.mx.Unlock()
	wg.Done()

}

func (c *counter) getCount() int{
	var result int
	c.mx.RLock()
	result = c.i
	c.mx.RUnlock()
	return result
}

func main() {
	var m sync.RWMutex
	c := counter{0, m}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go c.increment(&wg)

	}
	wg.Wait()
	fmt.Println(c.getCount())
}
