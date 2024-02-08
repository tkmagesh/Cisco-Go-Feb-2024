/*
	Refactor the below using functions
	Ensure that each function follows SRP
*/

package main

import "fmt"

func main() {
	var (
		userChoice, n1, n2, result int
		operation                  func(int, int) int
	)

	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		n1, n2 = getOperands()
		operation = getOperation(userChoice)
		result = operation(n1, n2)
		fmt.Println("Result :", result)
	}
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice :")
	fmt.Scanln(&userChoice)
	return userChoice
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the operands :")
	fmt.Scanln(&n1, &n2)
	return
}

func getOperation(userChoice int) func(int, int) int {
	switch userChoice {
	case 1:
		return add
	case 2:
		return subtract
	case 3:
		return multiply
	case 4:
		return divide
	default:
		return func(i1, i2 int) int { return 0 }
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
