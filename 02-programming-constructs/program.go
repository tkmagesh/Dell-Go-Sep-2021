package main

import "fmt"

func main() {

	//if else construct
	no := 21
	if no%2 == 0 {
		fmt.Println("no is even")
	} else {
		fmt.Println("no is odd")
	}
	fmt.Println("no is", no)

	/*
		if no := 20; no%2 == 0 {
			fmt.Println("no is even")
		} else {
			fmt.Println("no is odd")
		}
		fmt.Println("no is", no)
	*/

	//for construct
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//simulating a while construct
	nosSum := 1
	for nosSum < 100 {
		nosSum += nosSum
	}
	fmt.Println(nosSum)

	//indefinite loop
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println(x)
}
