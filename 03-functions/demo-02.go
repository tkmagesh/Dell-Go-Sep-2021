package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(10, 20, 30, 40))
	fmt.Println(sum(10, 20, 30, 40, 50, 60))
}

func sum(nos ...int) int {
	//fmt.Printf("nos = %#v\n", nos)
	result := 0
	/*
		for idx := 0; idx < len(nos); idx++ {
			result += nos[idx]
		}
	*/
	for _, no := range nos {
		result += no
	}
	return result
}
