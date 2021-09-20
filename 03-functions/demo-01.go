package main

import "fmt"

func main() {
	fmt.Println(add(100, 200))
	/*
		q, r := divide(100, 7)
		fmt.Println(q, r)
	*/
	q, _ := divide(100, 7)
	fmt.Println(q)
}

/*
func add(x int, y int) int {
	return x + y
}
*/

/*
func add(x, y int) int {
	return x + y
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
