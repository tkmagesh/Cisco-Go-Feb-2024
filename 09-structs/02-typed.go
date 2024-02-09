package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	var product Product
	// fmt.Println(product)
	// fmt.Printf("%#v\n", product)
	// fmt.Printf("%+v\n", product)
	fmt.Println(FormatProduct(product))

	var pen Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// fmt.Printf("%+v\n", pen)
	fmt.Println(FormatProduct(pen))

	/*
		var penPtr *Product = &pen
		fmt.Println(penPtr.Id)
	*/

	fmt.Println("After applying 10% discount")
	ApplyDiscount(&pen, 10)
	fmt.Println(FormatProduct(pen))

	p1 := Product{Id: 100, Name: "Pencil", Cost: 5}
	p2 := Product{Id: 100, Name: "Pencil", Cost: 5}
	fmt.Printf("&p1 = %p\n", &p1)
	fmt.Printf("&p2 = %p\n", &p2)
	fmt.Println("p1 == p2 ?", p1 == p2)
}

// factory function
func NewProduct(id int, name string, cost float32) *Product {
	newProduct := &Product{
		Id : id,
		Name : name,
		Cost : cost
	}
	return newProduct
}

func FormatProduct(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
}

func ApplyDiscount(productPtr *Product, discountPercentage float32) {
	// attributes can be accessed using the "." notation even with a pointer
	productPtr.Cost = productPtr.Cost * ((100 - discountPercentage) / 100)
}
