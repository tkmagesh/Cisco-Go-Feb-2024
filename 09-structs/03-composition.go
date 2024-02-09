package main

import "fmt"

type Product struct {
	id   int
	Name string
	Cost float32
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product // composition
	Dummy   // composition
	Expiry  string
}

// factory
func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	// var grapes Product
	/*
		var grapes PerishableProduct
		grapes = PerishableProduct{
			Product: Product{Id: 100, Name: "Grapes", Cost: 50},
			Expiry:  "2 Days",
		}
	*/
	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days")
	fmt.Println(grapes.Expiry)
	fmt.Println(grapes.Product.Id)
	fmt.Println(grapes.Id)

	// fmt.Println(grapes.Name) //will not work due to conflict
	fmt.Println(grapes.Product.Name)
	fmt.Println(grapes.Dummy.Name)

}
