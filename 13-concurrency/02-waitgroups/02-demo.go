package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go f1(wg)
	f2()
	wg.Wait() // block until the wg counter becomes 0 (default)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
