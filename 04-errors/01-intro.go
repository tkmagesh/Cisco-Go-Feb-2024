package main

import (
	"errors"
	"fmt"
)

var ErrFileOpen error = errors.New("error opening the file")
var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	// create an error
	var err error
	/*
		err = errors.New("dummy error")
		if err != nil {
			fmt.Println("some error occurred")
		} else {
			fmt.Println("All good!")
		}
	*/

	// err = ErrDivideByZero
	// err = ErrFileOpen
	err = errors.New("dummy error")
	if err == nil {
		fmt.Println("All good!")
	} else if err == ErrFileOpen {
		fmt.Println("error opening the file")
	} else if err == ErrDivideByZero {
		fmt.Println("divide by zero error")
	} else {
		fmt.Printf("Unknown error, error : %s\n", err.Error())
	}
}
