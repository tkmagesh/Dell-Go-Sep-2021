package main

import (
	"fmt"
	"sync"
	"time"
)

//communicate by sharing memory [ NOT ADVISABLE IN GOLANG]
var result int

var mutex sync.Mutex = sync.Mutex{}

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	wg.Add(1)
	go add(100, 200)
	go add(100, 200)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int) {
	time.Sleep(time.Second * 4)
	//mutex.Lock()
	//{
	result = x + y
	//}
	//mutex.Unlock()
	wg.Done()
}
