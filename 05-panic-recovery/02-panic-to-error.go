package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {

	divisor := 0
	q, r, err := divideWrapper(100, divisor)
	if err != nil {
		fmt.Println("divide error, try again")
	} else {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	}
}

// wrapper for the 3rd party api that coverts the panic into an error
func divideWrapper(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error) // convert the "any" to "error"
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api (raising a panic)
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
