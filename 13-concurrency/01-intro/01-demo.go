package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling this function execution through the scheduler to be executed in future
	f2()

	// blocking the execution of this function so that the scheduler can look for other goroutines that are scheduled and execute them
	// DO NOT DO THE FOLLOWING IN PRODUCTION

	time.Sleep(2 * time.Second)

	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
