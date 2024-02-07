package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("Enter your name  :")
		fmt.Scanln(&name)
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	var firstName, lastName string
	fmt.Println("Enter your name [firstname lastname] :")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)

	var x, y int
	fmt.Println("Enter the numbers")
	fmt.Scanln(&x, &y)
	fmt.Println(x + y)
}
