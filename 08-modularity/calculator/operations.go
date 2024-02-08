package calculator

import "fmt"

// private
/*
var operationCount int

func GetOpCount() int {
	return operationCount
}
*/

var operationCount map[string]int

func init() {
	fmt.Println("calculator package initialized - [operations.go]")
	operationCount = make(map[string]int)
}

func GetOpCount() map[string]int {
	return operationCount
}
