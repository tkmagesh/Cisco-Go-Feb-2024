package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Done")
}

// producer
func genNos(ch chan int) {
	count := rand.Intn(20)
	for i := 1; i <= count; i++ {
		ch <- 10 * i
	}
	close(ch)
}
