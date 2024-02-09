package main

import "fmt"

func main() {
	/*
		// version-1
		ch := make(chan int)
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

	/*
		// version-2
		ch := make(chan int)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	// version-3
	/*
		ch := make(chan int)
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/

	// version-4
	ch := make(chan int)
	go func() {
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100

}
