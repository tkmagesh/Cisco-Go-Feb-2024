package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fn := getFn()
	fn()
}

func getFn() func() {
	random := rand.Intn(int(time.Now().Unix()))
	fmt.Println(random)
	if random%2 == 0 {
		return f1
	}
	return f2
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
