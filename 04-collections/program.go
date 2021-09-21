package main

import "fmt"

func main() {
	//Array

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
}
