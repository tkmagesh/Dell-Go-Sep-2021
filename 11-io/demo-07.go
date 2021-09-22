package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data3.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}
