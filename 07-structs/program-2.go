package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

//Ver 1.0
/*
type PerishableProduct struct {
	Prod   Product
	Expiry string
}
*/

//Ver 2.0
type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, units int, cost float64, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Units:    units,
			Cost:     cost,
			Category: category,
		},
		Expiry: expiry,
	}
}

func main() {
	//Ver 1.0
	/*
		p := PerishableProduct{
			Prod: Product{
				Id:       1,
				Name:     "Cake",
				Units:    2,
				Cost:     1.50,
				Category: "Confectionary",
			},
			Expiry: "Dec 1, 2021",
		}
		println(p.Prod.Name)
		println(p.Prod.Units)
		println(p.Prod.Cost)
		println(p.Prod.Category)
		println(p.Expiry)
	*/

	//Ver 2.0
	/* p := PerishableProduct{
		Product: Product{
			Id:       1,
			Name:     "Cake",
			Units:    2,
			Cost:     1.50,
			Category: "Confectionary",
		},
		Expiry: "Dec 1, 2021",
	} */

	// println(p.Product.Name)
	// println(p.Product.Units)
	// println(p.Product.Cost)
	// println(p.Product.Category)
	// println(p.Expiry)

	p := NewPerishableProduct(1, "Cake", 2, 1.50, "Confectionary", "Dec 1, 2021")
	println(p.Name)
	println(p.Units)
	println(p.Cost)
	println(p.Category)
	println(p.Expiry)

	fmt.Println(p)
}
