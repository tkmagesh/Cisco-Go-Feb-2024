package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 2, "marker": 3}
	// var productRanks = map[string]int{"pen": 5, "pencil": 2, "marker": 3}
	// productRanks := map[string]int{"pen": 5, "pencil": 2, "marker": 3}
	/*
		productRanks := map[string]int{
			"pen":    5,
			"pencil": 2,
			"marker": 3,
		}
	*/

	// var productRanks map[string]int = map[string]int{}
	var productRanks map[string]int = make(map[string]int)
	productRanks["pen"] = 5
	productRanks["pencil"] = 2
	productRanks["marker"] = 3
	fmt.Println(productRanks)
	fmt.Println("len(productRanks) :", len(productRanks))

	fmt.Println("iterating a map")
	for key, value := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, value)
	}

	// check for the existence of a key
	// keyToCheck := "fountain-pen"
	keyToCheck := "pen"
	// fmt.Println(productRanks[keyToCheck])
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key [%q] does not exist!\n", keyToCheck)
	}

	// remove an item
	// keyToRemove := "pen"
	keyToRemove := "fountain-pen"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)

}
