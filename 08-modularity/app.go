package main

import (
	"fmt"

	"github.com/tkmagesh/cisco-go-feb-2024/08-modularity/calculator"
	ut "github.com/tkmagesh/cisco-go-feb-2024/08-modularity/calculator/utils"
)

func main() {
	fmt.Println("Application executed")

	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.GetOpCount())

	// fmt.Println(utils.IsEven(20))
	fmt.Println(ut.IsOdd(21))
}
