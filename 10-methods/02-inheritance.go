package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (product *Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
}

func (productPtr *Product) ApplyDiscount(discountPercentage float32) {
	// attributes can be accessed using the "." notation even with a pointer
	productPtr.Cost = productPtr.Cost * ((100 - discountPercentage) / 100)
}

// factory function
func NewProduct(id int, name string, cost float32) *Product {
	newProduct := &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
	return newProduct
}

/* Perishable Product */

type PerishableProduct struct {
	Product // composition
	Expiry  string
}

// method overriding
func (pp PerishableProduct) Format() string {
	// return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Expiry = %q", pp.Id, pp.Name, pp.Cost, pp.Expiry)
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {
	grapes := PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Grapes",
			Cost: 50,
		},
		Expiry: "2 Days",
	}
	fmt.Println(grapes.Id)
	fmt.Println(grapes.Format())

	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())
}
