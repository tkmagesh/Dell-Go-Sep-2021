package main

import (
	"fmt"
	"time"
)

func main() {
	//share memory for communication
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
	fmt.Println("Exiting main")
}

func add(x, y int, ch chan int) {
	time.Sleep(time.Second * 4)
	result := x + y
	ch <- result
	fmt.Println("Exiting add")
}
