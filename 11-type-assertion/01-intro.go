package main

import "fmt"

func main() {
	// var x interface{}
	var x any // any => type alias for interface{}
	x = 100
	fmt.Println(x)

	x = "Ut deserunt anim duis velit amet cillum id voluptate."
	fmt.Println(x)

	x = true
	fmt.Println(x)

	x = 99.99
	fmt.Println(x)

	x = struct{}{}
	fmt.Println(x)

	// x = 200
	x = "Ex culpa magna eiusmod voluptate cillum proident Lorem quis incididunt veniam incididunt ea ex."
	// y := x.(int) * 10
	// type assertion - 1
	if val, ok := x.(int); ok {
		y := val * 10
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// type assertion - 2 (using 'type switch')
	// x = 100
	// x = "Ut deserunt anim duis velit amet cillum id voluptate."
	// x = true
	// x = 99.99
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case float64:
		fmt.Println("x is a float64, x * 0.09 =", val*0.09)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	default:
		fmt.Println("x is an unknown type")
	}

}
