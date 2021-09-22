package main

import (
	"io"
	"log"
	"os"
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
	fileReader, err := os.Open("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	alphaReader := AlphaReader{Src: fileReader}
	io.Copy(os.Stdout, alphaReader)
}
