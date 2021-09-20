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

	//switch construct
	/*
		num := 21
		switch num % 2 {
		case 0:
			fmt.Println("num is even")
		case 1:
			fmt.Println("num is odd")
		}
	*/

	switch num := 21; num % 2 {
	case 0:
		fmt.Println("num is even")
	case 1:
		fmt.Println("num is odd")
	}

	/*
		comment by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	score := 18
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	//using fallthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("n is 0")
		fallthrough
	case 1:
		fmt.Println("n is <= 1")
		fallthrough
	case 2:
		fmt.Println("n is <= 2")
		fallthrough
	case 3:
		fmt.Println("n is <= 3")
		fallthrough
	case 4:
		fmt.Println("n is <= 4")
		fallthrough
	case 5:
		fmt.Println("n is <= 5")
		fallthrough
	case 6:
		fmt.Println("n is <= 6")
	case 7:
		fmt.Println("n is <=7")
		fallthrough
	case 8:
		fmt.Println("n is <= 8")
		fallthrough
	case 9:
		fmt.Println("n is <= 9")
	}
}
