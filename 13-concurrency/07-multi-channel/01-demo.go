package main

import (
	"fmt"
	"time"
)

// modify in such a way that the data is printed in the order in which they are produced

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
