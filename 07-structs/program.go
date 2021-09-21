package main

import "fmt"

type Product struct {
	Id              int
	Name            string
	Units           int
	Cost            float64
	ProductCategory *Category
}

type Category struct {
	Name     string
	Seasonal bool
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

	//var p *Product = &Product{Id: 101, Name: "MacBook Pro", Units: 2, Cost: 12500, ProductCategory: &Category{Name: "Computers", Seasonal: false}}
	var p *Product = &Product{101, "Macbook Pro", 2, 12500, &Category{"Computers", false}}
	fmt.Printf("%#v\n", p)
	fmt.Println(p.ProductCategory.Name)

	fmt.Println("Applying discount of 10%")
	applyDiscount(p, 10)
	fmt.Printf("%#v\n", p)
}

func applyDiscount(p *Product, discount int) {
	p.Cost = p.Cost * ((100 - float64(discount)) / 100)
}
