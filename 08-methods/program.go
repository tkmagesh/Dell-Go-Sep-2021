package main

import "fmt"

type Calculator struct{}

/*
type Calculator struct {
	x      int
	y      int
	result int
}
*/

func (c Calculator) add(x, y int) int {
	return x + y
}

func (c Calculator) subtract(x, y int) int {
	return x - y
}

func main() {
	var calc Calculator
	fmt.Println(calc.add(42, 13))
	fmt.Println(calc.subtract(42, 13))
}
