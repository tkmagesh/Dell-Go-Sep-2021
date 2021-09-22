package main

import (
	"io"
	"os"
	"strings"
)

type AlphaReader struct {
	Src io.Reader
}

func (alphaReader AlphaReader) Read(p []byte) (int, error) {
	inputData := make([]byte, len(p))
	count, err := alphaReader.Src.Read(inputData)
	if err != nil {
		return count, err
	}
	dataCount := 0
	for idx := 0; idx < len(inputData); idx++ {
		if (inputData[idx] >= 'a' && inputData[idx] <= 'z') || (inputData[idx] >= 'A' && inputData[idx] <= 'Z') {
			p[dataCount] = inputData[idx]
			dataCount++
		}
	}
	return dataCount, io.EOF
}

func main() {
	stringReader := strings.NewReader("this is a sample text with numbers like 123344667 and special characters like &*^$#@&()&^. The alphareader is expected to filter all the numbers & special character")
	alphaReader := AlphaReader{Src: stringReader}
	io.Copy(os.Stdout, alphaReader)
}
