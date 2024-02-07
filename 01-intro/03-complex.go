package main

import "fmt"

func main() {
	var c1 complex128
	c1 = 3 + 5i
	fmt.Println(c1)

	c2 := 4 + 5i
	fmt.Println(c2)

	c3 := c1 + c2
	fmt.Println(c3)

	fmt.Printf("real(c3) = %v\n", real(c3))
	fmt.Printf("imag(c3) = %v\n", imag(c3))

}
