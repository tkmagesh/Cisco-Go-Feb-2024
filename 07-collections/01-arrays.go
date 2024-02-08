package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// using :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("iteration - Using the indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("iteration - Using the range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	newNos := nos // a new copy of the array is created (coz everything is a value)
	newNos[0] = 9999

	fmt.Println(newNos[0]) //=> 9999
	fmt.Println(nos[0])

	// array pointers
	var nosPtr *[5]int

	// reference
	nosPtr = &nos

	// deference
	fmt.Println((*nosPtr)[0])
	fmt.Println(nosPtr[0])

	bubbleSort(&nos) //sort the nos array
	fmt.Println(nos)
}

func bubbleSort(nos *[5]int) /* no return values */ {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if nos[i] > nos[j] {
				nos[i], nos[j] = nos[j], nos[i]
			}
		}
	}
}
