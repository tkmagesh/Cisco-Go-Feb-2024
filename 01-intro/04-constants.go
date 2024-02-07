package main

import "fmt"

// unused constant at package scope (allowed)
const y = 200

func main() {
	// const app_version string = "1.0"

	// type inference
	// const app_version = "1.0"

	const (
		pi          = 3.14
		app_version = "1.0"
	)
	fmt.Println(pi, app_version)

	// unused constant at function scope (allowed)
	const x = 100

	myVar := 100

}
