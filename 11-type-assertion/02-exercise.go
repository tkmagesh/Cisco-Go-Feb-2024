package main

import "fmt"

func main() {
	fmt.Println(sum(10))                      //=> 10
	fmt.Println(sum())                        //=> 0
	fmt.Println(sum(10, 20, 30, 40))          //=> 100
	fmt.Println(sum(10, 20, "30", 40))        //=> 100 (use "strconv" package apis)
	fmt.Println(sum(10, 20, "30", "abc", 40)) //=> 100 ("abc" should be ignored by the "sum" function)
}

func sum( /*  */ ) int {
	var result int
	/* ? */
	return result
}
