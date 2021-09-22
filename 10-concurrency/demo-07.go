package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	wg.Add(2)
	go print("Hello", 3*time.Second, ch1, ch2, wg)
	go print("World", 1*time.Second, ch2, ch1, wg)
	ch1 <- "start"
	wg.Wait()
	fmt.Println("Exiting main")
}

func print(s string, delay time.Duration, x, y chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-x
		time.Sleep(delay)
		fmt.Println(s)
		y <- "done"
	}
	wg.Done()
}

/*
Output:
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
Exiting main
*/
