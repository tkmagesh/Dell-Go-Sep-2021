package main

import "fmt"

func main() {
	//Array
	fmt.Println("Array")

	// Create an array of integers with the default value of int
	//var nos [5]int

	//Create an array of integers with data
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos [5]int = [...]int{3, 1, 4, 2, 5}
	nos := [...]int{3, 1, 4, 2, 5}
	fmt.Printf("%T -> %v\n", nos, nos)

	//iterating ovar an array (using the indexer)
	/*
		for idx := 0; idx < len(nos); idx++ {
			fmt.Printf("%d \n", nos[idx])
		}
	*/

	//iterating ovar an array (using the range construct)
	/*
		for idx, value := range nos {
			fmt.Printf("%d -> %d \n", idx, value)
		}
	*/
	for _, value := range nos {
		fmt.Printf("%d \n", value)
	}

	//Slice
	//Create a slice of integers with the default value of int
	fmt.Println("Slice")
	//var products []string
	var products []string = make([]string, 0, 3)
	fmt.Printf("len = %d, capacity = %d\n", len(products), cap(products))
	fmt.Printf("%T -> %v\n", products, products)
	fmt.Printf("Capacity of products => %v\n", cap(products))

	//adding items to the slice
	fmt.Println("Appending to the slice")
	products = append(products, "iPod")
	products = append(products, "iPad")
	products = append(products, "iPhone")
	products = append(products, "Macbook")
	fmt.Printf("%T -> %v\n", products, products)
	fmt.Printf("len => %d, capacity => %d\n", len(products), cap(products))

	//slicing elements
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to the end
		[ : hi] => from the beginning to hi-1
		[ lo : lo ] => []
		[lo : lo+1] == [lo]
		[:] => copy of the slice
	*/
	fmt.Println("slicing elements", products[1:3])

	//Creating a copy of slice
	fmt.Println("Creating a copy of slice")
	copyOfProducts := products[:]
	products[0] = "Printer"
	fmt.Printf("%v , %v\n", products, copyOfProducts)

	//appending a value to the 'products' slice
	fmt.Println()
	fmt.Println("appending a value to the 'products' slice")
	products = append(products, "Marker")

	fmt.Printf("products, len => %d, capacity => %d\n", len(products), cap(products))
	fmt.Printf("copyOfProducts, len => %d, capacity => %d\n", len(copyOfProducts), cap(copyOfProducts))

	//??? What happens when you append an item to the copyOfProducts slice?

	//Map
	fmt.Printf("\nMap\n")
	//var productRanks = make(map[string]int, 5)
	var productRanks = map[string]int{
		"Pen":    1,
		"Pencil": 2,
		"Marker": 5,
		"Eraser": 3,
	}
	fmt.Println(productRanks, len(productRanks))
	fmt.Printf("\nIterating a map\n")
	for key, value := range productRanks {
		fmt.Printf("Key => %v, Value => %v\n", key, value)
	}

	fmt.Printf("\nAdding a new value\n")
	productRanks["Scribble Pad"] = 4
	fmt.Println(productRanks, len(productRanks))

	fmt.Printf("\nRemoving an item\n")
	delete(productRanks, "Pencil")
	fmt.Println(productRanks, len(productRanks))

	fmt.Printf("\nVerifying the existence of a key\n")
	rank, exists := productRanks["Mouse"]
	if exists {
		fmt.Println("Mouse is in the map with rank ", rank)
	} else {
		fmt.Println("Mouse is not in the map")
	}
	fmt.Println(rank)
}
