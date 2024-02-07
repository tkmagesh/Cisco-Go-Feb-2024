package main

import "fmt"

func main() {
	fmt.Println("v1.0")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("v2.0 [while]")
	/*
		i := 0
		for i < 10 {
			fmt.Println(i)
			i++
		}
	*/

	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	fmt.Println("v3.0 [infinite]")
	no := 1
	for {
		fmt.Println(no)
		no += no
		if no > 128 {
			break
		}
	}

	fmt.Println("v4.0 [label]")
I_LOOP:
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				fmt.Println("====================")
				continue I_LOOP
				// break
			}
		}

	}
}
