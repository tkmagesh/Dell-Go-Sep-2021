package main

import (
	"fmt"
	"time"
)

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 10, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	resultCh := make(chan int)
	firstSet := data[:len(data)/2]
	secondSet := data[len(data)/2:]
	go func() {
		result := sum(firstSet...)
		fmt.Println("Sum completed. Writing the first result")
		resultCh <- result
		fmt.Println("first result written")
	}()
	go func() {
		result := sum(secondSet...)
		fmt.Println("Sum completed. Writing the second result")
		resultCh <- result
		fmt.Println("second result written")
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Attempting to read the results")
	result := <-resultCh + <-resultCh
	fmt.Println(result)
}

func sum(nos ...int) int {
	result := 0
	for _, no := range nos {
		result += no
	}
	return result

}
