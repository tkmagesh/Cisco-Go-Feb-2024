package main

import "fmt"

func main() {
	var x int16 = 100

	// type conversion syntax -> use the typename like a function
	var f float64 = float64(x)
	fmt.Println(f)
}
