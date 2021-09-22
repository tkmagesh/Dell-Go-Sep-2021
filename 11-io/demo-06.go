package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data2.txt")
	defer fileReader.Close()
	if err != nil {
		log.Fatalln(err)
	}
	bufferedReader := bufio.NewReader(fileReader)
	for {
		sentence, err := bufferedReader.ReadString('.')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(sentence)
	}
}
