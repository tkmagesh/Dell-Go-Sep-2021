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
		{104, "S-Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(Format(products))

	fmt.Printf("\nIndexOf\n")
	stationaryProduct := Product{104, "S-Pad", 20, 20, "Stationary"}
	fmt.Println(IndexOf(products, stationaryProduct))

	fmt.Printf("\nIncludes\n")
	fmt.Println(Includes(products, stationaryProduct))

	fmt.Printf("\nFilter\n")
	fmt.Println("Stationary Products")
	//stationaryProducts := FilterStationaryProducts(products)
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := Filter(products, stationaryProductPredicate)
	fmt.Println(Format(stationaryProducts))
	fmt.Println("Costly Products")
	//costlyProducts := FilterCostlyProducts(products)
	costlyProducts := Filter(products, func(product Product) bool {
		return product.Cost > 100
	})
	fmt.Println(Format(costlyProducts))

	fmt.Printf("\nAny\n")
	fmt.Println("Are there any stationary products ? : ", Any(products, stationaryProductPredicate))

	fmt.Printf("\nAll\n")
	fmt.Println("Are all the products stationary products ? : ", All(products, stationaryProductPredicate))

}
func Format(products []Product) string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%d\t %s\t\t %.2f\t %d\t %s\n", p.Id, p.Name, p.Cost, p.Units, p.Category)
	}
	return result
}

/*
IndexOf() => index of the given product object
*/
func IndexOf(products []Product, product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

/*
Includes() => true/false depending on the presence of the given product in the collection
*/

func Includes(products []Product, product Product) bool {
	return IndexOf(products, product) != -1
}

/*
Filter() => filter the products based on the the given condition
			ex., filter all stationary products
				 filter all costly products (cost > 100)
*/

/* func FilterStationaryProducts(products []Product) []Product {
	result := make([]Product, 0)
	for _, p := range products {
		if p.Category == "Stationary" {
			result = append(result, p)
		}
	}
	return result
}

func FilterCostlyProducts(products []Product) []Product {
	result := make([]Product, 0)
	for _, p := range products {
		if p.Cost > 100 {
			result = append(result, p)
		}
	}
	return result
} */

func Filter(products []Product, predicate func(Product) bool) []Product {
	result := make([]Product, 0)
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

/*
Any() => true/false depending on the presence of any product in the collection based on the given condition
*/

func Any(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

/*
All() => true/false depending on all products in the collection satifying the given condition
*/
func All(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}
