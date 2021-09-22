package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	//share memory for communication
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, ch)
	result := <-ch
	wg.Wait()
	fmt.Println(result)

}

func add(x, y int, ch chan int) {
	time.Sleep(time.Second * 4)
	result := x + y
	ch <- result
	wg.Done()
}
