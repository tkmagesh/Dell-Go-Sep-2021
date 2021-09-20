//package declaration
package main

//import other packages
import "fmt"

//package level variables & types
/*
	var msg string = "Hello World!"
*/
var msg = "Hello World!"

//package init function

//main function
func main() {

	/*
		var msg string
		msg = "Hello, World!"
	*/

	/*
		var msg string = "Hello, World!"
	*/

	//type inference
	/*
		var msg = "Hello, World!"
	*/

	/* The following syntax is application only in a function (NOT at package level) */
	//msg := "Hello, World!"

	fmt.Println(msg)
}

//other functions
