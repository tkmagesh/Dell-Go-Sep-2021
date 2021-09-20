package main

import "fmt"

func main() {
	increment := getIncrementer()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}

func getIncrementer() func() int { //step 1
	var counter int = 0       // step 2
	increment := func() int { //step 3
		counter++ //step 4
		return counter
	}
	return increment //step 5
}
