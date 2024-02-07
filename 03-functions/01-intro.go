package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	greetUser("Magesh", "Kuppan")
	fmt.Printf(getGreetMsg("Suresh", "Kannan"))

	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("quotient = %d, remainder = %d\n", q, r)

	// use only quotient
	/*
		q, _ := divide(100, 7)
		fmt.Printf("quotient = %d\n", q)
	*/
}

func sayHi() {
	fmt.Println("Hi there!")
}

// function with 1 parameter
func greet(name string) {
	fmt.Printf("Hi %s, Have a nice day!\n", name)
}

// function with 2 parameters
/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

// function with 2 parameters & 1 return result
func getGreetMsg(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

// function with 2 parameters & 2 return results
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// using named results
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
