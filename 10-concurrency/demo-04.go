package main

import (
	"fmt"
	"sync"
	"time"
)

//communicate by sharing memory [ NOT ADVISABLE IN GOLANG]
type Result struct {
	value int
	sync.Mutex
}

func (r *Result) SetValue(v int) {
	r.Lock()
	r.value = v
	r.Unlock()
}

func (r *Result) GetValue() int {
	return r.value
}

var result Result

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	wg.Add(1)
	add(100, 200)
	wg.Wait()
	fmt.Println(result.GetValue())
}

func add(x, y int) {
	time.Sleep(time.Second * 4)
	result.SetValue(x + y)
	wg.Done()
}
