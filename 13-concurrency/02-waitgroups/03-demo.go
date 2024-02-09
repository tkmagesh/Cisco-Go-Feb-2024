package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	count := flag.Int("count", 0, "# of goroutines to spin")
	flag.Parse()
	fmt.Println("count :", *count)
	fmt.Printf("Starting %d goroutines, Hit ENTER to start\n", *count)
	fmt.Scanln()
	for i := 1; i <= *count; i++ {
		wg.Add(1)
		go fn(wg, i)
	}
	wg.Wait() // block until the wg counter becomes 0 (default)
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
