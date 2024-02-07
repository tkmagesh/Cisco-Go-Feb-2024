/*
Build an interactive calculator
- Display the following menu
	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit
- Accept the user choice
- if user choice is 5
	then exit
- if user choice > 5 OR < 1
	print "invalid input"
	display the menu again
- if user choice >=1 OR <= 4
	accept 2 operands
	perform the corresponding operation based on the user choice
	print the result
	display the menu again

Assumption:
	The user will always enter valid data
*/

package main

import "fmt"

func main() {
	var n1, n2, userChoice, result int
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&userChoice)
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		fmt.Println("Enter the operands :")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result :", result)
	}
}
