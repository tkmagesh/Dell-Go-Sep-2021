package main

import (
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(fileContents))
}
