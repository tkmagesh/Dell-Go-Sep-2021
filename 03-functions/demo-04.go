package main

import "fmt"

func main() {

	/*
		fmt.Println("before invocation")
		fmt.Println(add(100, 200))
		fmt.Println("after invocation")

		fmt.Println("before invocation")
		fmt.Println(subtract(100, 200))
		fmt.Println("after invocation")
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		multiply := func(x, y int) int {
			return x * y
		}
		logOperation(multiply, 100, 200)
	*/

	loggerAdd := logOperation(add)
	loggerAdd(100, 200)
	loggerSubtract := logOperation(subtract)
	loggerSubtract(100, 200)
}

func logOperation(oper func(int, int) int) func(int, int) {
	return func(x, y int) {
		fmt.Println("before invocation")
		fmt.Println(oper(x, y))
		fmt.Println("after invocation")
	}
}

func add(x, y int) int {
	result := x + y
	return result
}

func subtract(x, y int) int {
	return x - y
}
