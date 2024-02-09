package main

import (
	"fmt"
	"math"
)

// sprint-1
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// sprint-2
type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

/*
func printArea(x interface{}) {
	switch s := x.(type) {
	case Circle:
		fmt.Println("Area :", s.Area())
	case Rectangle:
		fmt.Println("Area :", s.Area())
	default:
		fmt.Println("argument has not Area() method")
	}

}
*/

/*
func printArea(x interface{}) {
	switch s := x.(type) {
	case interface{ Area() float64 }: // any object that has an Area() method
		fmt.Println("Area :", s.Area())
	default:
		fmt.Println("argument has not Area() method")
	}
}
*/

func printArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

// sprint-3
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

func printPerimeter(x interface{ Perimeter() float64 }) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// sprint-4
/*
func printStats(x interface {
	Area() float64
	Perimeter() float64
}) {
	fmt.Println("Area :", x.Area())
	fmt.Println("Perimeter :", x.Perimeter())
}
*/

// interface composition
func printStats(x interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}) {
	printArea(x)
	printPerimeter(x)
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		printArea(c)
		printPerimeter(c)
	*/
	printStats(c)

	r := Rectangle{Length: 10, Breadth: 12}
	// fmt.Println("Area :", r.Area())
	/*
		printArea(r)
		printPerimeter(r)
	*/
	printStats(r)
}
