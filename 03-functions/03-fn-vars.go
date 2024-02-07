package main

import "fmt"

func main() {

	var myPrint func(a ...any) (n int, err error)
	myPrint = fmt.Println

	var sayHi func()
	sayHi = func() {
		// fmt.Println("Hi there!")
		myPrint("Hi there!")
	}
	sayHi()

	sayHi = func() {
		// fmt.Println("Hello there!")
		myPrint("Hello there!")
	}
	sayHi()

	// function with 1 parameter
	var greet func(string)
	greet = func(name string) {
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	}
	greet("Magesh")

	// function with 2 parameters
	// var greetUser func(fn string, ln string)
	var greetUser func(string, string)
	greetUser = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greetUser("Magesh", "Kuppan")
	/*




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
	*/
}
