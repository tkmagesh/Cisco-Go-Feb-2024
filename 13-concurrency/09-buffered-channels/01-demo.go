package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))

	ch <- 100
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))

	ch <- 200
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))

	data := <-ch
	fmt.Println(data)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))

	data = <-ch
	fmt.Println(data)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
}
