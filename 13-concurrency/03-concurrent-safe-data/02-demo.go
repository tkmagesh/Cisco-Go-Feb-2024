package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var c Counter

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println(c.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.increment()
}
