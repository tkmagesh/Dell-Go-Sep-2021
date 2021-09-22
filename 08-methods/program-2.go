package main

import "fmt"

type Calculator struct {
	x      int
	y      int
	result int
}

//function
func add(calc *Calculator) {
	calc.result = calc.x + calc.y
}

func format(c Calculator) string {
	return fmt.Sprintf("x=%d, y=%d, result=%d", c.x, c.y, c.result)
}

//method

func (calc *Calculator) add() {
	calc.result = calc.x + calc.y
}

func (c Calculator) format() string {
	return fmt.Sprintf("x=%d, y=%d, result=%d", c.x, c.y, c.result)
}

func main() {

	var c = Calculator{}
	c.x = 100
	c.y = 200
	//add(&c)
	c.add()

	//fmt.Println(format(c))
	fmt.Println(c.format())
}
