package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func fibonacci(ch chan int) {
	x, y := 0, 1
	for i := 0; i < 20; i++ {
		ch <- x
		time.Sleep(500 * time.Millisecond)
		x, y = y, x+y
	}
	close(ch)
}
