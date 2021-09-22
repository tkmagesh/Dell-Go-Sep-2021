package main

import (
	"fmt"
)

func main() {
	go f1()
	f2()
	//time.Sleep(100 * time.Millisecond)
	/*
		var input string
		fmt.Scanln(&input)
	*/
}

func f1() {
	fmt.Println("f1 is invoked")
}

func f2() {
	fmt.Println("f2 is invoked")
}
