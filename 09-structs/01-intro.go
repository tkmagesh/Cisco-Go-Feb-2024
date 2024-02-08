package main

import "fmt"

func main() {
	var product struct {
		Id   int
		Name string
		Cost float32
	}
	// fmt.Println(product)
	// fmt.Printf("%#v\n", product)
	// fmt.Printf("%+v\n", product)
	fmt.Println(FormatProduct(product))

	var pen struct {
		Id   int
		Name string
		Cost float32
	} = struct {
		Id   int
		Name string
		Cost float32
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// fmt.Printf("%+v\n", pen)
	fmt.Println(FormatProduct(pen))
}

func FormatProduct(product struct {
	Id   int
	Name string
	Cost float32
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
}
