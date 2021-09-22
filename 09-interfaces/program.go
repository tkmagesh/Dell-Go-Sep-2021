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

//implementing the fmt.Stringer interface
func (c Circle) String() string {
	return fmt.Sprintf("Circle: radius = %f", c.Radius)
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * 2
}

type Rectange struct {
	Width  float64
	Height float64
}

func (r Rectange) Area() float64 {
	return r.Width * r.Height
}

func (r Rectange) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectange) String() string {
	return fmt.Sprintf("Rectange: width = %f, height = %f", r.Width, r.Height)
}

type ShapeWithArea interface {
	Area() float64
}

func printArea(sa ShapeWithArea) {
	fmt.Printf("Area = %f\n", sa.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func PrintPerimeter(sp ShapeWithPerimeter) {
	fmt.Printf("Perimeter = %f\n", sp.Perimeter())
}

/*
type Shape interface {
	Area() float64
	Perimeter() float64
}
*/

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func printShape(s Shape) {
	printArea(s)
	PrintPerimeter(s)
}

func main() {
	c := Circle{Radius: 5}
	/*
		printArea(c)
		PrintPerimeter(c)
	*/
	fmt.Println(c)
	printShape(c)

	r := Rectange{Width: 10, Height: 5}
	/*
		printArea(r)
		PrintPerimeter(r)
	*/
	fmt.Println(r)
	printShape(r)
}
