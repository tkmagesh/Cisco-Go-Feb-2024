package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {

	defer func() {
		// fmt.Println("[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("application panicked : ", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[divide] Calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	fmt.Println("[divide] Calculating remainder")
	remainder = x % y
	return
}
