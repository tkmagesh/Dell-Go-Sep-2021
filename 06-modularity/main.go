package main

import (
	"fmt"
	calc "myapp/calculator"
	"myapp/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calc.Add(10, 20))
	fmt.Println(calc.Subtract(10, 20))
	fmt.Println(calc.GetOpCount())
	fmt.Println(utils.IsEven(100))
	color.Red("This is a string in red!")
}
