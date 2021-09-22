package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40))                                         //=> 100
	fmt.Println(sum(10))                                                     //=> 10
	fmt.Println(sum())                                                       //=> 0
	fmt.Println(sum(10, "20", 30, "40"))                                     //=> 100
	fmt.Println(sum(10, "20", 30, "40", "abc"))                              //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                                  //=> 100
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{10, 20}}))           //=> 130
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{10, "20"}})) //=> 130
}

func sum(nos ...interface{}) int {
	result := 0
	for _, no := range nos {
		switch v := no.(type) {
		case int:
			result += v
		case string:
			if val, err := strconv.Atoi(v); err == nil {
				result += val
			}
		case []int:
			intfList := make([]interface{}, len(v))
			for idx, x := range v {
				intfList[idx] = x
			}
			result += sum(intfList...)
		case []interface{}:
			result += sum(v...)
		}

	}
	return result
}
