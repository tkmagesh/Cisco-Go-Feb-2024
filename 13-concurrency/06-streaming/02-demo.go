package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}

// producer
func genNos(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- 10 * i
		time.Sleep(500 * time.Millisecond)
	}

}
