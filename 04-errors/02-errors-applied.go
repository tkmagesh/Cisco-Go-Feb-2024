package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	divisor := 7
	q, r, err := divide(100, divisor)
	if err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	} else if err == ErrDivideByZero {
		fmt.Println("Please do not attempt to divide by zero")
	} else {
		fmt.Println("something went wrong :", err.Error())
	}
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		// err := errors.New("divide by zero error")
		err := ErrDivideByZero
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

// using named results
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
