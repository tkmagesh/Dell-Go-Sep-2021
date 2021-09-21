package collections

import "myapp/model"

type Products []model.Product

func (products Products) Format() string {
	result := ""
	for _, p := range products {
		result += p.Format()
	}
	return result
}

/*
IndexOf() => index of the given product object
*/
func (products Products) IndexOf(product model.Product) int {
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

func (products Products) Includes(product model.Product) bool {
	return products.IndexOf(product) != -1
}

/*
Filter() => filter the products based on the the given condition
			ex., filter all stationary products
				 filter all costly products (cost > 100)
*/

func (products Products) Filter(predicate func(model.Product) bool) Products {
	result := make(Products, 0)
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

func (products Products) Any(predicate func(model.Product) bool) bool {
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
func (products Products) All(predicate func(model.Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}
