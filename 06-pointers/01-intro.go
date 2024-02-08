package main

import "fmt"

func main() {
	var i int
	i = 100

	var iPtr *int
	iPtr = &i // reference (address) of a value

	//dereferencing - accessing the value in the underlying address
	fmt.Println(*iPtr)

	fmt.Println(i == *(&i))

	fmt.Println("Before incrementing, i =", i)
	fmt.Println("[main] &i =", &i)
	increment(&i)
	fmt.Println("After incrementing, i =", i)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(no *int) /* no return values */ {
	fmt.Println("[increment] &no =", no)
	// increment the given value by 1
	*no++
}

func swap(x, y *int) /* no return values */ {
	*x, *y = *y, *x
}
