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

}
