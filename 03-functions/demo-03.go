package main

import "fmt"

func main() {
	/* functions assigned to variables */

	/* fn := func() {
		fmt.Println("fn is invoked!")
	} */

	var fn func()
	fn = func() {
		fmt.Println("fn is invoked!")
	}
	fn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	/* anonymous functions */
	func() {
		fmt.Println("anonymous function invoked")
	}()

	/* anonymous function with arguments */
	func(x int, y int) {
		fmt.Println(x, y)
	}(100, 200)

	/* anonymous function with return value */
	/*
		result := func(x, y int) int {
			return x + y
		}(100, 200)
		fmt.Println(result)
	*/

	fmt.Println(func(x, y int) int {
		return x + y
	}(100, 200))
}
