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

func main() {

	/*
		var pen Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	pen := NewProduct(100, "Pen", 10)

	// fmt.Println(FormatProduct(pen))
	fmt.Println(pen.Format())

	fmt.Println("After applying 10% discount")
	// ApplyDiscount(&pen, 10)
	pen.ApplyDiscount(10)
	// fmt.Println(FormatProduct(pen))
	fmt.Println(pen.Format())

}
