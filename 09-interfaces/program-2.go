package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type ShapeWithArea interface {
	Area() float64
}

func main() {
	var x interface{}
	x = 1.0
	//fmt.Println(int(x))
	/*
		x = "A string value"
		x = true
		x = struct{}{}
	*/

	/* val, ok := x.(int); ok {
		y := val + 100
		fmt.Println(y)
	} */

	//x = []int{3, 1, 4, 2, 5}
	var shapeWithArea ShapeWithArea = Circle{Radius: 10}
	x = shapeWithArea

	switch v := x.(type) {
	case int:
		fmt.Println("x is an int", v)
	case string:
		fmt.Println("x is a string", v)
	case bool:
		fmt.Println("x is a bool", v)
	case float64:
		fmt.Println("x is a float34", v)
	case []int:
		fmt.Println("x is a slice of int", v)
	case ShapeWithArea:
		fmt.Println("x is a ShapeWithArea", v)
	default:
		fmt.Println("x is of a different type", v)

	}

}
