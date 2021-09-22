package main

import "fmt"

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	resultCh := make(chan int)
	firstSet := data[:len(data)/2]
	secondSet := data[len(data)/2:]
	go sum(resultCh, firstSet...)
	go sum(resultCh, secondSet...)
	result := <-resultCh + <-resultCh
	fmt.Println(result)
}

func sum(ch chan int, nos ...int) {
	result := 0
	for _, no := range nos {
		result += no
	}
	ch <- result

}
