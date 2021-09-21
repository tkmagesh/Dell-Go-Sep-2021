package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	products := []Product{
		{105, "Pen", 5, 50, "Stationary"},
		{107, "Pencil", 2, 100, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5000, 5, "Utencil"},
		{101, "Kettle", 2500, 10, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println(products)
}

/*
IndexOf() => index of the given product object
Includes() => true/false depending on the presence of the given product in the collection
Filter() => filter the products based on the the given condition
			ex., filter all stationary products
				 filter all costly products (cost > 100)
Any() => true/false depending on the presence of any product in the collection based on the given condition
All() => true/false depending on all products in the collection satifying the given condition

*/
