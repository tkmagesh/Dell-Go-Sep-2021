package main

import (
	"io"
	"os"
)

type AlphaReader string

func (alphaReader AlphaReader) Read(p []byte) (n int, err error) {
	count := 0
	for idx := 0; idx < len(alphaReader); idx++ {
		if (alphaReader[idx] >= 'a' && alphaReader[idx] <= 'z') || (alphaReader[idx] >= 'A' && alphaReader[idx] <= 'Z') {
			p[count] = alphaReader[idx]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	alphaReader := AlphaReader("this is a sample text with numbers like 123344667 and special characters like &*^$#@&()&^. The alphareader is expected to filter all the numbers & special character")
	io.Copy(os.Stdout, alphaReader)
}

/*
modify the above program to consider the following
	1. this input string is bigger than the buffer size
	2. Write your own implementation of the Copy function
*/
