package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for no := range ch {
		fmt.Println(no)
	}
}

func fn(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
	close(ch)

}
