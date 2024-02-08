package calculator

import "fmt"

func init() {
	fmt.Println("calculator package initialized - [subtract.go]")
}

func Subtract(x, y int) int {
	// operationCount++
	operationCount["Subtract"]++
	return x - y
}
