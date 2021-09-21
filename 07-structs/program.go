package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

func main() {
	/*
		var p Product
		p.Id = 101
		p.Name = "MacBook Pro"
		p.Units = 2
		p.Cost = 12500
		p.Category = "Computers"
	*/

	/*
		var p Product = Product{Id: 101, Name: "MacBook Pro", Units: 2, Cost: 12500, Category: "Computers"}
	*/

	/*
		var p *Product
		p = new(Product)
		(*p).Id = 101
		p.Name = "MacBook Pro"
		p.Units = 2
		p.Cost = 12500
		p.Category = "Computers"
		fmt.Println(p)
	*/

	var p *Product = &Product{Id: 101, Name: "MacBook Pro", Units: 2, Cost: 12500, Category: "Computers"}
	fmt.Println(p)
}
