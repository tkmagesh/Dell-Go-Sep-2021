package main

import (
	"fmt"
	"time"
)

func main() {
	print("Hello", 3*time.Second)
	print("World", 1*time.Second)
	fmt.Println("Exiting main")
}

func print(s string, delay time.Duration) {
	for i := 0; i < 5; i++ {
		time.Sleep(delay)
		fmt.Println(s)
	}
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
