package calculator

import "fmt"

var opCount int

func init() {
	fmt.Println("calculator intialized")
}

func GetOpCount() int {
	return opCount
}
