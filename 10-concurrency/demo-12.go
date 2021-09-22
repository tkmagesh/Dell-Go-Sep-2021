package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stopCh := make(chan bool)
	go fibonacci(ch, stopCh)
	go func() {
		for no := range ch {
			fmt.Println(no)
		}
	}()
	fmt.Println("hit ENTER to stop!")
	var input string
	fmt.Scanln(&input)
	stopCh <- true
	fmt.Println("Done")
}

func fibonacci(ch chan int, stopCh chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-stopCh:
			fmt.Println("Stopped")
			close(ch)
			return
		}
	}
}
