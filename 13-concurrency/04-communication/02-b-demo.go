package main

import (
	"fmt"
)

// share memory by communicating

func main() {
	ch := make(chan int)
	go add(ch, 100, 200)
	result := <-ch
	fmt.Println(result)
}

func add(ch chan int, x, y int) {
	result := x + y
	ch <- result
}
