package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// size of the slice
	fmt.Println("len(nos) :", len(nos))

	// appending new items
	nos = append(nos, 10)
	fmt.Println(nos)

	nos = append(nos, 20, 30, 40)
	fmt.Println(nos)

	// appending another slice
	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	fmt.Println("iteration - Using the indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("iteration - Using the range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		newNos := nos
		newNos[0] = 9999

		fmt.Println(newNos[0])
		fmt.Println(nos[0])
	*/

	// slicing
	s2 := nos[0:3]
	fmt.Println("s2[0] :", s2[0])
	fmt.Println("s2[1] :", s2[1])
	fmt.Println("s2[2] :", s2[2])
	// fmt.Println("s2[3] :", s2[3])

	fmt.Println("nos[:3] :", nos[:3])
	fmt.Println("nos[3:] :", nos[3:])

	bubbleSort(nos)
	fmt.Println(nos)

}

func bubbleSort(nos []int) /* no return values */ {
	for i := 0; i < len(nos)-1; i++ {
		for j := i + 1; j < len(nos); j++ {
			if nos[i] > nos[j] {
				nos[i], nos[j] = nos[j], nos[i]
			}
		}
	}
}
