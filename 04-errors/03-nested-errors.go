package main

import (
	"errors"
	"fmt"
)

var ErrF1 = errors.New("f1 error")
var ErrF2 = errors.New("f2 error")

func main() {
	err := f1()
	if err == nil {
		fmt.Println("Operation successful")
		return
	}
	fmt.Println(err)
	if errors.Is(err, ErrF1) {
		fmt.Println("f1 error occurred")
	}
	if errors.Is(err, ErrF2) {
		fmt.Println("f2 error occurred")
	}
}

func f1() error {
	f2Err := f2()
	return fmt.Errorf("%w, %w", ErrF1, f2Err)
}

func f2() error {
	return ErrF2
}
