package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hi there!")
	}()

	// function with 1 parameter
	func(name string) {
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	}("Magesh")

	// function with 2 parameters
	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	// function with 2 parameters & 1 return result
	msg := func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}("Suresh", "Kannan")
	fmt.Print(msg)

	// using named results
	q, r := func(x, y int) (quotient int, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("quotient = %d, remainder = %d\n", q, r)
}
