package main

import (
	"fmt"
	"myapp/collections"
	"myapp/model"
)

type MyProduct model.Product

func (p MyProduct) WhoAmI() {
	fmt.Println("I am a product")
}

func main() {
	stationaryProduct := model.Product{104, "S-Pad", 20, 20, "Stationary"}
	fmt.Println(stationaryProduct.Format())
	stationaryProduct.ApplyDiscount(20)
	fmt.Println(stationaryProduct.Format())

	MyProduct(stationaryProduct).WhoAmI()

	//collections
	products := collections.Products{
		model.Product{105, "Pen", 5, 50, "Stationary"},
		model.Product{107, "Pencil", 2, 100, "Stationary"},
		model.Product{103, "Marker", 50, 20, "Utencil"},
		model.Product{102, "Stove", 5000, 5, "Utencil"},
		model.Product{101, "Kettle", 2500, 10, "Utencil"},
		model.Product{104, "S-Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products.Format())

	fmt.Printf("\nIndexOf\n")
	fmt.Println(products.IndexOf(stationaryProduct))

	fmt.Printf("\nIncludes\n")
	fmt.Println(products.Includes(stationaryProduct))

	fmt.Printf("\nFilter\n")
	fmt.Println("Stationary Products")
	//stationaryProducts := FilterStationaryProducts(products)
	stationaryProductPredicate := func(product model.Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts.Format())
	fmt.Println("Costly Products")
	//costlyProducts := FilterCostlyProducts(products)
	costlyProducts := products.Filter(func(product model.Product) bool {
		return product.Cost > 100
	})
	fmt.Println(costlyProducts.Format())

	fmt.Printf("\nAny\n")
	fmt.Println("Are there any stationary products ? : ", products.Any(stationaryProductPredicate))

	fmt.Printf("\nAll\n")
	fmt.Println("Are all the products stationary products ? : ", products.All(stationaryProductPredicate))

	//methods accessibility through composition
	grapes := model.PerishableProduct{model.Product{101, "Grapes", 10, 10, "Food"}, "2 Days"}
	fmt.Println("Before discount => ", grapes.Format())
	grapes.ApplyDiscount(10)
	fmt.Println("After discount => ", grapes.Format())

}
