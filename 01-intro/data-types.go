package main

import "fmt"

func main() {
	/*
		var x int
		var y int
		var msg string

		x = 10
		y = 20
		msg = "result = "
	*/

	/*
		var x,y int
		var msg string

		x = 10
		y = 20
		msg = "result = "
	*/

	/*
		var (
			x, y int
			msg  string
		)
		x = 10
		y = 20
		msg = "result = "
	*/

	/*
		var (
			x   int = 10
			y   int = 20
			msg  string   = "result = "
		)
	*/

	/*
		var (
			x, y int    = 10, 20
			msg  string = "result = "
		)
	*/

	/* var (
		x, y = 10, 20
		msg  = "result = "
	) */

	/*
		var x, y, msg = 10, 20, "result = "
	*/

	x, y, msg := 10, 20, "result = "

	//fmt.Println(x, y, msg)
	fmt.Printf("x = %d, y = %d and %s %d\n", x, y, msg, x+y)

	//type conversion
	var noInt int = 100
	var noFloat float64
	noFloat = float64(noInt)
	fmt.Println(noInt, noFloat)
}
