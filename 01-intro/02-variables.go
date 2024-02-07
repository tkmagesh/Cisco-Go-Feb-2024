package main

import "fmt"

/* package scope */
// var app_name string = "hello_world"
var app_name = "hello_world"

// app_name := "hello_world" // invalid

func main() {
	/*
		var username string
		username = "Magesh"
	*/

	/*
		var username string = "Magesh"
	*/

	// type inference
	// var username = "Magesh"

	// using :=
	username := "Magesh"
	fmt.Println(username)

	/* handling multiple variables */
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	// multi assignment
	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y int    = 100, 200
			str  string = "Sum of %d and %d is %d\n"
		)
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	// type inference

	/*
		var x = 100
		var y = 200
		var str = "Sum of %d and %d is %d\n"
		var result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x      = 100
			y      = 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// using :=
	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)

}
