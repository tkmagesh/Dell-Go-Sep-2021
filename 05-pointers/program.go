package main

import "fmt"

func main() {
	var no int = 10
	//noPtr := &no
	var noPtr *int = &no
	fmt.Println(noPtr, no)

	//deferencing the pointer to get the underlying value
	var n = *noPtr
	fmt.Println(n)

	x, y := 10, 20
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
