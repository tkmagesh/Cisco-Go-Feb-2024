package calculator

import "fmt"

func init() {
	fmt.Println("calculator package initialized - [add.go]")
}

// public
func Add(x, y int) int {
	// operationCount++
	operationCount["Add"]++
	return x + y
}
