package main

import "fmt"

func main() {
	adder := getAdder()
	fmt.Println(adder(100, 200))
}

func getAdder() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}
